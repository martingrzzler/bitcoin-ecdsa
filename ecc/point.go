package ecc

import (
	"fmt"
	"math/big"
)

type Point interface {
	X() FE
	Y() FE
	A() FE
	B() FE
	Scale(coefficient *big.Int) Point
	Add(other Point) Point
	OnSameCurve(other Point) bool
	AdditiveInverse(other Point) bool
	Equals(other Point) bool
	String() string
	OnCurve() bool
}

// y^2=x^3+ax+b -> Bitcoin uses secp256k1 = y^2=x^3+7
type point struct {
	x, y, a, b FE
}

func NewPoint(x, y, a, b FE) Point {
	p := point{x: x, y: y, a: a, b: b}

	if p.x.Num().Cmp(INFINITY) == 0 && p.y.Num().Cmp(INFINITY) == 0 {
		return p
	}

	if !p.OnCurve() {
		errorMsg := fmt.Sprintf("(0x%x, 0x%x) is not on 0x%x", x.Num(), y.Num(), p.Curve())
		panic(errorMsg)
	}

	return p
}

func NewInfinityPoint(a, b FE) Point {
	return point{x: NewFE(INFINITY, a.Prime()), y: NewFE(INFINITY, a.Prime()), a: a, b: b}
}

func (p point) Scale(coefficient *big.Int) Point {
	coeff := coefficient
	var current Point = p
	result := NewInfinityPoint(p.a, p.b)

	for coeff.Cmp(big.NewInt(0)) != 0 {
		if new(big.Int).And(coeff, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 {
			result = result.Add(current)
		}
		current = current.Add(current)
		coeff.Rsh(coeff, 1)
	}
	return result
}

func (p point) Add(other Point) Point {
	// 1.  points are in a vertical line or using the identity point
	if !p.OnSameCurve(other) {
		errorMsg := fmt.Sprintf("(0x%x, 0x%x) is not on %s", other.X(), other.Y(), p.Curve())
		panic(errorMsg)
	}

	if p.x.Num().Cmp(INFINITY) == 0 {
		return other
	}

	if other.X().Num().Cmp(INFINITY) == 0 {
		return p
	}

	if p.AdditiveInverse(other) {
		return NewInfinityPoint(p.a, p.b)
	}
	// 2. the two points are the same
	if p.Equals(other) {
		// special case - tangent line is vertical
		if p.y.IsZero() {
			return NewInfinityPoint(p.a, p.b)
		}
		// s = (3x1^2 + a)/(2y1)
		// x3 = s^2 - 2x1
		// y3 = s(x1 - x3) - y1
		s := p.x.Pow(big.NewInt(2)).Mul(NewFE(big.NewInt(3), p.x.Prime())).Add(p.a).Div(NewFE(big.NewInt(2), p.x.Prime()).Mul(p.y))
		x3 := s.Pow(big.NewInt(2)).Sub(NewFE(big.NewInt(2), p.x.Prime()).Mul(p.x))
		y3 := s.Mul(p.x.Sub(x3)).Sub(p.y)
		return point{x: x3, y: y3, a: p.a, b: p.b}

	}
	// 3. points are not in a vertical line, but are different
	// s = (y2 - y1)/(x2 - x1)
	// x3 = s^2 - x1 - x2
	// y3 = s(x1 - x3) - y1
	s := other.Y().Sub(p.y).Div(other.X().Sub(p.x))
	x3 := s.Pow(big.NewInt(2)).Sub(p.x).Sub(other.X())
	y3 := s.Mul(p.x.Sub(x3)).Sub(p.y)

	return point{x: x3, y: y3, a: p.a, b: p.b}
}

func (p point) Equals(other Point) bool {
	return p.x.Equals(other.X()) && p.y.Equals(other.Y()) && p.a.Equals(other.A()) && p.b.Equals(other.B())
}

func (p point) OnCurve() bool {
	left := p.y.Pow(big.NewInt(2))
	right := Add(p.x.Pow(big.NewInt(3)), p.a.Mul(p.x), p.b)
	return left.Equals(right)
}

func (p point) OnSameCurve(other Point) bool {
	return p.a.Equals(other.A()) && p.b.Equals(other.B())
}

func (p point) AdditiveInverse(other Point) bool {
	return p.x.Equals(other.X()) && !p.y.Equals(other.Y())
}

func (p *point) Curve() string {
	b := fmt.Sprintf("+ 0x%x", p.b.Num())
	if p.b.IsZero() {
		b = ""
	}

	ax := fmt.Sprintf("+ 0x%x ", p.a.Num())
	if p.a.IsZero() {
		ax = ""
	} else if p.a.Num().Cmp(big.NewInt(1)) == 0 {
		ax = "+ x"
	}

	return fmt.Sprintf("y^2 = x^3 %s%s", ax, b)
}

func (p point) String() string {
	return fmt.Sprintf("(0x%x, 0x%x)", p.x.Num(), p.y.Num())
}

func (p point) X() FE {
	return p.x
}

func (p point) Y() FE {
	return p.y
}

func (p point) A() FE {
	return p.a
}

func (p point) B() FE {
	return p.b
}
