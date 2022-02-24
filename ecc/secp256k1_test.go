package ecc

import (
	"math/big"
	"testing"
)

func TestSecp256k1(t *testing.T) {
	gx, ok := new(big.Int).SetString("0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 0)
	if !ok {
		t.Errorf("failed to initialize gx")
	}
	gy, ok := new(big.Int).SetString("0x483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 0)
	if !ok {
		t.Errorf("failed to initialize gy")
	}

	n, ok := new(big.Int).SetString("0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 0)
	if !ok {
		t.Errorf("failed to initialize n (order)")
	}

	point := NewSecp256k1Point(NewSecp256k1FE(gx), NewSecp256k1FE(gy))
	if !point.OnCurve() {
		t.Errorf("Generator point is not on curve")
	}

	if point.Scale(n) != NewInfinityPoint(SECP256K1A, SECP256K1B) {
		t.Errorf("Expected infinity point")
	}
}
