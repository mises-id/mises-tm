package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mises-id/mises-tm/x/misestm/types"
)

func TestDidRegistryMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateDidRegistry(ctx, &types.MsgCreateDidRegistry{Creator: creator})
		require.NoError(t, err)
		assert.Equal(t, i, int(resp.Id))
	}
}

func TestDidRegistryMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDidRegistry
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateDidRegistry{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateDidRegistry{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateDidRegistry{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateDidRegistry(ctx, &types.MsgCreateDidRegistry{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateDidRegistry(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDidRegistryMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteDidRegistry
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteDidRegistry{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteDidRegistry{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteDidRegistry{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateDidRegistry(ctx, &types.MsgCreateDidRegistry{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteDidRegistry(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
