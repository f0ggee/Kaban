package serverManagment

import (
	"MasterServer_/DomainLevel"
	"encoding/hex"
	"log"
	"log/slog"
	"os"
)

type Pack2 struct {
	RsaKey DomainLevel.RsaKeyManipulation
}

type ServerManagement struct {
	S Pack2
}

func (s *ServerManagement) SayHi() string {
	//TODO implement me
	return "Hello World!"
}

func NewServerManagement(s Pack2) *ServerManagement {
	return &ServerManagement{S: s}
}

func (s *ServerManagement) GetServerKey(Num int) []byte {
	slog.Info("start getting the server key", "ServerNumber", Num)
	switch Num {
	case 1:

		ea := os.Getenv("server1")
		slog.Info("E:", "E", ea)
		//KeyInBytes := s.RsaKey.ConvertRsaKeyToBytes(ea)
		////if KeyInBytes == nil {
		////	log.Println("Rsa key does not exist")
		////	return nil
		////}

		KeyInBytes, err := hex.DecodeString(ea)
		if err != nil {
			log.Println("Error converting rsa key to bytes", "Error", err.Error())
			return nil
		}
		slog.Info("Got server key", "ServerNumber", Num)
		return KeyInBytes

	}

	slog.Info("Couldn't find the server key", "ServerNumber", Num)
	return nil
}
