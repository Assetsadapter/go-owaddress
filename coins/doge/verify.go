package doge

import (
	"github.com/Assetsadapter/go-owaddress/address"
	"github.com/Assetsadapter/go-owaddress/utils"
	"github.com/blocktree/go-owcrypt"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "doge"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {
	var (
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
		bech32Alphabet = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

		// pro
		P2PKHPrefix  = byte(0x1E)
		// test public key hash
		//P2PKHPrefix  = byte(0x6F)
		// script
		P2SHPrefix   = byte(0x05)
		Bech32Prefix = "bc"
	)

	if address == "" {
		return false
	}

	if strings.Index(address, Bech32Prefix) == 0 { // bech32 address
		hash, err := utils.Bech32Decode(address, bech32Alphabet)
		if err != nil || len(hash) != 20 {
			return false
		}

		return true
	}

	decodeBytes, err := utils.Base58Decode(address, base58Alphabet)
	if err != nil || len(decodeBytes) != 25 {
		return false
	}

	if decodeBytes[0] != P2PKHPrefix && decodeBytes[0] != P2SHPrefix {
		return false
	}

	check := owcrypt.Hash(decodeBytes[:21], 0, owcrypt.HASh_ALG_DOUBLE_SHA256)[:4]

	for i := 0; i < 4; i++ {
		if check[i] != decodeBytes[21+i] {
			return false
		}
	}

	return true
}
