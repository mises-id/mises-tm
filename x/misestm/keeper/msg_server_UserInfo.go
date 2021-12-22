package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func (k msgServer) UpdateUserInfo(goCtx context.Context, msg *types.MsgUpdateUserInfo) (*types.MsgUpdateUserInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userMgr := NewUserMgrImpl(k.Keeper)

	uidAddr, uidOk := types.CheckDid(msg.Uid, types.DIDTypeUser)
	if !uidOk || uidAddr != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect uid")
	}

	// query first
	misesAcc, err := userMgr.GetUserAccount(ctx, msg.Uid)
	if misesAcc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s not exists", msg.Uid)
	}
	if err != nil {
		return nil, err
	}

	// Checks that the element exists
	if !k.HasUserInfo(ctx, misesAcc.InfoID) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("user info key %d doesn't exist", misesAcc.InfoID))
	}

	//check sign

	oldUserInfo := k.GetUserInfo(ctx, misesAcc.InfoID)
	if msg.Creator != oldUserInfo.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if msg.Version != oldUserInfo.Version+1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect version")
	}

	var UserInfo = types.UserInfo{
		Creator: msg.Creator,
		Id:      misesAcc.InfoID,
		Uid:     msg.Uid,
		PriInfo: msg.PriInfo,
		Version: msg.Version,
	}

	k.SetUserInfo(ctx, UserInfo)

	return &types.MsgUpdateUserInfoResponse{}, nil
}
