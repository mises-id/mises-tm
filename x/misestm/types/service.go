package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
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
	InvalidID            = ^uint64(0)
	RelTypeBitFollow     = uint64(1)
	RelTypeBitBlock      = uint64(2)
	RelTypeBitReferredBy = uint64(4)

	DIDPrefixForUser = "did:mises:"
	DIDPrefixForApp  = "did:misesapp:"
	DIDTypeUser      = uint64(1)
	DIDTypeApp       = uint64(2)
)

type AppMgr interface {
	GetAppAccount(ctx sdk.Context, did string) (*MisesAccount, error)
	GetAppFeeGrant(ctx sdk.Context, appDID string, userDID string) (ret *AppFeeGrant, err error)
}

type UserMgr interface {
	GetUserAccount(ctx sdk.Context, did string) (*MisesAccount, error)
	GetUserRelation(ctx sdk.Context, didFrom string, didTo string) (*UserRelation, error)
	GetUserRelations(ctx sdk.Context, relType uint64, didFrom string, lastDidTo string, limit int) ([]*UserRelation, error)
}

func CheckDid(did string, didType uint64) (string, bool) {
	if didType == DIDTypeUser && strings.HasPrefix(did, DIDPrefixForUser) {
		addrStr := strings.Replace(did, DIDPrefixForUser, "", 1)
		return addrStr, true
	} else if didType == DIDTypeApp && strings.HasPrefix(did, DIDPrefixForApp) {
		addrStr := strings.Replace(did, DIDPrefixForApp, "", 1)
		return addrStr, true
	}
	return "", false
}
func AddrFromDid(did string) (sdk.AccAddress, uint64, error) {
	var addrStr string
	var didType uint64
	if strings.HasPrefix(did, DIDPrefixForUser) {
		addrStr = strings.Replace(did, DIDPrefixForUser, "", 1)
		didType = DIDTypeUser
	} else if strings.HasPrefix(did, DIDPrefixForApp) {
		addrStr = strings.Replace(did, DIDPrefixForApp, "", 1)
		didType = DIDTypeApp
	} else {
		return nil, didType, fmt.Errorf("unsupported did prefix")
	}
	addr, error := sdk.AccAddressFromBech32(addrStr)

	return addr, didType, error
}

type MsgReqBase struct {
	MisesID string `json:"mises_id,omitempty"`
}

type MsgCreateMisesID struct {
	MsgReqBase
	PubKey string `json:"pub_key,omitempty"`
}
