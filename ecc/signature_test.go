package ecc

import (
	"math/big"
	"testing"
)

func TestVerification(t *testing.T) {
	z, ok := new(big.Int).SetString("0xbc62d4b80d9e36da29c16c5d4d9f11731f36052c72401a76c23c0fb5a9b74423", 0)
	if !ok {
		t.Errorf("Failed to create z")
	}
	r, ok := new(big.Int).SetString("0x37206a0610995c58074999cb9767b87af4c4978db68c06e8e6e81d282047a7c6", 0)
	if !ok {
		t.Errorf("Failed to create r")
	}
	s, ok := new(big.Int).SetString("0x8ca63759c1157ebeaec0d03cecca119fc9a75bf8e6d0fa65c841c8e2738cdaec", 0)
	if !ok {
		t.Errorf("Failed to create s")
	}
	px, ok := new(big.Int).SetString("0x04519fac3d910ca7e7138f7013706f619fa8f033e6ec6e09370ea38cee6a7574", 0)
	if !ok {
		t.Errorf("Failed to create px")
	}
	py, ok := new(big.Int).SetString("0x82b51eab8c27c66e26c858a079bcdf4f1ada34cec420cafc7eac1a42216fb6c4", 0)
	if !ok {
		t.Errorf("Failed to create py")
	}
	point := NewSecp256k1Point(NewSecp256k1FE(px), NewSecp256k1FE(py))
	sig := NewSignature(r, s)

	if !Verify(point, z, sig) {
		t.Errorf("Expected signature to be valid")
	}

}
