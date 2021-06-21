package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func (k msgServer) CreateUserInfo(goCtx context.Context, msg *types.MsgCreateUserInfo) (*types.MsgCreateUserInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var UserInfo = types.UserInfo{
		Creator: msg.Creator,
		Uid:     msg.Uid,
		EncData: msg.EncData,
		Iv:      msg.Iv,
		Version: msg.Version,
	}

	id := k.AppendUserInfo(
		ctx,
		UserInfo,
	)

	return &types.MsgCreateUserInfoResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateUserInfo(goCtx context.Context, msg *types.MsgUpdateUserInfo) (*types.MsgUpdateUserInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var UserInfo = types.UserInfo{
		Creator: msg.Creator,
		Id:      msg.Id,
		Uid:     msg.Uid,
		EncData: msg.EncData,
		Iv:      msg.Iv,
		Version: msg.Version,
	}

	// Checks that the element exists
	if !k.HasUserInfo(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetUserInfoOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetUserInfo(ctx, UserInfo)

	return &types.MsgUpdateUserInfoResponse{}, nil
}

func (k msgServer) DeleteUserInfo(goCtx context.Context, msg *types.MsgDeleteUserInfo) (*types.MsgDeleteUserInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasUserInfo(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetUserInfoOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveUserInfo(ctx, msg.Id)

	return &types.MsgDeleteUserInfoResponse{}, nil
}
