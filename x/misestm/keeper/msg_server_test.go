package keeper

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	keeper, ctx := setupKeeper(t)
	return NewMsgServerImpl(*keeper), sdk.WrapSDKContext(ctx)
}
