package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
	"github.com/stretchr/testify/assert"
)

func createNUserInfo(keeper *Keeper, ctx sdk.Context, n int) []types.UserInfo {
	items := make([]types.UserInfo, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendUserInfo(ctx, items[i])
	}
	return items
}

func TestUserInfoGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUserInfo(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetUserInfo(ctx, item.Id))
	}
}

func TestUserInfoExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUserInfo(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasUserInfo(ctx, item.Id))
	}
}

func TestUserInfoRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUserInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserInfo(ctx, item.Id)
		assert.False(t, keeper.HasUserInfo(ctx, item.Id))
	}
}

func TestUserInfoGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUserInfo(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllUserInfo(ctx))
}

func TestUserInfoCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUserInfo(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetUserInfoCount(ctx))
}
