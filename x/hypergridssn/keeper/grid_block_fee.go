package keeper

import (
	"context"
	"encoding/binary"

	"hypergrid-ssn/x/hypergridssn/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetGridBlockFeeCount get the total number of gridBlockFee
func (k Keeper) GetGridBlockFeeCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.GridBlockFeeCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetGridBlockFeeCount set the total number of gridBlockFee
func (k Keeper) SetGridBlockFeeCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.GridBlockFeeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendGridBlockFee appends a gridBlockFee in the store with a new id and update the count
func (k Keeper) AppendGridBlockFee(
	ctx context.Context,
	gridBlockFee types.GridBlockFee,
) uint64 {
	// Create the gridBlockFee
	count := k.GetGridBlockFeeCount(ctx)

	// Set the ID of the appended value
	gridBlockFee.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridBlockFeeKey))
	appendedValue := k.cdc.MustMarshal(&gridBlockFee)
	store.Set(GetGridBlockFeeIDBytes(gridBlockFee.Id), appendedValue)

	// Save the hash of the gridBlockFee to make blockhash unique.
	store2 := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridBlockhashKey))
	hashBytes := []byte(gridBlockFee.Blockhash)
	store2.Set(GetGridBlockFeeHashBytes(gridBlockFee.Blockhash), hashBytes)

	// Update gridBlockFee count
	k.SetGridBlockFeeCount(ctx, count+1)

	return count
}

// SetGridBlockFee set a specific gridBlockFee in the store
func (k Keeper) SetGridBlockFee(ctx context.Context, gridBlockFee types.GridBlockFee) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridBlockFeeKey))
	b := k.cdc.MustMarshal(&gridBlockFee)
	store.Set(GetGridBlockFeeIDBytes(gridBlockFee.Id), b)
}

// GetGridBlockFee returns a gridBlockFee from its id
func (k Keeper) GetGridBlockFee(ctx context.Context, id uint64) (val types.GridBlockFee, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridBlockFeeKey))
	b := store.Get(GetGridBlockFeeIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGridBlockFee removes a gridBlockFee from the store
func (k Keeper) RemoveGridBlockFee(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridBlockFeeKey))
	store.Delete(GetGridBlockFeeIDBytes(id))
}

// GetAllGridBlockFee returns all gridBlockFee
func (k Keeper) GetAllGridBlockFee(ctx context.Context) (list []types.GridBlockFee) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridBlockFeeKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GridBlockFee
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// / Check if the key exists in the store
func (k Keeper) HasGridBlockFeeHash(
	ctx context.Context,
	hash string,

) bool {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GridBlockhashKey))

	b := store.Get(GetGridBlockFeeHashBytes(
		hash,
	))
	return b != nil
}

// GetGridBlockFeeIDBytes returns the byte representation of the ID
func GetGridBlockFeeIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.GridBlockFeeKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

// GetGridBlockFeeIDBytes returns the byte representation of the ID
func GetGridBlockFeeHashBytes(
	hash string,
) []byte {
	var key []byte
	hashBytes := []byte(hash)
	key = append(key, hashBytes...)
	key = append(key, []byte("/")...)

	return key
}
