package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

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
const AIDE_LAST_SEND_SLOT = "/home/ubuntu/.hypergrid-ssn/last_slot.txt"
const AIDE_GET_BLOCKS_COUNT_LIMIT = 10

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
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
	var last_sent_slot uint64
	if CheckFileExist(AIDE_LAST_SEND_SLOT) {
		var f *os.File
		f, err = os.Open(AIDE_LAST_SEND_SLOT)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		fd, err := io.ReadAll(f)
		if err != nil {
			fmt.Println("read to fd fail", err)

		}
		last_sent_slot, err = strconv.ParseUint(string(fd), 10, 64) // 将fd从[]byte转换为string，然后转换为int
		if err != nil {
			fmt.Println("convert fd to int fail", err)

		}
	} else {
		last_sent_slot = 0
	}
	println("last_sent_slot: ", last_sent_slot)
	blocks, err := solana.GetBlocks(last_sent_slot+1, AIDE_GET_BLOCKS_COUNT_LIMIT)
	println("last_sent_slot2: ", last_sent_slot)
	if err != nil {
		println("GetBlocks fail")
		log.Fatal(err)
	}
	println("blocks: ", len(blocks))
	if len(blocks) > 0 {
		println("SendGridBlockFees")
		resp, err := cosmos.SendGridBlockFees(account, gridId, blocks)
		if err != nil {
			log.Fatal(err)
			println("SendGridBlockFees fail")
		} else {
			println("SendGridBlockFees success")
			last_sent_slot = blocks[len(blocks)-1].Slot
		}
		fmt.Print("MsgCreateGridTxFee:", resp)
	} else {
		last_sent_slot = last_sent_slot + AIDE_GET_BLOCKS_COUNT_LIMIT
	}

	var f *os.File
	if CheckFileExist(AIDE_LAST_SEND_SLOT) {
		f, err = os.OpenFile(AIDE_LAST_SEND_SLOT, os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f, err = os.Create(AIDE_LAST_SEND_SLOT)
		if err != nil {
			log.Fatal(err)
		}
	}
	println("last slot: ", last_sent_slot)
	_, err = f.WriteString(strconv.FormatUint(last_sent_slot, 10))
	defer f.Close()

	queryResp, err := cosmos.QueryAllGridBlockFees()
	if err != nil {
		log.Fatal(err)
	}

	// Print response from querying all the posts
	fmt.Print("\n\nAll grid tx fee:\n\n")
	fmt.Println(queryResp)
}
