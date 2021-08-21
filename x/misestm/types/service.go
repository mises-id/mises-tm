package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

// baseAppSimulateFn is the signature of the Baseapp#Simulate function.
type baseAppSimulateFn func(txBytes []byte) (sdk.GasInfo, *sdk.Result, error)

var Simulater baseAppSimulateFn
var Logger log.Logger

func RegisterBaseAppSimulateFn(
	logger log.Logger,
	simulateFn baseAppSimulateFn,
) {
	Logger = logger
	Simulater = simulateFn
}

const (
	InvalidID        = ^uint64(0)
	RelTypeBitFollow = uint64(1)
	RelTypeBitBlock  = uint64(2)
)

type UserMgr interface {
	dbm.TrackWriteListener

	GetUserAccount(ctx sdk.Context, did string) (*MisesAccount, error)
	GetUserRelation(ctx sdk.Context, didFrom string, didTo string) (*UserRelation, error)
	GetUserRelations(ctx sdk.Context, relType uint64, didFrom string, lastDidTo string, limit int) ([]*UserRelation, error)
}

func AddrFormDid(did string) (sdk.AccAddress, error) {
	addrStr := strings.Replace(did, "did:mises:", "", 1)

	return sdk.AccAddressFromBech32(addrStr)
}

type MsgReqBase struct {
	MisesID string `json:"mises_id,omitempty"`
}

type MsgCreateMisesID struct {
	MsgReqBase
	PubKey string `json:"pub_key,omitempty"`
}
