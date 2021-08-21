package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mises-id/mises-tm/x/misestm/types"
)

func TestUserRelationMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateUserRelation(ctx, &types.MsgCreateUserRelation{Creator: creator})
		require.NoError(t, err)
		assert.Equal(t, i, int(resp.Id))
	}
}

func TestUserRelationMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateUserRelation
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateUserRelation{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateUserRelation{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateUserRelation{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateUserRelation(ctx, &types.MsgCreateUserRelation{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateUserRelation(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserRelationMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteUserRelation
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteUserRelation{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteUserRelation{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteUserRelation{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateUserRelation(ctx, &types.MsgCreateUserRelation{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteUserRelation(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
