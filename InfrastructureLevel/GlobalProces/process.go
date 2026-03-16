package GlobalProces

import (
	"MasterServer_/Dto"
	"encoding/hex"
	"errors"
	"log/slog"
	"os"

	"github.com/awnumar/memguard"
)

func (psa *ControllingExchange) HandlingAndSendData(KeyOfServer []byte, RsaKeyNew []byte, NameServer string) error {

	slog.Info("Starting HandlingRequests")

	AesKey := memguard.NewBufferFromBytes(psa.E.CryptoGen.GenerateAesKey())
	if AesKey == nil {
		return errors.New("AesKey is nil")
	}
	defer AesKey.Destroy()

	EncryptedRsaKey, err := psa.E.Cryptos.EncrypterRsaKey(AesKey.Bytes(), RsaKeyNew)
	if err != nil {

		return err
	}

	EncryptedAesKey, err := psa.E.Cryptos.EncrypterAesKey(AesKey.Bytes(), KeyOfServer)
	if err != nil {
		return err
	}

	HashSha := psa.E.CryptoGen.GenerateHash(EncryptedRsaKey, EncryptedAesKey)

	MasterServerPrivateKey := os.Getenv("OurKey")
	BytesMasterServerPrivateKey, err := hex.DecodeString(MasterServerPrivateKey)
	if err != nil {
		slog.Error("Error decoding master server private key", "Error", err.Error())
		return err
	}

	Sign, err := psa.E.CryptoGen.SignerData(HashSha, BytesMasterServerPrivateKey)
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
