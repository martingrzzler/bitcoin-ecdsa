package ecc

import "testing"

func TestPointIsOnCurve(t *testing.T) {
	var prime int64 = 223
	a := NewFE(0, prime)
	b := NewFE(7, prime)
	testCases := []struct {
		x, y FE
		want bool
	}{
		{NewFE(192, prime), NewFE(105, prime), true},
	}

	for _, test := range testCases {
		p := Point{test.x, test.y, a, b}
		if p.OnCurve() != test.want {
			t.Errorf("got for %s %t, want %t", p.String(), p.OnCurve(), test.want)
		}
	}
}

func TestPointAdd(t *testing.T) {
	var prime int64 = 223
	a := NewFE(0, prime)
	b := NewFE(7, prime)
	testCases := []struct {
		x, y, want Point
	}{
		{NewPoint(NewFE(170, prime), NewFE(142, prime), a, b), NewPoint(NewFE(60, prime), NewFE(139, prime), a, b), NewPoint(NewFE(220, prime), NewFE(181, prime), a, b)},
		{NewPoint(NewFE(47, prime), NewFE(71, prime), a, b), NewPoint(NewFE(17, prime), NewFE(56, prime), a, b), NewPoint(NewFE(215, prime), NewFE(68, prime), a, b)},
		{NewPoint(NewFE(143, prime), NewFE(98, prime), a, b), NewPoint(NewFE(76, prime), NewFE(66, prime), a, b), NewPoint(NewFE(47, prime), NewFE(71, prime), a, b)},
	}

	for _, test := range testCases {
		result := test.x.Add(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got for %s + %s = %s, want %s", test.x.String(), test.y.String(), result.String(), test.want.String())
		}
	}
}
