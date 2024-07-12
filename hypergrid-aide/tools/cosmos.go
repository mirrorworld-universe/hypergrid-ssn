package tools

import (
	"context"
	"fmt"
	"hypergridssn/x/hypergridssn/types"
	"log"
	"strconv"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	// Importing the types package of your blog blockchain
)

const COSMOS_ADDRESS_PREFIX = "cosmos"

type CosmosClient struct {
	Context context.Context
	Client  cosmosclient.Client
}

func NewCosmosClient(options ...cosmosclient.Option) *CosmosClient {
	println("NewCosmosClient")
	// Create a Cosmos client instance
	ctx := context.Background()
	cosmos, err := cosmosclient.New(ctx, options...)
	if err != nil {
		log.Fatal(err)
	}

	return &CosmosClient{
		Context: ctx,
		Client:  cosmos,
	}
}

func (c *CosmosClient) Account(name string) (cosmosaccount.Account, error) {
	return c.Client.Account(name)
}

func (c *CosmosClient) SendGridBlockFees(account cosmosaccount.Account, gridId string, blocks []SolanaBlock) (*cosmosclient.Response, error) {
	address, err := account.Address(COSMOS_ADDRESS_PREFIX)
	if err != nil {
		log.Fatal(err)
	}
	println("Account: ", address)

	items := []*types.GridBlockFeeItem{}
	for _, block := range blocks {
		item := types.GridBlockFeeItem{
			Grid:      gridId,
			Slot:      strconv.FormatUint(block.Slot, 10),
			Blockhash: block.Blockhash,
			Blocktime: int32(block.BlockTime),
			Fee:       strconv.FormatUint(block.Fee, 10),
		}
		items = append(items, &item)
	}

	txResp := cosmosclient.Response{}
	if len(items) > 0 {
		// Define a message to create a post
		msg := &types.MsgCreateGridBlockFee{
			Creator: address,
			Items:   items,
		}

		// Broadcast a transaction from account `alice` with the message
		// to create a post store response in txResp
		txResp, err := c.Client.BroadcastTx(c.Context, account, msg)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		// Print response from broadcasting a transaction
		fmt.Print("MsgCreateGridTxFee:\n\n")
		fmt.Println(txResp)
	}
	return &txResp, nil
}

func (c *CosmosClient) SendGridInbox(account cosmosaccount.Account, gridId string, block SolanaBlock) (*cosmosclient.Response, error) {
	address, err := account.Address(COSMOS_ADDRESS_PREFIX)
	if err != nil {
		log.Fatal(err)
	}
	println("Account: ", address)

	// Define a message to create a grid inbox
	msg := types.MsgCreateGridInbox{
		Creator: address,
		Grid:    gridId,
		Slot:    strconv.FormatUint(block.Slot, 10),
		Hash:    block.Blockhash,
	}

	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := c.Client.BroadcastTx(c.Context, account, &msg)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Print response from broadcasting a transaction
	fmt.Print("MsgCreateGridTxFee:\n\n")
	fmt.Println(txResp)

	return &txResp, nil
}

func (c *CosmosClient) SyncStateAccount(account cosmosaccount.Account, source string, pubkey string) (*cosmosclient.Response, error) {
	address, err := account.Address(COSMOS_ADDRESS_PREFIX)
	if err != nil {
		log.Fatal(err)
	}
	println("Account: ", address)

	// Define a message to create a grid inbox
	msg := types.MsgCreateSolanaAccount{
		Creator: address,
		Address: pubkey,
		Version: "0",
		Source:  source,
	}

	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := c.Client.BroadcastTx(c.Context, account, &msg)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Print response from broadcasting a transaction
	fmt.Print("MsgCreateSolanaAccount:\n\n")
	fmt.Println(txResp)

	return &txResp, nil
}

func (c *CosmosClient) QueryAllGridBlockFees() (*types.QueryAllGridBlockFeeResponse, error) {
	// Instantiate a query client for your `blog` blockchain
	queryClient := types.NewQueryClient(c.Client.Context())

	// Query the blockchain using the client's `PostAll` method
	// to get all posts store all posts in queryResp
	// return queryClient.GridBlockFeeAll(c.Context, &types.QueryAllGridBlockFeeRequest{})
	queryResp, err := queryClient.GridBlockFeeAll(c.Context, &types.QueryAllGridBlockFeeRequest{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Print response from querying all the posts
	fmt.Print("\n\nAll grid tx fee:\n\n")
	fmt.Println(queryResp)
	return queryResp, err
}

func (c *CosmosClient) QueryGridBlockFee(_id uint64) (*types.QueryGetGridBlockFeeResponse, error) {
	// Instantiate a query client for your `blog` blockchain
	queryClient := types.NewQueryClient(c.Client.Context())

	// Query the blockchain using the client's `PostAll` method
	// to get all posts store all posts in queryResp
	// return queryClient.GridBlockFeeAll(c.Context, &types.QueryAllGridBlockFeeRequest{})
	queryResp, err := queryClient.GridBlockFee(c.Context, &types.QueryGetGridBlockFeeRequest{Id: _id})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Print response from querying all the posts
	fmt.Print("\n\nGet grid tx fee:\n\n")
	fmt.Println(queryResp)
	return queryResp, err
}

func (c *CosmosClient) QueryAllHypergridNodes() (*types.QueryAllHypergridNodeResponse, error) {
	// Instantiate a query client for your `blog` blockchain
	queryClient := types.NewQueryClient(c.Client.Context())

	queryResp, err := queryClient.HypergridNodeAll(c.Context, &types.QueryAllHypergridNodeRequest{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Print response from querying all the posts
	fmt.Print("\n\nAll Hypergrid Nodes:\n\n")
	fmt.Println(queryResp)
	return queryResp, err
}
