package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func (k msgServer) CreateDidRegistry(goCtx context.Context, msg *types.MsgCreateDidRegistry) (*types.MsgCreateDidRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var DidRegistry = types.DidRegistry{
		Creator:       msg.Creator,
		Did:           msg.Did,
		PkeyDid:       msg.PkeyDid,
		PkeyType:      msg.PkeyType,
		PkeyMultibase: msg.PkeyMultibase,
		Version:       msg.Version,
	}

	id := k.AppendDidRegistry(
		ctx,
		DidRegistry,
	)

	return &types.MsgCreateDidRegistryResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateDidRegistry(goCtx context.Context, msg *types.MsgUpdateDidRegistry) (*types.MsgUpdateDidRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var DidRegistry = types.DidRegistry{
		Creator:       msg.Creator,
		Id:            msg.Id,
		Did:           msg.Did,
		PkeyDid:       msg.PkeyDid,
		PkeyType:      msg.PkeyType,
		PkeyMultibase: msg.PkeyMultibase,
		Version:       msg.Version,
	}

	// Checks that the element exists
	if !k.HasDidRegistry(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetDidRegistryOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetDidRegistry(ctx, DidRegistry)

	return &types.MsgUpdateDidRegistryResponse{}, nil
}

func (k msgServer) DeleteDidRegistry(goCtx context.Context, msg *types.MsgDeleteDidRegistry) (*types.MsgDeleteDidRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasDidRegistry(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetDidRegistryOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDidRegistry(ctx, msg.Id)

	return &types.MsgDeleteDidRegistryResponse{}, nil
}
