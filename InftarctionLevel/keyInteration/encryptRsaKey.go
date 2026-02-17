package keyInteration

import (
	"crypto/aes"
	"crypto/cipher"
	"log/slog"
)

func (d *KeyInterationController) EncryptRsaKey(AesKey []byte, RsaKey []byte) ([]byte, error) {

	aesNewBlocke, err := aes.NewCipher(AesKey)
	if err != nil {
		slog.Error("error creating new aes cipher", "Error", err.Error())
		return nil, err

	}

	gcm, err := cipher.NewGCM(aesNewBlocke)
	if err != nil {
		slog.Error("error creating new gcm", "Error", err.Error())
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())

	text := gcm.Seal(nonce, nonce, RsaKey, nil)

	return text, nil
}
