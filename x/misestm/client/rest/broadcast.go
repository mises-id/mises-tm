package rest

import (
	"io/ioutil"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func prepareFactory(clientCtx client.Context, txf tx.Factory) (tx.Factory, error) {
	gasSetting := flags.GasSetting{true, 0}
	txf = txf.
		WithTxConfig(clientCtx.TxConfig).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithKeybase(clientCtx.Keyring).
		WithChainID(clientCtx.ChainID).
		WithGas(gasSetting.Gas).
		WithSimulateAndExecute(gasSetting.Simulate).
		WithTimeoutHeight(0).
		WithGasAdjustment(0).
		WithMemo("memo").
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT)

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

	txf, err := prepareFactory(clientCtx, txf)
	if err != nil {
		return nil, err
	}

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

	// Construct the SignatureV2 struct
	sigData := signing.SingleSignatureData{
		SignMode:  txf.SignMode(),
		Signature: []byte{},
	}
	sig := signing.SignatureV2{
		PubKey:   nil,
		Data:     &sigData,
		Sequence: txf.Sequence(),
	}
	txb.SetSignatures(sig)

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
		msg := types.NewMsgCreateDidRegistry(req.Did, req.Did, "key0", "user", req.Pkey, 0)
		if err := msg.ValidateBasic(); err != nil {
			return
		}
		clientCtx := clientCtx.
			WithChainID("misestm").
			WithBroadcastMode(flags.BroadcastSync)
		txf := tx.Factory{}.
			WithSequence(1).
			WithAccountNumber(0)

		res, err := BroadcastTx(clientCtx, txf, msg)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponseBare(w, clientCtx, res)
	}
}
