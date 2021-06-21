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

func (k Keeper) UserRelationAll(c context.Context, req *types.QueryAllUserRelationRequest) (*types.QueryAllUserRelationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var UserRelations []*types.UserRelation
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	UserRelationStore := prefix.NewStore(store, types.KeyPrefix(types.UserRelationKey))

	pageRes, err := query.Paginate(UserRelationStore, req.Pagination, func(key []byte, value []byte) error {
		var UserRelation types.UserRelation
		if err := k.cdc.UnmarshalBinaryBare(value, &UserRelation); err != nil {
			return err
		}

		UserRelations = append(UserRelations, &UserRelation)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserRelationResponse{UserRelation: UserRelations, Pagination: pageRes}, nil
}

func (k Keeper) UserRelation(c context.Context, req *types.QueryGetUserRelationRequest) (*types.QueryGetUserRelationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var UserRelation types.UserRelation
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasUserRelation(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserRelationKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetUserRelationIDBytes(req.Id)), &UserRelation)

	return &types.QueryGetUserRelationResponse{UserRelation: &UserRelation}, nil
}
