package xlm

import (
	"encoding/base32"
	"github.com/Assetsadapter/go-owaddress/address"
	"github.com/Assetsadapter/go-owaddress/utils"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "xlm"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {
	var (
		P2PKHPrefix = byte(0x30)
	)

	if address == "" {
		return false
	}

	decodeBytes, err := base32.StdEncoding.DecodeString(address)
	if err != nil || len(decodeBytes) != 35 {
		return false
	}

	if decodeBytes[0] != P2PKHPrefix && decodeBytes[0] != P2PKHPrefix {
		return false
	}

	//check := owcrypt.Hash(decodeBytes[:21], 0, owcrypt.HASh_ALG_DOUBLE_SHA256)[:4]

	check := utils.Crc16Checksum(decodeBytes[:33])

	for i := 0; i < 2; i++ {
		if check[i] != decodeBytes[33+i] {
			return false
		}
	}

	return true
}
