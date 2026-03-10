package GrpcImplementation

import (
	"MasterServer_/Dto"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
)

func (s *GrpcDataManagement) GenerateSignature() ([]byte, error) {

	PrivateKeyBytes, err := hex.DecodeString(Dto.Keys.MasterServerKey)
	if err != nil {
		slog.Error("GenerateSignature Error while trying to decode private key", "error", err.Error())
		return nil, err
	}

	PrivateKey, err := s.ServerDataManagement.ConvertPrivateKey(PrivateKeyBytes)
	if err != nil {
		slog.Error("GenerateSignature Error while trying to convert private key", "error", err.Error())
		return nil, err
	}

	hash := sha256.New().Sum(Dto.Keys.NewPrivateKey.Bytes())
	SignedMassage, err := rsa.SignPKCS1v15(rand.Reader, PrivateKey, crypto.SHA256, hash)
	if err != nil {
		slog.Error("GenerateSignature Error while trying to sign signature", "error", err.Error())
		return nil, err
	}
	return SignedMassage, nil

}
