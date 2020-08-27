package test

import (
	"github.com/Assetsadapter/go-owaddress/v2"
	"testing"
)

func Test_ntn_AddressVerify_Valid(t *testing.T) {

	coin := "ntn"
	expect := true

	Address := "0x1dcbc4eac58965d9d845442df859a2f5434fec7a"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}


func Test_ntn_AddressVerify_InValid(t *testing.T) {

	coin := "ntn"
	expect := false

	Address := "0xldcbc4eac58965d9d845442df859a2f5434fec7a"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
