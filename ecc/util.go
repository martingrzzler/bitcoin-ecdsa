package ecc

import "math"

var INFINITY int64 = int64(math.Inf(1))

// Modulus operations with negative numbers just return the negative number
func Mod(x, y int64) int64 {
	return ((x % y) + y) % y
}

func IntPow(x, y int) int {
	if y == 0 {
		return 1
	}

	result := x
	for i := 2; i <= y; i++ {
		result *= x
	}
	return result
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
