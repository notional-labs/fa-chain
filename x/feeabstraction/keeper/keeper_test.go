package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"

	"github.com/notional-labs/fa-chain/app"
	"github.com/notional-labs/fa-chain/app/apptesting"
	appparams "github.com/notional-labs/fa-chain/app/params"
	"github.com/notional-labs/fa-chain/x/interchainquery/keeper"
	"github.com/notional-labs/fa-chain/x/interchainquery/types"
	osmoparams "github.com/osmosis-labs/osmosis/v13/app/params"
)

const (
	HostChainId = "OSMO"
)

var (
	BaseDenomIBC = app.GetIBCDenom("channel-0", appparams.DefaultBondDenom).IBCDenom()
)

type KeeperTestSuite struct {
	apptesting.AppTestHelper
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.CreateTransferChannel(HostChainId)

	// check if Osmosis has pool
	pool, _ := s.HostApp.GAMMKeeper.GetPoolAndPoke(s.HostCtx, uint64(1))
	s.Require().NotNil(pool)
}

// ====== IBC ======

// sending ufa from osmosis to fachain
func (s *KeeperTestSuite) MockIBCTransferFromBtoA() error {
	timeoutHeight := clienttypes.NewHeight(0, 110)

	amount, _ := sdk.NewIntFromString("100000000") // 2^63 (one above int64)
	coinToSendToA := sdk.NewCoin(osmoparams.DefaultBondDenom, amount)

	// send from chainA to chainB
	msg := transfertypes.NewMsgTransfer(s.TransferPath.EndpointB.ChannelConfig.PortID, s.TransferPath.EndpointB.ChannelID, coinToSendToA, s.HostChain.SenderAccount.GetAddress().String(), s.Chain.SenderAccount.GetAddress().String(), timeoutHeight, 0)
	res, err := s.HostChain.SendMsgs(msg)
	if err != nil {
		return err
	}

	packet, err := ibctesting.ParsePacketFromEvents(res.GetEvents())
	if err != nil {
		return err
	}

	// relay send
	if err = s.TransferPath.RelayPacket(packet); err != nil {
		return err
	}

	return nil
}

func (s *KeeperTestSuite) GetMsgServer() types.MsgServer {
	return keeper.NewMsgServerImpl(s.App.InterchainqueryKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
