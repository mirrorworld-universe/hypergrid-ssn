package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	//import solana.go
	"hypergrid-aide/tools"

	// Importing the general purpose Cosmos blockchain client

	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	"gopkg.in/yaml.v3"
	// Importing the types package of your blog blockchain
)

// Default values for the global variables
var SOLANA_RPC_ENDPOINT = "http://localhost:8899"
var SOLANA_BASELAYER_RPC = "https://api.testnet.solana.com"
var SOLANA_SONIC_GRID_RPC = "http://api.grid-sonic.sonic.game"
var SOLANA_PRIVATE_KEY = "~/.config/solana/id.json"
var SOLANA_InboxProgramID = "FG8P631H9q5b53qsVM9aD71GZTWBKvujtqeWUGstpeka"
var SonicFeeProgramID = "SonicFeeSet1ement11111111111111111111111111"
var SonicFeeDataAccountID = "SonicFeeSet1ementData1111111111111111111112"
var COSMOS_RPC_ENDPOINT = "http://localhost:26657"
var COSMOS_ADDRESS_PREFIX = "cosmos"
var COSMOS_HOME = "~/.hypergrid-ssn"
var COSMOS_KEY = "my_key"
var COSMOS_GAS = uint64(100000000)

const AIDE_GET_BLOCKS_COUNT_LIMIT = uint64(200)

// read variables from yaml file
func readVariablesFromYaml(filename string) {
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
		log.Println(err)
		return
	}
	defer file.Close()

	// Read the file
	data, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	// Unmarshal the YAML
	var params map[string]interface{}
	err = yaml.Unmarshal(data, &params)
	if err != nil {
		log.Println(err)
		return
	}

	// Print the params
	log.Println(params)

	// Set the global variables
	solana_params := params["solana"].(map[string]interface{})
	SOLANA_RPC_ENDPOINT = solana_params["rpc"].(string)
	SOLANA_BASELAYER_RPC = solana_params["baselayer_rpc"].(string)
	SOLANA_PRIVATE_KEY = solana_params["private_key"].(string)
	SOLANA_InboxProgramID = solana_params["inbox_program_id"].(string)
	SOLANA_SONIC_GRID_RPC = solana_params["sonic_grid_rpc"].(string)
	SonicFeeProgramID = solana_params["fee_program_id"].(string)
	SonicFeeDataAccountID = solana_params["fee_data_account_id"].(string)

	cosmos_params := params["cosmos"].(map[string]interface{})
	COSMOS_RPC_ENDPOINT = cosmos_params["rpc"].(string)
	COSMOS_ADDRESS_PREFIX = cosmos_params["address_prefix"].(string)
	COSMOS_HOME = cosmos_params["home"].(string)
	COSMOS_KEY = cosmos_params["key"].(string)
	COSMOS_GAS = uint64(cosmos_params["gas"].(int))
	tools.COSMOS_ADDRESS_PREFIX = COSMOS_ADDRESS_PREFIX
}

func SendGridBlockFees(cosmos tools.CosmosClient, solana tools.SolanaClient, account cosmosaccount.Account, gridId string, limit uint64) {
	log.Println("SendGridBlockFees:", gridId)
	first_available_slot, err := solana.GetFirstBlock()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("first_available_slot: ", first_available_slot)

	last_sent_slot, err := tools.GetLastSentSlot()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("last_sent_slot: ", last_sent_slot)
	//choose the max of last_sent_slot and first_available_slot - 1
	start_slot := last_sent_slot + 1
	if last_sent_slot < first_available_slot {
		start_slot = first_available_slot
	}

	log.Println("start_slot: ", start_slot)
	blocks, latest_slot, err := solana.GetBlocks(start_slot, limit)
	log.Println("start_slot2: ", start_slot)
	if err != nil {
		log.Println("GetBlocks fail")
		log.Fatal(err)
	}
	log.Println("blocks: ", len(blocks))
	if len(blocks) > 0 {
		log.Println("SendGridBlockFees")
		resp, err_send := cosmos.SendGridBlockFees(account, gridId, blocks)
		if err_send != nil {
			log.Println("SendGridBlockFees fail")
			log.Fatal(err_send)
		} else {
			log.Println("SendGridBlockFees success")
			last_sent_slot = latest_slot //blocks[len(blocks)-1].Slot
			_, err = tools.SetLastSentSlot(last_sent_slot)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Print("MsgCreateGridTxFee:", resp)
	} else {
		last_sent_slot = latest_slot
		_, err = tools.SetLastSentSlot(last_sent_slot)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func SendGridInbox(solana tools.SolanaClient) {
	block, err := solana.GetLastBlock()
	if err != nil {
		log.Fatal(err)
	}

	genesis_hash, err := solana.GetGenesisHash()
	if err != nil {
		log.Fatal(err)
	}

	tools.SendTxInbox(SOLANA_PRIVATE_KEY, SOLANA_BASELAYER_RPC, SOLANA_InboxProgramID, block.Slot, block.Blockhash, genesis_hash)
}

func SyncStateAccount(cosmos tools.CosmosClient, account cosmosaccount.Account, source string, pubkey string, version string) {
	res, err := cosmos.SyncStateAccount(account, source, pubkey, version)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("SyncStateAccount:\n\n")
	log.Println(res)
}

func SettleFeeBill(cosmos tools.CosmosClient, account cosmosaccount.Account) {
	fromId := uint64(0)
	endId := uint64(0)

	res1, err1 := cosmos.QueryLastFeeSettlementBill()
	if err1 != nil {
		log.Fatal(err1)
	}
	if len(res1.FeeSettlementBill) > 0 {
		// Id = res1.FeeSettlementBill[0].Id
		fromId = res1.FeeSettlementBill[0].EndId
	}

	res2, err2 := cosmos.QueryLastGridBlockFee()
	if err2 != nil {
		log.Fatal(err1)
	}
	if len(res2.GridBlockFee) > 0 {
		// Id = res1.FeeSettlementBill[0].Id
		endId = res2.GridBlockFee[0].Id
	}

	if fromId == 0 && endId == 0 {
		log.Fatal("fromId and endId is 0")
		return
	}

	res, err := cosmos.SettleFeeBill(account, fromId, endId)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("SettleFeeBill:\n\n")
	log.Println(res)

	res3, err3 := cosmos.QueryLastFeeSettlementBill()
	if err3 != nil {
		log.Fatal(err3)
	}

	bills := res3.FeeSettlementBill[0].Bill
	//convert bills to bytes
	// billsBytes := []byte(bills)
	var Bills map[string]uint64
	json.Unmarshal([]byte(bills), &Bills)

	res4, err4 := tools.SendTxFeeSettlement(SOLANA_PRIVATE_KEY, SOLANA_SONIC_GRID_RPC, SonicFeeProgramID, SonicFeeDataAccountID, res3.FeeSettlementBill[0].FromId, res3.FeeSettlementBill[0].EndId, Bills)
	if err4 != nil {
		log.Fatal(err4)
	}

	log.Print("SendTxFeeSettlement:\n\n")
	log.Println(res4)
}

func InitFeeAccounts(cosmos tools.CosmosClient, account cosmosaccount.Account) {
	res1, err1 := cosmos.QueryAllHypergridNodes()
	if err1 != nil {
		log.Fatal(err1)
	}

	nodes := res1.HypergridNode
	has_hssn := false
	has_sonic_grid := false
	for _, node := range nodes {
		account_type := uint32(node.Role)
		if account_type == 4 {
			continue
		}
		//to make sure only one hssn account is created
		if account_type == 1 {
			if has_hssn {
				continue
			}
			has_hssn = true
		}
		//to make sure only one sonic_grid account is created
		if account_type == 2 {
			if has_sonic_grid {
				continue
			}
			has_sonic_grid = true
		}
		owner := node.DataAccount
		if owner == "" {
			owner = node.Pubkey
		}
		res4, err4 := tools.InitializeDataAccount(SOLANA_PRIVATE_KEY, SOLANA_SONIC_GRID_RPC, SonicFeeProgramID, SonicFeeDataAccountID, owner, account_type)
		if err4 != nil {
			log.Print(err4)
		}

		log.Print("InitializeDataAccount:\n\n")
		log.Println(res4)
	}
}

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//get program arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: hypergrid-aide <command>")
		os.Exit(1)
	}

	//read variables from yaml file
	readVariablesFromYaml(".hypergrid-aide.yaml")

	command := args[1]
	switch command {
	case "sync":
		if len(args) < 5 {
			fmt.Println("Usage: hypergrid-aide sync <source> <pubkey> <version>")
			os.Exit(1)
		}
		source := args[2]
		pubkey := args[3]
		version := args[4]

		//convert COSMOS_GAS to string
		// gas_str := strconv.FormatUint(COSMOS_GAS, 10)
		cosmos := tools.NewCosmosClient(
			cosmosclient.WithNodeAddress(COSMOS_RPC_ENDPOINT),
			cosmosclient.WithAddressPrefix(COSMOS_ADDRESS_PREFIX),
			cosmosclient.WithHome(COSMOS_HOME),
			cosmosclient.WithGas(strconv.FormatUint(COSMOS_GAS, 10)),
		)
		account, err := cosmos.Account(COSMOS_KEY)
		if err != nil {
			log.Fatal(err)
		}
		SyncStateAccount(*cosmos, account, source, pubkey, version)
		// break
	case "inbox":
		solana := tools.NewSolanaClient(SOLANA_RPC_ENDPOINT)
		SendGridInbox(*solana)
		// break
	case "block":
		limit := AIDE_GET_BLOCKS_COUNT_LIMIT
		if len(args) > 2 {
			//convert string to uint64
			limit_int, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			limit = limit_int
		}
		cosmos := tools.NewCosmosClient(
			cosmosclient.WithNodeAddress(COSMOS_RPC_ENDPOINT),
			cosmosclient.WithAddressPrefix(COSMOS_ADDRESS_PREFIX),
			cosmosclient.WithHome(COSMOS_HOME),
			cosmosclient.WithGas(strconv.FormatUint(COSMOS_GAS, 10)),
		)
		solana := tools.NewSolanaClient(SOLANA_RPC_ENDPOINT)
		account, err := cosmos.Account(COSMOS_KEY)
		if err != nil {
			log.Println("Account fail:", err)
			log.Fatal(err)
		}
		// resp, err := solana.GetIdentity()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// gridId := resp.Identity.String()

		gridId, err := solana.GetGenesisHash()
		if err != nil {
			log.Fatal(err)
		}
		SendGridBlockFees(*cosmos, *solana, account, gridId, limit)
		// break
	case "settle":
		cosmos := tools.NewCosmosClient(
			cosmosclient.WithNodeAddress(COSMOS_RPC_ENDPOINT),
			cosmosclient.WithAddressPrefix(COSMOS_ADDRESS_PREFIX),
			cosmosclient.WithHome(COSMOS_HOME),
			cosmosclient.WithGas(strconv.FormatUint(COSMOS_GAS, 10)),
		)
		account, err := cosmos.Account(COSMOS_KEY)
		if err != nil {
			log.Fatal(err)
		}
		SettleFeeBill(*cosmos, account)
	case "initFeeAccount":
		cosmos := tools.NewCosmosClient(
			cosmosclient.WithNodeAddress(COSMOS_RPC_ENDPOINT),
			cosmosclient.WithAddressPrefix(COSMOS_ADDRESS_PREFIX),
			cosmosclient.WithHome(COSMOS_HOME),
			cosmosclient.WithGas(strconv.FormatUint(COSMOS_GAS, 10)),
		)
		account, err := cosmos.Account(COSMOS_KEY)
		if err != nil {
			log.Fatal(err)
		}
		InitFeeAccounts(*cosmos, account)
	default:
		fmt.Println("Usage: hypergrid-aide <command>")
	}
}
