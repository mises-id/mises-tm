package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

// SetAppInfo set a specific AppInfo in the store
func (k Keeper) SetMisesAccount(ctx sdk.Context, acc types.MisesAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MisesAccountKey))
	b := k.cdc.MustMarshal(&acc)
	store.Set(GetMisesAccountKeyBytes(acc.MisesID), b)
}

// GetMisesAccount returns a AppInfo from its id
func (k Keeper) GetMisesAccount(ctx sdk.Context, misesID string) types.MisesAccount {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MisesAccountKey))
	var acc types.MisesAccount
	k.cdc.MustUnmarshal(store.Get(GetMisesAccountKeyBytes(misesID)), &acc)
	return acc
}

// HasMisesAccount returns a MisesAccount from its misesID
func (k Keeper) HasMisesAccount(ctx sdk.Context, misesID string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MisesAccountKey))
	return store.Has(GetMisesAccountKeyBytes(misesID))
}

// GetAllMisesAccount returns all MisesAccounts
func (k Keeper) GetAllMisesAccount(ctx sdk.Context) (list []types.MisesAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MisesAccountKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MisesAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func GetMisesAccountKeyBytes(misesID string) []byte {
	return []byte(misesID)
}
