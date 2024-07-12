package main

import (
	"fmt"
	"log"

	//import solana.go
	"hypergrid-aide/tools"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	// Importing the types package of your blog blockchain
)

const SOLANA_RPC_ENDPOINT = "http://localhost:8899" //"https://devnet1.sonic.game" //
const COSMOS_RPC_ENDPOINT = "http://172.31.10.244:26657"
const COSMOS_ADDRESS_PREFIX = "cosmos"
const COSMOS_HOME = "/home/ubuntu/.hypergrid-ssn"
const COSMOS_KEY = "my_key"

const AIDE_GET_BLOCKS_COUNT_LIMIT = 10

func SendGridBlockFees(cosmos tools.CosmosClient, solana tools.SolanaClient, account cosmosaccount.Account, gridId string) {
	first_available_slot, err := solana.GetFirstBlock()
	if err != nil {
		log.Fatal(err)
	}

	last_sent_slot, err := tools.GetLastSentSlot()
	if err != nil {
		log.Fatal(err)
	}
	start_slot := max(first_available_slot-1, last_sent_slot) + 1
	println("last_sent_slot: ", start_slot)
	blocks, err := solana.GetBlocks(start_slot, AIDE_GET_BLOCKS_COUNT_LIMIT)
	println("last_sent_slot2: ", start_slot)
	if err != nil {
		println("GetBlocks fail")
		log.Fatal(err)
	}
	println("blocks: ", len(blocks))
	if len(blocks) > 0 {
		println("SendGridBlockFees")
		resp, err_send := cosmos.SendGridBlockFees(account, gridId, blocks)
		if err_send != nil {
			log.Fatal(err_send)
			println("SendGridBlockFees fail")
		} else {
			println("SendGridBlockFees success")
			last_sent_slot = blocks[len(blocks)-1].Slot
			_, err = tools.SetLastSentSlot(last_sent_slot)
			if err != nil {
				log.Fatal(err)
			}

		}
		fmt.Print("MsgCreateGridTxFee:", resp)
	} else {
		last_sent_slot = start_slot + AIDE_GET_BLOCKS_COUNT_LIMIT - 1
		_, err = tools.SetLastSentSlot(last_sent_slot)
		if err != nil {
			log.Fatal(err)
		}

	}

	queryResp, err := cosmos.QueryAllGridBlockFees()
	if err != nil {
		log.Fatal(err)
	}

	// Print response from querying all the posts
	fmt.Print("\n\nAll grid tx fee:\n\n")
	fmt.Println(queryResp)
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
	fmt.Print("SyncStateAccount:\n\n")
	fmt.Println(res)
}

func main() {
	//get program arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: hypergrid-aide <command>")
		os.Exit(1)
	}

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
			cosmosclient.WithHome(COSMOS_HOME),
			cosmosclient.WithGas("100000000"),
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
			cosmosclient.WithHome(COSMOS_HOME),
			cosmosclient.WithGas("100000000"),
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
		cosmos := tools.NewCosmosClient(
			cosmosclient.WithNodeAddress(COSMOS_RPC_ENDPOINT),
			cosmosclient.WithAddressPrefix(COSMOS_ADDRESS_PREFIX),
			cosmosclient.WithHome(COSMOS_HOME),
			cosmosclient.WithGas("100000000"),
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
		SendGridBlockFees(*cosmos, *solana, account, gridId)
		// break
	default:
		fmt.Println("Usage: hypergrid-aide <command>")
	}

	// println("Hypergrid Aide")

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
	// println("Grid ID: ", gridId)

	// SendGridBlockFees(*cosmos, *solana, account, gridId)

	// SendGridInbox(*cosmos, *solana, account, gridId)

}
