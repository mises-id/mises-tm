package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
	"github.com/stretchr/testify/assert"
)

func createNDidRegistry(keeper *Keeper, ctx sdk.Context, n int) []types.DidRegistry {
	items := make([]types.DidRegistry, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendDidRegistry(ctx, items[i])
	}
	return items
}

func TestDidRegistryGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDidRegistry(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetDidRegistry(ctx, item.Id))
	}
}

func TestDidRegistryExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDidRegistry(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasDidRegistry(ctx, item.Id))
	}
}

func TestDidRegistryRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDidRegistry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDidRegistry(ctx, item.Id)
		assert.False(t, keeper.HasDidRegistry(ctx, item.Id))
	}
}

func TestDidRegistryGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDidRegistry(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllDidRegistry(ctx))
}

func TestDidRegistryCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDidRegistry(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetDidRegistryCount(ctx))
}
