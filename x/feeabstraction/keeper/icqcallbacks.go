package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	icqtypes "github.com/nghuyenthevinh2000/fa-chain/x/interchainquery/types"
	gammtypes "github.com/osmosis-labs/osmosis/v13/x/gamm/types"
	twapquery "github.com/osmosis-labs/osmosis/v13/x/twap/client/queryproto"
)

const (
	ICQCallbackID_FeeRate = "fee_rate"
	ICQCallbackID_Pool    = "pool"
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
		AddICQCallback(ICQCallbackID_Pool, ICQCallback(PoolCallBack))
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

	denomJuno := k.GetDenomTrack(ctx, twapReq.BaseAsset)

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

	poolReq := &gammtypes.QueryPoolsWithFilterRequest{}
	if err := poolReq.Unmarshal(query.Request); err != nil {
		return err
	}

	var pool gammtypes.PoolI
	err := k.cdc.UnpackAny(poolRes.Pools[0], &pool)
	if err != nil {
		panic(err)
	}

	k.SetPool(ctx, poolReq.MinLiquidity.GetDenomByIndex(0), pool.GetId())

	return nil
}
