package ecc

import (
	"math/big"
	"testing"
)

func TestFieldElementAdd(t *testing.T) {
	testCases := []struct{ x, y, want FE }{
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(4), big.NewInt(5)), NewFE(big.NewInt(2), big.NewInt(5))},
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(1), big.NewInt(5)), NewFE(big.NewInt(4), big.NewInt(5))},
		{NewFE(big.NewInt(1), big.NewInt(11)), NewFE(big.NewInt(9), big.NewInt(11)), NewFE(big.NewInt(10), big.NewInt(11))},
		{NewFE(big.NewInt(5), big.NewInt(11)), NewFE(big.NewInt(6), big.NewInt(11)), NewFE(big.NewInt(0), big.NewInt(11))},
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
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(4), big.NewInt(5)), NewFE(big.NewInt(0), big.NewInt(5)), NewFE(big.NewInt(2), big.NewInt(5))},
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(1), big.NewInt(5)), NewFE(big.NewInt(2), big.NewInt(5)), NewFE(big.NewInt(1), big.NewInt(5))},
		{NewFE(big.NewInt(1), big.NewInt(11)), NewFE(big.NewInt(9), big.NewInt(11)), NewFE(big.NewInt(8), big.NewInt(11)), NewFE(big.NewInt(7), big.NewInt(11))},
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
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(4), big.NewInt(5)), NewFE(big.NewInt(4), big.NewInt(5))},
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(1), big.NewInt(5)), NewFE(big.NewInt(2), big.NewInt(5))},
		{NewFE(big.NewInt(1), big.NewInt(11)), NewFE(big.NewInt(9), big.NewInt(11)), NewFE(big.NewInt(3), big.NewInt(11))},
		{NewFE(big.NewInt(10), big.NewInt(11)), NewFE(big.NewInt(6), big.NewInt(11)), NewFE(big.NewInt(4), big.NewInt(11))},
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
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(4), big.NewInt(5)), NewFE(big.NewInt(0), big.NewInt(5)), NewFE(big.NewInt(4), big.NewInt(5))},
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(1), big.NewInt(5)), NewFE(big.NewInt(2), big.NewInt(5)), NewFE(big.NewInt(0), big.NewInt(5))},
		{NewFE(big.NewInt(1), big.NewInt(11)), NewFE(big.NewInt(9), big.NewInt(11)), NewFE(big.NewInt(8), big.NewInt(11)), NewFE(big.NewInt(6), big.NewInt(11))},
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
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(4), big.NewInt(5)), NewFE(big.NewInt(2), big.NewInt(5))},
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(1), big.NewInt(5)), NewFE(big.NewInt(3), big.NewInt(5))},
		{NewFE(big.NewInt(1), big.NewInt(11)), NewFE(big.NewInt(9), big.NewInt(11)), NewFE(big.NewInt(9), big.NewInt(11))},
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
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(4), big.NewInt(5)), NewFE(big.NewInt(0), big.NewInt(5)), NewFE(big.NewInt(0), big.NewInt(5))},
		{NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(1), big.NewInt(5)), NewFE(big.NewInt(2), big.NewInt(5)), NewFE(big.NewInt(1), big.NewInt(5))},
		{NewFE(big.NewInt(1), big.NewInt(11)), NewFE(big.NewInt(9), big.NewInt(11)), NewFE(big.NewInt(8), big.NewInt(11)), NewFE(big.NewInt(6), big.NewInt(11))},
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
		exp     *big.Int
		x, want FE
	}{
		{big.NewInt(3), NewFE(big.NewInt(3), big.NewInt(5)), NewFE(big.NewInt(2), big.NewInt(5))},
		{big.NewInt(6), NewFE(big.NewInt(2), big.NewInt(5)), NewFE(big.NewInt(4), big.NewInt(5))},
		{big.NewInt(7), NewFE(big.NewInt(5), big.NewInt(11)), NewFE(big.NewInt(3), big.NewInt(11))},
		{big.NewInt(-3), NewFE(big.NewInt(7), big.NewInt(13)), NewFE(big.NewInt(8), big.NewInt(13))},
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
		{NewFE(big.NewInt(2), big.NewInt(19)), NewFE(big.NewInt(7), big.NewInt(19)), NewFE(big.NewInt(3), big.NewInt(19))},
		{NewFE(big.NewInt(7), big.NewInt(19)), NewFE(big.NewInt(5), big.NewInt(19)), NewFE(big.NewInt(9), big.NewInt(19))},
	}

	for _, test := range testCases {
		result := test.x.Div(test.y)
		if !result.Equals(test.want) {
			t.Errorf("got %s, want %s", result.String(), test.want)
		}
	}
}
