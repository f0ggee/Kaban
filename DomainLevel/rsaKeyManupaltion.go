package DomainLevel

type RsaKeyManipulation interface {
	GenerateRsaKey() []byte
	ConvertRsaKeyToBytes(string) []byte
}
