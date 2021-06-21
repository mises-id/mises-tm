package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
	"github.com/stretchr/testify/assert"
)

func createNAppInfo(keeper *Keeper, ctx sdk.Context, n int) []types.AppInfo {
	items := make([]types.AppInfo, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendAppInfo(ctx, items[i])
	}
	return items
}

func TestAppInfoGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAppInfo(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetAppInfo(ctx, item.Id))
	}
}

func TestAppInfoExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAppInfo(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasAppInfo(ctx, item.Id))
	}
}

func TestAppInfoRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAppInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAppInfo(ctx, item.Id)
		assert.False(t, keeper.HasAppInfo(ctx, item.Id))
	}
}

func TestAppInfoGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAppInfo(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllAppInfo(ctx))
}

func TestAppInfoCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAppInfo(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetAppInfoCount(ctx))
}
