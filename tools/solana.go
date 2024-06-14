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
	"github.com/gagliardetto/solana-go/text"
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

func SendTransaction(rpcUrl string, programID string) {
	// Create a new RPC client:
	rpcClient := rpc.New(rpcUrl)

	// Create a new WS client (used for confirming transactions)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		panic(err)
	}

	// Load the account that you will send funds FROM:
	accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile("/path/to/.config/solana/id.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("accountFrom private key:", accountFrom)
	fmt.Println("accountFrom public key:", accountFrom.PublicKey())

	recent, err := rpcClient.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			solana.NewInstruction(
				solana.MustPublicKeyFromBase58(programID),
				solana.AccountMetaSlice{
					solana.NewAccountMeta(accountFrom.PublicKey(), false, true),
				},
				[]byte{1, 2, 3, 4}, // data
			),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(accountFrom.PublicKey()),
	)
	if err != nil {
		panic(err)
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
		panic(fmt.Errorf("unable to sign transaction: %w", err))
	}
	spew.Dump(tx)
	// Pretty print the transaction:
	tx.EncodeTree(text.NewTreeEncoder(os.Stdout, "Transfer SOL"))

	// Send transaction, and wait for confirmation:
	sig, err := confirm.SendAndConfirmTransaction(
		context.TODO(),
		rpcClient,
		wsClient,
		tx,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(sig)
}

// func main() {
// 	GetAccountInfo("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
// }
