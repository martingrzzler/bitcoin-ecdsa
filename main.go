package main

import (
	"fmt"
	"moos/ecc"
)

func main() {

	var gpoint ecc.Point = ecc.NewSecp256k1Point(ecc.SECP256K1GPointX, ecc.SECP256K1GPointY)

	res := gpoint.Scale(ecc.SECP256K1Order)

	fmt.Println(res.String())
}
