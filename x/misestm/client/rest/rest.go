package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rest"
)

// REST query and parameter values
const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

// RegisterRoutes registers all transaction routes on the provided router.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	r.HandleFunc("/mises/did", RestCreateDidRequest(clientCtx)).Methods(MethodPost)
}
