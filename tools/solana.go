package tools

import (
	"context"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/near/borsh-go"
)

func GetAccountInfo(rpcUrl string, address string) (*rpc.GetAccountInfoResult, error) {
	// endpoint := rpc.DevNet_RPC //MainNetBeta_RPC
	client := rpc.New(rpcUrl)
	pubKey := solana.MustPublicKeyFromBase58(address) // serum token

	// Get the account
	return client.GetAccountInfoWithOpts(
		context.TODO(),
		pubKey,
		// You can specify more options here:
		&rpc.GetAccountInfoOpts{
			Encoding:   solana.EncodingBase64Zstd,
			Commitment: rpc.CommitmentFinalized,
			// You can get just a part of the account data by specify a DataSlice:
			// DataSlice: &rpc.DataSlice{
			//  Offset: pointer.ToUint64(0),
			//  Length: pointer.ToUint64(1024),
			// },
		},
	)

	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(resp)
	// }
	// // spew.Dump(resp)
	// //convert to json
	// jsonBytes, err := json.Marshal(resp)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(jsonBytes))

	// return jsonBytes, err

}

func RequestAirdrop(rpcUrl string, address string, amount uint64) {
	// endpoint := rpc.DevNet_RPC //MainNetBeta_RPC
	client := rpc.New(rpcUrl)
	pubKey := solana.MustPublicKeyFromBase58(address) // serum token
	out, err := client.RequestAirdrop(context.TODO(), pubKey, amount, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}

type SettlementBillParam struct {
	Key    solana.PublicKey
	Amount uint64
}

type SettleFeeBillParams struct {
	FromID uint64
	EndID  uint64
	Bills  []SettlementBillParam
}

const ProgramID = "SonicFeeSet11111111111111111111111111111111"

func SendTxFeeSettlement(rpcUrl string, data_accounts []string, FromId uint64, EndID uint64, bills map[string]uint64) (*solana.Signature, error) {
	// Create a new RPC client:
	rpcClient := rpc.New(rpcUrl)

	// Create a new WS client (used for confirming transactions)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		// panic(err)
		return nil, err
	}

	//get home path "~/"
	home, err := os.UserHomeDir()
	if err != nil {
		// panic(err)
		return nil, err
	}

	// Load the account that you will send funds FROM:
	accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile(home + "/.config/solana/id.json")
	if err != nil {
		// panic(err)
		return nil, err
	}
	fmt.Println("accountFrom private key:", accountFrom)
	fmt.Println("accountFrom public key:", accountFrom.PublicKey())

	recent, err := rpcClient.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		// panic(err)
		return nil, err
	}

	Bills := []SettlementBillParam{}
	// convert bills to []SettlementBillParam
	for key, value := range bills {
		Bills = append(Bills, SettlementBillParam{
			Key:    solana.MustPublicKeyFromBase58(key),
			Amount: value,
		})
	}

	instructionData := SettleFeeBillParams{
		FromID: FromId,
		EndID:  EndID,
		Bills:  Bills,
	}

	// Serialize to bytes using Borsh
	serializedData, err := borsh.Serialize(instructionData)
	if err != nil {
		// panic(err)
		return nil, err
	}

	accounts := solana.AccountMetaSlice{}
	for _, data_account := range data_accounts {
		accounts = append(accounts, solana.NewAccountMeta(solana.MustPublicKeyFromBase58(data_account), true, false))
	}
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			solana.NewInstruction(
				solana.MustPublicKeyFromBase58(ProgramID),
				accounts,
				serializedData, // data
			),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(accountFrom.PublicKey()),
	)
	if err != nil {
		// panic(err)
		return nil, err
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if accountFrom.PublicKey().Equals(key) {
				return &accountFrom
			}
			return nil
		},
	)
	if err != nil {
		// panic(fmt.Errorf("unable to sign transaction: %w", err))
		return nil, err
	}
	spew.Dump(tx)

	// Send transaction, and wait for confirmation:
	sig, err := confirm.SendAndConfirmTransaction(
		context.TODO(),
		rpcClient,
		wsClient,
		tx,
	)
	if err != nil {
		// panic(err)
		return nil, err
	}
	spew.Dump(sig)
	return &sig, nil
}

// func main() {
// 	GetAccountInfo("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
// }
