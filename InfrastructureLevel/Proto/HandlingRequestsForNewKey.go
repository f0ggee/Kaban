package Proto

import (
	"MasterServer_/DomainLevel"
	"MasterServer_/Dto"
	InftarctionLevel "MasterServer_/InfrastructureLevel"
	pb "MasterServer_/InfrastructureLevel/Proto/protoFiles"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/awnumar/memguard"
)

type HandlingRequestsForNewKey struct {
	Grpc             DomainLevel.GrpcHandleData
	ServerManagement DomainLevel.GettingServersInfo
	Encryption       DomainLevel.Encryption
	Checker          DomainLevel.PacketChecker
	Decrypting       DomainLevel.Decryptor
	CryptoGenerating DomainLevel.CryptoGenerator
	ConverterJson    DomainLevel.ConverterData
}

type GrpcHandlerGettingNewKey struct {
	pb.UnimplementedSendingGettingServer
	S HandlingRequestsForNewKey
}

func NewGrpcHandlerGettingNewKey(unimplementedSendingGettingServer *pb.UnimplementedSendingGettingServer, s *HandlingRequestsForNewKey) *GrpcHandlerGettingNewKey {
	return &GrpcHandlerGettingNewKey{UnimplementedSendingGettingServer: *unimplementedSendingGettingServer, S: *s}
}

func (s GrpcHandlerGettingNewKey) GettingNewKey(ctx context.Context, data *pb.InputSendData) (*pb.OutputSendData, error) {

	NewContext, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*10))
	defer cancel()
	go func() {
		select {
		case <-NewContext.Done():
			slog.Info("The request has been cancelled")
			return
		}
	}()
	slog.Info("Start exchanging a key")
	if data == nil {
		slog.Error("Data was getting empty")
		return &pb.OutputSendData{}, errors.New("data is nil")
	}

	slog.Info("Here")

	if s.S.Checker.FindHash([32]byte(sha256.New().Sum(data.SendData[:]))) {
		slog.Error("Hash has already been used")
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	s.S.Grpc.SaveHash(data.SendData)

	DataIncomingLook := Dto.GrpcPacket{
		AesKeyData: nil,
		CipherData: nil,
	}

	err := json.Unmarshal(data.SendData, &DataIncomingLook)
	if err != nil {
		slog.Error("Unmarshal Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	DecryptedAesKey, err := s.S.Decrypting.GrpcDecrypterAesKey(DataIncomingLook.AesKeyData)
	if err != nil {
		slog.Error("Decrypt Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	Data, err := s.S.Decrypting.DecrypterCipherData(DecryptedAesKey, DataIncomingLook.CipherData)
	if err != nil {
		slog.Error("Decrypt Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	DataIntoPacket := Dto.GrpcDataIncomingPacket{
		Time:             time.Time{},
		ServerName:       nil,
		SignedServerName: nil,
	}

	err = json.Unmarshal(Data, &DataIntoPacket)
	if err != nil {
		slog.Error("Unmarshal Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}
	ResultComparingTime := s.S.Checker.CheckLifePacket(DataIntoPacket.Time)
	if ResultComparingTime {
		slog.Info("ResultComparingTime", "Time", DataIntoPacket.Time.String())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	serversKey := os.Getenv(string(DataIntoPacket.ServerName))

	if serversKey == "" {
		slog.Error("Server Key is empty")
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	ServerKey, err1 := hex.DecodeString(serversKey)
	if err1 != nil {
		slog.Error("Server Key is invalid", "Error", err1.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	err = s.S.Checker.CheckSignature(DataIntoPacket.SignedServerName, ServerKey, DataIntoPacket.ServerName)
	if err != nil {
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	SignedKey, err := s.S.CryptoGenerating.GrpcSignerKey()
	if err != nil {
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	Dto.Keys.Mu.Lock()
	OutcomingDataJson, err := s.S.ConverterJson.ConvertDataToJsonType(Dto.GrpcOutcomingDataPacket{
		Sign:   SignedKey,
		RsaKey: Dto.Keys.NewPrivateKey.Bytes(),
		T1:     time.Now(),
		T2:     InftarctionLevel.TimeForSwapping,
	})
	Dto.Keys.Mu.Unlock()

	if err != nil {
		slog.Error("Marshal Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	AesKey, err := memguard.NewBufferFromReader(rand.Reader, 32)
	if err != nil {
		slog.Error("Generate AesKey Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}
	defer AesKey.Destroy()

	var wg sync.WaitGroup
	plainText := []byte{}
	encryptedAesKey := []byte{}

	ChanError := make(chan error, 2)

	wg.Add(2)
	go func() {
		defer wg.Done()
		PlainText, err12 := s.S.Encryption.EncrypterRsaKey(AesKey.Bytes(), OutcomingDataJson)
		if err != nil {
			slog.Error("Encrypt Error", "Error", err.Error())
			ChanError <- err12
			return
		}
		plainText = PlainText
	}()

	go func() {
		defer wg.Done()
		EncryptedAesKey, err13 := s.S.Encryption.EncrypterAesKey(AesKey.Bytes(), ServerKey)
		if err != nil {
			slog.Error("Error encrypt the aesKey", "Error", err)
			ChanError <- err13
			return
		}
		encryptedAesKey = EncryptedAesKey
	}()

	wg.Wait()
	close(ChanError)
	for v := range ChanError {
		if v != nil {
			return &pb.OutputSendData{}, errors.New("something gone wrong")
		}
	}

	OutcomingPacket, err := s.S.ConverterJson.ConvertDataToJsonType(Dto.GrpcPacket{
		AesKeyData: encryptedAesKey,
		CipherData: plainText,
	})
	if err != nil {
		slog.Error("Marshal Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	slog.Info("finished the exchange")
	return &pb.OutputSendData{
		BytesOutput: OutcomingPacket,
		Error:       nil,
	}, nil
}
