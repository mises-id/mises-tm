package rest

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"
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
		fmt.Println(err, hashHexStr)
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

// HandleQueryDidRequest the QueryDidRequest http handler
func HandleQueryDidRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		misesIDStr := r.Form.Get("mises_id")

		queryClient := types.NewRestQueryClient(clientCtx)

		params := &types.RestQueryDidRequest{
			MisesId: misesIDStr,
		}

		resp, err := queryClient.QueryDid(context.Background(), params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		PostProcessResponseBare(w, clientCtx, resp)
	}
}

// HandleQueryUserRequest the QueryUserRequest http handler
func HandleQueryUserRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		misesIDStr := r.Form.Get("mises_id")

		queryClient := types.NewRestQueryClient(clientCtx)

		params := &types.RestQueryUserRequest{
			MisesUid: misesIDStr,
		}

		resp, err := queryClient.QueryUser(context.Background(), params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		PostProcessResponseBare(w, clientCtx, resp)
	}
}

// HandleQueryUserRelationRequest the QueryUserRelationRequest http handler
func HandleQueryUserRelationRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		misesIDStr := r.Form.Get("mises_id")
		filterStr := r.Form.Get("filter")
		keyStr := r.Form.Get("pagination.key")
		offsetStr := r.Form.Get("pagination.offset")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = 0
		}
		limitStr := r.Form.Get("pagination.limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = 10
		}
		countTotalStr := r.Form.Get("pagination.countTotal")

		queryClient := types.NewRestQueryClient(clientCtx)

		params := &types.RestQueryUserRelationRequest{
			MisesUid: misesIDStr,
			Filter:   filterStr,
			Pagination: &query.PageRequest{
				Key:        []byte(keyStr),
				Offset:     uint64(offset),
				Limit:      uint64(limit),
				CountTotal: countTotalStr == "true",
			},
		}

		resp, err := queryClient.QueryUserRelation(context.Background(), params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		PostProcessResponseBare(w, clientCtx, resp)
	}
}

// HandleQueryAppRequest the QueryAppRequest http handler
func HandleQueryAppRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		misesIDStr := r.Form.Get("mises_appid")

		queryClient := types.NewRestQueryClient(clientCtx)

		params := &types.RestQueryAppRequest{
			MisesAppid: misesIDStr,
		}

		resp, err := queryClient.QueryApp(context.Background(), params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		PostProcessResponseBare(w, clientCtx, resp)
	}
}

// HandleQueryAppFeeGrantRequest the QueryAppFeeGrantRequest http handler
func HandleQueryAppFeeGrantRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		misesAppIDStr := r.Form.Get("mises_appid")
		misesUIDStr := r.Form.Get("mises_uid")
		queryClient := types.NewRestQueryClient(clientCtx)

		params := &types.RestQueryAppFeeGrantRequest{
			MisesAppid: misesAppIDStr,
			MisesUid:   misesUIDStr,
		}

		resp, err := queryClient.QueryAppFeeGrant(context.Background(), params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		PostProcessResponseBare(w, clientCtx, resp)
	}
}
