package main

import (
	"fmt"
	"moos/ecc"
)

func main() {
	p := ecc.NewPoint(-1, -1, 5, 7)
	fmt.Println(p.Curve())
}
