package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_dos_AddressVerify_Valid(t *testing.T) {

	coin := "dos"
	expect := true

	p2pkhAddress := "ScEcoTiRhN6pzEyPuwsTdLjLUXewiYTAzVW"
	//p2shAddress := "3BYx8ciMdywxd2bbn5h9V7EAZtzLg2RhhX"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}

	//valid, err = owaddress.Verify(coin, p2shAddress)
	//
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//if valid != expect {
	//	t.Error("Failed to verify P2SH valid address")
	//}
}

func Test_dos_AddressVerify_InValid(t *testing.T) {

	coin := "hc"
	expect := false

	p2pkhAddress := "SXbFiUns2fyVerngv4qGPyCeh9Rxz29pLX"
	//p2shAddress := "3BYx8ciMdywxd2bbn5h9V7EAZtzLg2R0hX"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}

	//valid, err = owaddress.Verify(coin, p2shAddress)
	//
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//if valid != expect {
	//	t.Error("Failed to verify P2SH invalid address")
	//}
}
