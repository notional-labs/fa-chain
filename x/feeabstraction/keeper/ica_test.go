package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/notional-labs/fa-chain/x/feeabstraction/keeper"
)

func (s *KeeperTestSuite) TestSendIBCFee() {
	s.SetupTest()
	feeAccountOwner := fmt.Sprintf("%s.%s", keeper.HOST_ZONE_CHAIN_ID, "FEE")
	s.CreateICAChannel(feeAccountOwner)

	// fund nn fee decorator (assume that TestAccs[2] is nn fee decorator)
	ibcDenom := s.GetIBCDenomTrace("uosmo")
	s.FundAppAccount(s.TestAccs[2], sdk.NewInt64Coin(ibcDenom.IBCDenom(), 10000))

	// send from chainA to chainB
	nnFeeAddress := s.TestAccs[2]
	tokens := s.App.BankKeeper.GetAllBalances(s.Ctx, nnFeeAddress)

	hostFeeAddress := s.App.FAKeeper.GetFeeICAAddress(s.Ctx)
	fmt.Printf("hostFeeAddress = %v \n", hostFeeAddress)

	for _, token := range tokens {
		timeout, err := s.App.FAKeeper.GetTtl(s.Ctx)
		s.Require().NoError(err)

		msg := transfertypes.NewMsgTransfer(s.TransferPath.EndpointA.ChannelConfig.PortID, s.TransferPath.EndpointA.ChannelID, token, nnFeeAddress.String(), hostFeeAddress, clienttypes.Height{}, timeout)
		res, err := s.Chain.SendMsgs(msg)
		s.Require().NoError(err)

		packet, err := ibctesting.ParsePacketFromEvents(res.GetEvents())
		s.Require().NoError(err)

		// relay send
		err = s.TransferPath.RelayPacket(packet)
		s.Require().NoError(err)
	}

	// check if ibc denom is on hostFeeAddress
}
