package rest

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func prepareFactory(clientCtx client.Context, txf tx.Factory) (tx.Factory, error) {
	gasSetting := flags.GasSetting{true, 10000}
	txf = txf.
		WithTxConfig(clientCtx.TxConfig).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithKeybase(clientCtx.Keyring).
		WithChainID(clientCtx.ChainID).
		WithGas(gasSetting.Gas).
		WithSimulateAndExecute(gasSetting.Simulate).
		WithTimeoutHeight(0).
		WithGasAdjustment(2.0).
		WithMemo("memo").
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT)

	from := clientCtx.GetFromAddress()
	if from != nil {
		if err := txf.AccountRetriever().EnsureExists(clientCtx, from); err != nil {
			return txf, err
		}

		initNum, initSeq := txf.AccountNumber(), txf.Sequence()
		if initNum == 0 || initSeq == 0 {
			num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, from)
			if err != nil {
				return txf, err
			}

			if initNum == 0 {
				txf = txf.WithAccountNumber(num)
			}

			if initSeq == 0 {
				txf = txf.WithSequence(seq)
			}
		}
	}

	return txf, nil
}

// CalculateGas simulates the execution of a transaction and returns the
// simulation response obtained by the query and the adjusted gas amount.
func CalculateGas(
	clientCtx client.Context, txf tx.Factory, msgs ...sdk.Msg,
) (*sdk.Result, uint64, error) {
	txBytes, err := tx.BuildSimTx(txf, msgs...)
	if err != nil {
		return nil, 0, err
	}

	gasInfo, result, err := types.Simulater(txBytes)
	if err != nil {
		return nil, 0, err
	}

	return result, uint64(txf.GasAdjustment() * float64(gasInfo.GasUsed)), nil
}

func BroadcastTx(clientCtx client.Context, txf tx.Factory, msgs ...sdk.Msg) (*sdk.TxResponse, error) {

	if txf.SimulateAndExecute() || clientCtx.Simulate {
		_, adjusted, err := CalculateGas(clientCtx, txf, msgs...)
		if err != nil {
			return nil, err
		}

		txf = txf.WithGas(adjusted)
	}
	if clientCtx.Simulate {
		return nil, nil
	}

	txb, err := tx.BuildUnsignedTx(txf, msgs...)
	if err != nil {
		return nil, err
	}

	txb.SetFeeGranter(clientCtx.GetFeeGranterAddress())

	err = authclient.SignTx(txf, clientCtx, clientCtx.GetFromName(), txb, clientCtx.Offline, true)
	if err != nil {
		return nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txb.GetTx())
	if err != nil {
		return nil, err
	}

	res, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AddrFormDid(did string) string {
	return strings.Replace(did, "did:mises:", "", 1)
}

func prepareSigner(clientCtx client.Context) (client.Context, error) {
	if clientCtx.ChainID == "" {
		clientCtx = clientCtx.WithChainID("mises")
	}
	if clientCtx.Keyring == nil {
		keyring, err := client.NewKeyringFromBackend(clientCtx, keyring.BackendTest)
		if err != nil {
			return clientCtx, err
		}
		clientCtx = clientCtx.WithKeyring(keyring)
	}

	keyring := clientCtx.Keyring
	// keyring, err := client.NewKeyringFromBackend(clientCtx, keyring.BackendTest)
	// if err != nil {
	// 	return clientCtx, err
	// }
	keyname := "signer"
	var keyaddr sdk.AccAddress
	if key, err := keyring.Key(keyname); err != nil {
		return clientCtx, err
	} else {
		keyaddr = key.GetAddress()
	}

	clientCtx = clientCtx.WithKeyring(keyring).
		WithFromAddress(keyaddr).
		WithFromName(keyname)
	return clientCtx, nil
}
func RestCreateDidRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RestCreateDidRequest
		body, err := ioutil.ReadAll(r.Body)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		// NOTE: amino is used intentionally here, don't migrate it!
		err = clientCtx.Codec.UnmarshalJSON(body, &req)
		if err != nil {
			if rest.CheckBadRequestError(w, err) {
				return
			}
		}

		clientCtx := clientCtx.
			WithBroadcastMode(flags.BroadcastSync)

		clientCtx, err = prepareSigner(clientCtx)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		msg := types.NewMsgCreateDidRegistry(
			clientCtx.FromAddress.String(),
			req.Did,
			req.Did+"#key0",
			"EcdsaSecp256k1VerificationKey2019", // will shift to Ed25519VerificationKey2020
			req.Pkey,
			0,
		)
		if err := msg.ValidateBasic(); err != nil {
			if rest.CheckBadRequestError(w, err) {
				return
			}
		}
		txf := tx.Factory{}
		txf, err = prepareFactory(clientCtx, txf)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		res, err := BroadcastTx(clientCtx, txf, msg)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		resp := types.RestCreateDidResponse{TxResponse: res}

		rest.PostProcessResponseBare(w, clientCtx, resp)
	}
}
