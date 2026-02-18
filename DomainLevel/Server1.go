package DomainLevel

type Proccess interface {
	ProsessAndSendData([]byte, []byte, string) error
}
