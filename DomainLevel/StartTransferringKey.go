package DomainLevel

type Process interface {
	HandlingAndSendData([]byte, []byte, string) error
}
