package ecc

import "math/big"

type Secp256k1FE struct {
	FE
}

type Secp256k1Point struct {
	Point
}

func NewSecp256k1FE(num *big.Int) Secp256k1FE {
	return Secp256k1FE{NewFE(num, SECP256K1Prime)}
}

func NewSecp256k1Point(x, y Secp256k1FE) Secp256k1Point {
	if !x.FieldEquals(y) {
		panic("Invalid x or y for Secp256k1 point")
	}
	return Secp256k1Point{NewPoint(x, y, SECP256K1A, SECP256K1B)}
}

func (p Secp256k1Point) Scale(coefficient *big.Int) Point {
	coefficient = coefficient.Mod(coefficient, SECP256K1Order)
	coeff := coefficient
	var current Point = p
	result := NewInfinityPoint(p.A(), p.B())

	for coeff.Cmp(big.NewInt(0)) != 0 {
		if new(big.Int).And(coeff, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 {
			result = result.Add(current)
		}
		current = current.Add(current)
		coeff.Rsh(coeff, 1)
	}
	return result
}
