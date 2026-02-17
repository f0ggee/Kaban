package DomainLevel

type RedisUse interface {
	SendData([]byte, []byte, []byte, string) error
}
