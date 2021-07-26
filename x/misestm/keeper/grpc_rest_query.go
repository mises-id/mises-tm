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

var _ types.RestQueryServer = Keeper{}

func (k Keeper) QueryDid(c context.Context, req *types.RestQueryDidRequest) (*types.RestQueryDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var DidRegistry types.DidRegistry
	ctx := sdk.UnwrapSDKContext(c)
	userMgr := NewUserMgrImpl(k)
	misesAcc, err := userMgr.GetUserAccount(ctx, req.Did)
	if err != nil {
		return nil, err
	}
	if misesAcc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "mises id %s not exists", req.Did)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))
	k.cdc.Unmarshal(store.Get(GetDidRegistryIDBytes(misesAcc.DidRegistryID)), &DidRegistry)

	return &types.RestQueryDidResponse{DidRegistry: &DidRegistry}, nil
}

func (k Keeper) QueryUser(c context.Context, req *types.RestQueryUserRequest) (*types.RestQueryUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var UserInfo types.UserInfo
	ctx := sdk.UnwrapSDKContext(c)
	userMgr := NewUserMgrImpl(k)
	misesAcc, err := userMgr.GetUserAccount(ctx, req.Did)
	if err != nil {
		return nil, err
	}

	if misesAcc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "mises id %s not exists", req.Did)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))
	k.cdc.Unmarshal(store.Get(GetUserInfoIDBytes(misesAcc.UserInfoID)), &UserInfo)

	return &types.RestQueryUserResponse{UserInfo: &UserInfo}, nil
}

// query user relations
func (k Keeper) QueryUserRelation(c context.Context, req *types.RestQueryUserRelationRequest) (*types.RestQueryUserRelationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	userMgr := NewUserMgrImpl(k)
	misesAcc, err := userMgr.GetUserAccount(ctx, req.Did)
	if err != nil {
		return nil, err
	}

	if misesAcc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "mises id %s not exists", req.Did)
	}
	UserRelations, err := userMgr.GetUserRelations(ctx, req.Did, string(req.Pagination.Key), int(req.Pagination.Limit))
	if err != nil {
		return nil, err
	}
	nextKey := ""
	if len(UserRelations) > 0 {
		nextKey = UserRelations[len(UserRelations)-1].UidTo
	}
	pageRes := &query.PageResponse{NextKey: []byte(nextKey)}

	resp := types.RestQueryUserRelationResponse{
		UserRelation: UserRelations,
		Pagination:   pageRes,
	}
	return &resp, nil
}
