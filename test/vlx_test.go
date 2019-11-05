package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_vlx_AddressVerify_Valid(t *testing.T) {

	coin := "vlx"
	expect := true

	p2pkhAddress := "VLhxQ3SKdcm3xwUQyazfjmKQp3LL5PmQY5K"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}

}

func Test_vlx_AddressVerify_InValid(t *testing.T) {

	coin := "vlx"
	expect := false

	p2pkhAddress := "VLhxQ3SKdcm3xwUQyazfjmKQp3LL5PmQY5K"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}

}
