package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
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
	ak := k.ak
	userMgr := NewUserMgrImpl(k.Keeper)
	misesAcc, err := userMgr.GetUserAccount(ctx, DidRegistry.Did)
	if misesAcc != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", msg.Did)
	}
	if err != nil {
		return nil, err
	}
	addr, err := types.AddrFormDid(DidRegistry.Did)
	if err != nil {
		return nil, err
	}
	baseAccount := ak.NewAccountWithAddress(ctx, addr)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
	}
	var acc authtypes.AccountI
	acc = baseAccount
	ak.SetAccount(ctx, acc)

	defer func() {
		telemetry.IncrCounter(1, "new", "account")
	}()

	regID := k.AppendDidRegistry(
		ctx,
		DidRegistry,
	)

	info := types.UserInfo{}
	info.Creator = DidRegistry.Creator

	infoID := k.AppendUserInfo(
		ctx,
		info,
	)

	newMisesAcc := types.MisesAccount{
		MisesID:       DidRegistry.Did,
		DidRegistryID: regID,
		UserInfoID:    infoID,
	}
	k.SetMisesAccount(ctx, newMisesAcc)

	return &types.MsgCreateDidRegistryResponse{
		Id: regID,
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
