package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	strideapp "github.com/notional-labs/fa-chain/app"
	"github.com/notional-labs/fa-chain/x/icacallbacks/keeper"
)

func IcacallbacksKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	app := strideapp.InitTestApp(true)
	icacallbackskeeper := app.IcacallbacksKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "stride-1", Time: time.Now().UTC()})

	return &icacallbackskeeper, ctx
}
