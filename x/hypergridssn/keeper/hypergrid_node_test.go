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

func createNHypergridNode(keeper keeper.Keeper, ctx context.Context, n int) []types.HypergridNode {
	items := make([]types.HypergridNode, n)
	for i := range items {
		items[i].Pubkey = strconv.Itoa(i)

		keeper.SetHypergridNode(ctx, items[i])
	}
	return items
}

func TestHypergridNodeGet(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNHypergridNode(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetHypergridNode(ctx,
			item.Pubkey,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestHypergridNodeRemove(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNHypergridNode(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHypergridNode(ctx,
			item.Pubkey,
		)
		_, found := keeper.GetHypergridNode(ctx,
			item.Pubkey,
		)
		require.False(t, found)
	}
}

func TestHypergridNodeGetAll(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNHypergridNode(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHypergridNode(ctx)),
	)
}
