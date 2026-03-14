package rsaKeyManipulation

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
)

type RsaKeyManipulation struct{}

func (r *RsaKeyManipulation) GenerateRsaKey() []byte {
	log.Println("RsaKeyManipulation.GenerateRsaKey()", "Start", true)

	RsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Println("Getting an error while generating key RSA", "Error", err)
		log.Println("Generating Rsa key", "Success", false)

		return nil
	}
	log.Println("Generating Rsa key", "Success", true)
	RsaKeyInBytes := (x509.MarshalPKCS1PrivateKey(RsaKey))

	return RsaKeyInBytes

}
