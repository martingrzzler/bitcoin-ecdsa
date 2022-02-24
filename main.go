package main

import (
	"math/big"
	"moos/ecc"
)

func main() {
	gx, ok := new(big.Int).SetString("0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 0)
	if !ok {
		panic("not ok")
	}
	gy, ok := new(big.Int).SetString("0x483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 0)
	if !ok {
		panic("not ok")
	}
	p, ok := new(big.Int).SetString("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f", 0)
	if !ok {
		panic("not ok")
	}

	point := ecc.NewPoint(ecc.NewFE(gx, p), ecc.NewFE(gy, p), ecc.NewFE(new(big.Int), p), ecc.NewFE(big.NewInt(7), p))
	if !point.OnCurve() {
		panic("not on curve")
	}
}
