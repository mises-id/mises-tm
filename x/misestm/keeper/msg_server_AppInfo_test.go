package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mises-id/mises-tm/x/misestm/types"
)

func TestAppInfoMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateAppInfo(ctx, &types.MsgCreateAppInfo{Creator: creator})
		require.NoError(t, err)
		assert.Equal(t, i, int(resp.Id))
	}
}

func TestAppInfoMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateAppInfo
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAppInfo{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAppInfo{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAppInfo{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateAppInfo(ctx, &types.MsgCreateAppInfo{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateAppInfo(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAppInfoMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteAppInfo
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAppInfo{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAppInfo{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteAppInfo{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateAppInfo(ctx, &types.MsgCreateAppInfo{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAppInfo(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
