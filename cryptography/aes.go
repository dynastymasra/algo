package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"time"
)

// This is using SHA256
var key = "CA978112CA1BBDCAFAC231B39A23DC4DA786EFF8147C4E72B9807785AFEE48BB"

func main() {
	start := time.Now()

	ciphertext, err := encrypt()
	if err != nil {
		fmt.Println("Failed to encrypt", err)
		return
	}

	if err := decrypt(ciphertext); err != nil {
		fmt.Println("Failed to decrypt", err)
		return
	}

	end := time.Now()
	fmt.Println("Start =", start.Format(time.RFC3339))
	fmt.Println("End =", start.Format(time.RFC3339))
	fmt.Println("Delta =", end.Sub(start))
}

func encrypt() (string, error) {
	start := time.Now().UTC()
	fmt.Println("Encryption AES-GCM")

	text := []byte("This is very long string text to test AES")

	// 32 Byte long key, you can use another key as long as 32/16 bytes(AES-256/AES-128)
	secret, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	fmt.Println("Text to Encrypt =", string(text))
	ciphertext := fmt.Sprintf("%x", gcm.Seal(nonce, nonce, text, nil))
	fmt.Println("Encryption with AES-256 =", ciphertext)

	end := time.Now().UTC()

	fmt.Println("Start Encryption =", start.Format(time.RFC3339))
	fmt.Println("End Encryption =", start.Format(time.RFC3339))
	fmt.Println("Delta Encryption =", end.Sub(start))

	return ciphertext, nil
}

func decrypt(ciphertext string) error {
	start := time.Now().UTC()
	fmt.Println("Decryption AES-GCM")

	// 32 Byte long key, you can use another key as long as 32/16 bytes(AES-256/AES-128)
	secret, err := hex.DecodeString(key)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(secret)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	ciphertextByte, err := hex.DecodeString(ciphertext)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertextByte) < nonceSize {
		return errors.New("error")
	}

	nonce, ciphertextByte := ciphertextByte[:nonceSize], ciphertextByte[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertextByte, nil)
	if err != nil {
		return err
	}

	fmt.Println("Decryption with AES-256 =", string(plaintext))

	end := time.Now().UTC()

	fmt.Println("Start Decryption =", start.Format(time.RFC3339))
	fmt.Println("End Decryption =", start.Format(time.RFC3339))
	fmt.Println("Delta Decryption =", end.Sub(start))

	return nil
}
