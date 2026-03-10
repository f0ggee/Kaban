package serverManagment

import (
	"crypto/rsa"
	"crypto/x509"
)

func (s *ServerManagement) ConvertPublicKey(bytes []byte) (*rsa.PublicKey, error) {

	return x509.ParsePKCS1PublicKey(bytes)
}
