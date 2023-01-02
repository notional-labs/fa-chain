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

	// calculate fee
	amt, err := k.ConvertToBaseToken(ctx, req.Fee)
	if err != nil {
		return &types.QueryFeeRateResponse{}, err
	}

	return &types.QueryFeeRateResponse{FeeRate: amt}, nil
}
