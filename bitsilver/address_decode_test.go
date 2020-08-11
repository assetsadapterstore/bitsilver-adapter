/*
 * Copyright 2018 The openwallet Authors
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

package bitsilver

import (
	"encoding/hex"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"testing"
)

func TestAddressDecoder_AddressEncode(t *testing.T) {
	p2pk, _ := hex.DecodeString("48cc2963cb56c3d8e8c82eadf51abfb2a3864c2c")
	p2pkAddr, _ := tw.DecoderV2.AddressEncode(p2pk)
	t.Logf("p2pkAddr: %s", p2pkAddr)

	p2sh, _ := hex.DecodeString("5c0a2143624c9f9cf410976ee24f608e916d69c9")
	p2shAddr, _ := tw.DecoderV2.AddressEncode(p2sh, addressEncoder.BTC_mainnetAddressP2SH)
	t.Logf("p2shAddr: %s", p2shAddr)
}

func TestAddressDecoder_AddressDecode(t *testing.T) {

	p2pkAddr := "17dvEDQfwkGjwLdFfDx8zCwUg5fAGz63En"
	p2pkHash, _ := tw.DecoderV2.AddressDecode(p2pkAddr)
	t.Logf("p2pkHash: %s", hex.EncodeToString(p2pkHash))

	p2shAddr := "3A5gE2q2ziDDLVs6UkBY2naSn2e9DhykBt"

	p2shHash, _ := tw.DecoderV2.AddressDecode(p2shAddr, addressEncoder.BTC_mainnetAddressP2SH)
	t.Logf("p2shHash: %s", hex.EncodeToString(p2shHash))
}

func TestAddressDecoder_ScriptPubKeyToBech32Address(t *testing.T) {

	scriptPubKey, _ := hex.DecodeString("002079db247b3da5d5e33e036005911b9341a8d136768a001e9f7b86c5211315e3e1")

	addr, err := tw.Decoder.ScriptPubKeyToBech32Address(scriptPubKey)
	if err != nil {
		t.Errorf("ScriptPubKeyToBech32Address failed unexpected error: %v\n", err)
		return
	}
	t.Logf("addr: %s", addr)


	t.Logf("addr: %s", addr)
}