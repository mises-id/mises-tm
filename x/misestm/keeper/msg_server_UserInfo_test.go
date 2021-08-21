package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mises-id/mises-tm/x/misestm/types"
)

func TestUserInfoMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateUserInfo(ctx, &types.MsgCreateUserInfo{Creator: creator})
		require.NoError(t, err)
		assert.Equal(t, i, int(resp.Id))
	}
}

func TestUserInfoMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateUserInfo
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateUserInfo{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateUserInfo{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateUserInfo{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateUserInfo(ctx, &types.MsgCreateUserInfo{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateUserInfo(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserInfoMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteUserInfo
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteUserInfo{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteUserInfo{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteUserInfo{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateUserInfo(ctx, &types.MsgCreateUserInfo{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteUserInfo(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
