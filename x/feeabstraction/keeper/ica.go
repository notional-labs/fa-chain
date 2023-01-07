package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/notional-labs/fa-chain/utils"
	"github.com/notional-labs/fa-chain/x/feeabstraction/types"

	gammtypes "github.com/osmosis-labs/osmosis/v13/x/gamm/types"
)

// SubmitTxs submits an ICA transaction containing multiple messages
// Will submit tx from fee account
func (k Keeper) SubmitTxs(
	ctx sdk.Context,
	connectionId string,
	msgs []sdk.Msg,
	timeoutTimestamp uint64,
	callbackId string,
	callbackArgs []byte,
) (uint64, error) {
	chainId, err := k.GetChainID(ctx, connectionId)
	if err != nil {
		return 0, err
	}
	owner := types.GetFeeICAAccountOwner(chainId)
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return 0, err
	}

	k.Logger(ctx).Info(utils.LogWithHostZone(chainId, "  Submitting ICA Tx on %s, %s with TTL: %d", portID, connectionId, timeoutTimestamp))
	for _, msg := range msgs {
		k.Logger(ctx).Info(utils.LogWithHostZone(chainId, "    Msg: %+v", msg))
	}

	channelID, found := k.IcaControllerKeeper.GetActiveChannelID(ctx, connectionId, portID)
	if !found {
		return 0, sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	data, err := icatypes.SerializeCosmosTx(k.cdc, msgs)
	if err != nil {
		return 0, err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	sequence, err := k.IcaControllerKeeper.SendTx(ctx, chanCap, connectionId, portID, packetData, timeoutTimestamp)
	if err != nil {
		return 0, err
	}

	return sequence, nil
}

// step 1: send money over to fee account on Osmosis
func (k Keeper) SendIBCFee(ctx sdk.Context) error {
	nnFeeAddress := k.accountKeeper.GetModuleAddress(types.NonNativeFeeCollectorName)
	tokens := k.bankKeeper.GetAllBalances(ctx, nnFeeAddress)

	hostFeeAddress := k.GetFeeICAAddress(ctx)

	for _, token := range tokens {
		timeout, err := k.GetTtl(ctx)
		if err != nil {
			return err
		}

		if err := k.transferKeeper.SendTransfer(
			ctx,
			ibctransfertypes.PortID,
			juno_osmo_channel_id,
			token,
			nnFeeAddress,
			hostFeeAddress,
			clienttypes.Height{},
			timeout,
		); err != nil {
			return err
		}
	}

	// save tokens sent to ica address
	k.SetTempFee(ctx, tokens)

	return nil
}

// step 2: execute ICA swap
func (k Keeper) ICASwap(ctx sdk.Context, coinOsmo sdk.Coin) error {

	// message creation
	hostFeeAddress := k.GetFeeICAAddress(ctx)
	baseDenom, err := k.GetBaseDenom(ctx)
	if err != nil {
		return err
	}

	poolId := k.GetPool(ctx, coinOsmo.GetDenom())

	msgs := []sdk.Msg{
		&gammtypes.MsgSwapExactAmountIn{
			Sender: hostFeeAddress,
			Routes: []gammtypes.SwapAmountInRoute{
				{
					PoolId:        poolId,
					TokenOutDenom: GetIBCDenom(osmo_juno_channel_id, baseDenom).IBCDenom(),
				},
			},
			TokenIn:           coinOsmo,
			TokenOutMinAmount: sdk.OneInt(),
		},
	}

	icaTimeoutNanos, err := k.GetTtl(ctx)
	if err != nil {
		return err
	}

	swapCallback := types.SwapCallback{
		TokenIn: coinOsmo,
	}
	b, err := swapCallback.Marshal()
	if err != nil {
		return err
	}

	_, err = k.SubmitTxs(ctx, JUNO_OSMO_CONNECTION_ID, msgs, icaTimeoutNanos, ICACallbackID_SWAP, b)
	if err != nil {
		return sdkerrors.Wrapf(err, "failed to submit txs")
	}

	return nil
}

// step 3: execute ICA IBC transfer from Osmosis back to native fee collector
func (k Keeper) ICATransferToFeeCollector(ctx sdk.Context) error {
	return nil
}
