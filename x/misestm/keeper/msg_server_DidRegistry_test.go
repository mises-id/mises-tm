package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mises-id/mises-tm/x/misestm/types"
)

func TestDidRegistryMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		_, err := srv.CreateDidRegistry(ctx, &types.MsgCreateDidRegistry{Creator: creator})
		require.NoError(t, err)
	}
}
