package keeper

import (
	"fmt"

	icacallbacks "github.com/notional-labs/fa-chain/x/icacallbacks"
	icacallbackstypes "github.com/notional-labs/fa-chain/x/icacallbacks/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"

	gammtypes "github.com/osmosis-labs/osmosis/v13/x/gamm/types"
)

const (
	ICACallbackID_SWAP        = "swap"
	ICACallbackID_FEE_RECEIVE = "fee_receive"
)

// ICACallbacks wrapper struct for stakeibc keeper
type ICACallback func(Keeper, sdk.Context, channeltypes.Packet, *channeltypes.Acknowledgement, []byte) error

type ICACallbacks struct {
	k            Keeper
	icacallbacks map[string]ICACallback
}

var _ icacallbackstypes.ICACallbackHandler = ICACallbacks{}

func (k Keeper) ICACallbackHandler() ICACallbacks {
	return ICACallbacks{k, make(map[string]ICACallback)}
}

func (c ICACallbacks) CallICACallback(ctx sdk.Context, id string, packet channeltypes.Packet, ack *channeltypes.Acknowledgement, args []byte) error {
	return c.icacallbacks[id](c.k, ctx, packet, ack, args)
}

func (c ICACallbacks) HasICACallback(id string) bool {
	_, found := c.icacallbacks[id]
	return found
}

func (c ICACallbacks) AddICACallback(id string, fn interface{}) icacallbackstypes.ICACallbackHandler {
	c.icacallbacks[id] = fn.(ICACallback)
	return c
}

func (c ICACallbacks) RegisterICACallbacks() icacallbackstypes.ICACallbackHandler {
	a := c.
		AddICACallback(ICACallbackID_SWAP, ICACallback(SwapCallback)).
		AddICACallback(ICACallbackID_FEE_RECEIVE, ICACallback(FeeReceiveCallback))
	return a.(ICACallbacks)
}

// get result of token outs and execute ICATransferToFeeCollector
func SwapCallback(k Keeper, ctx sdk.Context, packet channeltypes.Packet, ack *channeltypes.Acknowledgement, args []byte) error {
	k.Logger(ctx).Info(fmt.Sprintf("executing SwapCallback on connection = %v", packet))

	// handle nil ack
	if ack == nil {
		k.Logger(ctx).Error(fmt.Sprintf("SwapCallback ack is nil, packet %v", packet))
		// after a nil ack, there should be a recovery mechanism for this by doing ICA Swap again
		return nil
	}

	txMsgData, err := icacallbacks.GetTxMsgData(ctx, *ack, k.Logger(ctx))
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("failed to unmarshal txMsgData, packet %v", packet))
		return sdkerrors.Wrap(icacallbackstypes.ErrTxMsgData, err.Error())
	}
	k.Logger(ctx).Info("SwapCallback executing", "packet", packet, "txMsgData", txMsgData, "args", args)
	// handle failed tx on host chain
	if len(txMsgData.Data) == 0 {
		k.Logger(ctx).Error(fmt.Sprintf("SwapCallback failed, packet %v", packet))
		// after an error, there should be a recovery mechanism for this by doing ICA Swap again
		return nil
	}

	// unmarshall to MsgSwapExactAmountInResponse
	res := &gammtypes.MsgSwapExactAmountInResponse{}
	b := txMsgData.GetData()[0].GetData()
	if err := res.Unmarshal(b); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Unmarshall MsgSwapExactAmountInResponse failed, packet %v", packet))
		// after an error, there should be a recovery mechanism for this by doing ICA Swap again
		return nil
	}

	// send ica transfer
	k.Logger(ctx).Info(fmt.Sprintf("TokenOutAmount = %v", res.TokenOutAmount.Uint64()))
	k.ICATransferToFeeCollector(ctx, sdk.NewInt64Coin(k.MustGetBaseIBCDenomOnOsmo(ctx).IBCDenom(), res.TokenOutAmount.Int64()))

	return nil
}

func FeeReceiveCallback(k Keeper, ctx sdk.Context, packet channeltypes.Packet, ack *channeltypes.Acknowledgement, args []byte) error {
	k.Logger(ctx).Info(fmt.Sprintf("executing FeeReceiveCallback on connection = %v", packet))
	return nil
}
