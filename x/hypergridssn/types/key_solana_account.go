package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SolanaAccountKeyPrefix is the prefix to retrieve all SolanaAccount
	SolanaAccountKeyPrefix = "SolanaAccount/value/"
)

// SolanaAccountKey returns the store key to retrieve a SolanaAccount from the index fields
func SolanaAccountKey(
	address string,
	version string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	versionBytes := []byte(version)
	key = append(key, versionBytes...)
	key = append(key, []byte("/")...)

	return key
}
