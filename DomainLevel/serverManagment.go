package DomainLevel

type GettingServersInfo interface {
	GetServerKey(int) []byte
	GetServerName(int) string
}

type ConverterData interface {
	ConvertDataToJsonType(any) ([]byte, error)
}
