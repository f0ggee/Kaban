package serverManagment

import (
	"crypto/rsa"
	"crypto/x509"
)

func (s *ServerManagement) ConvertPrivateKey(Data []byte) (*rsa.PrivateKey, error) {

	return x509.ParsePKCS1PrivateKey(Data)
}
