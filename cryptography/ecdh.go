package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now().UTC()
	fmt.Println("Go with ECDH (Elliptic Curve Diffie Hellman)")
	privA, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}
	pubA := privA.PublicKey

	privB, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}
	pubB := privB.PublicKey

	fmt.Println("Private Key A = ", privA.D)
	fmt.Printf("Public Key A X(%v) Y(%v)\n", pubA.X, pubA.Y)

	fmt.Println("Private Key B = ", privB.D)
	fmt.Printf("Public Key B X(%v) Y(%v)\n", pubB.X, pubB.Y)

	keyAA, keyAB := pubA.Curve.ScalarMult(pubA.X, pubA.Y, privB.D.Bytes())
	keyBA, keyBB := pubB.Curve.ScalarMult(pubB.X, pubB.Y, privA.D.Bytes())

	fmt.Println("A key A = ", keyAA)
	fmt.Println("A key B = ", keyAB)
	fmt.Println("B key A = ", keyBA)
	fmt.Println("B key B = ", keyBB)

	sharedAA := sha256.Sum256(keyAA.Bytes())
	sharedAB := sha256.Sum256(keyAB.Bytes())
	sharedBA := sha256.Sum256(keyBA.Bytes())
	sharedBB := sha256.Sum256(keyBB.Bytes())

	fmt.Printf("Shared AA = %x \n", sharedAA)
	fmt.Printf("Shared AB = %x \n", sharedAB)
	fmt.Printf("Shared BA = %x \n", sharedBA)
	fmt.Printf("Shared BB = %x \n", sharedBB)

	end := time.Now().UTC()
	delta := end.Sub(start)
	fmt.Println("Start time = ", start)
	fmt.Println("End time = ", end)
	fmt.Println("Delta time = ", delta)
}
