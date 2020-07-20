package eos

import (
	"strings"

	"github.com/Assetsadapter/go-owaddress/address"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "eos"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {
	if len(address) == 0 {
		return false
	}
	// 判空
	isEmpty := strings.Trim(address, " \n\t") == ""

	if isEmpty {
		return false
	}
	return true
}
