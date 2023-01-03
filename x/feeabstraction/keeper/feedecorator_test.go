package keeper_test

import (
	client "github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	testdata "github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/notional-labs/fa-chain/app"
	"github.com/notional-labs/fa-chain/x/feeabstraction/keeper"
	"github.com/notional-labs/fa-chain/x/feeabstraction/types"
)

func (suite *KeeperTestSuite) TestFeeDecorator() {
	suite.SetupTest()

	// setting up fee-rate store (uosmo ibc denom, fee-rate)
	err := suite.App.FAKeeper.SetFeeRate(suite.Ctx, BaseOsmoDenomIBC, sdk.MustNewDecFromStr("0.5"))
	suite.Require().NoError(err)

	// build transaction
	encodingConfig := app.MakeEncodingConfig()
	clientCtx := client.Context{}.
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithJSONCodec(encodingConfig.Marshaler)

	txBuilder := clientCtx.TxConfig.NewTxBuilder()

	priv, _, addr := testdata.KeyTestPubAddr()
	acc := suite.App.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr)
	suite.App.AccountKeeper.SetAccount(suite.Ctx, acc)
	suite.FundAppAccount(suite.TestAccs[0], sdk.NewCoin(BaseOsmoDenomIBC, sdk.NewInt(1000000000)))
	suite.FundAppAccount(suite.TestAccs[0], sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000000000)))

	msgs := []sdk.Msg{
		&banktypes.MsgSend{
			FromAddress: suite.TestAccs[0].String(),
			ToAddress:   suite.TestAccs[1].String(),
			Amount:      sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000000))),
		},
	}
	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv}, []uint64{0}, []uint64{0}
	signerData := authsigning.SignerData{
		ChainID:       suite.Ctx.ChainID(),
		AccountNumber: accNums[0],
		Sequence:      accSeqs[0],
	}

	sigV2, _ := clienttx.SignWithPrivKey(
		1,
		signerData,
		txBuilder,
		privs[0],
		clientCtx.TxConfig,
		accSeqs[0],
	)

	// paying with not base denom fee
	txFee := sdk.NewCoins(sdk.NewInt64Coin(BaseOsmoDenomIBC, 15000))
	gasLimit := uint64(10000)

	tx := suite.BuildTx(txBuilder, msgs, sigV2, "", txFee, gasLimit)
	// build ante handler and set transaction through it
	mfd := keeper.NewMempoolFeeDecorator(suite.App.FAKeeper)
	dfd := keeper.NewDeductFeeDecorator(suite.App.FAKeeper, suite.App.AccountKeeper, suite.App.BankKeeper, suite.App.FeeGrantKeeper)
	antehandlerMFD := sdk.ChainAnteDecorators(mfd, dfd)
	_, err = antehandlerMFD(suite.Ctx, tx, false)
	suite.Require().NoError(err)

	// check if non native fee collector has collected fee
	moduleAddr := suite.App.AccountKeeper.GetModuleAddress(types.NonNativeFeeCollectorName)
	suite.Require().Equal(txFee[0], suite.App.BankKeeper.GetBalance(suite.Ctx, moduleAddr, txFee[0].Denom))
}
