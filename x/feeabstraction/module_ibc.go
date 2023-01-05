package feeabstraction

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"

	"github.com/notional-labs/fa-chain/x/feeabstraction/keeper"
	"github.com/notional-labs/fa-chain/x/feeabstraction/types"
)

// IBCModule implements the ICS26 interface for interchain accounts controller chains
type IBCModule struct {
	keeper keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper) IBCModule {
	return IBCModule{
		keeper: k,
	}
}

func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) error {
	im.keeper.Logger(ctx).Info(fmt.Sprintf("OnChanOpenAck: portID %s, channelID %s", portID, channelID))
	// Note: The channel capability must be claimed by the authentication module in OnChanOpenInit otherwise the
	// authentication module will not be able to send packets on the channel created for the associated interchain account.
	if err := im.keeper.ClaimCapability(ctx, channelCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return err
	}
	im.keeper.Logger(ctx).Info(fmt.Sprintf("%s claimed the channel capability %v", types.ModuleName, channelCap))
	return nil
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	im.keeper.Logger(ctx).Info(fmt.Sprintf("OnChanOpenAck: portID %s, channelID %s, counterpartyChannelID %s, counterpartyVersion %s", portID, channelID, counterpartyChannelID, counterpartyVersion))
	controllerConnectionId, err := im.keeper.GetConnectionId(ctx, portID)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("Unable to get connection for port: %s", portID))
	}
	address, found := im.keeper.IcaControllerKeeper.GetInterchainAccountAddress(ctx, controllerConnectionId, portID)
	if !found {
		ctx.Logger().Error(fmt.Sprintf("Expected to find an address for %s/%s", controllerConnectionId, portID))
		return nil
	}
	// get host chain id from connection
	// fetch counterparty connection
	hostChainId, err := im.keeper.GetChainID(ctx, controllerConnectionId)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("Unable to obtain counterparty chain for connection: %s, port: %s, err: %s", controllerConnectionId, portID, err.Error()))
		return nil
	}

	// addresses
	feeAddress, err := icatypes.NewControllerPortID(types.GetFeeICAAccountOwner(hostChainId))
	if err != nil {
		return err
	}

	if portID == feeAddress {
		im.keeper.SetFeeICAAddress(ctx, address)
	} else {
		ctx.Logger().Error(fmt.Sprintf("Missing portId: %s", portID))
	}

	return nil
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	im.keeper.Logger(ctx).Info(fmt.Sprintf("OnAcknowledgementPacket: packet %v, relayer %v", modulePacket, relayer))
	ackInfo := fmt.Sprintf("sequence #%d, from %s %s, to %s %s",
		modulePacket.Sequence, modulePacket.SourceChannel, modulePacket.SourcePort, modulePacket.DestinationChannel, modulePacket.DestinationPort)

	var ack channeltypes.Acknowledgement
	err := channeltypes.SubModuleCdc.UnmarshalJSON(acknowledgement, &ack)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to unmarshal ack from feeabstraction OnAcknowledgePacket | Sequence %d, from %s %s, to %s %s",
			modulePacket.Sequence, modulePacket.SourceChannel, modulePacket.SourcePort, modulePacket.DestinationChannel, modulePacket.DestinationPort)
		im.keeper.Logger(ctx).Error(errMsg)
		return sdkerrors.Wrapf(types.ErrMarshalFailure, errMsg)
	}
	im.keeper.Logger(ctx).Info(fmt.Sprintf("Acknowledgement was successfully unmarshalled: ackInfo: %s", ackInfo))
	eventType := "ack"
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			eventType,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyAck, fmt.Sprintf("%v", ackInfo)),
		),
	)
	return nil
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {

	// WARNING: For some reason, in IBCv3 the ICA controller module does not call the underlying OnChanCloseConfirm (this function)
	// So, we need to put logic that _should_ execute upon channel closure in the OnTimeoutPacket function
	// This works because ORDERED channels are always closed when a timeout occurs, but if we migrate to using ORDERED channels that don't
	// close on timeout, we will need to move this logic to the OnChanCloseConfirm function
	// relevant IBCv3 code: https://github.com/cosmos/ibc-go/blob/5c0bf8b8a0f79643e36be98fb9883ea163d2d93a/modules/apps/27-interchain-accounts/controller/ibc_module.go#L123
	return nil
}

// ###################################################################################
// 	Helper functions
// ###################################################################################

func (im IBCModule) NegotiateAppVersion(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionID string,
	portID string,
	counterparty channeltypes.Counterparty,
	proposedVersion string,
) (version string, err error) {
	return proposedVersion, nil
}

// ###################################################################################
// 	Required functions to satisfy interface but not implemented for ICA auth modules
// ###################################################################################

// OnChanOpenTry implements the IBCModule interface
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	panic("UNIMPLEMENTED")
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	panic("UNIMPLEMENTED")
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	panic("UNIMPLEMENTED")
}

// OnRecvPacket implements the IBCModule interface
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	panic("UNIMPLEMENTED")
}
