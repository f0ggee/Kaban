package InftarctionLevel

import (
	"MasterServer_/InftarctionLevel/keyInteration"
	"encoding/hex"
	"errors"
	"log/slog"
	"os"

	"github.com/awnumar/memguard"
)

func init() {

}

type ProcessContoller struct{}

func (p ProcessContoller) ProsessAndSendData(KeyOfServer []byte, RsaKeyNew []byte, NameServer string) error {

	keysInteraction := *keyInteration.Connection()
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

	slog.Info("Generated master server signature", "Signature", Sign)
}
