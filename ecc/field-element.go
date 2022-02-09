package ecc

import (
	"fmt"
	"math/big"
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

func (e *FieldElement) FieldEquals(other FieldElement) bool {
	return e.Prime == other.Prime
}
