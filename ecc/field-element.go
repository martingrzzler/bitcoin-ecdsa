package ecc

import (
	"fmt"
	"math"
	"math/big"
)

type FieldElement struct {
	Num   float64
	Prime float64
}

func NewFieldElement(num float64, prime float64) FieldElement {
	if !big.NewInt(int64(prime)).ProbablyPrime(0) {
		errorMsg := fmt.Sprintf("%.2f is probably not prime", prime)
		panic(errorMsg)
	}
	if num >= prime || num < 0 {
		errorMsg := fmt.Sprintf("Num %.2f not in field range 0 to %.2f", num, prime)
		panic(errorMsg)
	}

	return FieldElement{Num: num, Prime: prime}
}

func (e *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%.2f(%.2f)", e.Prime, e.Num)
}

func (e *FieldElement) Equals(other FieldElement) bool {
	return e.Num == other.Num && e.Prime == other.Prime
}

func (e *FieldElement) Add(other FieldElement) FieldElement {
	if !e.FieldEquals(other) {
		panic("Cannot add two numbers in different Field")
	}
	num := Mod((e.Num + other.Num), e.Prime)
	return FieldElement{Num: num, Prime: e.Prime}
}

func Add(values ...FieldElement) FieldElement {
	result := values[0]
	for _, e := range values[1:] {
		result = result.Add(e)
	}
	return result
}

func (e *FieldElement) Sub(other FieldElement) FieldElement {
	if !e.FieldEquals(other) {
		panic("Cannot subtract two numbers in different Field")
	}
	num := Mod((e.Num - other.Num), e.Prime)
	return FieldElement{Num: num, Prime: e.Prime}
}

func Sub(values ...FieldElement) FieldElement {
	result := values[0]
	for _, e := range values[1:] {
		result = result.Sub(e)
	}
	return result
}

func (e *FieldElement) Mul(other FieldElement) FieldElement {
	if !e.FieldEquals(other) {
		panic("Cannot subtract two numbers in different Field")
	}
	num := Mod((e.Num * other.Num), e.Prime)
	return FieldElement{Num: num, Prime: e.Prime}
}

func Mul(values ...FieldElement) FieldElement {
	result := values[0]
	for _, e := range values[1:] {
		result = result.Mul(e)
	}
	return result
}

func (e *FieldElement) Pow(exp float64) FieldElement {
	// Fermat's Little Thereom 1=n^(n-1) mod p; where p is prime
	// negative exponents can be made positive by a^-3 = a^-2 * a^(p-1) = a^(p-4)
	// doing this repeatedly will turn the `exp` positive
	exp = Mod(exp, e.Prime-1)

	num := Mod(math.Pow(e.Num, exp), e.Prime)

	return FieldElement{Num: num, Prime: e.Prime}
}

// trick is to turn division float64o exponentiation a * a^-1 = 1
// a^-1 = a^-1 * a^(p-1) mod p = a^(p-2) mod p
func (e *FieldElement) Div(other FieldElement) FieldElement {
	if !e.FieldEquals(other) {
		panic("Cannot divide two numbers in different Field")
	}

	fermatsInverse := Mod(math.Pow(other.Num, other.Prime-2), other.Prime)

	return FieldElement{Num: Mod((e.Num * fermatsInverse), e.Prime), Prime: e.Prime}
}

func (e *FieldElement) FieldEquals(other FieldElement) bool {
	return e.Prime == other.Prime
}
