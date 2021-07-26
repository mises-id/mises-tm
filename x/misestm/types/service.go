package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"

	dbm "github.com/tendermint/tm-db"
)

// baseAppSimulateFn is the signature of the Baseapp#Simulate function.
type baseAppSimulateFn func(txBytes []byte) (sdk.GasInfo, *sdk.Result, error)

var Simulater baseAppSimulateFn

func RegisterBaseAppSimulateFn(
	simulateFn baseAppSimulateFn,
) {
	Simulater = simulateFn
}

const (
	InvalidID        = ^uint64(0)
	RelTypeBitFollow = uint64(1)
	RelTypeBitBlock  = uint64(2)
)

type UserMgr interface {
	dbm.TrackWriteListener
	AddrFormDid(did string) (types.AccAddress, error)
	GetUserAccount(ctx sdk.Context, did string) (*MisesAccount, error)
	GetUserRelation(ctx sdk.Context, didFrom string, didTo string) (*UserRelation, error)
	GetUserRelations(ctx sdk.Context, didFrom string, lastDidTo string, limit int) ([]*UserRelation, error)
}
