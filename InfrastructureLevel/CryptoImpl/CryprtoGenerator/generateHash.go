package CryprtoGenerator

import (
	"crypto/sha256"
)

func (s *CryprtoGenerating) GenerateHash(DataToHash []byte) []byte {
	ShaHash := sha256.New().Sum(DataToHash)
	return ShaHash
}
