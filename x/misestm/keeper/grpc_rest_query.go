package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	if !k.HasDidRegistry(ctx, 0) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))
	k.cdc.Unmarshal(store.Get(GetDidRegistryIDBytes(0)), &DidRegistry)

	return &types.RestQueryDidResponse{DidRegistry: &DidRegistry}, nil
}
