package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

// GetUserInfoCount get the total number of UserInfo
func (k Keeper) GetUserInfoCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInfoCountKey))
	byteKey := types.KeyPrefix(types.UserInfoCountKey)
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

// SetUserInfoCount set the total number of UserInfo
func (k Keeper) SetUserInfoCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInfoCountKey))
	byteKey := types.KeyPrefix(types.UserInfoCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendUserInfo appends a UserInfo in the store with a new id and update the count
func (k Keeper) AppendUserInfo(
	ctx sdk.Context,
	UserInfo types.UserInfo,
) uint64 {
	// Create the UserInfo
	count := k.GetUserInfoCount(ctx)

	// Set the ID of the appended value
	UserInfo.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInfoKey))
	appendedValue := k.cdc.MustMarshal(&UserInfo)
	store.Set(GetUserInfoIDBytes(UserInfo.Id), appendedValue)

	// Update UserInfo count
	k.SetUserInfoCount(ctx, count+1)

	return count
}

// SetUserInfo set a specific UserInfo in the store
func (k Keeper) SetUserInfo(ctx sdk.Context, UserInfo types.UserInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInfoKey))
	b := k.cdc.MustMarshal(&UserInfo)
	store.Set(GetUserInfoIDBytes(UserInfo.Id), b)
}

// GetUserInfo returns a UserInfo from its id
func (k Keeper) GetUserInfo(ctx sdk.Context, id uint64) types.UserInfo {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInfoKey))
	var UserInfo types.UserInfo
	k.cdc.MustUnmarshal(store.Get(GetUserInfoIDBytes(id)), &UserInfo)
	return UserInfo
}

// HasUserInfo checks if the UserInfo exists in the store
func (k Keeper) HasUserInfo(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInfoKey))
	return store.Has(GetUserInfoIDBytes(id))
}

// GetUserInfoOwner returns the creator of the UserInfo
func (k Keeper) GetUserInfoOwner(ctx sdk.Context, id uint64) string {
	return k.GetUserInfo(ctx, id).Creator
}

// RemoveUserInfo removes a UserInfo from the store
func (k Keeper) RemoveUserInfo(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInfoKey))
	store.Delete(GetUserInfoIDBytes(id))
}

// GetAllUserInfo returns all UserInfo
func (k Keeper) GetAllUserInfo(ctx sdk.Context) (list []types.UserInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInfoKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetUserInfoIDBytes returns the byte representation of the ID
func GetUserInfoIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetUserInfoIDFromBytes returns ID in uint64 format from a byte array
func GetUserInfoIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
