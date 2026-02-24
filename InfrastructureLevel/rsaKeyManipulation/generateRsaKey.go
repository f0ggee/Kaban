package rsaKeyManipulation

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
)

type RsaKeyManipulation struct{}

func (r *RsaKeyManipulation) ConvertRsaKeyToBytes(RsaKeyString string) []byte {

	log.Println("Start converting rsa key to bytes")

	//Rsa, err := hex.DecodeString(RsaKeyString)
	//if err != nil {
	//	log.Println("Error converting rsa key to bytes", "Error", err.Error())
	//	return nil
	//}

	return nil
}

func (r *RsaKeyManipulation) GenerateRsaKey() []byte {
	log.Println("RsaKeyManipulation.GenerateRsaKey()", "Start", true)

	RsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Println("Getting an error while generating key RSA", "Error", err)
		log.Println("Generating Rsa key", "Success", false)

		return nil
	}
	log.Println("Generating Rsa key", "Success", true)
	RsaKeyInBytes := x509.MarshalPKCS1PrivateKey(RsaKey)

	return RsaKeyInBytes

}
