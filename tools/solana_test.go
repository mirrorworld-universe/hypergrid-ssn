package tools

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestSendTxFeeSettlement(t *testing.T) {
	// FromId := uint64(1)
	// EndID := uint64(100)
	// bills := map[string]uint64{
	// 	"EyYFxQ2FRcSkR8rdvefEDNy69KWHi2xTzbuVKxuBVueS": 100,
	// 	"AzVnQCpY2rqQmxJz6PQzxWj9HHQQc5qFuL89wvov6cL4": 101,
	// }

	// sig, err := SendTxFeeSettlement("http://localhost:8899", []string{"AzVnQCpY2rqQmxJz6PQzxWj9HHQQc5qFuL89wvov6cL4"}, FromId, EndID, bills)
	// if err != nil {
	// 	panic(err)
	// }
	// spew.Dump(sig)

	sig, key, err := SendTxInbox("https://api.devnet.solana.com", 53893, "2YmCa9RBVN9CZjAcN3o431UhTEU8BwmaiGoCnbeK1Sgr")
	if err != nil {
		panic(err)
	}
	spew.Dump(sig)
	spew.Dump(key)

	// Bills := []SettlementBillParam{}
	// // convert bills to []SettlementBillParam
	// for key, value := range bills {
	// 	Bills = append(Bills, SettlementBillParam{
	// 		Key:    solana.MustPublicKeyFromBase58(key),
	// 		Amount: value,
	// 	})
	// }

	// instructionData := SettleFeeBillParams{
	// 	Instruction: 1,
	// 	FromID:      FromId,
	// 	EndID:       EndID,
	// 	Bills:       Bills,
	// }

	// // Serialize to bytes using Borsh
	// // serializedData, err := borsh.Serialize(instructionData)
	// serializedData, err := instructionData.BorshEncode()
	// if err != nil {
	// 	panic(err)
	// }

	// println("serializedData", serializedData)
}
