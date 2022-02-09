package ecc

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct{ x, y, want FieldElement }{
		{NewFieldElement(3, 5), NewFieldElement(4, 5), NewFieldElement(2, 5)},
		{NewFieldElement(3, 5), NewFieldElement(1, 5), NewFieldElement(4, 5)},
		{NewFieldElement(1, 11), NewFieldElement(9, 11), NewFieldElement(10, 11)},
		{NewFieldElement(5, 11), NewFieldElement(6, 11), NewFieldElement(0, 11)},
	}

	for _, test := range testCases {
		result := test.x.Add(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), &test.want)
		}
	}
}
func TestAddMulti(t *testing.T) {
	testCases := []struct{ x, y, z, want FieldElement }{
		{NewFieldElement(3, 5), NewFieldElement(4, 5), NewFieldElement(0, 5), NewFieldElement(2, 5)},
		{NewFieldElement(3, 5), NewFieldElement(1, 5), NewFieldElement(2, 5), NewFieldElement(1, 5)},
		{NewFieldElement(1, 11), NewFieldElement(9, 11), NewFieldElement(8, 11), NewFieldElement(7, 11)},
	}

	for _, test := range testCases {
		result := Add(test.x, test.y, test.z)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), &test.want)
		}
	}
}
func TestSub(t *testing.T) {
	testCases := []struct{ x, y, want FieldElement }{
		{NewFieldElement(3, 5), NewFieldElement(4, 5), NewFieldElement(4, 5)},
		{NewFieldElement(3, 5), NewFieldElement(1, 5), NewFieldElement(2, 5)},
		{NewFieldElement(1, 11), NewFieldElement(9, 11), NewFieldElement(3, 11)},
		{NewFieldElement(10, 11), NewFieldElement(6, 11), NewFieldElement(4, 11)},
	}

	for _, test := range testCases {
		result := test.x.Sub(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), &test.want)
		}
	}
}
func TestSubMulti(t *testing.T) {
	testCases := []struct{ x, y, z, want FieldElement }{
		{NewFieldElement(3, 5), NewFieldElement(4, 5), NewFieldElement(0, 5), NewFieldElement(4, 5)},
		{NewFieldElement(3, 5), NewFieldElement(1, 5), NewFieldElement(2, 5), NewFieldElement(0, 5)},
		{NewFieldElement(1, 11), NewFieldElement(9, 11), NewFieldElement(8, 11), NewFieldElement(6, 11)},
	}

	for _, test := range testCases {
		result := Sub(test.x, test.y, test.z)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), &test.want)
		}
	}
}

func TestMul(t *testing.T) {
	testCases := []struct{ x, y, want FieldElement }{
		{NewFieldElement(3, 5), NewFieldElement(4, 5), NewFieldElement(2, 5)},
		{NewFieldElement(3, 5), NewFieldElement(1, 5), NewFieldElement(3, 5)},
		{NewFieldElement(1, 11), NewFieldElement(9, 11), NewFieldElement(9, 11)},
	}

	for _, test := range testCases {
		result := test.x.Mul(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), &test.want)
		}
	}
}

func TestMulMulti(t *testing.T) {
	testCases := []struct{ x, y, z, want FieldElement }{
		{NewFieldElement(3, 5), NewFieldElement(4, 5), NewFieldElement(0, 5), NewFieldElement(0, 5)},
		{NewFieldElement(3, 5), NewFieldElement(1, 5), NewFieldElement(2, 5), NewFieldElement(1, 5)},
		{NewFieldElement(1, 11), NewFieldElement(9, 11), NewFieldElement(8, 11), NewFieldElement(6, 11)},
	}

	for _, test := range testCases {
		result := Mul(test.x, test.y, test.z)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), &test.want)
		}
	}
}

func TestPow(t *testing.T) {
	testCases := []struct {
		exp     int
		x, want FieldElement
	}{
		{3, NewFieldElement(3, 5), NewFieldElement(2, 5)},
		{6, NewFieldElement(2, 5), NewFieldElement(4, 5)},
		{7, NewFieldElement(5, 11), NewFieldElement(3, 11)},
		{-3, NewFieldElement(7, 13), NewFieldElement(8, 13)},
	}

	for _, test := range testCases {
		result := test.x.Pow(test.exp)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), &test.want)
		}
	}
}

func TestDiv(t *testing.T) {
	testCases := []struct{ x, y, want FieldElement }{
		{NewFieldElement(2, 19), NewFieldElement(7, 19), NewFieldElement(3, 19)},
		{NewFieldElement(7, 19), NewFieldElement(5, 19), NewFieldElement(9, 19)},
	}

	for _, test := range testCases {
		result := test.x.Div(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), &test.want)
		}
	}
}
