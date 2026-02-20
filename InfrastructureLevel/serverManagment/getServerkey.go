package serverManagment

import (
	"MasterServer_/DomainLevel"
	"log"
	"log/slog"
	"os"
	"time"
)

type Pack2 struct {
	RsaKey DomainLevel.RsaKeyManipulation
}
type ServerManagement struct{}

func (s Pack2) GetServerKey(Num int) []byte {

	slog.Time("Start searching for server key", time.Now())
	switch Num {
	case 1:

		slog.Info("start getting the server key", "ServerNumber", Num)

		KeyInBytes := s.RsaKey.ConvertRsaKeyToBytes(os.Getenv("server1"))
		if KeyInBytes == nil {
			log.Println("Rsa key does not exist")
			return nil
		}
		return KeyInBytes

	}

	slog.Info("Couldn't find the server key", "ServerNumber", Num)
	return nil
}
