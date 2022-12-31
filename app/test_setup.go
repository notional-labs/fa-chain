package app

import (
	"encoding/json"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	appparams "github.com/nghuyenthevinh2000/fa-chain/app/params"
	osmoapp "github.com/osmosis-labs/osmosis/v13/app"
	osmoparams "github.com/osmosis-labs/osmosis/v13/app/params"
	"github.com/osmosis-labs/osmosis/v13/x/gamm/pool-models/balancer"
	gammtypes "github.com/osmosis-labs/osmosis/v13/x/gamm/types"
)

func init() {
	appparams.SetAddressPrefixes()
}

// Initializes a new StrideApp without IBC functionality
func InitTestApp(initChain bool) *App {
	db := dbm.NewMemDB()
	codec := MakeEncodingConfig()
	app := New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		5,
		codec,
		simapp.EmptyAppOptions{},
	)
	if initChain {
		genesisState := NewDefaultGenesisState(codec.Marshaler)
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

// Initializes a new Stride App casted as a TestingApp for IBC support
func InitStrideIBCTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	app := InitTestApp(false)
	return ibctesting.TestingApp(app), NewDefaultGenesisState(app.appCodec)
}

// initializes a new Osmosis App casted as a TestingApp for IBC support
func InitOsmosisIBCTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	app := OsmoSetup(false)
	return ibctesting.TestingApp(app), OsmoGenesisStateWithPools(app)
}

// setup Osmo App with pools in genesis
// Setup initializes a new OsmosisApp.
func OsmoSetup(isCheckTx bool) *osmoapp.OsmosisApp {
	db := dbm.NewMemDB()
	app := osmoapp.NewOsmosisApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, simapp.EmptyAppOptions{}, osmoapp.GetWasmEnabledProposals(), osmoapp.EmptyWasmOpts)
	if !isCheckTx {
		// creating genesis with pools
		stateBytes, _ := json.MarshalIndent(OsmoGenesisStateWithPools(app), "", " ")

		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func OsmoGenesisStateWithPools(app *osmoapp.OsmosisApp) osmoapp.GenesisState {
	gen := osmoapp.NewDefaultGenesisState()

	// ufac ibc format
	ibcDenom := GetIBCDenom("channel-0", appparams.DefaultBondDenom).IBCDenom()

	balancerPool, _ := balancer.NewBalancerPool(1, balancer.PoolParams{
		SwapFee: sdk.NewDecWithPrec(1, 2),
		ExitFee: sdk.NewDecWithPrec(1, 2),
	}, []balancer.PoolAsset{
		{
			Weight: sdk.NewInt(1),
			Token:  sdk.NewInt64Coin(ibcDenom, 1000000000),
		},
		{
			Weight: sdk.NewInt(1),
			Token:  sdk.NewInt64Coin(osmoparams.DefaultBondDenom, 1000000000),
		},
	}, "", time.Now())

	any, _ := codectypes.NewAnyWithValue(&balancerPool)
	gammGen := gammtypes.GenesisState{
		Pools:          []*codectypes.Any{any},
		NextPoolNumber: 2,
		Params: gammtypes.Params{
			PoolCreationFee: sdk.Coins{sdk.NewInt64Coin(osmoparams.DefaultBondDenom, 1000_000_000)},
		},
	}
	gen[gammtypes.ModuleName] = app.AppCodec().MustMarshalJSON(&gammGen)

	return gen
}

func GetIBCDenom(channelId, baseDenom string) transfertypes.DenomTrace {
	sourcePrefix := transfertypes.GetDenomPrefix("transfer", channelId)
	prefixedDenom := sourcePrefix + baseDenom

	return transfertypes.ParseDenomTrace(prefixedDenom)
}
