package dot

import (
	"github.com/Assetsadapter/go-owaddress/address"
	"github.com/Assetsadapter/go-owaddress/utils"
	"github.com/blocktree/go-owcrypt"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "dot"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {
	var (
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
		SSPrefix       = []byte{0x53, 0x53, 0x35, 0x38, 0x50, 0x52, 0x45}
		AddrPrefix     = byte(0x00)
	)

	if address == "" {
		return false
	}

	decodeBytes, err := utils.Base58Decode(address, base58Alphabet)
	if err != nil || len(decodeBytes) != 35 {
		return false
	}

	if decodeBytes[0] != AddrPrefix {
		return false
	}

	checkMate := append(SSPrefix, decodeBytes[:len(decodeBytes)-2]...)

	check := owcrypt.Hash(checkMate, 64, owcrypt.HASH_ALG_BLAKE2B)[:2]

	return equalsArray(check, decodeBytes[len(decodeBytes)-2:])
}

func equalsArray(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
