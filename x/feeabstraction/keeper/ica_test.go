package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
)

func (s *KeeperTestSuite) TestSendIBCFee() {
	s.SetupTest()
	feeAccountOwner := fmt.Sprintf("%s.%s", HostChainId, "FEE")
	s.CreateICAChannel(feeAccountOwner)
	s.MockIBCTransferFromBtoA()

	// assume that s.Chain.SenderAccount is nn fee decorator
	// send from chainA to chainB
	nnFeeAddress := s.Chain.SenderAccount.GetAddress()
	tokens := s.App.BankKeeper.GetAllBalances(s.Ctx, nnFeeAddress)

	hostFeeAddress := s.App.FAKeeper.GetFeeICAAddress(s.Ctx)
	acc, err := sdk.AccAddressFromBech32(hostFeeAddress)
	s.Require().NoError(err)

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
	coins := s.HostApp.BankKeeper.GetAllBalances(s.HostCtx, acc)
	s.Require().Equal(sdk.NewIntFromUint64(100000000), coins.AmountOf("uosmo"))
}
