package keyInteration

import (
	"crypto/rand"
)

type KeyInterationController struct {
}

func (*KeyInterationController) AesKey() []byte {

	AesKey := make([]byte, 32)
	if _, err := rand.Read(AesKey); err != nil {
		return nil
	}
	return AesKey
}
