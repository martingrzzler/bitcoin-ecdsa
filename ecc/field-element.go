package ecc

import (
	"fmt"
	"math/big"
	"strconv"
)

type FieldElement struct {
	Num   int
	Prime int
}

func NewFieldElement(num int, prime int) FieldElement {
	if !big.NewInt(int64(prime)).ProbablyPrime(0) {
		errorMsg := fmt.Sprintf("%d is probably not prime", prime)
		panic(errorMsg)
	}
	if num >= prime || num < 0 {
		errorMsg := fmt.Sprintf("Num %d not in field range 0 to %d", num, prime)
		panic(errorMsg)
	}

	return FieldElement{Num: num, Prime: prime}
}

func (e *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", e.Prime, e.Num)
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

func (e *FieldElement) Pow(exp int) FieldElement {
	// Fermat's Little Thereom 1=n^(n-1) mod p; where p is prime
	// negative exponents can be made positive by a^-3 = a^-2 * a^(p-1) = a^(p-4)
	// doing this repeatedly will turn the `exp` positive
	exp = Mod(exp, e.Prime-1)

	bigNum := new(big.Int).Exp(big.NewInt(int64(e.Num)), big.NewInt(int64(exp)), big.NewInt(int64(e.Prime)))
	num, err := strconv.Atoi(bigNum.String())
	if err != nil {
		panic("Conversion from big.Int to int failed")
	}

	return FieldElement{Num: num, Prime: e.Prime}
}

// trick is to turn division into exponentiation a * a^-1 = 1
// a^-1 = a^-1 * a^(p-1) mod p = a^(p-2) mod p
func (e *FieldElement) Div(other FieldElement) FieldElement {
	if !e.FieldEquals(other) {
		panic("Cannot subtract two numbers in different Field")
	}

	// num := (e.Num * int(math.Pow(float64(other.Num), float64(other.Prime-2)))) % e.Prime

	bigNum := new(big.Int).Exp(big.NewInt(int64(other.Num)), big.NewInt(int64(other.Prime-2)), big.NewInt(int64(other.Prime)))
	fermatsInverse, err := strconv.Atoi(bigNum.String())
	if err != nil {
		panic("Conversion from big.Int to int failed")
	}

	return FieldElement{Num: (e.Num * fermatsInverse) % e.Prime, Prime: e.Prime}
}

func (e *FieldElement) FieldEquals(other FieldElement) bool {
	return e.Prime == other.Prime
}
