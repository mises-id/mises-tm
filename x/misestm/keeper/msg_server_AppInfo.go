package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func (k msgServer) UpdateAppInfo(goCtx context.Context, msg *types.MsgUpdateAppInfo) (*types.MsgUpdateAppInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	appMgr := NewAppMgrImpl(k.Keeper)

	// query first
	misesAcc, err := appMgr.GetAppAccount(ctx, msg.Appid)
	if misesAcc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s not exists", msg.Appid)
	}
	if err != nil {
		return nil, err
	}

	// Checks that the element exists
	if !k.HasAppInfo(ctx, misesAcc.InfoID) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("app info key %d doesn't exist", misesAcc.InfoID))
	}
	oldAppInfo := k.GetAppInfo(ctx, misesAcc.InfoID)
	if msg.Creator != oldAppInfo.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	if msg.Version != oldAppInfo.Version+1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect version")
	}
	var AppInfo = oldAppInfo
	AppInfo.PubInfo = &types.PublicAppInfo{
		Name:      msg.Name,
		Domains:   msg.Domains,
		Developer: msg.Developer,
		HomeUrl:   msg.HomeUrl,
		IconUrl:   msg.IconUrl,
	}
	AppInfo.Version = msg.Version
	k.SetAppInfo(ctx, AppInfo)

	return &types.MsgUpdateAppInfoResponse{}, nil
}
