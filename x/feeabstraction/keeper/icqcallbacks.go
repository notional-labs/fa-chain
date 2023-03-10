package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/notional-labs/fa-chain/x/feeabstraction/types"
	icqtypes "github.com/notional-labs/fa-chain/x/interchainquery/types"
	gammtypes "github.com/osmosis-labs/osmosis/v13/x/gamm/types"
	twapquery "github.com/osmosis-labs/osmosis/v13/x/twap/client/queryproto"
)

const (
	ICQCallbackID_FeeRate        = "fee_rate"
	ICQCallbackID_Pool           = "pool"
	ICQCallbackID_ConfirmReceive = "ibc_transfer_receive"
)

// ICQCallbacks wrapper struct for stakeibc keeper
type ICQCallback func(Keeper, sdk.Context, []byte, icqtypes.Query) error

type ICQCallbacks struct {
	k         Keeper
	callbacks map[string]ICQCallback
}

var _ icqtypes.QueryCallbacks = ICQCallbacks{}

func (k Keeper) ICQCallbackHandler() ICQCallbacks {
	return ICQCallbacks{k, make(map[string]ICQCallback)}
}

func (c ICQCallbacks) CallICQCallback(ctx sdk.Context, id string, args []byte, query icqtypes.Query) error {
	return c.callbacks[id](c.k, ctx, args, query)
}

func (c ICQCallbacks) HasICQCallback(id string) bool {
	_, found := c.callbacks[id]
	return found
}

func (c ICQCallbacks) AddICQCallback(id string, fn interface{}) icqtypes.QueryCallbacks {
	c.callbacks[id] = fn.(ICQCallback)
	return c
}

func (c ICQCallbacks) RegisterICQCallbacks() icqtypes.QueryCallbacks {
	return c.
		AddICQCallback(ICQCallbackID_FeeRate, ICQCallback(FeeRateCallBack)).
		AddICQCallback(ICQCallbackID_Pool, ICQCallback(PoolCallBack)).
		AddICQCallback(ICQCallbackID_ConfirmReceive, ICQCallback(ConfirmReceiveCallback))
}

func FeeRateCallBack(k Keeper, ctx sdk.Context, args []byte, query icqtypes.Query) error {
	k.Logger(ctx).Info(fmt.Sprintf("FeeRateCallBack executing, QueryId: %vs, Host: %s, QueryType: %s, Connection: %s",
		query.Id, query.ChainId, query.QueryType, query.ConnectionId))

	twapRes := &twapquery.ArithmeticTwapToNowResponse{}
	if err := twapRes.Unmarshal(args); err != nil {
		return err
	}

	// save result to store
	twapReq := &twapquery.ArithmeticTwapToNowRequest{}
	if err := twapReq.Unmarshal(query.Request); err != nil {
		return err
	}

	denomJuno := k.GetOsmoDenomTrack(ctx, twapReq.BaseAsset)

	if err := k.SetFeeRate(ctx, denomJuno, twapRes.ArithmeticTwap); err != nil {
		return err
	}

	return nil
}

func PoolCallBack(k Keeper, ctx sdk.Context, args []byte, query icqtypes.Query) error {
	k.Logger(ctx).Info(fmt.Sprintf("PoolCallBack executing, QueryId: %vs, Host: %s, QueryType: %s, Connection: %s",
		query.Id, query.ChainId, query.QueryType, query.ConnectionId))

	poolRes := &gammtypes.QueryPoolsWithFilterResponse{}
	if err := poolRes.Unmarshal(args); err != nil {
		return err
	}

	// empty query
	if len(poolRes.Pools) == 0 {
		k.Logger(ctx).Info("Empty pool response")
		return nil
	}

	poolReq := &gammtypes.QueryPoolsWithFilterRequest{}
	if err := poolReq.Unmarshal(query.Request); err != nil {
		return err
	}

	var pool gammtypes.PoolI
	err := k.cdc.UnpackAny(poolRes.Pools[0], &pool)
	if err != nil {
		panic(err)
	}

	denom := poolReq.MinLiquidity.GetDenomByIndex(0)
	if !k.HasOsmoDenomTrack(ctx, denom) {
		denom = poolReq.MinLiquidity.GetDenomByIndex(1)
	}

	k.SetPool(ctx, denom, pool.GetId())

	return nil
}

func ConfirmReceiveCallback(k Keeper, ctx sdk.Context, args []byte, query icqtypes.Query) error {
	k.Logger(ctx).Info(fmt.Sprintf("ConfirmReceiveCallback executing, QueryId: %vs, Host: %s, QueryType: %s, Connection: %s",
		query.Id, query.ChainId, query.QueryType, query.ConnectionId))

	// Unmarshal the CB args into a coin type
	fee := sdk.Coin{}
	err := k.cdc.Unmarshal(args, &fee)
	if err != nil {
		errMsg := fmt.Sprintf("unable to unmarshal balance in callback args, err: %s", err.Error())
		k.Logger(ctx).Error(errMsg)
		return sdkerrors.Wrapf(types.ErrMarshalFailure, errMsg)
	}

	// Check if the coin is nil (which would indicate the fee has not come to ica address)
	// If already swapped, this should fail also
	if fee.IsNil() {
		k.Logger(ctx).Info(fmt.Sprintf("ConfirmReceiveCallback: ica address does not have this fee"))
		return nil
	}

	// execute ICASwap
	if err := k.ICASwap(ctx, fee); err != nil {
		errMsg := fmt.Sprintf("fail to ica swap, err :%s", err.Error())
		k.Logger(ctx).Error(errMsg)
		return sdkerrors.Wrapf(types.ErrMarshalFailure, errMsg)
	}

	junoDenom := k.GetOsmoDenomTrack(ctx, fee.GetDenom())
	fees, _ := k.GetTempFee(ctx)
	fees = fees.Sub(sdk.NewCoins(sdk.NewCoin(junoDenom, fee.Amount)))
	k.SetTempFee(ctx, fees)

	return nil
}
