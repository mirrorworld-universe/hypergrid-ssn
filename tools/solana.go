package tools

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/near/borsh-go"
	"gopkg.in/yaml.v3"
)

var SonicFeeProgramID = "SonicFeeSet1ement11111111111111111111111111"
var SonicFeeDataAccountID = "SonicFeeSet1ementData1111111111111111111112"
var L1InboxProgramID = "FG8P631H9q5b53qsVM9aD71GZTWBKvujtqeWUGstpeka"
var SonicStateOracleURL = "https://nisaba-hssn.sonia.game"
var SonicPrivateKey = "~/.config/solana/id.json"
var loaded = false

// read variables from yaml file
func ReadVariablesFromYaml(filename string) {
	if loaded {
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}

	filepath := home + "/" + filename
	//check if file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		filepath = "./" + filename
	}

	fmt.Println("Reading variables from", filepath)

	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the file
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the YAML
	var params map[string]interface{}
	err = yaml.Unmarshal(data, &params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the params
	fmt.Println(params)

	// Set the global variables
	solana_params := params["sonic"].(map[string]interface{})
	SonicFeeProgramID = solana_params["fee_program_id"].(string)
	SonicFeeDataAccountID = solana_params["fee_data_account_id"].(string)
	SonicPrivateKey = solana_params["private_key"].(string)
	SonicStateOracleURL = solana_params["state_oracle_url"].(string)
	L1InboxProgramID = solana_params["inbox_program_id"].(string)

	loaded = true
}

// Get account info from oracle
func GetAccountFromOracle(rpcUrl string, address string, version string) (*rpc.GetAccountInfoResult, error) {
	// read variables from yaml file
	ReadVariablesFromYaml(".hypergrid.yaml")

	// call http client to get account info from oracle
	client := &http.Client{}
	reqBody := fmt.Sprintf(`{"rpc": "%s", "address": "%s", "version": "%s"}`, rpcUrl, address, version)
	req, err := http.NewRequest("POST", SonicStateOracleURL, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get account info: %s", resp.Status)
	}

	var result rpc.GetAccountInfoResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

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

// BorshEncode encodes the InstructionData using Borsh
func (d *InitializedParams) BorshEncode() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, d.Instruction)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.Owner.Bytes())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.AccountType)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func getLocalPrivateKey() (solana.PrivateKey, error) {
	// Load the account that you will send funds FROM:
	accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile(SonicPrivateKey)
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

	recent, err := rpcClient.GetLatestBlockhash(context.TODO(), rpc.CommitmentFinalized)
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

func SendTxFeeSettlement(rpcUrl string /*data_accounts []string,*/, FromId uint64, EndID uint64, bills map[string]uint64) (*solana.Signature, error) {
	Bills := []SettlementBillParam{}
	// convert bills to []SettlementBillParam
	for key, value := range bills {
		Bills = append(Bills, SettlementBillParam{
			Key:    solana.MustPublicKeyFromBase58(key),
			Amount: value,
		})
	}

	//sort bills by key
	sort.Slice(Bills, func(i, j int) bool {
		return Bills[i].Key.String() < Bills[j].Key.String()
	})

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

	accounts := solana.AccountMetaSlice{
		solana.NewAccountMeta(solana.MustPublicKeyFromBase58(SonicFeeDataAccountID), true, false),
	}
	// for _, data_account := range data_accounts {
	// 	accounts = append(accounts, solana.NewAccountMeta(solana.MustPublicKeyFromBase58(data_account), true, false))
	// }
	signer, err := getLocalPrivateKey()
	if err != nil {
		// panic(err)
		return nil, err
	}

	signers := []solana.PrivateKey{signer}
	return sendSonicTx(rpcUrl, SonicFeeProgramID, accounts, serializedData, signers)
}

func InitializeDataAccount(rpcUrl string, owner string /*data_account string,*/, account_type uint32) (*solana.Signature, error) {
	instructionData := InitializedParams{
		Instruction: 0,
		Owner:       solana.MustPublicKeyFromBase58(owner),
		AccountType: account_type,
	}

	// Serialize to bytes using Borsh
	serializedData, err := instructionData.BorshEncode()
	if err != nil {
		// panic(err)
		return nil, err
	}

	accounts := solana.AccountMetaSlice{
		solana.NewAccountMeta(solana.MustPublicKeyFromBase58(SonicFeeDataAccountID), true, false),
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
