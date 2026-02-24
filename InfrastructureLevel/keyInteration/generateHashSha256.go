package keyInteration

import "crypto/sha256"

func (s *KeyInterationController) GenerateHash(EncryptedRsaKey []byte, EncryptedAesKey []byte) []byte {
	ShaHash := sha256.New()
	ShaHash.Write(EncryptedRsaKey)
	ShaHash.Write(EncryptedAesKey)
	return ShaHash.Sum(nil)
}
