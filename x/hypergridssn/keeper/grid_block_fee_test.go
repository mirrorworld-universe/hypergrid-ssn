package keeper_test

import (
	"context"
	"testing"

	keepertest "hypergrid-ssn/testutil/keeper"
	"hypergrid-ssn/testutil/nullify"
	"hypergrid-ssn/x/hypergridssn/keeper"
	"hypergrid-ssn/x/hypergridssn/types"

	"github.com/stretchr/testify/require"
)

func createNGridBlockFee(keeper keeper.Keeper, ctx context.Context, n int) []types.GridBlockFee {
	items := make([]types.GridBlockFee, n)
	for i := range items {
		items[i].Id = keeper.AppendGridBlockFee(ctx, items[i])
	}
	return items
}

func TestGridBlockFeeGet(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNGridBlockFee(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetGridBlockFee(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestGridBlockFeeRemove(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNGridBlockFee(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGridBlockFee(ctx, item.Id)
		_, found := keeper.GetGridBlockFee(ctx, item.Id)
		require.False(t, found)
	}
}

func TestGridBlockFeeGetAll(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNGridBlockFee(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGridBlockFee(ctx)),
	)
}

func TestGridBlockFeeCount(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNGridBlockFee(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetGridBlockFeeCount(ctx))
}
