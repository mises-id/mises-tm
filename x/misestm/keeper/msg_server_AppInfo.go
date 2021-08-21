package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func (k msgServer) CreateAppInfo(goCtx context.Context, msg *types.MsgCreateAppInfo) (*types.MsgCreateAppInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var AppInfo = types.AppInfo{
		Creator:   msg.Creator,
		Appid:     msg.Appid,
		Domain:    msg.Domain,
		Name:      msg.Name,
		Developer: msg.Developer,
		IconDid:   msg.IconDid,
		IconThumb: msg.IconThumb,
		Quota:     msg.Quota,
		Version:   msg.Version,
	}

	id := k.AppendAppInfo(
		ctx,
		AppInfo,
	)

	return &types.MsgCreateAppInfoResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAppInfo(goCtx context.Context, msg *types.MsgUpdateAppInfo) (*types.MsgUpdateAppInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var AppInfo = types.AppInfo{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Appid:     msg.Appid,
		Domain:    msg.Domain,
		Name:      msg.Name,
		Developer: msg.Developer,
		IconDid:   msg.IconDid,
		IconThumb: msg.IconThumb,
		Quota:     msg.Quota,
		Version:   msg.Version,
	}

	// Checks that the element exists
	if !k.HasAppInfo(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetAppInfoOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAppInfo(ctx, AppInfo)

	return &types.MsgUpdateAppInfoResponse{}, nil
}

func (k msgServer) DeleteAppInfo(goCtx context.Context, msg *types.MsgDeleteAppInfo) (*types.MsgDeleteAppInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasAppInfo(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetAppInfoOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAppInfo(ctx, msg.Id)

	return &types.MsgDeleteAppInfoResponse{}, nil
}
