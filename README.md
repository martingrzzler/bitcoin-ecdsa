# bitcoin-ecdsa
implementation of ECDSA algorithm used in Bitcoin

This implementation is solely for educational purposes.
I've used the book Programming *Bitcoin: Learn How to Program Bitcoin from Scratch* by *Jimmy Song* as a learning resource.

### Overview

`FE` stands for finite element and implements the math for finite fields. Moreover elliptic curves are represented by a simple `Point` structure. 
Bitcoin's elliptic curve (y^2 = x^3 + 7) is implemented as `SECP256k1` struct. The library enables signing and verfification operations. See `main.go` for an example.

### Run example
- `go run main.go`
