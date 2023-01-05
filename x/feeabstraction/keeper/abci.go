package keeper

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/notional-labs/fa-chain/x/feeabstraction/types"
	gammtypes "github.com/osmosis-labs/osmosis/v13/x/gamm/types"
	twapquery "github.com/osmosis-labs/osmosis/v13/x/twap/client/queryproto"
)

func (k Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	k.Logger(ctx).Info("Discovering ibc tokens from osmosis")

	// due to ibc denom hash, we can only accept denom directly from Osmosis
	// is there a way to get all assets on a channel
	k.transferKeeper.IterateDenomTraces(ctx, func(denomTrace ibctransfertypes.DenomTrace) bool {
		k.Logger(ctx).Info(fmt.Sprintf("Found token pair: (%s, %s) on channel %s",
			denomTrace.GetBaseDenom(), denomTrace.IBCDenom(), osmo_juno_channel_id))

		// if an ibc denom exists, skip
		if k.HasDenomTrack(ctx, denomTrace.GetBaseDenom()) {
			return true
		}

		// if found out that denom belongs to osmosis channel_id, register denom trace
		if strings.Contains(denomTrace.GetPath(), osmo_juno_channel_id) {
			k.Logger(ctx).Info("Registering token pair")
			k.SetDenomTrack(ctx, denomTrace.GetBaseDenom(), denomTrace.IBCDenom())
		}

		return true
	})

	feeAccount := types.GetFeeICAAccountOwner(HOST_ZONE_CHAIN_ID)
	_, connexist := k.IbcKeeper.ConnectionKeeper.GetConnection(ctx, JUNO_OSMO_CONNECTION_ID)
	_, exist := k.IcaControllerKeeper.GetInterchainAccountAddress(ctx, JUNO_OSMO_CONNECTION_ID, feeAccount)
	if connexist && !exist {
		if err := k.IcaControllerKeeper.RegisterInterchainAccount(ctx, JUNO_OSMO_CONNECTION_ID, feeAccount); err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("unable to register fee account, err: %s", err.Error()))
		}
	}

	// get pools information from Osmosis
	// make request for pools
	k.IterateDenomTrack(ctx, func(denomOsmo, _ string) bool {

		// check if it has pool
		if k.HasPool(ctx, denomOsmo) {
			return true
		}

		baseDenom, err := k.GetBaseDenom(ctx)
		if err != nil {
			k.Logger(ctx).Error(err.Error())
			return true
		}

		req := gammtypes.QueryPoolsWithFilterRequest{
			MinLiquidity: sdk.NewCoins(sdk.NewCoin(denomOsmo, sdk.OneInt()), sdk.NewCoin(GetIBCDenom(osmo_juno_channel_id, baseDenom).IBCDenom(), sdk.OneInt())),
			PoolType:     "Balancer",
		}

		k.Logger(ctx).Info(fmt.Sprintf("Preparing msg = %v", req))

		data, err := req.Marshal()
		if err != nil {
			k.Logger(ctx).Error("failed to marshall request", "error", err)
			return true
		}
		ttl, err := k.GetTtl(ctx)
		if err != nil {
			k.Logger(ctx).Error("failed to cast value", "error", err)
			return true
		}

		if err := k.icqKeeper.MakeRequest(ctx,
			types.ModuleName,
			ICQCallbackID_Pool,
			HOST_ZONE_CHAIN_ID,
			JUNO_OSMO_CONNECTION_ID,
			types.POOL_STORE_QUERY,
			data,
			ttl,
		); err != nil {
			k.Logger(ctx).Error("failed to make request", "error", err)
			return true
		}

		return true
	})
}

// EndBlocker of feeabstraction module
func (k Keeper) EndBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	k.Logger(ctx).Info("Fetching fee rate for ibc tokens from osmosis")

	// update fee
	k.IteratePool(ctx, func(denomOsmo string, poolId uint64) bool {
		k.Logger(ctx).Info(fmt.Sprintf("Found pool: (%s, %d)", denomOsmo, poolId))

		baseDenom, err := k.GetBaseDenom(ctx)
		if err != nil {
			k.Logger(ctx).Error(err.Error())
			return true
		}

		// make request for twap
		// TODO: better handling of start time
		req := twapquery.ArithmeticTwapToNowRequest{
			PoolId:     poolId,
			BaseAsset:  denomOsmo,
			QuoteAsset: GetIBCDenom(osmo_juno_channel_id, baseDenom).IBCDenom(),
			StartTime:  time.Now().Add(-time.Second * 10),
		}
		data, err := req.Marshal()
		if err != nil {
			k.Logger(ctx).Error("failed to marshall request", "error", err)
			return true
		}
		ttl, err := k.GetTtl(ctx)
		if err != nil {
			k.Logger(ctx).Error("failed to cast value", "error", err)
			return true
		}

		if err := k.icqKeeper.MakeRequest(ctx,
			types.ModuleName,
			ICQCallbackID_FeeRate,
			HOST_ZONE_CHAIN_ID,
			JUNO_OSMO_CONNECTION_ID,
			types.TWAP_STORE_QUERY,
			data,
			ttl,
		); err != nil {
			k.Logger(ctx).Error("failed to make request", "error", err)
			return true
		}

		return true
	})

	// temporary condition for coin
	addr := k.accountKeeper.GetModuleAddress(types.NonNativeFeeCollectorName)
	coins := k.bankKeeper.GetAllBalances(ctx, addr)
	if !coins.IsZero() {
		k.Logger(ctx).Info("Execute transfering all ibc tokens from nn fee collector")
		//execute cross - chain swap
		if err := k.SendIBCFee(ctx); err != nil {
			k.Logger(ctx).Error("failed to ibc transfer fee for nn fee collector", "error", err)
		}
	}
}
