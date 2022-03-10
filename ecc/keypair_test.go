package ecc

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"
)

func TestSEC(t *testing.T) {
	kp := NewKeyPair(big.NewInt(5000))

	result := kp.SEC(false)

	want, err := hex.DecodeString("04ffe558e388852f0120e46af2d1b370f85854a8eb0841811ece0e3e03d282d57c315dc72890a4f10a1481c031b03b351b0dc79901ca18a00cf009dbdb157a1d10")

	if err != nil {
		t.Fatal("Decoding Failed")
	}

	if !bytes.Equal(result, want) {
		t.Fatalf("Slices aren't equal")
	}
}

func TestParse(t *testing.T) {
	kp := NewKeyPair(big.NewInt(0xdeadbeef54321))

	sec := kp.SEC(true)

	result, err := Parse(sec)

	if err != nil {
		t.Errorf("unxepected parsing error")
	}

	if !kp.Address.Equals(result) {
		t.Fatalf("Parsing did not work")
	}

	sec = kp.SEC(true)

	sec[0] = 0x08

	result, err = Parse(sec)

	if err == nil {
		t.Errorf("should have thrown error")
	}
}
