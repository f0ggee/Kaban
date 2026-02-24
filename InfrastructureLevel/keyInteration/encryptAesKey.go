package keyInteration

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log/slog"
)

func (s *KeyInterationController) EncryptAesKey(AesKey []byte, RsaKey []byte) ([]byte, error) {

	RsaKeyInPublic, err := x509.ParsePKCS1PublicKey(RsaKey)
	if err != nil {
		slog.Error("Error Parsing RsaKey", "Error", err.Error())
		return nil, err
	}

	encryptedText, err := rsa.EncryptPKCS1v15(rand.Reader, RsaKeyInPublic, AesKey)
	if err != nil {
		slog.Error("Error Encrypting", "Error", err.Error())
		return nil, err
	}

	return encryptedText, nil
}
