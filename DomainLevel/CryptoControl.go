package DomainLevel

type Encryption interface {
	EncrypterRsaKey([]byte, []byte) ([]byte, error)
	EncrypterAesKey([]byte, []byte) ([]byte, error)
}

type CryptoGenerator interface {
	SignerData([]byte, []byte) ([]byte, error)
	GenerateHash([]byte) []byte
	GenerateAesKey() []byte
	GrpcSignerKey() ([]byte, error)
}

type Decryptor interface {
	DecrypterCipherData([]byte, []byte) ([]byte, error)
	GrpcDecrypterAesKey([]byte) ([]byte, error)
}
