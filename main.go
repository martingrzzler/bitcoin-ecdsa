package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"moos/ecc"
)

func main() {
	secret := "My secret phrase"

	hash := sha256.New()
	hash.Write([]byte(secret))

	privateNumber := new(big.Int).SetBytes(hash.Sum(nil))

	keypair := ecc.NewKeyPair(privateNumber)

	message := "Bob authorizes Alice to use his car"
	hash.Reset()
	hash.Write([]byte(message))
	messageHash := hash.Sum(nil)
	z := new(big.Int).SetBytes(messageHash)

	signature := ecc.Sign(keypair, z)

	fmt.Println(signature)

	valid := ecc.Verify(keypair.Address, z, signature)

	fmt.Println(valid)

}
