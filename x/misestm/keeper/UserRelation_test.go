package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
	"github.com/stretchr/testify/assert"
)

func createNUserRelation(keeper *Keeper, ctx sdk.Context, n int) []types.UserRelation {
	items := make([]types.UserRelation, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendUserRelation(ctx, items[i])
	}
	return items
}

func TestUserRelationGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUserRelation(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetUserRelation(ctx, item.Id))
	}
}

func TestUserRelationExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUserRelation(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasUserRelation(ctx, item.Id))
	}
}

func TestUserRelationRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUserRelation(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserRelation(ctx, item.Id)
		assert.False(t, keeper.HasUserRelation(ctx, item.Id))
	}
}

func TestUserRelationGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUserRelation(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllUserRelation(ctx))
}

func TestUserRelationCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUserRelation(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetUserRelationCount(ctx))
}
