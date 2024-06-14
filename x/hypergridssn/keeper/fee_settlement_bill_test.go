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

func createNFeeSettlementBill(keeper keeper.Keeper, ctx context.Context, n int) []types.FeeSettlementBill {
	items := make([]types.FeeSettlementBill, n)
	for i := range items {
		items[i].Id = keeper.AppendFeeSettlementBill(ctx, items[i])
	}
	return items
}

func TestFeeSettlementBillGet(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNFeeSettlementBill(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetFeeSettlementBill(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestFeeSettlementBillRemove(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNFeeSettlementBill(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFeeSettlementBill(ctx, item.Id)
		_, found := keeper.GetFeeSettlementBill(ctx, item.Id)
		require.False(t, found)
	}
}

func TestFeeSettlementBillGetAll(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNFeeSettlementBill(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFeeSettlementBill(ctx)),
	)
}

func TestFeeSettlementBillCount(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	items := createNFeeSettlementBill(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetFeeSettlementBillCount(ctx))
}
