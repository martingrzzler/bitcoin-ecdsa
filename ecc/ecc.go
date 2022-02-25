package ecc

import "math/big"

// initialized to 2^256
var INFINITY *big.Int
var SECP256K1Order *big.Int
var SECP256K1Prime *big.Int
var SECP256K1A FE
var SECP256K1B FE
var SECP256K1GPointX Secp256k1FE
var SECP256K1GPointY Secp256k1FE

func init() {
	if num, ok := new(big.Int).SetString("0x10000000000000000000000000000000000000000000000000000000000000000", 0); ok {
		INFINITY = num
	} else {
		panic("INFINITY creation was unsuccessful")
	}

	if num, ok := new(big.Int).SetString("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f", 0); ok {
		SECP256K1Prime = num
	} else {
		panic("SECP256K1Prime creation was unsuccessful")
	}

	if num, ok := new(big.Int).SetString("0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 0); ok {
		SECP256K1Order = num
	} else {
		panic("SECP256K1Order creation was unsuccessful")
	}

	x, ok := new(big.Int).SetString("0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 0)
	if !ok {
		panic("Genrator point x creation failed")
	}
	SECP256K1GPointX = NewSecp256k1FE(x)

	y, ok := new(big.Int).SetString("0x483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 0)
	if !ok {
		panic("Genrator point y creation failed")
	}
	SECP256K1GPointY = NewSecp256k1FE(y)

	SECP256K1A = NewFE(new(big.Int), SECP256K1Prime)
	SECP256K1B = NewFE(big.NewInt(7), SECP256K1Prime)
}
