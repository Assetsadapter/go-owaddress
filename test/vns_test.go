package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_vns_AddressVerify_Valid(t *testing.T) {

	coin := "vns"
	expect := true

	Address := "0x97edc8c4e5ca09fa5094d24639e09947fe39829a"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_vns_AddressVerify_InValid(t *testing.T) {

	coin := "vns"
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
