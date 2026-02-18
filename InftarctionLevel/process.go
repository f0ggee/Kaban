package InftarctionLevel

import (
	"MasterServer_/Dto"
	"MasterServer_/InftarctionLevel/keyInteration"
	"encoding/hex"
	"errors"
	"log/slog"
	"os"

	"MasterServer_/InftarctionLevel/RedisUse"
	"github.com/awnumar/memguard"
)

type ProcessContoller struct{}

func (p ProcessContoller) ProsessAndSendData(KeyOfServer []byte, RsaKeyNew []byte, NameServer string) error {

	keysInteraction := *keyInteration.Connection()
	redisConnect := *RedisUse.ConnectionRedis()
	AesKey := memguard.NewBufferFromBytes(keysInteraction.Tokens.AesKey())
	if AesKey == nil {
		return errors.New("AesKey is nil")
	}
	defer AesKey.Destroy()

	EncryptedRsaKey, err := keysInteraction.Tokens.EncryptRsaKey(AesKey.Bytes(), RsaKeyNew)
	if err != nil {

		return err
	}

	EncryptedAesKey, err := keysInteraction.Tokens.EncryptAesKey(AesKey.Bytes(), KeyOfServer)
	if err != nil {
		return err
	}

	HashSha := keysInteraction.Tokens.GenerateHash(EncryptedRsaKey, EncryptedAesKey)

	MasterServerPrivateKey := os.Getenv("OurKey")
	BytesMasterServerPrivateKey, err := hex.DecodeString(MasterServerPrivateKey)
	if err != nil {
		slog.Error("Error decoding master server private key", "Error", err.Error())
		return err
	}

	Sign, err := keysInteraction.Tokens.GenerateSignature(HashSha, BytesMasterServerPrivateKey)
	if err != nil {
		return err
	}

	RedisDat := &Dto.RedisDataLooksLike{
		AesKey:    EncryptedAesKey,
		PlainText: EncryptedRsaKey,
		Signature: Sign,
	}

	err = redisConnect.Client.SendData(RedisDat, NameServer)
	if err != nil {
		return err
	}

	return nil

}
