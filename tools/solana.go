package tools

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"

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
	Instruction uint32
	FromID      uint64
	EndID       uint64
	Bills       []SettlementBillParam
}

type InitializedParams struct {
	Instruction uint32
	Owner       solana.PublicKey
	AccountType uint32
}

// BorshEncode encodes the InstructionData using Borsh
func (d *SettleFeeBillParams) BorshEncode() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, d.Instruction)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.FromID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.EndID)
	if err != nil {
		return nil, err
	}
	billCount := uint64(len(d.Bills))
	err = binary.Write(buf, binary.LittleEndian, billCount)
	if err != nil {
		return nil, err
	}
	for _, bill := range d.Bills {
		err = binary.Write(buf, binary.LittleEndian, bill.Key[:])
		if err != nil {
			return nil, err
		}
		err = binary.Write(buf, binary.LittleEndian, bill.Amount)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

const SonicFeeProgramID = "SonicFeeSet1ement11111111111111111111111111"
const L1InboxProgramID = "5XJ1wZkTwAw9mc5FbM3eBgAT83TKgtAGzKos9wVxC6my"

func getLocalPrivateKey() (solana.PrivateKey, error) {
	//get home path "~/"
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	// panic(err)
	// 	return nil, err
	// }
	// Load the account that you will send funds FROM:
	// accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile(home + "/.config/solana/id.json")

	// Load the account that you will send funds FROM:
	accountFrom, err := solana.PrivateKeyFromBase58("5gA6JTpFziXu7py2j63arRUq1H29p6pcPMB74LaNuzcSqULPD6s1SZUS3UMPvFEE9oXmt1kk6ez3C6piTc3bwpJ6")
	if err != nil {
		// panic(err)
		return nil, err
	}
	fmt.Println("accountFrom private key:", accountFrom)
	fmt.Println("accountFrom public key:", accountFrom.PublicKey())

	return accountFrom, nil
}

func sendSonicTx(rpcUrl string, programId string, accounts solana.AccountMetaSlice, instructionData []byte, signers []solana.PrivateKey) (*solana.Signature, error) {
	// Create a new RPC client:
	rpcClient := rpc.New(rpcUrl)

	// Create a new WS client (used for confirming transactions)
	//replace http or https with ws
	rpcWsUrl := strings.Replace(rpcUrl, "http://", "ws://", 1)
	rpcWsUrl = strings.Replace(rpcWsUrl, "https://", "wss://", 1)
	rpcWsUrl = strings.Replace(rpcWsUrl, ":8899", ":8900", 1)

	wsClient, err := ws.Connect(context.Background(), rpcWsUrl)
	if err != nil {
		// panic(err)
		return nil, err
	}

	recent, err := rpcClient.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		// panic(err)
		return nil, err
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			solana.NewInstruction(
				solana.MustPublicKeyFromBase58(programId),
				accounts,
				instructionData, // data
			),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(signers[0].PublicKey()),
	)
	if err != nil {
		// panic(err)
		return nil, err
	}

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		//check key is in signers
		for _, signer := range signers {
			if key.Equals(signer.PublicKey()) {
				return &signer
			}
		}
		// if accountFrom.PublicKey().Equals(key) {
		// 	return &accountFrom
		// }
		return nil
	})
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

func SendTxFeeSettlement(rpcUrl string, data_accounts []string, FromId uint64, EndID uint64, bills map[string]uint64) (*solana.Signature, error) {
	Bills := []SettlementBillParam{}
	// convert bills to []SettlementBillParam
	for key, value := range bills {
		Bills = append(Bills, SettlementBillParam{
			Key:    solana.MustPublicKeyFromBase58(key),
			Amount: value,
		})
	}

	instructionData := SettleFeeBillParams{
		Instruction: 1,
		FromID:      FromId,
		EndID:       EndID,
		Bills:       Bills,
	}

	// Serialize to bytes using Borsh
	serializedData, err := instructionData.BorshEncode() // borsh.Serialize(instructionData)
	if err != nil {
		// panic(err)
		return nil, err
	}

	accounts := solana.AccountMetaSlice{}
	for _, data_account := range data_accounts {
		accounts = append(accounts, solana.NewAccountMeta(solana.MustPublicKeyFromBase58(data_account), true, false))
	}
	signer, err := getLocalPrivateKey()
	if err != nil {
		// panic(err)
		return nil, err
	}

	signers := []solana.PrivateKey{signer}
	return sendSonicTx(rpcUrl, SonicFeeProgramID, accounts, serializedData, signers)
}

func InitializeDataAccount(rpcUrl string, owner string, data_account string, account_type uint32) (*solana.Signature, error) {
	instructionData := InitializedParams{
		Instruction: 0,
		Owner:       solana.MustPublicKeyFromBase58(owner),
		AccountType: account_type,
	}

	// Serialize to bytes using Borsh
	serializedData, err := borsh.Serialize(instructionData)
	if err != nil {
		// panic(err)
		return nil, err
	}

	accounts := solana.AccountMetaSlice{
		solana.NewAccountMeta(solana.MustPublicKeyFromBase58(data_account), true, false),
	}

	signer, err := getLocalPrivateKey()
	if err != nil {
		// panic(err)
		return nil, err
	}
	signers := []solana.PrivateKey{signer}

	return sendSonicTx(rpcUrl, SonicFeeProgramID, accounts, serializedData, signers)
}

type InboxProgrmParams struct {
	Instruction [8]byte
	Slot        uint64
	Hash        string
}

func hashInstructionMethod(method string) [8]byte {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("global:%s", method)))
	result := hasher.Sum(nil)

	var hash [8]byte
	copy(hash[:], result[:8])
	return hash
}

func SendTxInbox(rpcUrl string, slot uint64, hash string) (*solana.Signature, *solana.PublicKey, error) {
	instructionData := InboxProgrmParams{
		Instruction: hashInstructionMethod("addblock"),
		Slot:        slot,
		Hash:        hash,
	}

	// Serialize to bytes using Borsh
	serializedData, err := borsh.Serialize(instructionData)
	if err != nil {
		// panic(err)
		return nil, nil, err
	}

	//create a new keypair
	data_account, err := solana.NewRandomPrivateKey()
	if err != nil {
		// panic(err)
		return nil, nil, err
	}
	data_key := data_account.PublicKey()
	fmt.Println("data_account:", data_key)

	signer, err := getLocalPrivateKey()
	if err != nil {
		// panic(err)
		return nil, nil, err
	}

	accounts := solana.AccountMetaSlice{
		solana.NewAccountMeta(data_account.PublicKey(), true, true),
		solana.NewAccountMeta(signer.PublicKey(), true, true),
		solana.NewAccountMeta(solana.MustPublicKeyFromBase58("11111111111111111111111111111111"), false, false),
	}

	signers := []solana.PrivateKey{signer, data_account}

	sig, err := sendSonicTx(rpcUrl, L1InboxProgramID, accounts, serializedData, signers)
	if err != nil {
		// panic(err)
		return nil, nil, err
	}
	fmt.Println("signature: ", sig)

	return sig, &data_key, nil
}

// func main() {
// 	GetAccountInfo("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
// }
