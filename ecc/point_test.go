package ecc

import (
	"math/big"
	"testing"
)

func TestPointIsOnCurve(t *testing.T) {
	prime := big.NewInt(223)
	a := NewFE(big.NewInt(0), prime)
	b := NewFE(big.NewInt(7), prime)
	testCases := []struct {
		x, y FE
		want bool
	}{
		{NewFE(big.NewInt(192), prime), NewFE(big.NewInt(105), prime), true},
	}

	for _, test := range testCases {
		p := Point{test.x, test.y, a, b}
		if p.OnCurve() != test.want {
			t.Errorf("got for %s %t, want %t", p.String(), p.OnCurve(), test.want)
		}
	}
}

func TestPointAdd(t *testing.T) {
	prime := big.NewInt(223)
	a := NewFE(big.NewInt(0), prime)
	b := NewFE(big.NewInt(7), prime)
	testCases := []struct {
		x, y, want Point
	}{
		{NewPoint(NewFE(big.NewInt(170), prime), NewFE(big.NewInt(142), prime), a, b), NewPoint(NewFE(big.NewInt(60), prime), NewFE(big.NewInt(139), prime), a, b), NewPoint(NewFE(big.NewInt(220), prime), NewFE(big.NewInt(181), prime), a, b)},
		{NewPoint(NewFE(big.NewInt(47), prime), NewFE(big.NewInt(71), prime), a, b), NewPoint(NewFE(big.NewInt(17), prime), NewFE(big.NewInt(56), prime), a, b), NewPoint(NewFE(big.NewInt(215), prime), NewFE(big.NewInt(68), prime), a, b)},
		{NewPoint(NewFE(big.NewInt(143), prime), NewFE(big.NewInt(98), prime), a, b), NewPoint(NewFE(big.NewInt(76), prime), NewFE(big.NewInt(66), prime), a, b), NewPoint(NewFE(big.NewInt(47), prime), NewFE(big.NewInt(71), prime), a, b)},
	}

	for _, test := range testCases {
		result := test.x.Add(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got for %s + %s = %s, want %s", test.x.String(), test.y.String(), result.String(), test.want.String())
		}
	}
}

func TestPointEquals(t *testing.T) {
	prime := big.NewInt(223)
	a := NewFE(big.NewInt(0), prime)
	b := NewFE(big.NewInt(7), prime)
	testCases := []struct {
		x, y Point
		want bool
	}{
		{NewInfinityPoint(a, b), NewInfinityPoint(a, b), true},
	}

	for _, test := range testCases {
		result := test.x.Equals(test.y)
		if !result == test.want {
			t.Errorf("got for %s + %s = %t, want %t", test.x.String(), test.y.String(), result, test.want)
		}
	}

}

func TestAddToItself(t *testing.T) {
	prime := big.NewInt(223)
	a := NewFE(big.NewInt(0), prime)
	b := NewFE(big.NewInt(7), prime)
	testCases := []struct {
		point, want Point
		scalar      int
	}{
		{NewPoint(NewFE(big.NewInt(192), prime), NewFE(big.NewInt(105), prime), a, b), NewPoint(NewFE(big.NewInt(49), prime), NewFE(big.NewInt(71), prime), a, b), 2},
		{NewPoint(NewFE(big.NewInt(143), prime), NewFE(big.NewInt(98), prime), a, b), NewPoint(NewFE(big.NewInt(64), prime), NewFE(big.NewInt(168), prime), a, b), 2},
		{NewPoint(NewFE(big.NewInt(47), prime), NewFE(big.NewInt(71), prime), a, b), NewPoint(NewFE(big.NewInt(194), prime), NewFE(big.NewInt(51), prime), a, b), 4},
		{NewPoint(NewFE(big.NewInt(47), prime), NewFE(big.NewInt(71), prime), a, b), NewPoint(NewFE(big.NewInt(116), prime), NewFE(big.NewInt(55), prime), a, b), 8},
		{NewPoint(NewFE(big.NewInt(47), prime), NewFE(big.NewInt(71), prime), a, b), NewInfinityPoint(a, b), 21},
		{NewPoint(NewFE(big.NewInt(15), prime), NewFE(big.NewInt(86), prime), a, b), NewInfinityPoint(a, b), 7},
	}

	for _, test := range testCases {
		result := test.point.Add(test.point)
		for i := 2; i < test.scalar; i++ {
			result = result.Add(test.point)
		}
		if !result.Equals(test.want) {
			t.Errorf("got for %d * %s = %s, want %s", test.scalar, test.point.String(), result.String(), test.want.String())
		}

	}

}
func TestScale(t *testing.T) {
	prime := big.NewInt(223)
	a := NewFE(big.NewInt(0), prime)
	b := NewFE(big.NewInt(7), prime)
	testCases := []struct {
		point, want Point
		coefficient *big.Int
	}{
		{NewPoint(NewFE(big.NewInt(192), prime), NewFE(big.NewInt(105), prime), a, b), NewPoint(NewFE(big.NewInt(49), prime), NewFE(big.NewInt(71), prime), a, b), big.NewInt(2)},
		{NewPoint(NewFE(big.NewInt(143), prime), NewFE(big.NewInt(98), prime), a, b), NewPoint(NewFE(big.NewInt(64), prime), NewFE(big.NewInt(168), prime), a, b), big.NewInt(2)},
		{NewPoint(NewFE(big.NewInt(47), prime), NewFE(big.NewInt(71), prime), a, b), NewPoint(NewFE(big.NewInt(194), prime), NewFE(big.NewInt(51), prime), a, b), big.NewInt(4)},
		{NewPoint(NewFE(big.NewInt(47), prime), NewFE(big.NewInt(71), prime), a, b), NewPoint(NewFE(big.NewInt(116), prime), NewFE(big.NewInt(55), prime), a, b), big.NewInt(8)},
		{NewPoint(NewFE(big.NewInt(47), prime), NewFE(big.NewInt(71), prime), a, b), NewInfinityPoint(a, b), big.NewInt(21)},
		{NewPoint(NewFE(big.NewInt(15), prime), NewFE(big.NewInt(86), prime), a, b), NewInfinityPoint(a, b), big.NewInt(7)},
	}

	for _, test := range testCases {
		result := test.point.Scale(test.coefficient)
		if !result.Equals(test.want) {
			t.Errorf("got for %d * %s = %s, want %s", test.coefficient, test.point.String(), result.String(), test.want.String())
		}

	}

}

func TestSec256k1GeneratorPoint(t *testing.T) {
	gx, ok := new(big.Int).SetString("0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 0)
	if !ok {
		t.Errorf("failed to initialize gx")
	}
	gy, ok := new(big.Int).SetString("0x483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 0)
	if !ok {
		t.Errorf("failed to initialize gy")
	}
	p, ok := new(big.Int).SetString("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f", 0)
	if !ok {
		t.Errorf("failed to initialize p")
	}

	n, ok := new(big.Int).SetString("0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 0)
	if !ok {
		t.Errorf("failed to initialize n (order)")
	}

	a := NewFE(new(big.Int), p)
	b := NewFE(big.NewInt(7), p)

	point := NewPoint(NewFE(gx, p), NewFE(gy, p), a, b)
	if !point.OnCurve() {
		t.Errorf("Generator point is not on curve")
	}

	if point.Scale(n) != NewInfinityPoint(a, b) {
		t.Errorf("Expected infinity point")
	}
}
