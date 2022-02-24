package ecc

import (
	"fmt"
	"math/big"
)

// Field Element
type FE struct {
	Num   *big.Int
	Prime *big.Int
}

func NewFE(num *big.Int, prime *big.Int) FE {
	if !prime.ProbablyPrime(0) {
		errorMsg := fmt.Sprintf("0x%x is probably not prime", prime)
		panic(errorMsg)
	}

	if (num.Cmp(prime) >= 0 || num.Cmp(new(big.Int)) < 0) && (num.Cmp(INFINITY) != 0) {
		errorMsg := fmt.Sprintf("Num 0x%x not in field range 0 to 0x%x", num, prime)
		panic(errorMsg)
	}

	return FE{Num: num, Prime: prime}
}

func (e FE) String() string {
	return fmt.Sprintf("FieldElement_0x%x(0x%x)", e.Prime, e.Num)
}

func (e FE) Equals(other FE) bool {
	return e.Num.Cmp(other.Num) == 0 && e.Prime.Cmp(other.Prime) == 0
}

func (e FE) Add(other FE) FE {
	if !e.FieldEquals(other) {
		panic("Cannot add two numbers in different Field")
	}
	num := new(big.Int).Add(e.Num, other.Num)
	num = num.Mod(num, e.Prime)

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
	num := new(big.Int).Sub(e.Num, other.Num)
	num = num.Mod(num, e.Prime)
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
		panic("Cannot multiply two numbers in different Field")
	}
	num := new(big.Int).Mul(e.Num, other.Num)
	num = num.Mod(num, e.Prime)
	return FE{Num: num, Prime: e.Prime}
}

func Mul(values ...FE) FE {
	result := values[0]
	for _, e := range values[1:] {
		result = result.Mul(e)
	}
	return result
}

func (e FE) Pow(exp *big.Int) FE {
	// Fermat's Little Thereom 1=n^(n-1) mod p; where p is prime
	// negative exponents can be made positive by a^-3 = a^-2 * a^(p-1) = a^(p-4)
	// doing this repeatedly will turn the `exp` positive
	exp = exp.Mod(exp, new(big.Int).Sub(e.Prime, big.NewInt(1)))

	num := new(big.Int).Exp(e.Num, exp, e.Prime)

	return FE{Num: num, Prime: e.Prime}
}

// trick is to turn division float64o exponentiation a * a^-1 = 1
// a^-1 = a^-1 * a^(p-1) mod p = a^(p-2) mod p
func (e FE) Div(other FE) FE {
	if !e.FieldEquals(other) {
		panic("Cannot divide two numbers in different Field")
	}

	fermatsInverse := new(big.Int).Exp(other.Num, new(big.Int).Sub(other.Prime, big.NewInt(2)), other.Prime)
	num := new(big.Int).Mul(e.Num, fermatsInverse)
	num = num.Mod(num, e.Prime)

	return FE{Num: num, Prime: e.Prime}
}

func (e FE) FieldEquals(other FE) bool {
	return e.Prime.Cmp(other.Prime) == 0
}
func (e FE) IsZero() bool {
	return e.Num.Cmp(big.NewInt(0)) == 0
}
