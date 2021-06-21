package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/mises-id/mises-tm/x/misestm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AppInfoAll(c context.Context, req *types.QueryAllAppInfoRequest) (*types.QueryAllAppInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var AppInfos []*types.AppInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	AppInfoStore := prefix.NewStore(store, types.KeyPrefix(types.AppInfoKey))

	pageRes, err := query.Paginate(AppInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var AppInfo types.AppInfo
		if err := k.cdc.UnmarshalBinaryBare(value, &AppInfo); err != nil {
			return err
		}

		AppInfos = append(AppInfos, &AppInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAppInfoResponse{AppInfo: AppInfos, Pagination: pageRes}, nil
}

func (k Keeper) AppInfo(c context.Context, req *types.QueryGetAppInfoRequest) (*types.QueryGetAppInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var AppInfo types.AppInfo
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasAppInfo(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppInfoKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAppInfoIDBytes(req.Id)), &AppInfo)

	return &types.QueryGetAppInfoResponse{AppInfo: &AppInfo}, nil
}
