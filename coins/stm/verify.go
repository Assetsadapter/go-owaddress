package stm

import (
	"github.com/Assetsadapter/go-owaddress/address"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "stm"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {

	if address == "" {
		return false
	}
	return true
}
