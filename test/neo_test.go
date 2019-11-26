package test

import (
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_neo_AddressVerify_Valid(t *testing.T) {
	coin := "neo"

	accountAddress := "AM6zzjGeU6nENBushUN5Rw46BkkKuCDyUn"

	valid, err := owaddress.Verify(coin, accountAddress)
	if err != nil {
		t.Error(err.Error())
	}

	if !valid {
		t.Error("Failed to verify account valid address")
	}
}
