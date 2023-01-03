package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/feeabstraction module sentinel errors
var (
	ErrInvalidFeeToken = sdkerrors.Register(ModuleName, 1, "invalid fee token")
	ErrTooManyFeeCoins = sdkerrors.Register(ModuleName, 2, "too many fee coins. only accepts fees in one denom")
	ErrNoBaseDenom     = sdkerrors.Register(ModuleName, 3, "no base denom was set")
)
