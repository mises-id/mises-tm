package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func (k msgServer) UpdateUserRelation(goCtx context.Context, msg *types.MsgUpdateUserRelation) (*types.MsgUpdateUserRelationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userMgr := NewUserMgrImpl(k.Keeper)

	fromAddr, fromOk := types.CheckDid(msg.UidFrom, types.DIDTypeUser)
	if !fromOk || fromAddr != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect from")
	}
	_, toOk := types.CheckDid(msg.UidTo, types.DIDTypeUser)
	if !toOk {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect from")
	}

	oldRelation, err := userMgr.GetUserRelation(ctx, msg.UidFrom, msg.UidTo)
	if err != nil {
		return nil, err
	}
	var newRelation types.UserRelation
	if oldRelation == nil {

		newRelation = types.UserRelation{
			Creator:      msg.Creator,
			UidFrom:      msg.UidFrom,
			UidTo:        msg.UidTo,
			IsFollowing:  msg.IsFollowing,
			IsBlocking:   msg.IsBlocking,
			IsReferredBy: msg.IsReferredBy,
			Version:      0,
		}

		id := k.AppendUserRelation(
			ctx,
			newRelation,
		)
		newRelation.Id = id
	} else {
		if msg.Creator != oldRelation.Creator {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
		}

		newRelation = *oldRelation
		newRelation.IsFollowing = msg.IsFollowing
		newRelation.IsBlocking = msg.IsBlocking
		newRelation.Version++
		k.SetUserRelation(
			ctx,
			newRelation,
		)
	}
	return &types.MsgUpdateUserRelationResponse{}, nil
}
