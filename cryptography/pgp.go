package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"

	"golang.org/x/crypto/openpgp/packet"
)

type PGP struct {
	message  []byte
	password []byte
	config   *packet.Config
}

func NewPGP(message, password []byte, config *packet.Config) *PGP {
	return &PGP{
		message:  message,
		password: password,
		config:   config,
	}
}

func main() {
	start := time.Now().UTC()
	fmt.Println("Hello PGP")
	fmt.Println("Time start = ", start)

	message := []byte("Hello PGP")
	password := []byte("password")
	config := &packet.Config{
		DefaultCipher: packet.CipherAES256,
	}

	pgp := NewPGP(message, password, config)
	encrypted, err := pgp.Encrypt()
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now().UTC()
	fmt.Println("Encrypted:", string(encrypted))
	fmt.Println("Time now = ", now)
	fmt.Println("Time Encrypt = ", now.Sub(start))

	decrypted, err := pgp.Decrypt(encrypted)
	if err != nil {
		log.Fatalln(err)
	}

	end := time.Now().UTC()
	fmt.Println("Time end = ", end)
	fmt.Println("Time Decrypt = ", end.Sub(now))
	fmt.Println("Time total = ", end.Sub(start))
	fmt.Println("Decrypted:", string(decrypted))
}

func (p *PGP) Encrypt() ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	w, err := armor.Encode(buf, "PGP Message", nil)
	if err != nil {
		return nil, err
	}
	defer w.Close()

	pt, err := openpgp.SymmetricallyEncrypt(w, p.password, nil, p.config)
	if err != nil {
		return nil, err
	}
	defer pt.Close()

	_, err = pt.Write(p.message)
	if err != nil {
		return nil, err
	}

	// Close writers to force-flush their buffer
	pt.Close()
	w.Close()

	return buf.Bytes(), nil
}

func (p *PGP) Decrypt(encrypted []byte) ([]byte, error) {
	buf := bytes.NewBuffer(encrypted)

	block, err := armor.Decode(buf)
	if err != nil {
		return nil, err
	}

	failed := false
	prompt := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		if failed {
			return nil, errors.New("decryption failed")
		}
		failed = true
		return p.password, nil
	}

	md, err := openpgp.ReadMessage(block.Body, nil, prompt, p.config)
	if err != nil {
		return nil, err
	}

	plaintext, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
