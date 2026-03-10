package DomainLevel

import "crypto/rsa"

type ServerDataManagement interface {
	GetServerKey(int) []byte
	GetServerName(int) string
	ConvertPrivateKey([]byte) (*rsa.PrivateKey, error)
	ConvertPublicKey([]byte) (*rsa.PublicKey, error)
}
