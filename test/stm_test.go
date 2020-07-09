package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_stm_AddressVerify_Valid(t *testing.T) {

	coin := "stm"
	expect := true

	address := "exchange"

	valid, err := owaddress.Verify(coin, address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}

func Test_stm_AddressVerify_InValid(t *testing.T) {

	coin := "stm"
	expect := false

	address := ""

	valid, err := owaddress.Verify(coin, address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}
}
