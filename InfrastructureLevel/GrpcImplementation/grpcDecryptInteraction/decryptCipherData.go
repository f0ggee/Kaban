package grpcDecryptInteraction

import (
	"crypto/aes"
	"crypto/cipher"
	"log/slog"
)

func (g GrpcDecryptRealization) DecryptCipherData(AesKey []byte, plainText []byte) ([]byte, error) {

	AesGcm, err := aes.NewCipher(AesKey)
	if err != nil {
		slog.Error("Error creating new Aes block", "error", err.Error())
		return nil, err
	}

	NewGcm, err := cipher.NewGCM(AesGcm)
	if err != nil {
		slog.Error("Error creating new Gcm", "error", err.Error())
		return nil, err
	}

	PlainText, err := NewGcm.Open(nil, plainText[:NewGcm.NonceSize()], plainText[NewGcm.NonceSize():], nil)
	if err != nil {
		slog.Error("Error decrypting ciphertext", "error", err.Error())
		return nil, err
	}
	return PlainText, nil
}
