package fm

import (
	"encoding/hex"
	"github.com/Assetsadapter/go-owaddress/address"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "fm"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {

	if address == "" {
		return false
	}

	if strings.Index(address, "FM") != 0 {
		return false
	}

	if len(address) != 42 {
		return false
	}

	_, err := hex.DecodeString(address[2:])
	if err != nil {
		return false
	}

	return true
}
