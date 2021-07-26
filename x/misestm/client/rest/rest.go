package rest

import (
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// REST query and parameter values
const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

// RegisterRoutes registers all transaction routes on the provided router.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := clientrest.WithHTTPDeprecationHeaders(rtr)
	r.HandleFunc("/mises/did", HandleCreateDidRequest(clientCtx)).Methods(MethodPost)
	r.HandleFunc("/mises/user/{did}", HandleUpdateUserInfoRequest(clientCtx)).Methods(MethodPost)
	r.HandleFunc("/mises/user/{did}/relation", HandleUpdateUserRelationRequest(clientCtx)).Methods(MethodPost)

	r.HandleFunc("/mises/tx/{txhash}", HandleQueryTxRequest(clientCtx)).Methods(MethodGet)
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
