package test

import (
	"testing"

	"github.com/Assetsadapter/go-owaddress"
)

func Test_fil_AddressVerify_Valid(t *testing.T) {

	coin := "fil"
	expect := true

	secp256k1Address := "t1n7rgcphp5dtjgt45vkjh2pf3ihembcvft7kf2ua"
	blsAddress := "f3w5lcylsgdoj5jtoqokoviyu7ud7ujte3dlvqok65fcbtnjgriczex7bauiyfwinudo3ggcqtbmvzrp7nmtmq"

	valid, err := owaddress.Verify(coin, secp256k1Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify secp256k1 valid address")
	}

	valid, err = owaddress.Verify(coin, blsAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify bls valid address")
	}

}

func Test_fil_AddressVerify_InValid(t *testing.T) {

	coin := "fil"
	expect := false

	secp256k1Address := "t1n7rgcphp5dtjgt45vkjh2pf3ihembcvft7kf2u2"
	blsAddress := "f3w5lcylsgdoj5jtoqokoviyu7ud7ujte3dlvqok65fcbtnjgriczex7bauiyfwinudo3ggcqtbmvzrp7nmtma"

	valid, err := owaddress.Verify(coin, secp256k1Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify secp256k1 invalid address")
	}

	valid, err = owaddress.Verify(coin, blsAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify bls invalid address")
	}

}
