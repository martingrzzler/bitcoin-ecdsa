package ecc

import "testing"

func TestFieldElementAdd(t *testing.T) {
	testCases := []struct{ x, y, want FE }{
		{NewFE(3, 5), NewFE(4, 5), NewFE(2, 5)},
		{NewFE(3, 5), NewFE(1, 5), NewFE(4, 5)},
		{NewFE(1, 11), NewFE(9, 11), NewFE(10, 11)},
		{NewFE(5, 11), NewFE(6, 11), NewFE(0, 11)},
	}

	for _, test := range testCases {
		result := test.x.Add(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), test.want)
		}
	}
}
func TestFieldElementAddMulti(t *testing.T) {
	testCases := []struct{ x, y, z, want FE }{
		{NewFE(3, 5), NewFE(4, 5), NewFE(0, 5), NewFE(2, 5)},
		{NewFE(3, 5), NewFE(1, 5), NewFE(2, 5), NewFE(1, 5)},
		{NewFE(1, 11), NewFE(9, 11), NewFE(8, 11), NewFE(7, 11)},
	}

	for _, test := range testCases {
		result := Add(test.x, test.y, test.z)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), test.want)
		}
	}
}
func TestFieldElementSub(t *testing.T) {
	testCases := []struct{ x, y, want FE }{
		{NewFE(3, 5), NewFE(4, 5), NewFE(4, 5)},
		{NewFE(3, 5), NewFE(1, 5), NewFE(2, 5)},
		{NewFE(1, 11), NewFE(9, 11), NewFE(3, 11)},
		{NewFE(10, 11), NewFE(6, 11), NewFE(4, 11)},
	}

	for _, test := range testCases {
		result := test.x.Sub(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), test.want)
		}
	}
}
func TestSubMulti(t *testing.T) {
	testCases := []struct{ x, y, z, want FE }{
		{NewFE(3, 5), NewFE(4, 5), NewFE(0, 5), NewFE(4, 5)},
		{NewFE(3, 5), NewFE(1, 5), NewFE(2, 5), NewFE(0, 5)},
		{NewFE(1, 11), NewFE(9, 11), NewFE(8, 11), NewFE(6, 11)},
	}

	for _, test := range testCases {
		result := Sub(test.x, test.y, test.z)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), test.want)
		}
	}
}

func TestFieldElementMul(t *testing.T) {
	testCases := []struct{ x, y, want FE }{
		{NewFE(3, 5), NewFE(4, 5), NewFE(2, 5)},
		{NewFE(3, 5), NewFE(1, 5), NewFE(3, 5)},
		{NewFE(1, 11), NewFE(9, 11), NewFE(9, 11)},
	}

	for _, test := range testCases {
		result := test.x.Mul(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), test.want)
		}
	}
}

func TestFieldElementMulMulti(t *testing.T) {
	testCases := []struct{ x, y, z, want FE }{
		{NewFE(3, 5), NewFE(4, 5), NewFE(0, 5), NewFE(0, 5)},
		{NewFE(3, 5), NewFE(1, 5), NewFE(2, 5), NewFE(1, 5)},
		{NewFE(1, 11), NewFE(9, 11), NewFE(8, 11), NewFE(6, 11)},
	}

	for _, test := range testCases {
		result := Mul(test.x, test.y, test.z)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), test.want)
		}
	}
}

func TestFieldElementPow(t *testing.T) {
	testCases := []struct {
		exp     int64
		x, want FE
	}{
		{3, NewFE(3, 5), NewFE(2, 5)},
		{6, NewFE(2, 5), NewFE(4, 5)},
		{7, NewFE(5, 11), NewFE(3, 11)},
		{-3, NewFE(7, 13), NewFE(8, 13)},
	}

	for _, test := range testCases {
		result := test.x.Pow(test.exp)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), test.want)
		}
	}
}

func TestFieldElementDiv(t *testing.T) {
	testCases := []struct{ x, y, want FE }{
		{NewFE(2, 19), NewFE(7, 19), NewFE(3, 19)},
		{NewFE(7, 19), NewFE(5, 19), NewFE(9, 19)},
	}

	for _, test := range testCases {
		result := test.x.Div(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), test.want)
		}
	}
}
