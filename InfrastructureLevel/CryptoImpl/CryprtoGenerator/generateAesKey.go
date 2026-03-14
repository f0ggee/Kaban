package CryprtoGenerator

import (
	"crypto/rand"
	"log/slog"
)

func (*CryprtoGenerating) GenerateAesKey() []byte {

	AesKey := make([]byte, 32)
	if _, err := rand.Read(AesKey); err != nil {

		slog.Error("Error generating AesKey")
		return []byte{}
	}

	return AesKey
}
