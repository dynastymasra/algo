package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"golang.org/x/crypto/openpgp/elgamal"
)

const (
	// This is the 1024-bit MODP group from RFC 5114
	primeHex     = "B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4371"
	generatorHex = "A4D1CBD5C3FD34126765A442EFB99905F8104DD258AC507FD6406CFF14266D31266FEA1E5C41564B777E690F5504F213160217B4B01B886A5E91547F9E2749F4D7FBD7D3B9A92EE1909D0D2263F80A76A6A24C087A091F531DBF0A0169B6A28AD662A4D18E73AFA32D779D5918D08BC8858F4DCEF97C2A24855E6EEB22B3B2E5"
)

func main() {
	start := time.Now().UTC()
	fmt.Println("Time start = ", start)
	fmt.Println("Go with ElGamal")

	G, _ := new(big.Int).SetString(generatorHex, 16)
	P, _ := new(big.Int).SetString(primeHex, 16)
	X, _ := new(big.Int).SetString("42", 16)

	fmt.Println("G = ", G)
	fmt.Println("P = ", P)
	fmt.Println("X = ", X)

	priv := &elgamal.PrivateKey{
		PublicKey: elgamal.PublicKey{
			G: G,
			P: P,
			Y: new(big.Int).Exp(G, X, P),
		},
		X: X,
	}

	fmt.Println("Y = ", priv.Y)

	message := []byte("Hello ElGamal")
	c1, c2, err := elgamal.Encrypt(rand.Reader, &priv.PublicKey, message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("C1 = ", c1)
	fmt.Println("C2 = ", c2)

	now := time.Now().UTC()
	fmt.Println("Time Encrypt = ", now)
	fmt.Println("Time Encrypt duration = ", now.Sub(start))

	result, err := elgamal.Decrypt(priv, c1, c2)
	if err != nil {
		log.Fatalln(err)
	}

	end := time.Now().UTC()
	fmt.Println("Result = ", string(result))
	fmt.Println("Time End = ", end)
	fmt.Println("Time Decrypt = ", end.Sub(now))
	fmt.Println("Total Time = ", end.Sub(start))
}
