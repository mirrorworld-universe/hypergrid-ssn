package hypergridssn

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"hypergrid-ssn/testutil/sample"
	hypergridssnsimulation "hypergrid-ssn/x/hypergridssn/simulation"
	"hypergrid-ssn/x/hypergridssn/types"
)

// avoid unused import issue
var (
	_ = hypergridssnsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateSolanaAccount = "op_weight_msg_solana_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSolanaAccount int = 100

	opWeightMsgUpdateSolanaAccount = "op_weight_msg_solana_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSolanaAccount int = 100

	opWeightMsgDeleteSolanaAccount = "op_weight_msg_solana_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSolanaAccount int = 100

	opWeightMsgCreateGridTxFee = "op_weight_msg_grid_tx_fee"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGridTxFee int = 100

	opWeightMsgCreateGridBlockFee = "op_weight_msg_grid_block_fee"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGridBlockFee int = 100

	opWeightMsgUpdateGridBlockFee = "op_weight_msg_grid_block_fee"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateGridBlockFee int = 100

	opWeightMsgDeleteGridBlockFee = "op_weight_msg_grid_block_fee"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteGridBlockFee int = 100

	opWeightMsgCreateHypergridNode = "op_weight_msg_hypergrid_node"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHypergridNode int = 100

	opWeightMsgUpdateHypergridNode = "op_weight_msg_hypergrid_node"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHypergridNode int = 100

	opWeightMsgDeleteHypergridNode = "op_weight_msg_hypergrid_node"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHypergridNode int = 100

	opWeightMsgCreateFeeSettlementBill = "op_weight_msg_fee_settlement_bill"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateFeeSettlementBill int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	hypergridssnGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		SolanaAccountList: []types.SolanaAccount{
			{
				Creator: sample.AccAddress(),
				Address: "0",
				Version: "0",
			},
			{
				Creator: sample.AccAddress(),
				Address: "1",
				Version: "1",
			},
		},
		GridBlockFeeList: []types.GridBlockFee{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		GridBlockFeeCount: 2,
		HypergridNodeList: []types.HypergridNode{
			{
				Creator: sample.AccAddress(),
				Pubkey:  "0",
			},
			{
				Creator: sample.AccAddress(),
				Pubkey:  "1",
			},
		},
		FeeSettlementBillList: []types.FeeSettlementBill{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		FeeSettlementBillCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&hypergridssnGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateSolanaAccount int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSolanaAccount, &weightMsgCreateSolanaAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSolanaAccount = defaultWeightMsgCreateSolanaAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSolanaAccount,
		hypergridssnsimulation.SimulateMsgCreateSolanaAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSolanaAccount int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateSolanaAccount, &weightMsgUpdateSolanaAccount, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSolanaAccount = defaultWeightMsgUpdateSolanaAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSolanaAccount,
		hypergridssnsimulation.SimulateMsgUpdateSolanaAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSolanaAccount int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteSolanaAccount, &weightMsgDeleteSolanaAccount, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSolanaAccount = defaultWeightMsgDeleteSolanaAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSolanaAccount,
		hypergridssnsimulation.SimulateMsgDeleteSolanaAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateGridBlockFee int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateGridBlockFee, &weightMsgCreateGridBlockFee, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGridBlockFee = defaultWeightMsgCreateGridBlockFee
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGridBlockFee,
		hypergridssnsimulation.SimulateMsgCreateGridBlockFee(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateHypergridNode int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateHypergridNode, &weightMsgCreateHypergridNode, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHypergridNode = defaultWeightMsgCreateHypergridNode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHypergridNode,
		hypergridssnsimulation.SimulateMsgCreateHypergridNode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHypergridNode int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateHypergridNode, &weightMsgUpdateHypergridNode, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHypergridNode = defaultWeightMsgUpdateHypergridNode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHypergridNode,
		hypergridssnsimulation.SimulateMsgUpdateHypergridNode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHypergridNode int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteHypergridNode, &weightMsgDeleteHypergridNode, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHypergridNode = defaultWeightMsgDeleteHypergridNode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHypergridNode,
		hypergridssnsimulation.SimulateMsgDeleteHypergridNode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateFeeSettlementBill int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateFeeSettlementBill, &weightMsgCreateFeeSettlementBill, nil,
		func(_ *rand.Rand) {
			weightMsgCreateFeeSettlementBill = defaultWeightMsgCreateFeeSettlementBill
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateFeeSettlementBill,
		hypergridssnsimulation.SimulateMsgCreateFeeSettlementBill(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateSolanaAccount,
			defaultWeightMsgCreateSolanaAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hypergridssnsimulation.SimulateMsgCreateSolanaAccount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateSolanaAccount,
			defaultWeightMsgUpdateSolanaAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hypergridssnsimulation.SimulateMsgUpdateSolanaAccount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteSolanaAccount,
			defaultWeightMsgDeleteSolanaAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hypergridssnsimulation.SimulateMsgDeleteSolanaAccount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateGridBlockFee,
			defaultWeightMsgCreateGridBlockFee,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hypergridssnsimulation.SimulateMsgCreateGridBlockFee(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateHypergridNode,
			defaultWeightMsgCreateHypergridNode,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hypergridssnsimulation.SimulateMsgCreateHypergridNode(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateHypergridNode,
			defaultWeightMsgUpdateHypergridNode,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hypergridssnsimulation.SimulateMsgUpdateHypergridNode(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteHypergridNode,
			defaultWeightMsgDeleteHypergridNode,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hypergridssnsimulation.SimulateMsgDeleteHypergridNode(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateFeeSettlementBill,
			defaultWeightMsgCreateFeeSettlementBill,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hypergridssnsimulation.SimulateMsgCreateFeeSettlementBill(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
