package ark

import (
	"github.com/Assetsadapter/go-owaddress/address"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "ark"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {

	return true
}
