package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_etz_AddressVerify_Valid(t *testing.T) {

	coin := "etz"
	expect := true

	Address := "0x9cf40caae2dbf43f201693ebe09eb8fc7496965b"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_etz_AddressVerify_InValid(t *testing.T) {

	coin := "etz"
	expect := false

	Address := "0x9cf40caae2dbf43f201693ebe09eb8fc7496965b"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
