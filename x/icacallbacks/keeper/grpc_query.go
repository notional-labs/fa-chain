package keeper

import (
	"github.com/notional-labs/fa-chain/x/icacallbacks/types"
)

var _ types.QueryServer = Keeper{}
