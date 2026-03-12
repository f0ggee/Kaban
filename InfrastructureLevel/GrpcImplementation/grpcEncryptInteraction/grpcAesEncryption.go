package grpcEncryptInteraction

import (
	"MasterServer_/DomainLevel"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log/slog"
)

type PackGrpcEncryptInteraction struct {
	DataManagement DomainLevel.ServerDataManagement
}

type GrpcEncryptInteraction struct {
	S PackGrpcEncryptInteraction
}

func (g GrpcEncryptInteraction) GrpcAesEncryption(Data []byte, aesKey []byte) ([]byte, error) {

	aesGCM, err := aes.NewCipher(aesKey)
	if err != nil {
		slog.Error("Error creating new AES cipher", "Error", err.Error())
		return nil, err
	}

	gcm, err := cipher.NewGCM(aesGCM)
	if err != nil {
		slog.Error("Error creating new GCM", "Error", err.Error())
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		slog.Error("Error creating new nonce", "Error", err.Error())
		return nil, err
	}

	return gcm.Seal(nonce, nonce, Data, nil), nil
}
