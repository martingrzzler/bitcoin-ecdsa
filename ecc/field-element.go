package ecc

import (
	"fmt"
	"math/big"
)

type FE interface {
	String() string
	Equals(other FE) bool
	Add(other FE) FE
	Sub(other FE) FE
	Mul(other FE) FE
	Pow(exp *big.Int) FE
	Div(other FE) FE
	FieldEquals(other FE) bool
	IsZero() bool
	Prime() *big.Int
	Num() *big.Int
}

// Field Element
type fe struct {
	num   *big.Int
	prime *big.Int
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

	return fe{num: num, prime: prime}
}

func (e fe) String() string {
	return fmt.Sprintf("FieldElement_0x%x(0x%x)", e.prime, e.num)
}

func (e fe) Equals(other FE) bool {
	return e.num.Cmp(other.Num()) == 0 && e.prime.Cmp(other.Prime()) == 0
}

func (e fe) Add(other FE) FE {
	if !e.FieldEquals(other) {
		panic("Cannot add two numbers in different Field")
	}
	num := new(big.Int).Add(e.num, other.Num())
	num = num.Mod(num, e.prime)

	return fe{num: num, prime: e.prime}
}

func Add(values ...FE) FE {
	result := values[0]
	for _, e := range values[1:] {
		result = result.Add(e)
	}
	return result
}

func (e fe) Sub(other FE) FE {
	if !e.FieldEquals(other) {
		panic("Cannot subtract two numbers in different Field")
	}
	num := new(big.Int).Sub(e.num, other.Num())
	num = num.Mod(num, e.prime)
	return fe{num: num, prime: e.prime}
}

func Sub(values ...FE) FE {
	result := values[0]
	for _, e := range values[1:] {
		result = result.Sub(e)
	}
	return result
}

func (e fe) Mul(other FE) FE {
	if !e.FieldEquals(other) {
		panic("Cannot multiply two numbers in different Field")
	}
	num := new(big.Int).Mul(e.num, other.Num())
	num = num.Mod(num, e.prime)
	return fe{num: num, prime: e.prime}
}

func Mul(values ...FE) FE {
	result := values[0]
	for _, e := range values[1:] {
		result = result.Mul(e)
	}
	return result
}

func (e fe) Pow(exp *big.Int) FE {
	// Fermat's Little Thereom 1=n^(n-1) mod p; where p is prime
	// negative exponents can be made positive by a^-3 = a^-2 * a^(p-1) = a^(p-4)
	// doing this repeatedly will turn the `exp` positive
	exp = exp.Mod(exp, new(big.Int).Sub(e.prime, big.NewInt(1)))

	num := new(big.Int).Exp(e.num, exp, e.prime)

	return fe{num: num, prime: e.prime}
}

// trick is to turn division float64o exponentiation a * a^-1 = 1
// a^-1 = a^-1 * a^(p-1) mod p = a^(p-2) mod p
func (e fe) Div(other FE) FE {
	if !e.FieldEquals(other) {
		panic("Cannot divide two numbers in different Field")
	}

	fermatsInverse := new(big.Int).Exp(other.Num(), new(big.Int).Sub(other.Prime(), big.NewInt(2)), other.Prime())
	num := new(big.Int).Mul(e.num, fermatsInverse)
	num = num.Mod(num, e.prime)

	return fe{num: num, prime: e.prime}
}

func (e fe) FieldEquals(other FE) bool {
	return e.prime.Cmp(other.Prime()) == 0
}
func (e fe) IsZero() bool {
	return e.num.Cmp(big.NewInt(0)) == 0
}

func (e fe) Prime() *big.Int {
	return e.prime
}

func (e fe) Num() *big.Int {
	return e.num
}
