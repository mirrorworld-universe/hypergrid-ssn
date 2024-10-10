package main

import (
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
var SOLANA_RPC_ENDPOINT = "http://localhost:8899" //"https://devnet1.sonic.game" //
var SOLANA_PRIVATE_KEY = "~/.config/solana/id.json"
var COSMOS_RPC_ENDPOINT = "http://172.31.10.244:26657"
var COSMOS_ADDRESS_PREFIX = "cosmos"
var COSMOS_HOME = ".hypergrid-ssn"
var COSMOS_KEY = "my_key"
var COSMOS_GAS = "100000000"

const AIDE_GET_BLOCKS_COUNT_LIMIT = uint64(200)

// read variables from yaml file
func readVariablesFromYaml(filename string) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// Read the file
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Unmarshal the YAML
	var params map[string]interface{}
	err = yaml.Unmarshal(data, &params)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Print the params
	log.Println(params)

	// Set the global variables
	SOLANA_RPC_ENDPOINT = params["solana_rpc"].(string)
	SOLANA_PRIVATE_KEY = params["solana_private_key"].(string)
	COSMOS_RPC_ENDPOINT = params["cosmos_rpc"].(string)
	COSMOS_ADDRESS_PREFIX = params["cosmos_address_prefix"].(string)
	COSMOS_HOME = params["cosmos_home"].(string)
	COSMOS_KEY = params["cosmos_key"].(string)
	COSMOS_GAS = params["cosmos_gas"].(string)

	tools.COSMOS_ADDRESS_PREFIX = COSMOS_ADDRESS_PREFIX
}

func SendGridBlockFees(cosmos tools.CosmosClient, solana tools.SolanaClient, account cosmosaccount.Account, gridId string, limit uint64) {
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
			log.Fatal(err_send)
			log.Println("SendGridBlockFees fail")
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

	// queryResp, err := cosmos.QueryAllGridBlockFees()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Print response from querying all the posts
	// fmt.Print("\n\nAll grid tx fee:\n\n")
	// fmt.Println(queryResp)
}

func SendGridInbox(cosmos tools.CosmosClient, solana tools.SolanaClient, account cosmosaccount.Account, gridId string) {
	block, err := solana.GetLastBlock()
	if err != nil {
		log.Fatal(err)
	}

	cosmos.SendGridInbox(account, gridId, block)
}

func SyncStateAccount(cosmos tools.CosmosClient, account cosmosaccount.Account, source string, pubkey string) {
	res, err := cosmos.SyncStateAccount(account, source, pubkey)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("SyncStateAccount:\n\n")
	log.Println(res)
}

func main() {
	//get program arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: hypergrid-aide <command>")
		os.Exit(1)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
		// os.Exit(1)
	}

	//read variables from yaml file
	readVariablesFromYaml(home + "/.hypergrid-aide.yaml")

	command := args[1]
	switch command {
	case "sync":
		if len(args) < 5 {
			fmt.Println("Usage: hypergrid-aide sync <source> <pubkey>")
			os.Exit(1)
		}
		source := args[2]
		pubkey := args[3]
		cosmos := tools.NewCosmosClient(
			cosmosclient.WithNodeAddress(COSMOS_RPC_ENDPOINT),
			cosmosclient.WithAddressPrefix(COSMOS_ADDRESS_PREFIX),
			cosmosclient.WithHome(home+"/"+COSMOS_HOME),
			cosmosclient.WithGas(COSMOS_GAS),
		)
		account, err := cosmos.Account(COSMOS_KEY)
		if err != nil {
			log.Fatal(err)
		}
		SyncStateAccount(*cosmos, account, source, pubkey)
		// break
	case "inbox":
		cosmos := tools.NewCosmosClient(
			cosmosclient.WithNodeAddress(COSMOS_RPC_ENDPOINT),
			cosmosclient.WithAddressPrefix(COSMOS_ADDRESS_PREFIX),
			cosmosclient.WithHome(home+"/"+COSMOS_HOME),
			cosmosclient.WithGas(COSMOS_GAS),
		)
		solana := tools.NewSolanaClient(SOLANA_RPC_ENDPOINT)
		account, err := cosmos.Account(COSMOS_KEY)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := solana.GetIdentity()
		if err != nil {
			log.Fatal(err)
		}
		gridId := resp.Identity.String()
		SendGridInbox(*cosmos, *solana, account, gridId)
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
			cosmosclient.WithHome(home+"/"+COSMOS_HOME),
			cosmosclient.WithGas(COSMOS_GAS),
		)
		solana := tools.NewSolanaClient(SOLANA_RPC_ENDPOINT)
		account, err := cosmos.Account(COSMOS_KEY)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := solana.GetIdentity()
		if err != nil {
			log.Fatal(err)
		}
		gridId := resp.Identity.String()
		SendGridBlockFees(*cosmos, *solana, account, gridId, limit)
		// break
	default:
		fmt.Println("Usage: hypergrid-aide <command>")
	}

	// fmt.Println("Hypergrid Aide")

	// cosmos := tools.NewCosmosClient(
	// 	cosmosclient.WithNodeAddress(COSMOS_RPC_ENDPOINT),
	// 	cosmosclient.WithAddressPrefix(COSMOS_ADDRESS_PREFIX),
	// 	cosmosclient.WithHome(COSMOS_HOME),
	// 	cosmosclient.WithGas("100000000"),
	// )

	// solana := tools.NewSolanaClient(SOLANA_RPC_ENDPOINT)

	// // Account `alice` was initialized during `ignite chain serve`
	// accountName := COSMOS_KEY

	// // Get account from the keyring
	// account, err := cosmos.Account(accountName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// resp, err := solana.GetIdentity()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// gridId := resp.Identity.String()
	// fmt.Println("Grid ID: ", gridId)

	// SendGridBlockFees(*cosmos, *solana, account, gridId)

	// SendGridInbox(*cosmos, *solana, account, gridId)

}
