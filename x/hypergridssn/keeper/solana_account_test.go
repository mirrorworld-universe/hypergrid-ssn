package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "hypergrid-ssn/testutil/keeper"
	"hypergrid-ssn/testutil/nullify"
	"hypergrid-ssn/x/hypergridssn/keeper"
	"hypergrid-ssn/x/hypergridssn/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSolanaAccount(keeper keeper.Keeper, ctx context.Context, n int) []types.SolanaAccount {
	items := make([]types.SolanaAccount, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].Version = strconv.Itoa(i)

		keeper.SetSolanaAccount(ctx, items[i])
	}
	return items
}

func TestSolanaAccountGet(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNSolanaAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSolanaAccount(ctx,
			item.Address,
			item.Version,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSolanaAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNSolanaAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSolanaAccount(ctx,
			item.Address,
			item.Version,
		)
		_, found := keeper.GetSolanaAccount(ctx,
			item.Address,
			item.Version,
		)
		require.False(t, found)
	}
}

func TestSolanaAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNSolanaAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSolanaAccount(ctx)),
	)
}
