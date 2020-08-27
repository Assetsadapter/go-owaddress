package eth

import (
	"encoding/hex"
	"github.com/Assetsadapter/go-owaddress/v2/address"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "eth"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {

	if address == "" {
		return false
	}

	if strings.Index(address, "0x") != 0 {
		return false
	}

	_, err := hex.DecodeString(address[2:])
	if err != nil {
		return false
	}

	return true
}

