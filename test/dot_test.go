/*
 * Copyright 2019 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package test

import (
	"fmt"
	"github.com/Assetsadapter/go-owaddress"
	"testing"
)

func Test_ksm_ed25519_AddressVerify_Valid(t *testing.T) {
	coin := "ksm"

	addressArr := make([]string, 0)
	addressArr = append(addressArr, "12tFo5exojfy4e9GLzzRQRP7Bdv4Zt8ZRL82rXCWW4Jq2Bub")	//正确

	for i := 0; i < len(addressArr); i++ {
		address := addressArr[i]
		valid, err := owaddress.Verify(coin, address)

		if err != nil {
			t.Error(err)
		}

		fmt.Println(address, " isvalid : ", valid)
	}
}

func Test_ksm_sr25519_AddressVerify_Valid(t *testing.T) {
	coin := "ksm"

	addressArr := make([]string, 0)
	addressArr = append(addressArr, "12tFo5exojfy4e9GLzzRQRP7Bdv4Zt8ZRL82rXCWW4Jq2Bub")	//正确

	for i := 0; i < len(addressArr); i++ {
		address := addressArr[i]
		valid, err := owaddress.Verify(coin, address)

		if err != nil {
			t.Error(err)
		}

		fmt.Println(address, " isvalid : ", valid)
	}
}
