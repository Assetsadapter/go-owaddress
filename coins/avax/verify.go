package avax

import (
	"github.com/Assetsadapter/go-owaddress/address"
	"github.com/Assetsadapter/go-owaddress/utils"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "avax"
)

type AddressVerify struct {
	address.AddressVerify
}

var (
	bech32Alphabet = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

	PChainPrefix      = "P"
	XChainPrefix      = "X"
	mainnetHRPPrefix  = "avax"
	testnetHRPPrefix  = "fuji"
	localnetHRPPrefix = "local"
)

func (b AddressVerify) IsValid(address string) bool {

	if address == "" {
		return false
	}

	// 检查地址属于那个链 p链以P开头， x链以X开头， c链以C开头，且p链与x链的地址编码格式一致。使用 bech32
	if !checkChainType(address[:1]) {
		return false
	}

	// 检查 hrp 并获取真是的 bech32 地址
	if !checkHrp(address) {
		return false
	}

	// bech32解码还原地址byte
	decodeBytes, err := utils.Bech32Decode(address[2:], bech32Alphabet)
	if err != nil || len(decodeBytes) != 20 {
		return false
	}

	return true
}

func checkChainType(chainType string) bool {
	return chainType == PChainPrefix || chainType == XChainPrefix
}

func checkHrp(address string) bool {
	if address[2:6] == mainnetHRPPrefix || address[2:6] == testnetHRPPrefix || address[2:7] == localnetHRPPrefix {
		return true
	}
	return false
}
