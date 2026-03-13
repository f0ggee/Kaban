package grpcEncryptInteraction

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"log/slog"
)

func (g GrpcEncryptInteraction) GrpcEncryptAesKey(AesKey []byte, RsaKey []byte) ([]byte, error) {

	//RsaKeyPrivateKeyType, err := g.S.DataManagement.ConvertPrivateKey(RsaKey)
	//if err != nil {
	//	slog.Error("Error parsing RsaKey", "Error", err.Error())
	//	return nil, err
	//}

	RsaKeyPublic, err := x509.ParsePKCS1PublicKey(RsaKey)
	if err != nil {
		slog.Error("Error Parsing RSA key", "Error", err.Error())
		return nil, err
	}
	aesEncryptedKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, RsaKeyPublic, AesKey, nil)
	if err != nil {
		slog.Error("Error encrypting AES key", "Error", err.Error())
		return nil, err
	}
	return aesEncryptedKey, nil
}
