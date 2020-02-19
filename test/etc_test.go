package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_etc_AddressVerify_Valid(t *testing.T) {

	coin := "etc"
	expect := true

	Address := "0x94fd70b3f601dbb51716f9ce351500a98372fcd4"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_etc_AddressVerify_InValid(t *testing.T) {

	coin := "etc"
	expect := false

	Address := "0x94fd70b3f601dbb51716f9ce351500a98372fcd41"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
