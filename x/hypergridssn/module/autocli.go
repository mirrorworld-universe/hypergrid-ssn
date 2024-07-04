package hypergridssn

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "hypergrid-ssn/api/hypergridssn/hypergridssn"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "SolanaAccountAll",
					Use:       "list-solana-account",
					Short:     "List all SolanaAccount",
				},
				{
					RpcMethod:      "SolanaAccount",
					Use:            "show-solana-account [id]",
					Short:          "Shows a SolanaAccount",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "version"}},
				},
				{
					RpcMethod: "GridBlockFeeAll",
					Use:       "list-grid-block-fee",
					Short:     "List all GridBlockFee",
				},
				{
					RpcMethod:      "GridBlockFee",
					Use:            "show-grid-block-fee [id]",
					Short:          "Shows a GridBlockFee by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "HypergridNodeAll",
					Use:       "list-hypergrid-node",
					Short:     "List all HypergridNode",
				},
				{
					RpcMethod:      "HypergridNode",
					Use:            "show-hypergrid-node [id]",
					Short:          "Shows a HypergridNode",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pubkey"}},
				},
				{
					RpcMethod: "FeeSettlementBillAll",
					Use:       "list-fee-settlement-bill",
					Short:     "List all FeeSettlementBill",
				},
				{
					RpcMethod:      "FeeSettlementBill",
					Use:            "show-fee-settlement-bill [id]",
					Short:          "Shows a FeeSettlementBill by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateSolanaAccount",
					Use:            "create-solana-account [address] [version] [source]",
					Short:          "Create a new SolanaAccount",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "version"}, {ProtoField: "source"}},
				},
				{
					RpcMethod:      "UpdateSolanaAccount",
					Use:            "update-solana-account [address] [version]",
					Short:          "Update SolanaAccount",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "version"}},
				},
				{
					RpcMethod:      "DeleteSolanaAccount",
					Use:            "delete-solana-account [address] [version]",
					Short:          "Delete SolanaAccount",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "version"}},
				},
				{
					RpcMethod:      "CreateGridBlockFee",
					Use:            "create-grid-block-fee [items]",
					Short:          "Create GridBlockFee",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "items"}},
				},
				{
					RpcMethod:      "CreateHypergridNode",
					Use:            "create-hypergrid-node [pubkey] [name] [rpc] [data_account] [role] [starttime]",
					Short:          "Create a new HypergridNode",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pubkey"}, {ProtoField: "name"}, {ProtoField: "rpc"}, {ProtoField: "role"}, {ProtoField: "data_account"}, {ProtoField: "starttime"}},
				},
				{
					RpcMethod:      "UpdateHypergridNode",
					Use:            "update-hypergrid-node [pubkey] [name] [rpc] [data_account] [role] [starttime]",
					Short:          "Update HypergridNode",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pubkey"}, {ProtoField: "name"}, {ProtoField: "rpc"}, {ProtoField: "role"}, {ProtoField: "data_account"}, {ProtoField: "starttime"}},
				},
				{
					RpcMethod:      "DeleteHypergridNode",
					Use:            "delete-hypergrid-node [pubkey]",
					Short:          "Delete HypergridNode",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pubkey"}},
				},
				{
					RpcMethod:      "CreateFeeSettlementBill",
					Use:            "create-fee-settlement-bill [fromId] [endId]",
					Short:          "Create FeeSettlementBill",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "fromId"}, {ProtoField: "endId"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
