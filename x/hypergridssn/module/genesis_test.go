package hypergridssn_test

import (
	"testing"

	keepertest "hypergrid-ssn/testutil/keeper"
	"hypergrid-ssn/testutil/nullify"
	hypergridssn "hypergrid-ssn/x/hypergridssn/module"
	"hypergrid-ssn/x/hypergridssn/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SolanaAccountList: []types.SolanaAccount{
			{
				Address: "0",
				Version: "0",
			},
			{
				Address: "1",
				Version: "1",
			},
		},
		GridBlockFeeList: []types.GridBlockFee{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		GridBlockFeeCount: 2,
		HypergridNodeList: []types.HypergridNode{
			{
				Pubkey: "0",
			},
			{
				Pubkey: "1",
			},
		},
		FeeSettlementBillList: []types.FeeSettlementBill{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		FeeSettlementBillCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HypergridssnKeeper(t)
	hypergridssn.InitGenesis(ctx, k, genesisState)
	got := hypergridssn.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SolanaAccountList, got.SolanaAccountList)
	require.ElementsMatch(t, genesisState.GridBlockFeeList, got.GridBlockFeeList)
	require.Equal(t, genesisState.GridBlockFeeCount, got.GridBlockFeeCount)
	require.ElementsMatch(t, genesisState.HypergridNodeList, got.HypergridNodeList)
	require.ElementsMatch(t, genesisState.FeeSettlementBillList, got.FeeSettlementBillList)
	require.Equal(t, genesisState.FeeSettlementBillCount, got.FeeSettlementBillCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
