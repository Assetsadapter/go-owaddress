package kto

import (
	"github.com/Assetsadapter/go-owaddress/address"
	"github.com/Assetsadapter/go-owaddress/utils"
	"strings"
)

// for register
var (
    DefaultStruct = &AddressVerify{}
    CoinName      = "kto"
)

type AddressVerify struct {
    address.AddressVerify
}

const (
    base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
    
    P2PKHPrefix  = byte(0x00)
    P2SHPrefix   = byte(0x05)
)

func (b AddressVerify) IsValid(address string) bool {
	
	if address == "" {
		return false
	}
	
	if strings.Index(address, "Kto") != 0 {
		return false
	}
	
	decodeBytes, err := utils.Base58Decode(address[3:], base58Alphabet)
	if err != nil || len(decodeBytes) != 32 {
		return false
	}
	
	return true
}
