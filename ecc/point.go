package ecc

import (
	"fmt"
)

// y^2=x^3+ax+b -> Bitcoin uses secp256k1 = y^2=x^3+7
type Point struct {
	X, Y, A, B FE
}

func NewPoint(x, y, a, b FE) Point {
	p := Point{X: x, Y: y, A: a, B: b}

	if p.X.Num == INFINITY && p.Y.Num == INFINITY {
		return p
	}

	if !p.OnCurve() {
		errorMsg := fmt.Sprintf("(%d, %d) is not on %s", x.Num, y.Num, p.Curve())
		panic(errorMsg)
	}

	return p
}

func NewInfinityPoint(a, b FE) Point {
	return Point{X: NewFE(INFINITY, a.Prime), Y: NewFE(INFINITY, a.Prime), A: a, B: b}
}

func (p *Point) Scale(coefficient int) Point {
	coeff := coefficient
	current := *p
	result := NewInfinityPoint(p.A, p.B)

	for coeff != 0 {
		if coeff&1 != 0 {
			result = result.Add(current)
		}
		current = current.Add(current)
		coeff >>= 1
	}
	return result
}

func (p *Point) Add(other Point) Point {
	// 1.  points are in a vertical line or using the identity point
	if !p.OnSameCurve(other) {
		errorMsg := fmt.Sprintf("(%d, %d) is not on %s", other.X, other.Y, p.Curve())
		panic(errorMsg)
	}

	if p.X.Num == INFINITY {
		return other
	}

	if other.X.Num == INFINITY {
		return *p
	}

	if p.AdditiveInverse(other) {
		return NewInfinityPoint(p.A, p.B)
	}
	// 2. the two points are the same
	if p.Equals(other) {
		// special case - tangent line is vertical
		if p.Y.Num == 0 {
			return NewInfinityPoint(p.A, p.B)
		}
		// s = (3x1^2 + a)/(2y1)
		// x3 = s^2 - 2x1
		// y3 = s(x1 - x3) - y1
		s := p.X.Pow(2).Mul(NewFE(3, p.X.Prime)).Add(p.A).Div(NewFE(2, p.X.Prime).Mul(p.Y))
		x3 := s.Pow(2).Sub(NewFE(2, p.X.Prime).Mul(p.X))
		y3 := s.Mul(p.X.Sub(x3)).Sub(p.Y)
		return Point{X: x3, Y: y3, A: p.A, B: p.B}

	}
	// 3. points are not in a vertical line, but are different
	// s = (y2 - y1)/(x2 - x1)
	// x3 = s^2 - x1 - x2
	// y3 = s(x1 - x3) - y1
	s := other.Y.Sub(p.Y).Div(other.X.Sub(p.X))
	x3 := s.Pow(2).Sub(p.X).Sub(other.X)
	y3 := s.Mul(p.X.Sub(x3)).Sub(p.Y)

	return Point{X: x3, Y: y3, A: p.A, B: p.B}
}

func (p *Point) Equals(other Point) bool {
	return p.X.Equals(other.X) && p.Y.Equals(other.Y) && p.A.Equals(other.A) && p.B.Equals(other.B)
}

func (p *Point) OnCurve() bool {
	left := p.Y.Pow(2)
	right := Add(p.X.Pow(3), p.A.Mul(p.X), p.B)
	return left.Equals(right)
}

func (p *Point) OnSameCurve(other Point) bool {
	return p.A.Equals(other.A) && p.B.Equals(other.B)
}

func (p *Point) AdditiveInverse(other Point) bool {
	return p.X.Equals(other.X) && !p.Y.Equals(other.Y)
}

func (p *Point) Curve() string {
	b := fmt.Sprintf("+ %d", p.B.Num)
	if p.B.Num == 0 {
		b = ""
	}

	ax := fmt.Sprintf("+ %dx ", p.A.Num)
	if p.A.Num == 0 {
		ax = ""
	} else if p.A.Num == 1 {
		ax = "+ x"
	}

	return fmt.Sprintf("y^2 = x^3 %s%s", ax, b)
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X.Num, p.Y.Num)
}
