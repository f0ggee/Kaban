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
	Process          DomainLevel.Process
	ServerManagement DomainLevel.ServerDataManagement
}

func NewProcessController(keyInteracting DomainLevel.KeyInteracting, redisInteracting DomainLevel.RedisUse, process DomainLevel.Process, serverManagement DomainLevel.ServerDataManagement) *ProcessController {
	return &ProcessController{KeyInteracting: keyInteracting, RedisInteracting: redisInteracting, Process: process, ServerManagement: serverManagement}
}

func (p ProcessController) HandlingAndSendData(KeyOfServer []byte, RsaKeyNew []byte, NameServer string) error {

	AesKey := memguard.NewBufferFromBytes(p.KeyInteracting.AesKey())
	if AesKey == nil {
		return errors.New("AesKey is nil")
	}
	defer AesKey.Destroy()

	EncryptedRsaKey, err := p.KeyInteracting.EncryptRsaKey(AesKey.Bytes(), RsaKeyNew)
	if err != nil {

		return err
	}

	EncryptedAesKey, err := p.KeyInteracting.EncryptAesKey(AesKey.Bytes(), KeyOfServer)
	if err != nil {
		return err
	}

	HashSha := p.KeyInteracting.GenerateHash(EncryptedRsaKey, EncryptedAesKey)

	MasterServerPrivateKey := os.Getenv("OurKey")
	BytesMasterServerPrivateKey, err := hex.DecodeString(MasterServerPrivateKey)
	if err != nil {
		slog.Error("Error decoding master server private key", "Error", err.Error())
		return err
	}

	Sign, err := p.KeyInteracting.GenerateSignature(HashSha, BytesMasterServerPrivateKey)
	if err != nil {
		return err
	}

	RedisDat := &Dto.RedisDataLooksLike{
		AesKey:    EncryptedAesKey,
		PlainText: EncryptedRsaKey,
		Signature: Sign,
	}

	err = p.RedisInteracting.SendData(RedisDat, NameServer)
	if err != nil {
		return err
	}

	return nil

}
