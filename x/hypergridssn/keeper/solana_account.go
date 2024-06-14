package keeper

import (
	"context"

	"hypergrid-ssn/x/hypergridssn/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetSolanaAccount set a specific solanaAccount in the store from its index
func (k Keeper) SetSolanaAccount(ctx context.Context, solanaAccount types.SolanaAccount) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SolanaAccountKeyPrefix))
	b := k.cdc.MustMarshal(&solanaAccount)
	store.Set(types.SolanaAccountKey(
		solanaAccount.Address,
		solanaAccount.Version,
	), b)
}

// GetSolanaAccount returns a solanaAccount from its index
func (k Keeper) GetSolanaAccount(
	ctx context.Context,
	address string,
	version string,

) (val types.SolanaAccount, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SolanaAccountKeyPrefix))

	b := store.Get(types.SolanaAccountKey(
		address,
		version,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSolanaAccount removes a solanaAccount from the store
func (k Keeper) RemoveSolanaAccount(
	ctx context.Context,
	address string,
	version string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SolanaAccountKeyPrefix))
	store.Delete(types.SolanaAccountKey(
		address,
		version,
	))
}

// GetAllSolanaAccount returns all solanaAccount
func (k Keeper) GetAllSolanaAccount(ctx context.Context) (list []types.SolanaAccount) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SolanaAccountKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SolanaAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
