package ecc

// Modulus operations with negative numbers just return the negative number
func Mod(x, y int) int {
	return ((x % y) + y) % y
}
