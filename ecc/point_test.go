package ecc

import "testing"

func TestPointIsOnCurve(t *testing.T) {
	var prime int64 = 223
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	testCases := []struct {
		x, y FieldElement
		want bool
	}{
		{NewFieldElement(192, prime), NewFieldElement(105, prime), true},
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
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	testCases := []struct {
		x, y, want Point
	}{
		{NewPoint(NewFieldElement(170, prime), NewFieldElement(142, prime), a, b), NewPoint(NewFieldElement(60, prime), NewFieldElement(139, prime), a, b), NewPoint(NewFieldElement(220, prime), NewFieldElement(181, prime), a, b)},
		// 1.  points are in a vertical line or using the identity point
		// {NewPoint(-1, -1, 5, 7), NewInfinityPoint(5, 7), NewPoint(-1, -1, 5, 7)},
		// {NewPoint(-1, 1, 5, 7), NewInfinityPoint(5, 7), NewPoint(-1, 1, 5, 7)},
		// {NewPoint(-1, -1, 5, 7), NewPoint(-1, 1, 5, 7), NewInfinityPoint(5, 7)},
		// // 2. the two points are the same
		// {NewPoint(-1, -1, 5, 7), NewPoint(-1, -1, 5, 7), NewPoint(18, 77, 5, 7)},
		// //    special case - tangent line is vertical; y=0
		// {NewPoint(-1, -1, 5, 7), NewPoint(-1, -1, 5, 7), NewInfinityPoint(5, 7)},
		// // 3. points are not in a vertical line, but are different
		// {NewPoint(2, 5, 5, 7), NewPoint(-1, -1, 5, 7), NewPoint(3, -7, 5, 7)},
	}

	for _, test := range testCases {
		result := test.x.Add(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got for %s + %s = %s, want %s", test.x.String(), test.y.String(), result.String(), test.want.String())
		}
	}
}
