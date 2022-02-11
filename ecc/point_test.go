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
		// 1.  points are in a vertical line or using the identity point
		{NewPoint(-1, -1, 5, 7), NewInfinityPoint(5, 7), NewPoint(-1, -1, 5, 7)},
		{NewPoint(-1, 1, 5, 7), NewInfinityPoint(5, 7), NewPoint(-1, 1, 5, 7)},
		{NewPoint(-1, -1, 5, 7), NewPoint(-1, 1, 5, 7), NewInfinityPoint(5, 7)},
		// 2. the two points are the same
		{NewPoint(-1, -1, 5, 7), NewPoint(-1, -1, 5, 7), NewPoint(18, 77, 5, 7)},
		// 3. points are not in a vertical line, but are different
		{NewPoint(2, 5, 5, 7), NewPoint(-1, -1, 5, 7), NewPoint(3, -7, 5, 7)},
	}

	for _, test := range testCases {
		result := test.x.Add(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got for %s + %s = %s, want %s", test.x.String(), test.y.String(), result.String(), test.want.String())
		}
	}
}
