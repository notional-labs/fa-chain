package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	epochstypes "github.com/notional-labs/fa-chain/x/epochs/types"
	"github.com/notional-labs/fa-chain/x/feeabstraction/types"
)

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	if epochInfo.Identifier == epochstypes.TEST_EPOCH {
		// temporary condition for coin
		// prevent accidental base denom
		// TODO: move accidental base denom to native fee collector
		addr := k.accountKeeper.GetModuleAddress(types.NonNativeFeeCollectorName)
		coins := k.bankKeeper.GetAllBalances(ctx, addr)
		baseDenom, _ := k.GetBaseDenom(ctx)
		if !coins.IsZero() && coins.AmountOf(baseDenom).IsZero() {
			k.Logger(ctx).Info("Execute transfering all ibc tokens from nn fee collector")
			//execute cross - chain swap
			if err := k.SendIBCFee(ctx); err != nil {
				k.Logger(ctx).Error("failed to ibc transfer fee for nn fee collector", "error", err)
			}
		}

		// ICQ check to confirm that ica address on Osmo has received temp fee
		fees, err := k.GetTempFee(ctx)
		if err != nil {
			k.Logger(ctx).Error("failed to get temp fee", "error", err)
		}

		// if there is temp fee waiting to be converted, execute icq
		// currently, it is best - effort, maybe I should add check ?
		for _, coin := range fees {
			k.Logger(ctx).Info(fmt.Sprintf("Trying to confirm that ica address has received fund for fee = %v", coin))
			denomOsmo := k.GetJunoDenomTrack(ctx, coin.Denom)

			_, addr, _ := bech32.DecodeAndConvert(k.GetFeeICAAddress(ctx))
			data := banktypes.CreateAccountBalancesPrefix(addr)

			ttl, err := k.GetTtl(ctx)
			if err != nil {
				k.Logger(ctx).Error("failed to cast value", "error", err)
			}

			err = k.icqKeeper.MakeRequest(
				ctx,
				types.ModuleName,
				ICQCallbackID_ConfirmReceive,
				HOST_ZONE_CHAIN_ID,
				JUNO_OSMO_CONNECTION_ID,
				types.BANK_STORE_QUERY_WITH_PROOF,
				append(data, []byte(denomOsmo)...),
				ttl, // ttl
			)
		}
	}
}

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {}

// Hooks wrapper struct for incentives keeper
type Hooks struct {
	k Keeper
}

var _ epochstypes.EpochHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// epochs hooks
func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	h.k.BeforeEpochStart(ctx, epochInfo)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	h.k.AfterEpochEnd(ctx, epochInfo)
}
