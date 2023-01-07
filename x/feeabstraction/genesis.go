package feeabstraction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/fa-chain/x/feeabstraction/keeper"
	"github.com/notional-labs/fa-chain/x/feeabstraction/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	err := k.SetBaseDenom(ctx, genState.BaseDenom)
	if err != nil {
		panic(err)
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.BaseDenom, _ = k.GetBaseDenom(ctx)

	return genesis
}
