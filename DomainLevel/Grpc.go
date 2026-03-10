package DomainLevel

type GrpcDataMalnutrition interface {
	FindHash([]byte) bool
	SaveHash([]byte)
	CheckSignature([]byte, []byte, []byte) error
	GenerateSignature() ([]byte, error)
}

type GrpcMechanism interface {
	HandlingAndSendData([]byte, []byte, string) error
}
type GrpcEncryptionInteraction interface {
	GrpcAesEncryption([]byte, []byte) ([]byte, error)
}
type GrpcDecryptInteraction interface {
	DecryptIncomingAesKey([]byte) ([]byte, error)
	DecryptCipherData([]byte, []byte) ([]byte, error)
}
