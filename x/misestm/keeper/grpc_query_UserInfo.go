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

func (k Keeper) UserInfoAll(c context.Context, req *types.QueryAllUserInfoRequest) (*types.QueryAllUserInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var UserInfos []*types.UserInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	UserInfoStore := prefix.NewStore(store, types.KeyPrefix(types.UserInfoKey))

	pageRes, err := query.Paginate(UserInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var UserInfo types.UserInfo
		if err := k.cdc.UnmarshalBinaryBare(value, &UserInfo); err != nil {
			return err
		}

		UserInfos = append(UserInfos, &UserInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserInfoResponse{UserInfo: UserInfos, Pagination: pageRes}, nil
}

func (k Keeper) UserInfo(c context.Context, req *types.QueryGetUserInfoRequest) (*types.QueryGetUserInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var UserInfo types.UserInfo
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasUserInfo(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInfoKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetUserInfoIDBytes(req.Id)), &UserInfo)

	return &types.QueryGetUserInfoResponse{UserInfo: &UserInfo}, nil
}
