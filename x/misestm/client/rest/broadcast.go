package rest

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	tmbtcec "github.com/tendermint/btcd/btcec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func prepareFactory(clientCtx client.Context, txf tx.Factory) tx.Factory {
	gasSetting := flags.GasSetting{
		Simulate: false,
		Gas:      100000,
	}
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

	seq := <-SeqInfoChan
	txf = txf.WithAccountNumber(seq.nextNum)
	txf = txf.WithSequence(seq.nextSeq)

	return txf
}

// CalculateGas simulates the execution of a transaction and returns the
// simulation response obtained by the query and the adjusted gas amount.
func calculateGas(
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

func broadcastTx(clientCtx client.Context, txf tx.Factory, msgs ...sdk.Msg) (*sdk.TxResponse, error) {

	if txf.SimulateAndExecute() || clientCtx.Simulate {
		_, adjusted, err := calculateGas(clientCtx, txf, msgs...)
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

	err = authclient.SignTx(txf, clientCtx, clientCtx.GetFromName(), txb, true, true)
	if err != nil {
		return nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txb.GetTx())
	if err != nil {
		return nil, err
	}

	//types.Logger.Error(fmt.Sprintf("BroadcastTx start with seq %v", txf.Sequence()))

	res, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return nil, err
	}

	//types.Logger.Error(fmt.Sprintf("BroadcastTx finish with code %v", res.Code))

	return res, nil
}

func prepareSigner(clientCtx client.Context) (client.Context, error) {
	if clientCtx.ChainID == "" {
		clientCtx = clientCtx.WithChainID("mises")
	}
	keyring, key, err := getLocalSignerKey(clientCtx)
	if err != nil {
		panic(err)
	}
	keyname := key.GetName()
	keyaddr := key.GetAddress()

	clientCtx = clientCtx.WithKeyring(keyring).
		WithFromAddress(keyaddr).
		WithFromName(keyname).
		WithBroadcastMode(flags.BroadcastSync)
	return clientCtx, nil
}

// HandleCreateDidRequest the CreateDidRequest http handler
func HandleCreateDidRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RestCreateDidRequest
		reqMsg, err := ParseReqeustBody(clientCtx, r, true)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		// NOTE: amino is used intentionally here, don't migrate it!
		err = clientCtx.Codec.UnmarshalJSON(reqMsg, &req)
		if err != nil {
			if rest.CheckBadRequestError(w, err) {
				return
			}
		}

		clientCtx, err = prepareSigner(clientCtx)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		pubKeyBytes, err := hex.DecodeString(req.PubKey)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		msg := types.NewMsgCreateDidRegistry(
			clientCtx.FromAddress.String(),
			req.MisesId,
			req.MisesId+"#key0",
			"EcdsaSecp256k1VerificationKey2019", // will shift to Ed25519VerificationKey2020
			base58.Encode(pubKeyBytes),
			0,
		)
		if err := msg.ValidateBasic(); err != nil {
			if rest.CheckBadRequestError(w, err) {
				return
			}
		}
		txf := tx.Factory{}
		txf = prepareFactory(clientCtx, txf)

		res, err := broadcastTx(clientCtx, txf, msg)

		postBroadcastTx(clientCtx, res, err, w)

	}
}

func postBroadcastTx(clientCtx client.Context, res *sdk.TxResponse, err error, w http.ResponseWriter) {
	if rest.CheckInternalServerError(w, err) {
		SeqCmdChan <- -1
		return
	}
	if res.Code == sdkerrors.ErrWrongSequence.ABCICode() {
		//reset cmdSeqChan
		SeqCmdChan <- 1
	} else {
		SeqCmdChan <- 0
	}

	resp := &types.RestTxResponse{TxResponse: res}
	if res.Code != 0 {
		resp.Code = res.Code
	}

	PostProcessResponseBare(w, clientCtx, resp)
}

func convertDERtoBER(signatureDER []byte) ([]byte, error) {
	sigDER, err := btcec.ParseDERSignature(signatureDER, btcec.S256())
	if err != nil {
		return nil, err
	}
	sigBER := tmbtcec.Signature{R: sigDER.R, S: sigDER.S}
	return sigBER.Serialize(), nil
}

func ParseReqeustBody(clientCtx client.Context, r *http.Request, isCreateDid bool) ([]byte, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	msg := r.Form.Get("msg")
	sig := r.Form.Get("sig")
	nonce := r.Form.Get("nonce")
	msgToSign := msg + "&" + nonce
	var pubKey cryptotypes.PubKey
	if !isCreateDid {
		var msgReq types.MsgReqBase
		err := json.Unmarshal([]byte(msg), &msgReq)
		if err != nil {
			return nil, err
		}
		addr, err := types.AddrFormDid(msgReq.MisesID)
		if err != nil {
			return nil, err
		}
		if err := clientCtx.AccountRetriever.EnsureExists(clientCtx, addr); err != nil {
			return nil, err
		}
		account, err := clientCtx.AccountRetriever.GetAccount(clientCtx, addr)
		if err != nil {
			return nil, err
		}
		pubKey = account.GetPubKey()
	} else {
		var msgReq types.MsgCreateMisesID
		err := json.Unmarshal([]byte(msg), &msgReq)
		if err != nil {
			return nil, err
		}
		addr, err := types.AddrFormDid(msgReq.MisesID)
		if err != nil {
			return nil, err
		}
		if err := clientCtx.AccountRetriever.EnsureExists(clientCtx, addr); err == nil {
			return nil, errors.New("user already exsit")
		}
		key, err := hex.DecodeString(msgReq.PubKey)
		if err != nil {
			return nil, err
		}
		pubKey = &secp256k1.PubKey{Key: key}
		if err != nil {
			return nil, err
		}
	}
	sigBytes, err := hex.DecodeString(sig)
	if err != nil {
		return nil, err
	}
	sigBytes, err = convertDERtoBER(sigBytes)
	if err != nil {
		return nil, err
	}
	if !pubKey.VerifySignature([]byte(msgToSign), sigBytes) {
		return nil, errors.New("wrong signature")
	}

	return []byte(msg), nil

}

// HandleUpdateUserInfoRequest the UpdateUserInfoRequest http handler
func HandleUpdateUserInfoRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RestUpdateUserInfoRequest
		reqMsg, err := ParseReqeustBody(clientCtx, r, false)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		err = clientCtx.Codec.UnmarshalJSON(reqMsg, &req)
		if err != nil {
			if rest.CheckBadRequestError(w, err) {
				return
			}
		}

		clientCtx, err = prepareSigner(clientCtx)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		msg := types.NewMsgUpdateUserInfo(
			clientCtx.FromAddress.String(),
			types.InvalidID,
			req.MisesId,
			req.PriInfo.EncData,
			req.PriInfo.Iv,
			0,
		)
		if err := msg.ValidateBasic(); err != nil {
			if rest.CheckBadRequestError(w, err) {
				return
			}
		}
		txf := tx.Factory{}
		txf = prepareFactory(clientCtx, txf)

		res, err := broadcastTx(clientCtx, txf, msg)
		postBroadcastTx(clientCtx, res, err, w)
	}
}

// HandleUpdateUserRelationRequest the UpdateUserRelationRequest http handler
func HandleUpdateUserRelationRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.RestUpdateUserRelationRequest
		reqMsg, err := ParseReqeustBody(clientCtx, r, false)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		// NOTE: amino is used intentionally here, don't migrate it!
		err = clientCtx.Codec.UnmarshalJSON(reqMsg, &req)
		if err != nil {
			if rest.CheckBadRequestError(w, err) {
				return
			}
		}

		clientCtx, err = prepareSigner(clientCtx)
		if rest.CheckInternalServerError(w, err) {
			return
		}
		var action uint64
		switch req.Action {
		case "follow":
			action = 0
		case "unfollow":
			action = 1
		case "block":
			action = 2
		case "unblock":
			action = 3
		}

		msg := types.NewMsgCreateUserRelation(
			clientCtx.FromAddress.String(),
			req.MisesId,
			req.TargetId,
			action,
			0,
		)
		if err := msg.ValidateBasic(); err != nil {
			if rest.CheckBadRequestError(w, err) {
				return
			}
		}
		txf := tx.Factory{}
		txf = prepareFactory(clientCtx, txf)

		res, err := broadcastTx(clientCtx, txf, msg)
		postBroadcastTx(clientCtx, res, err, w)
	}
}
