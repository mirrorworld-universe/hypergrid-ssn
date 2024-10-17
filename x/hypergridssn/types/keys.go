package types

const (
	// ModuleName defines the module name
	ModuleName = "hypergridssn"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hypergridssn"
)

var (
	ParamsKey = []byte("p_hypergridssn")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	GridTxFeeKey      = "GridTxFee/value/"
	GridTxFeeCountKey = "GridTxFee/count/"
)

const (
	GridBlockFeeKey      = "GridBlockFee/value/"
	GridBlockFeeCountKey = "GridBlockFee/count/"
	GridBlockhashKey     = "GridBlockFee/hash/"
)

const (
	FeeSettlementBillKey      = "FeeSettlementBill/value/"
	FeeSettlementBillCountKey = "FeeSettlementBill/count/"
)

const (
	GridInboxKey      = "GridInbox/value/"
	GridInboxCountKey = "GridInbox/count/"
)
