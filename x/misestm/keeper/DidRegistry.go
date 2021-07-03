package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

// GetDidRegistryCount get the total number of DidRegistry
func (k Keeper) GetDidRegistryCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryCountKey))
	byteKey := types.KeyPrefix(types.DidRegistryCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetDidRegistryCount set the total number of DidRegistry
func (k Keeper) SetDidRegistryCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryCountKey))
	byteKey := types.KeyPrefix(types.DidRegistryCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendDidRegistry appends a DidRegistry in the store with a new id and update the count
func (k Keeper) AppendDidRegistry(
	ctx sdk.Context,
	DidRegistry types.DidRegistry,
) uint64 {
	// Create the DidRegistry
	count := k.GetDidRegistryCount(ctx)

	// Set the ID of the appended value
	DidRegistry.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))
	appendedValue := k.cdc.MustMarshal(&DidRegistry)
	store.Set(GetDidRegistryIDBytes(DidRegistry.Id), appendedValue)

	// Update DidRegistry count
	k.SetDidRegistryCount(ctx, count+1)

	return count
}

// SetDidRegistry set a specific DidRegistry in the store
func (k Keeper) SetDidRegistry(ctx sdk.Context, DidRegistry types.DidRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))
	b := k.cdc.MustMarshal(&DidRegistry)
	store.Set(GetDidRegistryIDBytes(DidRegistry.Id), b)
}

// GetDidRegistry returns a DidRegistry from its id
func (k Keeper) GetDidRegistry(ctx sdk.Context, id uint64) types.DidRegistry {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))
	var DidRegistry types.DidRegistry
	k.cdc.MustUnmarshal(store.Get(GetDidRegistryIDBytes(id)), &DidRegistry)
	return DidRegistry
}

// HasDidRegistry checks if the DidRegistry exists in the store
func (k Keeper) HasDidRegistry(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))
	return store.Has(GetDidRegistryIDBytes(id))
}

// GetDidRegistryOwner returns the creator of the DidRegistry
func (k Keeper) GetDidRegistryOwner(ctx sdk.Context, id uint64) string {
	return k.GetDidRegistry(ctx, id).Creator
}

// RemoveDidRegistry removes a DidRegistry from the store
func (k Keeper) RemoveDidRegistry(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))
	store.Delete(GetDidRegistryIDBytes(id))
}

// GetAllDidRegistry returns all DidRegistry
func (k Keeper) GetAllDidRegistry(ctx sdk.Context) (list []types.DidRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidRegistryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDidRegistryIDBytes returns the byte representation of the ID
func GetDidRegistryIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDidRegistryIDFromBytes returns ID in uint64 format from a byte array
func GetDidRegistryIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
