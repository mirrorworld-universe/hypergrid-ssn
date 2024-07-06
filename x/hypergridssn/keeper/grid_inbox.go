package keeper

import (
	"context"
	"encoding/binary"

	"hypergrid-ssn/x/hypergridssn/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetGridInboxCount get the total number of gridInbox
func (k Keeper) GetGridInboxCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.GridInboxCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetGridInboxCount set the total number of gridInbox
func (k Keeper) SetGridInboxCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.GridInboxCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendGridInbox appends a gridInbox in the store with a new id and update the count
func (k Keeper) AppendGridInbox(
	ctx context.Context,
	gridInbox types.GridInbox,
) uint64 {
	// Create the gridInbox
	count := k.GetGridInboxCount(ctx)

	// Set the ID of the appended value
	gridInbox.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridInboxKey))
	appendedValue := k.cdc.MustMarshal(&gridInbox)
	store.Set(GetGridInboxIDBytes(gridInbox.Id), appendedValue)

	// Update gridInbox count
	k.SetGridInboxCount(ctx, count+1)

	return count
}

// SetGridInbox set a specific gridInbox in the store
func (k Keeper) SetGridInbox(ctx context.Context, gridInbox types.GridInbox) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridInboxKey))
	b := k.cdc.MustMarshal(&gridInbox)
	store.Set(GetGridInboxIDBytes(gridInbox.Id), b)
}

// GetGridInbox returns a gridInbox from its id
func (k Keeper) GetGridInbox(ctx context.Context, id uint64) (val types.GridInbox, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridInboxKey))
	b := store.Get(GetGridInboxIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGridInbox removes a gridInbox from the store
func (k Keeper) RemoveGridInbox(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridInboxKey))
	store.Delete(GetGridInboxIDBytes(id))
}

// GetAllGridInbox returns all gridInbox
func (k Keeper) GetAllGridInbox(ctx context.Context) (list []types.GridInbox) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridInboxKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GridInbox
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetGridInboxIDBytes returns the byte representation of the ID
func GetGridInboxIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.GridInboxKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
