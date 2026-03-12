package DomainLevel

type GrpcHandleData interface {
	FindHash([]byte) bool
	SaveHash([]byte)
	CheckSignature([]byte, []byte, []byte) error
	GenerateSignature() ([]byte, error)
}

type GrpcHandle interface {
	HandlingRequests([]byte, []byte, string) error
}
type GrpcEncryptor interface {
	GrpcAesEncryption([]byte, []byte) ([]byte, error)
	GrpcEncryptAesKeyByRsa(AesKey []byte, RsaKey []byte) ([]byte, error)
}
type GrpcDecryptor interface {
	DecryptIncomingAesKey([]byte) ([]byte, error)
	DecryptCipherData([]byte, []byte) ([]byte, error)
}
