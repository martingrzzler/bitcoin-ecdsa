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
