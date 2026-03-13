package GrpcImplementation

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"log/slog"
)

func (s *GrpcDataManagement) CheckSignature(Sign []byte, PublicServer []byte, ServerName []byte) error {
	PublicKey, err := s.ServerDataManagement.S.ConvertPublicKey(PublicServer)
	if err != nil {
		return err
	}
	Hash := sha256.Sum256(ServerName)
	err = rsa.VerifyPKCS1v15(PublicKey, crypto.SHA256, Hash[:], Sign)
	if err != nil {
		slog.Error("Error validating the signature of the calling function ", "Error", err.Error())
		return err
	}
	return nil
}
