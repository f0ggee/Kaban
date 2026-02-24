package DomainLevel

type ServerDataManagement interface {
	GetServerKey(int) []byte
	GetServerName(int) string
	SayHi() string
}
