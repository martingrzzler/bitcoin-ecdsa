package ecc

import (
	"bytes"
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

func (s Signature) DER() []byte {
	rbin := s.R.Bytes()
	rbin = bytes.TrimLeft(rbin, string(rune(0)))

	if rbin[0]&0x80 != 0 {
		rbin = append([]byte{0x00}, rbin...)
	}
	result := append([]byte{0x2, byte(len(rbin))}, rbin...)
	sbin := s.S.Bytes()
	sbin = bytes.TrimLeft(sbin, string(rune(0)))

	if sbin[0]&0x80 != 0 {
		sbin = append([]byte{0x00}, sbin...)
	}

	result = append(result, []byte{0x2, byte(len(sbin))}...)
	result = append(result, sbin...)

	return append([]byte{0x30, byte(len(result))}, result...)
}
