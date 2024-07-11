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

const SOLANA_RPC_ENDPOINT = "https://devnet1.sonic.game" //"http://localhost:8899" //
const COSMOS_RPC_ENDPOINT = "http://localhost:26657"
const COSMOS_ADDRESS_PREFIX = "cosmos"
const COSMOS_HOME = "/home/ubuntu/.hypergrid-ssn"

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

	data_account := "data_account" //todo: get data_account for solana L1

	cosmos.SendGridInbox(account, gridId, data_account, block)
}

func main() {
	println("Hypergrid Aide")

	cosmos := tools.NewCosmosClient(
		cosmosclient.WithNodeAddress(COSMOS_RPC_ENDPOINT),
		cosmosclient.WithAddressPrefix(COSMOS_ADDRESS_PREFIX),
		cosmosclient.WithHome(COSMOS_HOME),
		cosmosclient.WithGas("100000000"),
	)

	solana := tools.NewSolanaClient(SOLANA_RPC_ENDPOINT)

	// Account `alice` was initialized during `ignite chain serve`
	accountName := "alice" //my_validator"

	// Get account from the keyring
	account, err := cosmos.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := solana.GetIdentity()
	if err != nil {
		log.Fatal(err)
	}

	gridId := resp.Identity.String()
	println("Grid ID: ", gridId)

	SendGridBlockFees(*cosmos, *solana, account, gridId)

	SendGridInbox(*cosmos, *solana, account, gridId)

}
