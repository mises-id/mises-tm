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

var _ types.RestQueryServer = Keeper{}

func (k Keeper) QueryDid(c context.Context, req *types.RestQueryDidRequest) (*types.RestQueryDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var DidRegistry types.DidRegistry
	ctx := sdk.UnwrapSDKContext(c)

	_, didtype, err := types.AddrFromDid(req.MisesId)
	if err != nil {
		return nil, err
	}
	var misesAcc *types.MisesAccount = nil
	if didtype == types.DIDTypeApp {
		appMgr := NewAppMgrImpl(k)
		misesAcc, err = appMgr.GetAppAccount(ctx, req.MisesId)
		if err != nil {
			return nil, err
		}

	} else if didtype == types.DIDTypeUser {
		userMgr := NewUserMgrImpl(k)
		misesAcc, err = userMgr.GetUserAccount(ctx, req.MisesId)
		if err != nil {
			return nil, err
		}
	}

	if misesAcc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "mises id %s not exists", req.MisesId)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))

	if err := k.cdc.Unmarshal(store.Get(GetDidRegistryIDBytes(misesAcc.DidRegistryID)), &DidRegistry); err != nil {
		return nil, err
	}

	return &types.RestQueryDidResponse{DidRegistry: &DidRegistry}, nil
}

func (k Keeper) QueryUser(c context.Context, req *types.RestQueryUserRequest) (*types.RestQueryUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var UserInfo types.UserInfo
	ctx := sdk.UnwrapSDKContext(c)
	userMgr := NewUserMgrImpl(k)
	misesAcc, err := userMgr.GetUserAccount(ctx, req.MisesUid)
	if err != nil {
		return nil, err
	}

	if misesAcc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "mises id %s not exists", req.MisesUid)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInfoKey))

	if err := k.cdc.Unmarshal(store.Get(GetUserInfoIDBytes(misesAcc.InfoID)), &UserInfo); err != nil {
		return nil, err
	}

	pubInfo := types.PublicUserInfo{}
	priInfo := UserInfo.PriInfo

	return &types.RestQueryUserResponse{PubInfo: &pubInfo, PriInfo: priInfo}, nil
}

// query user relations
func (k Keeper) QueryUserRelation(c context.Context, req *types.RestQueryUserRelationRequest) (*types.RestQueryUserRelationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	userMgr := NewUserMgrImpl(k)
	misesAcc, err := userMgr.GetUserAccount(ctx, req.MisesUid)
	if err != nil {
		return nil, err
	}

	if misesAcc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "mises id %s not exists", req.MisesUid)
	}
	pagination := req.Pagination
	if pagination == nil {
		pagination = &query.PageRequest{
			Key:   []byte(""),
			Limit: 100,
		}
	}
	relType := uint64(0)
	if req.Filter == "following" {
		relType = types.RelTypeBitFollow
	}
	if req.Filter == "blocking" {
		relType = types.RelTypeBitBlock
	}
	UserRelations, err := userMgr.GetUserRelations(ctx, relType, req.MisesUid, string(pagination.Key), int(pagination.Limit))
	if err != nil {
		return nil, err
	}
	misesList := []*types.MisesID{}
	for _, r := range UserRelations {
		misesList = append(misesList, &types.MisesID{MisesId: r.UidTo})
	}
	nextKey := ""
	if len(misesList) > 0 {
		nextKey = misesList[len(misesList)-1].MisesId
	}
	pageRes := &query.PageResponse{NextKey: []byte(nextKey)}

	resp := types.RestQueryUserRelationResponse{
		MisesList:  misesList,
		Pagination: pageRes,
	}
	return &resp, nil
}

func (k Keeper) QueryApp(c context.Context, req *types.RestQueryAppRequest) (*types.RestQueryAppResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var AppInfo types.AppInfo
	ctx := sdk.UnwrapSDKContext(c)
	appMgr := NewAppMgrImpl(k)
	misesAcc, err := appMgr.GetAppAccount(ctx, req.MisesAppid)
	if err != nil {
		return nil, err
	}

	if misesAcc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "mises id %s not exists", req.MisesAppid)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppInfoKey))

	if err := k.cdc.Unmarshal(store.Get(GetAppInfoIDBytes(misesAcc.InfoID)), &AppInfo); err != nil {
		return nil, err
	}

	pubInfo := AppInfo.PubInfo

	return &types.RestQueryAppResponse{PubInfo: pubInfo}, nil
}

func (k Keeper) QueryAppFeeGrant(c context.Context, req *types.RestQueryAppFeeGrantRequest) (*types.RestQueryAppFeeGrantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	appMgr := NewAppMgrImpl(k)
	fg, err := appMgr.GetAppFeeGrant(ctx, req.MisesAppid, req.MisesUid)
	if err != nil {
		return nil, err
	}

	return &types.RestQueryAppFeeGrantResponse{
		Grant: fg,
	}, nil

}
