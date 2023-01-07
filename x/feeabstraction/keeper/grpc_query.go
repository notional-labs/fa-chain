package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/fa-chain/x/feeabstraction/types"
)

var _ types.QueryServer = Keeper{}

// calculate fee rate of a non - native coin to ujuno
func (k Keeper) FeeRate(goCtx context.Context, req *types.QueryFeeRateRequest) (*types.QueryFeeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// calculate fee
	feeRate, err := k.GetFeeRate(ctx, req.Denom)
	if err != nil {
		return &types.QueryFeeRateResponse{}, err
	}

	return &types.QueryFeeRateResponse{FeeRate: feeRate.String()}, nil
}
