package test

import (
	"testing"

	"github.com/Assetsadapter/go-owaddress"
)

func Test_eos_AddressVerify_Valid(t *testing.T) {

	coin := "eos"
	expect := true

	Address := "xxxxwithdraw"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_eos_AddressVerify_InValid(t *testing.T) {

	coin := "eos"
	expect := false

	Address := "\n \t"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
