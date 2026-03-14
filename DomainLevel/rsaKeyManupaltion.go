package DomainLevel

type RsaKeyManipulation interface {
	GenerateRsaKey() []byte
}
