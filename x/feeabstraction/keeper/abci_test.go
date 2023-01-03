package keeper_test

import (
	"github.com/notional-labs/fa-chain/app"
	osmoparams "github.com/osmosis-labs/osmosis/v13/app/params"
)

// go test -v -run ^TestKeeperTestSuite/TestIdentifyChain$ github.com/notional-labs/fa-chain/x/feeabstraction/keeper
func (s KeeperTestSuite) TestIdentifyChain() { //nolint:govet // it's fine to copy locks in a test
	s.SetupTest()

	// Send token
	err := s.MockIBCTransferFromBtoA()
	s.Suite.Require().NoError(err)
	// check if account on A has coin
	amt := s.App.BankKeeper.GetBalance(s.Ctx, s.Chain.SenderAccount.GetAddress(), BaseDenomIBC)
	s.Suite.Require().NotNil(amt)

	// run begin blocker
	s.App.FAKeeper.BeginBlocker(s.Ctx)

	// identify correct osmo denom in fee store
	res := s.App.FAKeeper.GetDenomTrack(s.Ctx, osmoparams.DefaultBondDenom)
	uosmoIbc := app.GetIBCDenom("channel-0", osmoparams.DefaultBondDenom).IBCDenom()
	s.Suite.Require().Equal(uosmoIbc, res)

	// check a query was created (a simple test; details about queries are covered in makeRequest's test)
	queries := s.App.InterchainqueryKeeper.AllQueries(s.Ctx)
	s.Require().Len(queries, 1, "one query should have been created")
}
