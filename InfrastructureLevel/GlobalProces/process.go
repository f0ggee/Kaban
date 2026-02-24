package GlobalProces

import (
	"MasterServer_/DomainLevel"
	"MasterServer_/Dto"
	"encoding/hex"
	"errors"
	"log/slog"
	"os"

	"github.com/awnumar/memguard"
)

type ProcessController struct {
	KeyInteracting   DomainLevel.KeyInteracting
	RedisInteracting DomainLevel.RedisUse
	ServerManagement DomainLevel.ServerDataManagement
	Process          DomainLevel.Process
}

//func NewProcessController(keyInteracting DomainLevel.KeyInteracting, redisInteracting DomainLevel.RedisUse, serverManagement DomainLevel.ServerDataManagement) *ProcessController {
//	return &ProcessController{KeyInteracting: keyInteracting, RedisInteracting: redisInteracting, ServerManagement: serverManagement}
//}

type AnotherProcessController struct {
	E ProcessController
}

func NewAnotherProcessController(e ProcessController) *AnotherProcessController {
	return &AnotherProcessController{E: e}
}

func (psa *AnotherProcessController) HandlingAndSendData(KeyOfServer []byte, RsaKeyNew []byte, NameServer string) error {

	slog.Info("Starting HandlingAndSendData")

	AesKey := memguard.NewBufferFromBytes(psa.E.KeyInteracting.AesKey())
	if AesKey == nil {
		return errors.New("AesKey is nil")
	}
	defer AesKey.Destroy()

	EncryptedRsaKey, err := psa.E.KeyInteracting.EncryptRsaKey(AesKey.Bytes(), RsaKeyNew)
	if err != nil {

		return err
	}

	EncryptedAesKey, err := psa.E.KeyInteracting.EncryptAesKey(AesKey.Bytes(), KeyOfServer)
	if err != nil {
		return err
	}

	HashSha := psa.E.KeyInteracting.GenerateHash(EncryptedRsaKey, EncryptedAesKey)

	MasterServerPrivateKey := os.Getenv("OurKey")
	BytesMasterServerPrivateKey, err := hex.DecodeString(MasterServerPrivateKey)
	if err != nil {
		slog.Error("Error decoding master server private key", "Error", err.Error())
		return err
	}

	Sign, err := psa.E.KeyInteracting.GenerateSignature(HashSha, BytesMasterServerPrivateKey)
	if err != nil {
		return err
	}

	RedisDat := &Dto.RedisDataLooksLike{
		AesKey:    EncryptedAesKey,
		PlainText: EncryptedRsaKey,
		Signature: Sign,
	}

	err = psa.E.RedisInteracting.SendData(RedisDat, NameServer)
	if err != nil {
		return err
	}

	return nil

}
