package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

type appMgr struct {
	Keeper
}

func NewAppMgrImpl(keeper Keeper) types.AppMgr {
	return &appMgr{Keeper: keeper}
}

var _ types.AppMgr = &appMgr{}

func (k *appMgr) getAccount(ctx sdk.Context, did string) (*types.MisesAccount, error) {
	ak := k.ak
	addr, didtype, err := types.AddrFromDid(did)
	if err != nil {
		return nil, err
	}

	if acc := ak.GetAccount(ctx, addr); acc == nil {
		return nil, nil
	}
	if !k.HasMisesAccount(ctx, did) {
		return nil, sdkerrors.ErrKeyNotFound
	}
	misesAcc := k.GetMisesAccount(ctx, did)

	if didtype != misesAcc.DidType {
		return nil, sdkerrors.ErrLogic
	}
	return &misesAcc, nil
}

func (k *appMgr) GetAppAccount(ctx sdk.Context, did string) (*types.MisesAccount, error) {

	misesAcc, err := k.getAccount(ctx, did)
	if err != nil {
		return nil, err
	}
	if misesAcc.DidType != types.DIDTypeApp {
		return nil, sdkerrors.ErrLogic
	}
	return misesAcc, nil
}

func (k *appMgr) GetAppFeeGrant(ctx sdk.Context, appDID string, userDID string) (ret *types.AppFeeGrant, err error) {

	_, err = k.GetAppAccount(ctx, appDID)
	if err != nil {
		return nil, err
	}
	appAddr, _, err := types.AddrFromDid(appDID)
	if err != nil {
		return nil, err
	}
	userAddr, _, err := types.AddrFromDid(userDID)
	if err != nil {
		return nil, err
	}
	feeAllowance, err := k.fk.GetAllowance(ctx, appAddr, userAddr)
	if err != nil {
		return nil, err
	}
	pa, ok := feeAllowance.(*feegrant.PeriodicAllowance)
	if !ok {
		return nil, sdkerrors.ErrInvalidType
	}
	if len(pa.Basic.SpendLimit) != 1 {
		return nil, sdkerrors.ErrInvalidType
	}

	ret = &types.AppFeeGrant{
		Period:     pa.Period,
		SpendLimit: &pa.Basic.SpendLimit[0],
		Expiration: pa.Basic.Expiration,
	}
	return ret, nil
}
