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

func (k Keeper) DidRegistryAll(c context.Context, req *types.QueryAllDidRegistryRequest) (*types.QueryAllDidRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var DidRegistrys []*types.DidRegistry
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	DidRegistryStore := prefix.NewStore(store, types.KeyPrefix(types.DidRegistryKey))

	pageRes, err := query.Paginate(DidRegistryStore, req.Pagination, func(key []byte, value []byte) error {
		var DidRegistry types.DidRegistry
		if err := k.cdc.UnmarshalBinaryBare(value, &DidRegistry); err != nil {
			return err
		}

		DidRegistrys = append(DidRegistrys, &DidRegistry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDidRegistryResponse{DidRegistry: DidRegistrys, Pagination: pageRes}, nil
}

func (k Keeper) DidRegistry(c context.Context, req *types.QueryGetDidRegistryRequest) (*types.QueryGetDidRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var DidRegistry types.DidRegistry
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasDidRegistry(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetDidRegistryIDBytes(req.Id)), &DidRegistry)

	return &types.QueryGetDidRegistryResponse{DidRegistry: &DidRegistry}, nil
}
