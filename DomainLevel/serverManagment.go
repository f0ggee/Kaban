package DomainLevel

type ServerDataManagement interface {
	GetServerKey(int) string
	GetServerName(int) string
}
