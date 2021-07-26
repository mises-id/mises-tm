package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func (k msgServer) CreateUserRelation(goCtx context.Context, msg *types.MsgCreateUserRelation) (*types.MsgCreateUserRelationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userMgr := NewUserMgrImpl(k.Keeper)

	oldRelation, err := userMgr.GetUserRelation(ctx, msg.UidFrom, msg.UidTo)
	if err != nil {
		return nil, err
	}
	action := msg.RelType
	var relType uint64
	if oldRelation != nil {
		relType = oldRelation.RelType
	} else {
		relType = 0
	}
	switch action {
	case 0:
		relType |= types.RelTypeBitFollow
	case 1:
		relType &= ^types.RelTypeBitFollow
	case 2:
		relType |= types.RelTypeBitBlock
	case 3:
		relType |= ^types.RelTypeBitBlock

	}
	var newRelation types.UserRelation
	if oldRelation == nil {

		newRelation = types.UserRelation{
			Creator: msg.Creator,
			UidFrom: msg.UidFrom,
			UidTo:   msg.UidTo,
			RelType: relType,
			Version: 0,
		}

		id := k.AppendUserRelation(
			ctx,
			newRelation,
		)
		newRelation.Id = id
	} else {
		newRelation = *oldRelation
		newRelation.RelType = relType
		newRelation.Version++
		k.SetUserRelation(
			ctx,
			newRelation,
		)
	}
	return &types.MsgCreateUserRelationResponse{
		Id: newRelation.Id,
	}, nil
}

func (k msgServer) UpdateUserRelation(goCtx context.Context, msg *types.MsgUpdateUserRelation) (*types.MsgUpdateUserRelationResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// var UserRelation = types.UserRelation{
	// 	Creator: msg.Creator,
	// 	Id:      msg.Id,
	// 	UidFrom: msg.UidFrom,
	// 	UidTo:   msg.UidTo,
	// 	RelType: msg.RelType,
	// 	Version: msg.Version,
	// }

	// // Checks that the element exists
	// if !k.HasUserRelation(ctx, msg.Id) {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	// }

	// // Checks if the the msg sender is the same as the current owner
	// if msg.Creator != k.GetUserRelationOwner(ctx, msg.Id) {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	// k.SetUserRelation(ctx, UserRelation)

	return &types.MsgUpdateUserRelationResponse{}, nil
}

func (k msgServer) DeleteUserRelation(goCtx context.Context, msg *types.MsgDeleteUserRelation) (*types.MsgDeleteUserRelationResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// if !k.HasUserRelation(ctx, msg.Id) {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	// }
	// if msg.Creator != k.GetUserRelationOwner(ctx, msg.Id) {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	// k.RemoveUserRelation(ctx, msg.Id)

	return &types.MsgDeleteUserRelationResponse{}, nil
}
