package keeper

import (
	"context"

	"hypergrid-ssn/x/hypergridssn/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetHypergridNode set a specific hypergridNode in the store from its index
func (k Keeper) SetHypergridNode(ctx context.Context, hypergridNode types.HypergridNode) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HypergridNodeKeyPrefix))
	b := k.cdc.MustMarshal(&hypergridNode)
	store.Set(types.HypergridNodeKey(
		hypergridNode.Pubkey,
	), b)
}

// GetHypergridNode returns a hypergridNode from its index
func (k Keeper) GetHypergridNode(
	ctx context.Context,
	pubkey string,

) (val types.HypergridNode, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HypergridNodeKeyPrefix))

	b := store.Get(types.HypergridNodeKey(
		pubkey,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) HasHypergridNode(
	ctx context.Context,
	pubkey string,

) bool {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HypergridNodeKeyPrefix))

	b := store.Get(types.HypergridNodeKey(
		pubkey,
	))
	return b != nil
}

// RemoveHypergridNode removes a hypergridNode from the store
func (k Keeper) RemoveHypergridNode(
	ctx context.Context,
	pubkey string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HypergridNodeKeyPrefix))
	store.Delete(types.HypergridNodeKey(
		pubkey,
	))
}

// GetAllHypergridNode returns all hypergridNode
func (k Keeper) GetAllHypergridNode(ctx context.Context) (list []types.HypergridNode) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HypergridNodeKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HypergridNode
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
