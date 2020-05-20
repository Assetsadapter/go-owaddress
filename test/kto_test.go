package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_kto_AddressVerify_Valid(t *testing.T) {

	coin := "kto"
	expect := true

	Address := "KtoHnmXmKpMYPPHzJeMz7Z2WuNPqSn1ha6zvJJ2yNUKka7v"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_kto_AddressVerify_InValid(t *testing.T) {

	coin := "kto"
	expect := false

	Address := "KtoHnmXmKpMYPPHzJeMz7Z2WuNPqSn1ha6zvJJ2yNUKka7v"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
