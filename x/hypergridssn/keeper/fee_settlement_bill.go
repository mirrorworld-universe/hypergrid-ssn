package keeper

import (
	"context"
	"encoding/binary"

	"hypergrid-ssn/x/hypergridssn/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetFeeSettlementBillCount get the total number of feeSettlementBill
func (k Keeper) GetFeeSettlementBillCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.FeeSettlementBillCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetFeeSettlementBillCount set the total number of feeSettlementBill
func (k Keeper) SetFeeSettlementBillCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.FeeSettlementBillCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendFeeSettlementBill appends a feeSettlementBill in the store with a new id and update the count
func (k Keeper) AppendFeeSettlementBill(
	ctx context.Context,
	feeSettlementBill types.FeeSettlementBill,
) uint64 {
	// Create the feeSettlementBill
	count := k.GetFeeSettlementBillCount(ctx)

	// Set the ID of the appended value
	feeSettlementBill.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeSettlementBillKey))
	appendedValue := k.cdc.MustMarshal(&feeSettlementBill)
	store.Set(GetFeeSettlementBillIDBytes(feeSettlementBill.Id), appendedValue)

	// Update feeSettlementBill count
	k.SetFeeSettlementBillCount(ctx, count+1)

	return count
}

// SetFeeSettlementBill set a specific feeSettlementBill in the store
func (k Keeper) SetFeeSettlementBill(ctx context.Context, feeSettlementBill types.FeeSettlementBill) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeSettlementBillKey))
	b := k.cdc.MustMarshal(&feeSettlementBill)
	store.Set(GetFeeSettlementBillIDBytes(feeSettlementBill.Id), b)
}

// GetFeeSettlementBill returns a feeSettlementBill from its id
func (k Keeper) GetFeeSettlementBill(ctx context.Context, id uint64) (val types.FeeSettlementBill, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeSettlementBillKey))
	b := store.Get(GetFeeSettlementBillIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFeeSettlementBill removes a feeSettlementBill from the store
func (k Keeper) RemoveFeeSettlementBill(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeSettlementBillKey))
	store.Delete(GetFeeSettlementBillIDBytes(id))
}

// GetAllFeeSettlementBill returns all feeSettlementBill
func (k Keeper) GetAllFeeSettlementBill(ctx context.Context) (list []types.FeeSettlementBill) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeSettlementBillKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FeeSettlementBill
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFeeSettlementBillIDBytes returns the byte representation of the ID
func GetFeeSettlementBillIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.FeeSettlementBillKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
