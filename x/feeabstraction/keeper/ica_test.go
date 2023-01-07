package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/gogo/protobuf/proto"
	"github.com/notional-labs/fa-chain/app"
	gammtypes "github.com/osmosis-labs/osmosis/v13/x/gamm/types"
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

func (s *KeeperTestSuite) TestICASwap() {
	s.SetupTest()
	feeAccountOwner := fmt.Sprintf("%s.%s", HostChainId, "FEE")
	s.CreateICAChannel(feeAccountOwner)
	s.MockIBCTransferFromBtoA()

	// assume that 10000uosmo is in ICA account
	acc := sdk.MustAccAddressFromBech32(s.IcaAddresses[feeAccountOwner])
	s.fundICAWallet(s.HostChain.GetContext(), sdk.NewCoins(sdk.NewInt64Coin(app.OsmoDefaultBondDenom, 10000)))
	old_balances := s.HostApp.BankKeeper.GetBalance(s.HostChain.GetContext(), acc, app.OsmoDefaultBondDenom)
	s.Require().Equal(uint64(10000), old_balances.Amount.Uint64())

	// make a swap msg
	msg := &gammtypes.MsgSwapExactAmountIn{
		Sender: s.IcaAddresses[feeAccountOwner],
		Routes: []gammtypes.SwapAmountInRoute{
			{
				PoolId:        uint64(1),
				TokenOutDenom: app.GetIBCDenom("channel-0", sdk.DefaultBondDenom).IBCDenom(),
			},
		},
		TokenIn:           sdk.NewInt64Coin(app.OsmoDefaultBondDenom, 10000),
		TokenOutMinAmount: sdk.OneInt(),
	}

	data, err := icatypes.SerializeCosmosTx(s.App.AppCodec(), []sdk.Msg{msg})
	s.Require().NoError(err)

	icaPacketData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	packetData := icaPacketData.GetBytes()

	// execute the msg
	packet := channeltypes.NewPacket(
		packetData,
		s.Chain.SenderAccount.GetSequence(),
		s.ICAPath.EndpointA.ChannelConfig.PortID,
		s.ICAPath.EndpointA.ChannelID,
		s.ICAPath.EndpointB.ChannelConfig.PortID,
		s.ICAPath.EndpointB.ChannelID,
		clienttypes.NewHeight(0, 100),
		0,
	)

	txResponse, err := s.HostApp.ICAHostKeeper.OnRecvPacket(s.HostChain.GetContext(), packet)
	txMsgData := &sdk.TxMsgData{}
	err = proto.Unmarshal(txResponse, txMsgData)
	s.Require().NoError(err)
	res := &gammtypes.MsgSwapExactAmountInResponse{}
	err = res.Unmarshal(txMsgData.Data[0].GetData())
	s.Require().NoError(err)

	// check if ica addr is updated accordingly
	balance := s.HostApp.BankKeeper.GetBalance(s.HostChain.GetContext(), acc, app.GetIBCDenom("channel-0", sdk.DefaultBondDenom).IBCDenom())
	s.Require().Equal(res.TokenOutAmount, balance.Amount)
}

func (s *KeeperTestSuite) fundICAWallet(ctx sdk.Context, amount sdk.Coins) {
	interchainAccountAddr, found := s.HostApp.ICAHostKeeper.GetInterchainAccountAddress(ctx, s.ICAPath.EndpointB.ConnectionID, s.ICAPath.EndpointA.ChannelConfig.PortID)
	s.Require().True(found)

	msgBankSend := &banktypes.MsgSend{
		FromAddress: s.HostChain.SenderAccount.GetAddress().String(),
		ToAddress:   interchainAccountAddr,
		Amount:      amount,
	}

	res, err := s.HostChain.SendMsgs(msgBankSend)
	s.Require().NotEmpty(res)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) TestICA_IBCTransfer() {
	s.SetupTest()
	feeAccountOwner := fmt.Sprintf("%s.%s", HostChainId, "FEE")
	s.CreateICAChannel(feeAccountOwner)
	s.MockIBCTransferFromBtoA()

	// assume that 10000uosmo is in ICA account
	acc := sdk.MustAccAddressFromBech32(s.IcaAddresses[feeAccountOwner])
	amt := sdk.NewInt64Coin(app.OsmoDefaultBondDenom, 10000)
	s.fundICAWallet(s.HostChain.GetContext(), sdk.NewCoins(amt))
	old_balances := s.HostApp.BankKeeper.GetBalance(s.HostChain.GetContext(), acc, app.OsmoDefaultBondDenom)
	s.Require().Equal(uint64(10000), old_balances.Amount.Uint64())

	// make a transfer message
	msgs, err := s.App.FAKeeper.MsgICATransferToFeeCollector(s.Ctx, amt, s.TransferPath.EndpointB.ChannelID)
	s.Require().NoError(err)

	data, err := icatypes.SerializeCosmosTx(s.App.AppCodec(), msgs)
	s.Require().NoError(err)

	icaPacketData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	packetData := icaPacketData.GetBytes()

	// execute the msg
	packet := channeltypes.NewPacket(
		packetData,
		s.Chain.SenderAccount.GetSequence(),
		s.ICAPath.EndpointA.ChannelConfig.PortID,
		s.ICAPath.EndpointA.ChannelID,
		s.ICAPath.EndpointB.ChannelConfig.PortID,
		s.ICAPath.EndpointB.ChannelID,
		clienttypes.NewHeight(0, 100),
		0,
	)

	txResponse, err := s.HostApp.ICAHostKeeper.OnRecvPacket(s.HostChain.GetContext(), packet)
	txMsgData := &sdk.TxMsgData{}
	err = proto.Unmarshal(txResponse, txMsgData)
	s.Require().NoError(err)
}
