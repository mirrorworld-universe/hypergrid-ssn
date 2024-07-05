package main

import (
	"fmt"
	"log"

	//import solana.go
	"hypergrid-aide/tools"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	// Importing the types package of your blog blockchain
)

const SOLANA_RPC_ENDPOINT = "https://devnet1.sonic.game" //"http://localhost:8899" //
const COSMOS_RPC_ENDPOINT = "http://localhost:26657"
const COSMOS_ADDRESS_PREFIX = "cosmos"
const COSMOS_HOME = "/home/ubuntu/.hypergrid-ssn"

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

	blocks, err := solana.GetBlocks(7169128, 10)
	if err != nil {
		log.Fatal(err)
	}

	// Print response from querying all the posts
	fmt.Print("\n\nAll blocks:\n\n")
	fmt.Println(blocks)
	println("Blocks: ", blocks)

	address, err := account.Address(COSMOS_ADDRESS_PREFIX)
	if err != nil {
		log.Fatal(err)
	}
	println("Account: ", address)

	if len(blocks) > 0 {

		resp, err := cosmos.SendGridBlockFees(account, gridId, blocks)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("MsgCreateGridTxFee:", resp)
	}

	queryResp, err := cosmos.QueryAllGridBlockFees()
	if err != nil {
		log.Fatal(err)
	}

	// Print response from querying all the posts
	fmt.Print("\n\nAll grid tx fee:\n\n")
	fmt.Println(queryResp)
}
