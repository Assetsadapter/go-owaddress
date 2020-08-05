package zec

import (
	"strings"

	"github.com/Assetsadapter/go-owaddress/address"
	"github.com/Assetsadapter/go-owaddress/utils"
	"github.com/blocktree/go-owcrypt"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "zec"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {
	var (
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
		bech32Alphabet = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

		MainnetP2PKHPrefix = []byte{0x1c, 0xb8}
		TestnetP2PKHPrefix = []byte{0x1d, 0x25}
		MainnetP2SHPrefix  = []byte{0x1c, 0xbd}
		TestnetP2SHPrefix  = []byte{0x1c, 0xba}
		p2pkhLen           = 26
		Bech32Prefix       = "bc"
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
	if err != nil || len(decodeBytes) != p2pkhLen {
		return false
	}

	prefix := decodeBytes[:2]
	p2pkh := !equSlice(prefix, MainnetP2PKHPrefix) && !equSlice(prefix, TestnetP2PKHPrefix)
	p2sh := !equSlice(prefix, MainnetP2SHPrefix) && !equSlice(prefix, TestnetP2SHPrefix)

	if p2pkh && p2sh {
		return false
	}

	check := owcrypt.Hash(decodeBytes[:p2pkhLen-4], 0, owcrypt.HASh_ALG_DOUBLE_SHA256)[:4]

	for i := 0; i < 4; i++ {
		if check[i] != decodeBytes[p2pkhLen-4+i] {
			return false
		}
	}

	return true
}

func equSlice(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
