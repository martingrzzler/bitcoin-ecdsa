package ecc

import (
	"math/big"
)

// initialized to 2^256
var INFINITY *big.Int
var SECP256K1Order *big.Int
var SECP256K1Prime *big.Int
var SECP256K1A FE
var SECP256K1B FE

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

	SECP256K1A = NewFE(new(big.Int), SECP256K1Prime)
	SECP256K1B = NewFE(big.NewInt(7), SECP256K1Prime)
}

// Modulus operations with negative numbers just return the negative number
func Mod(x, y int64) int64 {
	return ((x % y) + y) % y
}

// ModExpWithSquaring calculates modular exponentiation with exponentiation by squaring, O(log exponent).
func ModPow(base, exponent, modulus int64) int64 {
	if modulus == 1 {
		return 0
	}
	if exponent == 0 {
		return 1
	}

	result := ModPow(base, exponent/2, modulus)
	result = (result * result) % modulus
	if exponent&1 != 0 {
		return ((base % modulus) * result) % modulus
	}
	return result % modulus
}
