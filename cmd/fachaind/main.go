package main

import (
	"os"

	// IT IS HERE FOR A REASON: DO NOT DELETE
	// THIS IS TO ENFORCE CORRECT ORDERING OF init()
	_ "github.com/notional-labs/fa-chain/app/params"
	_ "github.com/osmosis-labs/osmosis/v13/app/params"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/notional-labs/fa-chain/app"
	"github.com/notional-labs/fa-chain/cmd/fachaind/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
