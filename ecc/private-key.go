package ecc

import (
	"crypto/hmac"
	"crypto/sha256"
	"math/big"
)

type Keypair struct {
	Private *big.Int
	Address Secp256k1Point
}

func NewKeyPair(private *big.Int) Keypair {
	address := NewSecp256k1Point(SECP256K1GPointX, SECP256K1GPointY).Scale(new(big.Int).Set(private))

	return Keypair{Private: private, Address: ToSecp256k1Point(address)}
}

// eG = P
// k - random big number
// R = kG -> r = R.x
// s = (z + re)/k
func (kp Keypair) Sign(z *big.Int) Signature {
	k := kp.Deterministic(z)
	r := NewSecp256k1Point(SECP256K1GPointX, SECP256K1GPointY).Scale(k).X().Num()
	kInv := new(big.Int).Exp(k, new(big.Int).Sub(SECP256K1Order, big.NewInt(2)), SECP256K1Order)
	s := new(big.Int).Add(z, new(big.Int).Mul(r, kp.Private))
	s = s.Mul(s, kInv).Mod(s, SECP256K1Order)

	// It turns out that using the low-s value will get nodes to relay our transactions.
	// This is for malleability reasons.
	if s.Cmp(new(big.Int).Div(SECP256K1Order, big.NewInt(2))) == 1 {
		s = s.Sub(SECP256K1Order, s)
	}

	return NewSignature(r, s)
}

// z - the double sha256
// u = z/s   v = r/s   R = uG + vP
func Verify(p Secp256k1Point, z *big.Int, sig Signature) bool {
	sInv := new(big.Int).Exp(sig.S, new(big.Int).Sub(SECP256K1Order, big.NewInt(2)), SECP256K1Order)
	u := new(big.Int).Mul(z, sInv)
	u = u.Mod(u, SECP256K1Order)
	v := new(big.Int).Mul(sig.R, sInv)
	v = v.Mod(v, SECP256K1Order)
	R := NewSecp256k1Point(SECP256K1GPointX, SECP256K1GPointY).Scale(u).Add(p.Scale(v))

	// x coordinate must match
	return R.X().Num().Cmp(sig.R) == 0
}

// generate a k which is garantueed not to be duplicated
func (kp Keypair) Deterministic(z *big.Int) *big.Int {
	k := make([]byte, 32)
	v := make([]byte, 32)

	for i := range v {
		v[i] = byte(0x1)
	}

	if z.Cmp(SECP256K1Order) == 1 {
		z = new(big.Int).Sub(z, SECP256K1Order)
	}

	msg := append(v, byte(0x1))
	msg = append(msg, kp.Private.Bytes()...)
	msg = append(msg, z.Bytes()...)

	kHash := hmac.New(sha256.New, k)
	kHash.Write(msg)
	k = kHash.Sum(nil)

	vHash := hmac.New(sha256.New, k)
	vHash.Write(v)
	v = vHash.Sum(nil)

	msg = append(v, byte(0x1))
	msg = append(msg, kp.Private.Bytes()...)
	msg = append(msg, z.Bytes()...)

	kHash = hmac.New(sha256.New, k)
	kHash.Write(msg)
	k = kHash.Sum(nil)

	vHash = hmac.New(sha256.New, k)
	vHash.Write(v)
	v = vHash.Sum(nil)

	for {
		vHash = hmac.New(sha256.New, k)
		vHash.Write(v)
		v = vHash.Sum(nil)
		candidate := new(big.Int).SetBytes(v)

		if GreaterOrEqual(candidate, big.NewInt(1)) && candidate.Cmp(SECP256K1Order) < 0 {
			return candidate
		}

		msg = append(v, byte(0x00))
		kHash = hmac.New(sha256.New, k)
		kHash.Write(msg)
		k = kHash.Sum(nil)

		vHash = hmac.New(sha256.New, k)
		vHash.Write(v)
		v = vHash.Sum(nil)
	}
}
