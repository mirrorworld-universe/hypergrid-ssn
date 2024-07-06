package hypergridssn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"hypergrid-ssn/x/hypergridssn/keeper"
	"hypergrid-ssn/x/hypergridssn/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the solanaAccount
	for _, elem := range genState.SolanaAccountList {
		k.SetSolanaAccount(ctx, elem)
	}

	// Set gridBlockFee count
	k.SetGridBlockFeeCount(ctx, genState.GridBlockFeeCount)
	// Set all the hypergridNode
	for _, elem := range genState.HypergridNodeList {
		k.SetHypergridNode(ctx, elem)
	}
	// Set all the feeSettlementBill
	for _, elem := range genState.FeeSettlementBillList {
		k.SetFeeSettlementBill(ctx, elem)
	}

	// Set feeSettlementBill count
	k.SetFeeSettlementBillCount(ctx, genState.FeeSettlementBillCount)
	// Set all the gridInbox
	for _, elem := range genState.GridInboxList {
		k.SetGridInbox(ctx, elem)
	}

	// Set gridInbox count
	k.SetGridInboxCount(ctx, genState.GridInboxCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.SolanaAccountList = k.GetAllSolanaAccount(ctx)
	genesis.GridBlockFeeList = k.GetAllGridBlockFee(ctx)
	genesis.GridBlockFeeCount = k.GetGridBlockFeeCount(ctx)
	genesis.HypergridNodeList = k.GetAllHypergridNode(ctx)
	genesis.FeeSettlementBillList = k.GetAllFeeSettlementBill(ctx)
	genesis.FeeSettlementBillCount = k.GetFeeSettlementBillCount(ctx)
	genesis.GridInboxList = k.GetAllGridInbox(ctx)
	genesis.GridInboxCount = k.GetGridInboxCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
