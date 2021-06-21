package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func (k msgServer) CreateUserRelation(goCtx context.Context, msg *types.MsgCreateUserRelation) (*types.MsgCreateUserRelationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var UserRelation = types.UserRelation{
		Creator: msg.Creator,
		UidFrom: msg.UidFrom,
		UidTo:   msg.UidTo,
		RelType: msg.RelType,
		Version: msg.Version,
	}

	id := k.AppendUserRelation(
		ctx,
		UserRelation,
	)

	return &types.MsgCreateUserRelationResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateUserRelation(goCtx context.Context, msg *types.MsgUpdateUserRelation) (*types.MsgUpdateUserRelationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var UserRelation = types.UserRelation{
		Creator: msg.Creator,
		Id:      msg.Id,
		UidFrom: msg.UidFrom,
		UidTo:   msg.UidTo,
		RelType: msg.RelType,
		Version: msg.Version,
	}

	// Checks that the element exists
	if !k.HasUserRelation(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetUserRelationOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetUserRelation(ctx, UserRelation)

	return &types.MsgUpdateUserRelationResponse{}, nil
}

func (k msgServer) DeleteUserRelation(goCtx context.Context, msg *types.MsgDeleteUserRelation) (*types.MsgDeleteUserRelationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasUserRelation(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetUserRelationOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveUserRelation(ctx, msg.Id)

	return &types.MsgDeleteUserRelationResponse{}, nil
}
