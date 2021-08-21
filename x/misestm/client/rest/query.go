package rest

import (
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

// HandleQueryTxRequest the QueryTxRequest http handler
func HandleQueryTxRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		hashHexStr := r.Form.Get("txhash")

		output, err := authtx.QueryTx(clientCtx, hashHexStr)
		if err != nil {
			if strings.Contains(err.Error(), hashHexStr) {
				rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
				return
			}
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resp := &types.RestTxResponse{
			TxResponse: output,
		}

		PostProcessResponseBare(w, clientCtx, resp)
	}
}
