package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nghuyenthevinh2000/fa-chain/app/apptesting"
	"github.com/nghuyenthevinh2000/fa-chain/x/interchainquery/keeper"
	"github.com/nghuyenthevinh2000/fa-chain/x/interchainquery/types"
)

type KeeperTestSuite struct {
	apptesting.AppTestHelper
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
}

func (s *KeeperTestSuite) GetMsgServer() types.MsgServer {
	return keeper.NewMsgServerImpl(s.App.InterchainqueryKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
