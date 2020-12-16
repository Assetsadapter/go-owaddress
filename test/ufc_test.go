package test

import (
	"testing"

	"github.com/Assetsadapter/go-owaddress"
)

func Test_ufc_AddressVerify_Valid(t *testing.T) {

	coin := "ufc"
	expect := true

	p2pkhAddress := "UFCNThZ6PykAWMXdMWxK3zwYvxq1GwkZo38Bb"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify XWC valid address")
	}

}

func Test_ufc_AddressVerify_InValid(t *testing.T) {

	coin := "ufc"
	expect := false

	p2pkhAddress := "znkvsEfqqiJ7r9MPiUnoH4bUdkBKDAGx3mm"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify XWC valid address")
	}

}
