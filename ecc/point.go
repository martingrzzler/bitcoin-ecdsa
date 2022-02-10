package ecc

import (
	"fmt"
	"math"
)

// y^2=x^3+ax+b -> Bitcoin uses secp256k1 = y^2=x^3+7
type Point struct {
	X, Y, A, B float64
}

func NewPoint(x, y, a, b float64) Point {
	p := Point{X: x, Y: y, A: a, B: b}
	if !p.OnCurve() {
		errorMsg := fmt.Sprintf("(%.2f, %.2f) is not on %s", x, y, p.Curve())
		panic(errorMsg)
	}

	return p
}

func (p *Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y && p.A == other.A && p.B == other.B
}

func (p *Point) OnCurve() bool {
	return math.Pow(p.Y, 2) == math.Pow(p.X, 3)+p.A*p.X+p.B
}

func (p *Point) Curve() string {
	b := fmt.Sprintf("+ %.2f", p.B)
	if p.B == 0 {
		b = ""
	}

	ax := fmt.Sprintf("+ %.2fx ", p.A)
	if p.A == 0 {
		ax = ""
	} else if p.A == 1 {
		ax = "+ x"
	}

	return fmt.Sprintf("y^2 = x^3 %s%s", ax, b)
}
