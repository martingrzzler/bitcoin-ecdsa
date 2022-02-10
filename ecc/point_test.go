package ecc

import "testing"

func TestIsOnCurve(t *testing.T) {
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
			t.Errorf("got for (%.2f, %.2f) %t, want %t", test.x, test.y, p.OnCurve(), test.want)
		}
	}
}
