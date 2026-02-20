package DomainLevel

type KeyInteracting interface {
	AesKey() []byte
	EncryptRsaKey([]byte, []byte) ([]byte, error)
	EncryptAesKey([]byte, []byte) ([]byte, error)
	GenerateHash([]byte, []byte) []byte
	GenerateSignature([]byte, []byte) ([]byte, error)
}
