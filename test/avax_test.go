package test

import (
	"testing"

	"github.com/Assetsadapter/go-owaddress"
)

func Test_avax_AddressVerify_Valid(t *testing.T) {

	coin := "avax"
	expect := true

	Address := "X-local1y6uxaspregf36wcnz0s256n8w7csru3954qfw7"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}

func Test_avax_AddressVerify_InValid(t *testing.T) {

	coin := "avax"
	expect := false
	Address := "X-local1y6uxaspregf36wcnz0s256n8w7csru3954qfw9"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}
}
