package tools

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/text"
)

type SolanaBlock struct {
	// Blockhash string
	Blockhash string
	// Slot uint64
	Slot uint64
	//BlockTime
	BlockTime int
	//Fee uint64
	Fee uint64
}

type SolanaClient struct {
	Endpoint string
	Client   *rpc.Client
	// wsClient *ws.Client
}

func NewSolanaClient(endpoint string) *SolanaClient {
	log.Println("NewSolanaClient")
	return &SolanaClient{
		Endpoint: endpoint,
		Client:   rpc.New(endpoint),
	}
}

func (s *SolanaClient) GetBalance(address string) (*rpc.GetBalanceResult, error) {
	pubKey := solana.MustPublicKeyFromBase58(address) // serum token

	// Get the balance
	balance, err := s.Client.GetBalance(context.TODO(), pubKey, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (s *SolanaClient) GetIdentity() (*rpc.GetIdentityResult, error) {
	// Get the identity
	return s.Client.GetIdentity(context.TODO())
}

func (s *SolanaClient) GetFirstBlock() (uint64, error) {
	resp, err := s.Client.GetBlocksWithLimit(context.TODO(), 0, 1, rpc.CommitmentFinalized)
	if err != nil {
		return 0, err
	}
	if len(*resp) > 0 {
		return (*resp)[0], nil
	}
	return 0, nil
}

func (s *SolanaClient) GetLastBlock() (SolanaBlock, error) {
	resp, err := s.Client.GetBlockHeight(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		return SolanaBlock{}, err
	}
	resp2, err := s.Client.GetBlock(context.TODO(), resp)
	if err != nil {
		return SolanaBlock{}, err
	}

	block := SolanaBlock{
		Blockhash: resp2.Blockhash.String(),
		Slot:      resp,
		BlockTime: resp2.BlockTime.Time().Second(),
	}

	return block, nil

}

func (s *SolanaClient) GetBlocks(start_slot uint64, limit uint64) ([]SolanaBlock, uint64, error) {
	log.Println("GetBlocks start_slot: ", start_slot)
	resp, err := s.Client.GetBlocksWithLimit(context.TODO(), start_slot, limit, rpc.CommitmentFinalized)

	if err != nil {
		return nil, 0, err
	}

	// Get the blocks
	blocks := []SolanaBlock{}
	rewards := true
	latest_slot := uint64(0)
	for _, block := range *resp {
		log.Println("block: ", block)
		latest_slot = block
		resp2, err := s.Client.GetBlockWithOpts(context.TODO(), block, &rpc.GetBlockOpts{
			// Encoding:           solana.EncodingJSONParsed,
			Commitment:         rpc.CommitmentFinalized,
			TransactionDetails: rpc.TransactionDetailsFull,
			Rewards:            &rewards,
		})
		if err != nil {
			log.Println("error: ", err.Error())
			continue
		}

		log.Println("blockhash: ", resp2.Blockhash.String())

		Fee := uint64(0)
		voteFee := uint64(0)
		// Calculate the fee
		for _, tx := range resp2.Transactions {
			isVote := false
			for _, account := range tx.MustGetTransaction().Message.AccountKeys {
				if account.String() == "Vote111111111111111111111111111111111111111" {
					isVote = true
					break
				}
			}
			if isVote {
				voteFee += tx.Meta.Fee
			} else {
				Fee += tx.Meta.Fee
			}
		}
		log.Println("voteFee: ", voteFee)
		log.Println("Fee: ", Fee)

		// // Calculate the fee
		// Rewards := uint64(0)
		// for _, reward := range resp2.Rewards {
		// 	if reward.RewardType == "Fee" {
		// 		Rewards += uint64(reward.Lamports)
		// 	}
		// }
		if Fee > 0 {
			blocks = append(blocks, SolanaBlock{
				Blockhash: resp2.Blockhash.String(),
				Slot:      block,
				BlockTime: resp2.BlockTime.Time().Second(),
				Fee:       Fee,
			})
		}
	}

	// Get the identity
	return blocks, latest_slot, nil
}

func (s *SolanaClient) GetAccountInfo(address string) (*rpc.GetAccountInfoResult, error) {
	pubKey := solana.MustPublicKeyFromBase58(address) // serum token

	// Get the account
	return s.Client.GetAccountInfoWithOpts(
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
}

func (s *SolanaClient) RequestAirdrop(address string, amount uint64) {
	pubKey := solana.MustPublicKeyFromBase58(address) // serum token
	out, err := s.Client.RequestAirdrop(context.TODO(), pubKey, amount, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}

func (s *SolanaClient) SendTransaction(LocalPrivateKey string, programID string) (*solana.Signature, error) {
	// Load the account that you will send funds FROM:
	accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile(LocalPrivateKey)
	if err != nil {
		panic(err)
	}
	log.Println("accountFrom private key:", accountFrom)
	log.Println("accountFrom public key:", accountFrom.PublicKey())

	recent, err := s.Client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		// panic(err)
		return nil, err
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
		// panic(err)
		return nil, err
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
		log.Println(fmt.Errorf("unable to sign transaction: %w", err))
		return nil, err
	}
	spew.Dump(tx)
	// Pretty print the transaction:
	tx.EncodeTree(text.NewTreeEncoder(os.Stdout, "Transfer SOL"))

	// Send transaction:
	sig, err := s.Client.SendTransaction(context.TODO(), tx)
	if err != nil {
		// panic(err)
		return nil, err
	}
	return &sig, nil
}

// get transaction info
func (s *SolanaClient) GetTransaction(signature string) (*rpc.GetTransactionResult, error) {
	txhash, err := solana.SignatureFromBase58(signature)
	if err != nil {
		log.Println(fmt.Errorf("unable to sign transaction: %w", err))
		return nil, err
	}
	return s.Client.GetTransaction(context.TODO(), txhash, &rpc.GetTransactionOpts{
		Encoding:   solana.EncodingJSON,
		Commitment: rpc.CommitmentFinalized,
	})
}

// func main() {
// 	GetAccountInfo("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
// }
