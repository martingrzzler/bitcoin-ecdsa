package ecc

import (
	"fmt"
	"math/big"
)

type Signature struct {
	// r - x coordinate of P = kG
	R *big.Int
	// s - is caclulated with the hash z and r
	S *big.Int
}

func NewSignature(r, s *big.Int) Signature {
	return Signature{R: r, S: s}
}

func (s Signature) String() string {
	return fmt.Sprintf("Signature(r: %s, s:%s)", s.R, s.S)
}
