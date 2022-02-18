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

func TestPointEquals(t *testing.T) {
	var prime int64 = 223
	a := NewFE(0, prime)
	b := NewFE(7, prime)
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
	var prime int64 = 223
	a := NewFE(0, prime)
	b := NewFE(7, prime)
	testCases := []struct {
		point, want Point
		scalar      int
	}{
		{NewPoint(NewFE(192, prime), NewFE(105, prime), a, b), NewPoint(NewFE(49, prime), NewFE(71, prime), a, b), 2},
		{NewPoint(NewFE(143, prime), NewFE(98, prime), a, b), NewPoint(NewFE(64, prime), NewFE(168, prime), a, b), 2},
		{NewPoint(NewFE(47, prime), NewFE(71, prime), a, b), NewPoint(NewFE(194, prime), NewFE(51, prime), a, b), 4},
		{NewPoint(NewFE(47, prime), NewFE(71, prime), a, b), NewPoint(NewFE(116, prime), NewFE(55, prime), a, b), 8},
		{NewPoint(NewFE(47, prime), NewFE(71, prime), a, b), NewInfinityPoint(a, b), 21},
		{NewPoint(NewFE(15, prime), NewFE(86, prime), a, b), NewInfinityPoint(a, b), 7},
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
	var prime int64 = 223
	a := NewFE(0, prime)
	b := NewFE(7, prime)
	testCases := []struct {
		point, want Point
		coefficient int
	}{
		{NewPoint(NewFE(192, prime), NewFE(105, prime), a, b), NewPoint(NewFE(49, prime), NewFE(71, prime), a, b), 2},
		{NewPoint(NewFE(143, prime), NewFE(98, prime), a, b), NewPoint(NewFE(64, prime), NewFE(168, prime), a, b), 2},
		{NewPoint(NewFE(47, prime), NewFE(71, prime), a, b), NewPoint(NewFE(194, prime), NewFE(51, prime), a, b), 4},
		{NewPoint(NewFE(47, prime), NewFE(71, prime), a, b), NewPoint(NewFE(116, prime), NewFE(55, prime), a, b), 8},
		{NewPoint(NewFE(47, prime), NewFE(71, prime), a, b), NewInfinityPoint(a, b), 21},
		{NewPoint(NewFE(15, prime), NewFE(86, prime), a, b), NewInfinityPoint(a, b), 7},
	}

	for _, test := range testCases {
		result := test.point.Scale(test.coefficient)
		if !result.Equals(test.want) {
			t.Errorf("got for %d * %s = %s, want %s", test.coefficient, test.point.String(), result.String(), test.want.String())
		}

	}

}
