package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_ria_AddressVerify_Valid(t *testing.T) {

	coin := "ria"
	expect := true

	p2pkhAddress := "GCSJMVRD43JLFR7GXG7EESMFVB53BHGODTZ2EY7C63I7EGW3JWAG2F6L"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}

}

func Test_ria_AddressVerify_InValid(t *testing.T) {

	coin := "amtc"
	expect := false

	p2pkhAddress := "19xD3nnvEiu7Uqd8irRvF3j5ExLb4ZtSj2"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}

}
