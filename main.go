package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"moos/ecc"
)

func main() {

	secret := "qwerty"
	hash := sha256.New()
	hash.Write([]byte(secret))
	privateKey := new(big.Int).SetBytes(hash.Sum(nil))
	kp := ecc.NewKeyPair(privateKey)
	hash = sha256.New()
	hash.Write([]byte("Mysdjkfnkjsdhfksdhfkjhk"))
	z := new(big.Int).SetBytes(hash.Sum(nil))

	sig := kp.Sign(z)

	publicKey := kp.Address

	fmt.Println(ecc.Verify(publicKey, z, sig))

}
