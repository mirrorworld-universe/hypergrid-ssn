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

func createNGridInbox(keeper keeper.Keeper, ctx context.Context, n int) []types.GridInbox {
	items := make([]types.GridInbox, n)
	for i := range items {
		items[i].Id = keeper.AppendGridInbox(ctx, items[i])
	}
	return items
}

func TestGridInboxGet(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNGridInbox(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetGridInbox(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestGridInboxRemove(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNGridInbox(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGridInbox(ctx, item.Id)
		_, found := keeper.GetGridInbox(ctx, item.Id)
		require.False(t, found)
	}
}

func TestGridInboxGetAll(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNGridInbox(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGridInbox(ctx)),
	)
}

func TestGridInboxCount(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNGridInbox(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetGridInboxCount(ctx))
}
