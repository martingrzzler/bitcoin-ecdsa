package ecc

import (
	"fmt"
	"math/big"
)

// Field Element
type FE struct {
	Num   int64
	Prime int64
}

func NewFE(num int64, prime int64) FE {
	if !big.NewInt(int64(prime)).ProbablyPrime(0) {
		errorMsg := fmt.Sprintf("%d is probably not prime", prime)
		panic(errorMsg)
	}
	if (num >= prime || num < 0) && (num != INFINITY) {
		errorMsg := fmt.Sprintf("Num %d not in field range 0 to %d", num, prime)
		panic(errorMsg)
	}

	return FE{Num: num, Prime: prime}
}

func (e FE) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", e.Prime, e.Num)
}

func (e FE) Equals(other FE) bool {
	return e.Num == other.Num && e.Prime == other.Prime
}

func (e FE) Add(other FE) FE {
	if !e.FieldEquals(other) {
		panic("Cannot add two numbers in different Field")
	}
	num := Mod((e.Num + other.Num), e.Prime)
	return FE{Num: num, Prime: e.Prime}
}

func Add(values ...FE) FE {
	result := values[0]
	for _, e := range values[1:] {
		result = result.Add(e)
	}
	return result
}

func (e FE) Sub(other FE) FE {
	if !e.FieldEquals(other) {
		panic("Cannot subtract two numbers in different Field")
	}
	num := Mod((e.Num - other.Num), e.Prime)
	return FE{Num: num, Prime: e.Prime}
}

func Sub(values ...FE) FE {
	result := values[0]
	for _, e := range values[1:] {
		result = result.Sub(e)
	}
	return result
}

func (e FE) Mul(other FE) FE {
	if !e.FieldEquals(other) {
		panic("Cannot subtract two numbers in different Field")
	}
	num := Mod((e.Num * other.Num), e.Prime)
	return FE{Num: num, Prime: e.Prime}
}

func Mul(values ...FE) FE {
	result := values[0]
	for _, e := range values[1:] {
		result = result.Mul(e)
	}
	return result
}

func (e FE) Pow(exp int64) FE {
	// Fermat's Little Thereom 1=n^(n-1) mod p; where p is prime
	// negative exponents can be made positive by a^-3 = a^-2 * a^(p-1) = a^(p-4)
	// doing this repeatedly will turn the `exp` positive
	exp = Mod(exp, e.Prime-1)

	num := ModPow(e.Num, exp, e.Prime)

	return FE{Num: num, Prime: e.Prime}
}

// trick is to turn division float64o exponentiation a * a^-1 = 1
// a^-1 = a^-1 * a^(p-1) mod p = a^(p-2) mod p
func (e FE) Div(other FE) FE {
	if !e.FieldEquals(other) {
		panic("Cannot divide two numbers in different Field")
	}

	fermatsInverse := ModPow(other.Num, other.Prime-2, other.Prime)

	return FE{Num: Mod((e.Num * fermatsInverse), e.Prime), Prime: e.Prime}
}

func (e FE) FieldEquals(other FE) bool {
	return e.Prime == other.Prime
}
