package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// HypergridNodeKeyPrefix is the prefix to retrieve all HypergridNode
	HypergridNodeKeyPrefix = "HypergridNode/value/"
)

// HypergridNodeKey returns the store key to retrieve a HypergridNode from the index fields
func HypergridNodeKey(
	pubkey string,
) []byte {
	var key []byte

	pubkeyBytes := []byte(pubkey)
	key = append(key, pubkeyBytes...)
	key = append(key, []byte("/")...)

	return key
}
