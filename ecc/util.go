package ecc

// Modulus operations with negative numbers just return the negative number
func Mod(x, y int) int {
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
