package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_etx_AddressVerify_Valid(t *testing.T) {

	coin := "etx"
	expect := true

	Address := "0xe124889dda32b7f8e26376250325ea8a3abfbe96"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_etx_AddressVerify_InValid(t *testing.T) {

	coin := "etx"
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
