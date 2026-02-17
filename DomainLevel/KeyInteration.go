package DomainLevel

type KeyInteration interface {
	AesKey() []byte
	EncryptRsaKey([]byte, []byte) ([]byte, error)
	EncryptAesKey([]byte, []byte) ([]byte, error)
	GenerateHash([]byte, []byte) []byte
	GenerateSignature([]byte, []byte) ([]byte, error)
}
