package rest

import (
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/gorilla/mux"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

// HandleQueryTxRequest the QueryTxRequest http handler
func HandleQueryTxRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hashHexStr := vars["txhash"]

		output, err := authtx.QueryTx(clientCtx, hashHexStr)
		if err != nil {
			if strings.Contains(err.Error(), hashHexStr) {
				rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
				return
			}
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resp := &types.RestQueryTxResponse{
			TxResponse: output,
		}

		PostProcessResponseBare(w, clientCtx, resp)
	}
}
