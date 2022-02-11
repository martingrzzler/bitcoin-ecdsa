package ecc

import "testing"

func TestPointIsOnCurve(t *testing.T) {
	testCases := []struct {
		x, y, a, b float64
		want       bool
	}{
		{2, 4, 5, 7, false},
		{-1, -1, 5, 7, true},
		{18, 77, 5, 7, true},
		{5, 7, 5, 7, false},
	}

	for _, test := range testCases {
		p := Point{test.x, test.y, test.a, test.b}
		if p.OnCurve() != test.want {
			t.Errorf("got for %s %t, want %t", p.String(), p.OnCurve(), test.want)
		}
	}
}

func TestPointAddCase1(t *testing.T) {
	testCases := []struct {
		x, y, want Point
	}{
		{NewPoint(-1, -1, 5, 7), NewInfinityPoint(5, 7), NewPoint(-1, -1, 5, 7)},
		{NewPoint(-1, 1, 5, 7), NewInfinityPoint(5, 7), NewPoint(-1, 1, 5, 7)},
		{NewPoint(-1, -1, 5, 7), NewPoint(-1, 1, 5, 7), NewInfinityPoint(5, 7)},
	}

	for _, test := range testCases {
		result := test.x.Add(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got for %s + %s, want %s", test.x.String(), test.y.String(), test.want.String())
		}
	}
}
