package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SolanaAccountList:     []SolanaAccount{},
		GridBlockFeeList:      []GridBlockFee{},
		HypergridNodeList:     []HypergridNode{},
		FeeSettlementBillList: []FeeSettlementBill{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in solanaAccount
	solanaAccountIndexMap := make(map[string]struct{})

	for _, elem := range gs.SolanaAccountList {
		index := string(SolanaAccountKey(elem.Address, elem.Version))
		if _, ok := solanaAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for solanaAccount")
		}
		solanaAccountIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in gridBlockFee
	gridBlockFeeIdMap := make(map[uint64]bool)
	gridBlockFeeCount := gs.GetGridBlockFeeCount()
	for _, elem := range gs.GridBlockFeeList {
		if _, ok := gridBlockFeeIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for gridBlockFee")
		}
		if elem.Id >= gridBlockFeeCount {
			return fmt.Errorf("gridBlockFee id should be lower or equal than the last id")
		}
		gridBlockFeeIdMap[elem.Id] = true
	}
	// Check for duplicated index in hypergridNode
	hypergridNodeIndexMap := make(map[string]struct{})

	for _, elem := range gs.HypergridNodeList {
		index := string(HypergridNodeKey(elem.Pubkey))
		if _, ok := hypergridNodeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for hypergridNode")
		}
		hypergridNodeIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in feeSettlementBill
	feeSettlementBillIdMap := make(map[uint64]bool)
	feeSettlementBillCount := gs.GetFeeSettlementBillCount()
	for _, elem := range gs.FeeSettlementBillList {
		if _, ok := feeSettlementBillIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for feeSettlementBill")
		}
		if elem.Id >= feeSettlementBillCount {
			return fmt.Errorf("feeSettlementBill id should be lower or equal than the last id")
		}
		feeSettlementBillIdMap[elem.Id] = true
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
