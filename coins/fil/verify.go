package fil

import (
	"encoding/base32"
	"fmt"

	"github.com/Assetsadapter/go-owaddress/address"
	"github.com/blocktree/go-owcrypt"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "fil"
)

var (
	BlsProtocol       = byte(0x3)
	Secp256k1Protocol = byte(0x1)
	BlsHashLength     = 48
	PayloadHashLength = 20

	TestNetPrefix = "t"
	MainNetPrefix = "f"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {
	var (
		EncodeStd = "abcdefghijklmnopqrstuvwxyz234567"
	)

	if address == "" {
		return false
	}

	if len(address) <= 2 {
		return false
	}

	prefix := address[:2]
	if !checkPrefix(prefix) {
		return false
	}

	protocol := Secp256k1Protocol
	hashLength := PayloadHashLength
	if prefix[1:] == fmt.Sprintf("%d", BlsProtocol) {
		protocol = BlsProtocol
		hashLength = BlsHashLength
	}

	addressEncoding := base32.NewEncoding(EncodeStd)
	addressEnd, _ := addressEncoding.WithPadding(-1).DecodeString(address[2:])

	decodeBytes := append([]byte{protocol}, addressEnd[:hashLength]...)
	check := owcrypt.Hash(decodeBytes, 4, owcrypt.HASH_ALG_BLAKE2B)

	for i := 0; i < 4; i++ {
		if check[i] != addressEnd[hashLength+i] {
			return false
		}
	}

	return true
}

func checkPrefix(prefix string) bool {
	return len(prefix) == 2 &&
		prefix[:1] == TestNetPrefix ||
		prefix[:1] == MainNetPrefix &&
			prefix[1:] == fmt.Sprintf("%d", BlsProtocol) ||
		prefix[1:] == fmt.Sprintf("%d", Secp256k1Protocol)
}
