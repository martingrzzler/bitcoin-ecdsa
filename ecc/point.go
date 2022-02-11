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

	if math.IsInf(p.X, 1) && math.IsInf(p.Y, 1) {
		return p
	}

	if !p.OnCurve() {
		errorMsg := fmt.Sprintf("(%.2f, %.2f) is not on %s", x, y, p.Curve())
		panic(errorMsg)
	}

	return p
}

func NewInfinityPoint(a, b float64) Point {
	return Point{X: math.Inf(1), Y: math.Inf(1), A: a, B: b}
}

func (p *Point) Add(other Point) Point {
	// 1.  points are in a vertical line or using the identity point
	if !p.OnSameCurve(other) {
		errorMsg := fmt.Sprintf("(%.2f, %.2f) is not on %s", other.X, other.Y, p.Curve())
		panic(errorMsg)
	}

	if math.IsInf(p.X, 1) {
		return other
	}

	if math.IsInf(other.X, 1) {
		return *p
	}

	if p.AdditiveInverse(other) {
		return NewInfinityPoint(p.A, p.B)
	}
	// 2. the two points are the same
	if p.Equals(other) {

	}
	// 3. points are not in a vertical line, but are different
	// s = (y2 - y1)/(x2 - x1)
	// x3 = s^2 - x1 - x2
	// y3 = s(x1 - x3) - y1
	s := (other.Y - p.Y) / (other.X - p.X)
	x3 := math.Pow(s, 2) - p.X - other.X
	y3 := s*(p.X-x3) - p.Y

	// TODO complete step 2 and 3
	return Point{X: x3, Y: y3, A: p.A, B: p.B}
}

func (p *Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y && p.A == other.A && p.B == other.B
}

func (p *Point) OnCurve() bool {
	return math.Pow(p.Y, 2) == math.Pow(p.X, 3)+p.A*p.X+p.B
}

func (p *Point) OnSameCurve(other Point) bool {
	return p.A == other.A && p.B == p.B
}

func (p *Point) AdditiveInverse(other Point) bool {
	return p.X == other.X && p.Y != other.Y
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

func (p *Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}
