package CryprtoGenerator

import (
	"crypto/sha256"
)

func (s *CryprtoGenerating) GenerateHash(DataToHash []byte, DataHash2 []byte) []byte {

	ShaHash := sha256.New()

	ShaHash.Write(DataToHash)
	ShaHash.Write(DataHash2)
	return ShaHash.Sum(nil)
}
