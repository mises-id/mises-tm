package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
)

func AddrFormDid(did string) string {
	return strings.Replace(did, "did:mises:", "", 1)
}

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
	addr, err := sdk.AccAddressFromBech32(AddrFormDid(DidRegistry.Did))
	if err != nil {
		return nil, err
	}

	if acc := ak.GetAccount(ctx, addr); acc != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", msg.Did)
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
