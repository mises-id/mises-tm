// +build cgo,tests

package rest_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	bankcli "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
	dbm "github.com/tendermint/tm-db"
	_ "github.com/tendermint/tm-db/metadb"

	"github.com/mises-id/mises-tm/app"
	"github.com/mises-id/mises-tm/app/params"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	stdTxRes sdk.TxResponse
}

// TestAppOptions is a stub implementing AppOptions
type TestAppOptions struct {
	db dbm.DB
}

// Get implements AppOptions
func (ao TestAppOptions) Get(o string) interface{} {
	if o == "misesdb" {
		return ao.db
	}
	return nil
}

func NewAppConstructor(encodingCfg params.EncodingConfig) network.AppConstructor {
	return func(val network.Validator) servertypes.Application {
		db := dbm.NewMemDB()
		return app.New(
			val.Ctx.Logger, db, nil, nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
			encodingCfg,
			TestAppOptions{db},
			baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices("0token"),
		)
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	app.SetConfig()
	encCfg := app.MakeEncodingConfig()
	cfg := network.DefaultConfig()
	cfg.AppConstructor = NewAppConstructor(encCfg)
	cfg.NumValidators = 2
	//cfg.EnableLogging = true
	s.cfg = cfg

	var err error
	s.network = network.New(s.T(), cfg)
	s.Require().NoError(err)

	kb := s.network.Validators[0].ClientCtx.Keyring
	_, _, err = kb.NewMnemonic("newAccount", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	account1, _, err := kb.NewMnemonic("newAccount1", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	account2, _, err := kb.NewMnemonic("newAccount2", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	multi := kmultisig.NewLegacyAminoPubKey(2, []cryptotypes.PubKey{account1.GetPubKey(), account2.GetPubKey()})
	_, err = kb.SaveMultisig("multi", multi)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.setupSigner()

	//s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestCreateDid() {
	val := s.network.Validators[0]
	kb := val.ClientCtx.Keyring
	account, _, err := kb.NewMnemonic("newAccount3", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	req := &types.MsgCreateMisesID{
		MsgReqBase: types.MsgReqBase{
			MisesID: types.DIDPrefixForUser + account.GetAddress().String(),
		},
		PubKey: hex.EncodeToString(account.GetPubKey().Bytes()),
	}
	msg, err := json.Marshal(req)
	s.Require().NoError(err)

	s.Require().NoError(err)
	nonce := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	msgSign := string(msg) + "&" + nonce
	sigBytes, _, err := kb.Sign("newAccount3", []byte(msgSign))
	s.Require().NoError(err)
	sig := hex.EncodeToString(sigBytes)

	// we just test with async mode because this tx will fail - all we care about is that it got encoded and broadcast correctly
	txRes, err := s.broadcastRestMsgRequest(string(msg), sig, nonce)

	s.Require().NoError(err)

	// NOTE: this uses amino explicitly, don't migrate it!

	// we just check for a non-empty TxHash here, the actual hash will depend on the underlying tx configuration
	s.Require().NotEmpty(txRes.TxHash)
}

func (s *IntegrationTestSuite) broadcastRestMsgRequest(msg string, sig string, nonce string) (*sdk.TxResponse, error) {
	val := s.network.Validators[0]

	// NOTE: this uses amino explicitly, don't migrate it!

	v := url.Values{}
	v.Set("msg", msg)
	v.Set("nonce", nonce)
	v.Set("sig", sig)
	fmt.Println(v.Encode())
	rt, err := rest.PostRequest(
		fmt.Sprintf("%s/mises/did", val.APIAddress), "application/x-www-form-urlencoded", []byte(v.Encode()),
	)
	s.Require().NoError(err)
	fmt.Println(string(rt))
	var resp types.RestTxResponse
	s.Require().NoError(s.cfg.Codec.UnmarshalJSON(rt, &resp))

	return resp.TxResponse, err

}

func (s *IntegrationTestSuite) setupSigner() {
	val := s.network.Validators[0]
	keyring := val.ClientCtx.Keyring
	keyaddr, _, err := server.GenerateSaveCoinKey(
		keyring, "signer", false, hd.Secp256k1,
	)
	s.Require().NoError(err)

	sendTokens := sdk.NewInt64Coin(s.cfg.BondDenom, 10)

	// Might need to wait a block to refresh sequences from previous setups.
	s.Require().NoError(s.network.WaitForNextBlock())

	out, err := bankcli.MsgSendExec(
		val.ClientCtx,
		val.Address,
		keyaddr,
		sdk.NewCoins(sendTokens),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		fmt.Sprintf("--gas=%d", flags.DefaultGasLimit),
	)

	s.Require().NoError(err)
	var txRes sdk.TxResponse
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txRes))
	s.Require().Equal(uint32(0), txRes.Code)

	s.Require().NoError(s.network.WaitForNextBlock())

	s.testQueryTx(txRes.Height, txRes.TxHash, keyaddr.String())
}

// Helper function to test querying txs. We will use it to query StdTx and service `Msg`s.
func (s *IntegrationTestSuite) testQueryTx(txHeight int64, txHash, txRecipient string) {
	val0 := s.network.Validators[0]

	testCases := []struct {
		desc     string
		malleate func() *sdk.TxResponse
	}{
		{
			"Query by hash",
			func() *sdk.TxResponse {
				txJSON, err := rest.GetRequest(fmt.Sprintf("%s/txs/%s", val0.APIAddress, txHash))
				s.Require().NoError(err)

				var txResAmino sdk.TxResponse
				s.Require().NoError(val0.ClientCtx.LegacyAmino.UnmarshalJSON(txJSON, &txResAmino))
				return &txResAmino
			},
		},
		{
			"Query by height",
			func() *sdk.TxResponse {
				txJSON, err := rest.GetRequest(fmt.Sprintf("%s/txs?limit=10&page=1&tx.height=%d", val0.APIAddress, txHeight))
				s.Require().NoError(err)

				var searchtxResult sdk.SearchTxsResult
				s.Require().NoError(val0.ClientCtx.LegacyAmino.UnmarshalJSON(txJSON, &searchtxResult))
				s.Require().Len(searchtxResult.Txs, 1)
				return searchtxResult.Txs[0]
			},
		},
		{
			"Query by event (transfer.recipient)",
			func() *sdk.TxResponse {
				txJSON, err := rest.GetRequest(fmt.Sprintf("%s/txs?transfer.recipient=%s", val0.APIAddress, txRecipient))
				s.Require().NoError(err)

				var searchtxResult sdk.SearchTxsResult
				s.Require().NoError(val0.ClientCtx.LegacyAmino.UnmarshalJSON(txJSON, &searchtxResult))
				s.Require().Len(searchtxResult.Txs, 1)
				return searchtxResult.Txs[0]
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			txResponse := tc.malleate()

			// Check that the height is correct.
			s.Require().Equal(txHeight, txResponse.Height)

			// Check that the events are correct.
			s.Require().Contains(
				txResponse.RawLog,
				fmt.Sprintf("{\"key\":\"recipient\",\"value\":\"%s\"}", txRecipient),
			)

			// Check that the Msg is correct.
			stdTx, ok := txResponse.Tx.GetCachedValue().(legacytx.StdTx)
			s.Require().True(ok)
			msgs := stdTx.GetMsgs()
			s.Require().Equal(len(msgs), 1)
			msg, ok := msgs[0].(*banktypes.MsgSend)
			s.Require().True(ok)
			s.Require().Equal(txRecipient, msg.ToAddress)
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
