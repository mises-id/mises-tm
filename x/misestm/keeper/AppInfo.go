package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

// GetAppInfoCount get the total number of AppInfo
func (k Keeper) GetAppInfoCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppInfoCountKey))
	byteKey := types.KeyPrefix(types.AppInfoCountKey)
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

// SetAppInfoCount set the total number of AppInfo
func (k Keeper) SetAppInfoCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppInfoCountKey))
	byteKey := types.KeyPrefix(types.AppInfoCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendAppInfo appends a AppInfo in the store with a new id and update the count
func (k Keeper) AppendAppInfo(
	ctx sdk.Context,
	AppInfo types.AppInfo,
) uint64 {
	// Create the AppInfo
	count := k.GetAppInfoCount(ctx)

	// Set the ID of the appended value
	AppInfo.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppInfoKey))
	appendedValue := k.cdc.MustMarshal(&AppInfo)
	store.Set(GetAppInfoIDBytes(AppInfo.Id), appendedValue)

	// Update AppInfo count
	k.SetAppInfoCount(ctx, count+1)

	return count
}

// SetAppInfo set a specific AppInfo in the store
func (k Keeper) SetAppInfo(ctx sdk.Context, AppInfo types.AppInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppInfoKey))
	b := k.cdc.MustMarshal(&AppInfo)
	store.Set(GetAppInfoIDBytes(AppInfo.Id), b)
}

// GetAppInfo returns a AppInfo from its id
func (k Keeper) GetAppInfo(ctx sdk.Context, id uint64) types.AppInfo {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppInfoKey))
	var AppInfo types.AppInfo
	k.cdc.MustUnmarshal(store.Get(GetAppInfoIDBytes(id)), &AppInfo)
	return AppInfo
}

// HasAppInfo checks if the AppInfo exists in the store
func (k Keeper) HasAppInfo(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppInfoKey))
	return store.Has(GetAppInfoIDBytes(id))
}

// GetAppInfoOwner returns the creator of the AppInfo
func (k Keeper) GetAppInfoOwner(ctx sdk.Context, id uint64) string {
	return k.GetAppInfo(ctx, id).Creator
}

// RemoveAppInfo removes a AppInfo from the store
func (k Keeper) RemoveAppInfo(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppInfoKey))
	store.Delete(GetAppInfoIDBytes(id))
}

// GetAllAppInfo returns all AppInfo
func (k Keeper) GetAllAppInfo(ctx sdk.Context) (list []types.AppInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppInfoKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AppInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAppInfoIDBytes returns the byte representation of the ID
func GetAppInfoIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAppInfoIDFromBytes returns ID in uint64 format from a byte array
func GetAppInfoIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
