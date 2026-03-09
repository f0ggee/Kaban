package serverManagment

import (
	"MasterServer_/DomainLevel"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"os"
	"time"
)

type Pack2 struct {
	RsaKey DomainLevel.RsaKeyManipulation
}

type ServerManagement struct {
	S Pack2
}

var MappingHash = make(map[[32]byte]time.Time)

func (s *ServerManagement) FindHash(data []byte) bool {
	sha := sha256.Sum256(data)

	if Time, ok := MappingHash[sha]; ok {

		if time.Now().After(Time) {
			return true
		}
	}

	return false

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
		rsaKey, err := hex.DecodeString(os.Getenv("server1"))
		if err != nil {
			slog.Error("Error in getting the server key", "ServerNumber", Num)
			return nil
		}

		slog.Info("Got server key", "ServerNumber", Num)
		return rsaKey

	case 2:

		//keyInBytes := s.S.RsaKey.ConvertRsaKeyToBytes(os.Getenv("server2"))

		rsaKey, err := hex.DecodeString(os.Getenv("server2"))
		if err != nil {
			slog.Error("Error in getting the server key", "ServerNumber", Num)
			return nil
		}
		slog.Info("Got server key", "ServerNumber", Num)
		return rsaKey

	}

	slog.Info("Couldn't find the server key", "ServerNumber", Num)
	return nil
}
