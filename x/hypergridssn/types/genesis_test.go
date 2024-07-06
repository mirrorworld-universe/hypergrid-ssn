package types_test

import (
	"testing"

	"hypergrid-ssn/x/hypergridssn/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				GridInboxList: []types.GridInbox{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				GridInboxCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated solanaAccount",
			genState: &types.GenesisState{
				SolanaAccountList: []types.SolanaAccount{
					{
						Address: "0",
						Version: "0",
					},
					{
						Address: "0",
						Version: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated gridBlockFee",
			genState: &types.GenesisState{
				GridBlockFeeList: []types.GridBlockFee{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid gridBlockFee count",
			genState: &types.GenesisState{
				GridBlockFeeList: []types.GridBlockFee{
					{
						Id: 1,
					},
				},
				GridBlockFeeCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated hypergridNode",
			genState: &types.GenesisState{
				HypergridNodeList: []types.HypergridNode{
					{
						Pubkey: "0",
					},
					{
						Pubkey: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated feeSettlementBill",
			genState: &types.GenesisState{
				FeeSettlementBillList: []types.FeeSettlementBill{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid feeSettlementBill count",
			genState: &types.GenesisState{
				FeeSettlementBillList: []types.FeeSettlementBill{
					{
						Id: 1,
					},
				},
				FeeSettlementBillCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated gridInbox",
			genState: &types.GenesisState{
				GridInboxList: []types.GridInbox{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid gridInbox count",
			genState: &types.GenesisState{
				GridInboxList: []types.GridInbox{
					{
						Id: 1,
					},
				},
				GridInboxCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
