package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_trx_AddressVerify_Valid(t *testing.T) {

	coin := "trx"
	expect := true

	Address := "TBkYoqf9mrAgfmw5Jagkvyr161Brf17CY3"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}

func Test_trx_AddressVerify_InValid(t *testing.T) {

	coin := "trx"
	expect := false

	Address := "TAJTMJuzvAqB8wmdUjRBVJW8CozfgrhpX2"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}

}
