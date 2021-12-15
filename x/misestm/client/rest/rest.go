package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// REST query and parameter values
const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

type SeqInfo struct {
	nextNum uint64
	nextSeq uint64
}

var (
	SeqCmdChan   chan int
	SeqInfoChan  chan SeqInfo
	LocalKeyring keyring.Keyring
	LocalKey     keyring.Info
)

func getTestSignerKey(clientCtx client.Context) (keyring.Keyring, keyring.Info, error) {

	if LocalKeyring == nil || LocalKey == nil {
		var err error
		LocalKeyring, err = client.NewKeyringFromBackend(clientCtx, keyring.BackendTest)
		if err != nil {
			return nil, nil, err
		}
		keyList, err := LocalKeyring.List()
		if err != nil {
			return nil, nil, err
		}
		if len(keyList) == 0 {
			return nil, nil, fmt.Errorf("no local keyring")
		}
		LocalKey = keyList[0]
	}

	return LocalKeyring, LocalKey, nil
}

func StarSeqGenerator(clientCtx client.Context) error {

	seqChan := make(chan SeqInfo, 1)
	cmdChan := make(chan int, 1)
	ar := clientCtx.AccountRetriever
	var key keyring.Info
	var err error
	if clientCtx.Keyring != nil {
		key, err = clientCtx.Keyring.KeyByAddress(clientCtx.FromAddress)
	} else {
		_, key, err = getTestSignerKey(clientCtx)
	}
	if err != nil {
		return err
	}
	keyaddr := key.GetAddress()

	go func() {
		var num, seq uint64
		for {
			if seq == 0 {
				var err error
				num, seq, err = ar.GetAccountNumberSequence(clientCtx, keyaddr)
				if err != nil {
					time.Sleep(2 * time.Second)
					continue
				}
			}

			seqChan <- SeqInfo{nextNum: num, nextSeq: seq}

			next := <-cmdChan
			if next == 1 {
				seq = 0
			} else {
				seq++
			}

		}

	}()
	SeqInfoChan = seqChan
	SeqCmdChan = cmdChan
	return nil
}

// RegisterRoutes registers all transaction routes on the provided router.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router, postapi bool) {

	r := clientrest.WithHTTPDeprecationHeaders(rtr)

	r.HandleFunc("/mises/did", HandleQueryDidRequest(clientCtx)).Methods(MethodGet)
	r.HandleFunc("/mises/user", HandleQueryUserRequest(clientCtx)).Methods(MethodGet)
	r.HandleFunc("/mises/user/relation", HandleQueryUserRelationRequest(clientCtx)).Methods(MethodGet)
	r.HandleFunc("/mises/app", HandleQueryAppRequest(clientCtx)).Methods(MethodGet)
	r.HandleFunc("/mises/app/feegrant", HandleQueryAppFeeGrantRequest(clientCtx)).Methods(MethodGet)

	r.HandleFunc("/mises/tx", HandleQueryTxRequest(clientCtx)).Methods(MethodGet)

	if postapi {
		// post apis is used in LCD/ClientSDK only
		_ = StarSeqGenerator(clientCtx)
		r.HandleFunc("/mises/did", HandleCreateDidRequest(clientCtx)).Methods(MethodPost)
		r.HandleFunc("/mises/user", HandleUpdateUserInfoRequest(clientCtx)).Methods(MethodPost)
		r.HandleFunc("/mises/user/relation", HandleUpdateUserRelationRequest(clientCtx)).Methods(MethodPost)
		r.HandleFunc("/mises/app", HandleUpdateAppInfoRequest(clientCtx)).Methods(MethodPost)
		r.HandleFunc("/mises/app/feegrant", HandleUpdateAppFeeGrantRequest(clientCtx)).Methods(MethodPost)
	}

}

// PostProcessResponseBare performs post processing for a REST response.
func PostProcessResponseBare(w http.ResponseWriter, ctx client.Context, resp proto.Message) {
	var (
		result []byte
		err    error
	)

	marshaler := ctx.Codec

	result, err = marshaler.MarshalJSON(resp)
	if rest.CheckInternalServerError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(result)
}
