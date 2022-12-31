package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nghuyenthevinh2000/fa-chain/x/feeabstraction/types"
)

var _ types.QueryServer = Keeper{}

// calculate fee rate of a non - native coin to ujuno
func (k Keeper) FeeRate(goCtx context.Context, req *types.QueryFeeRateRequest) (*types.QueryFeeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get newest fee rate
	feeRate, err := k.GetFeeRate(ctx, req.Fee.Denom)
	if err != nil {
		return &types.QueryFeeRateResponse{}, err
	}

	// calculate fee
	amt := sdk.NewCoin(req.Fee.GetDenom(), sdk.NewDecCoinFromCoin(req.Fee).Amount.Quo(feeRate).RoundInt())

	return &types.QueryFeeRateResponse{FeeRate: amt}, nil
}
