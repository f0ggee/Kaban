package Proto

import (
	"MasterServer_"
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
	pb.UnimplementedSendingGettingServer
	S                DomainLevel.GrpcHandleData
	DescriptionGrpc  DomainLevel.GrpcDecryptor
	ServerManagement DomainLevel.ServerDataManagement
	EncryptionGrpc   DomainLevel.GrpcEncryptor
}

func (s *HandlingRequestsForNewKey) GettingNewKey(ctx context.Context, data *pb.InputSendData) (*pb.OutputSendData, error) {

	slog.Info("Start exchanging a key")
	if data == nil {
		slog.Error("Data was getting empty")
		return &pb.OutputSendData{}, errors.New("data is nil")
	}

	if s.S.FindHash(data.SendData) {
		slog.Error("Hash has already been used")
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	s.S.SaveHash(data.SendData)

	DataIncomintLook := Dto.GrpcPacket{
		AesKeyData: nil,
		CipherData: nil,
	}

	err := json.Unmarshal(data.SendData, &DataIncomintLook)
	if err != nil {
		slog.Error("Unmarshal Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	DecryptedAesKey, err := s.DescriptionGrpc.DecryptIncomingAesKey(DataIncomintLook.AesKeyData)
	if err != nil {
		slog.Error("Decrypt Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	Data, err := s.DescriptionGrpc.DecryptCipherData(DecryptedAesKey, DataIncomintLook.CipherData)
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

	if DataIntoPacket.Time.Before(time.Now()) {
		slog.Error("Time is too early")
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

	err = s.S.CheckSignature(DataIntoPacket.SignedServerName, ServerKey, DataIntoPacket.ServerName)
	if err != nil {
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	SignedKey, err := s.S.GenerateSignature()
	if err != nil {
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}
	JsonData, err := json.Marshal(Dto.GrpcOutcomingDataPacket{
		Sign:   SignedKey,
		RsaKey: Dto.Keys.NewPrivateKey.Bytes(),
		T1:     time.Now(),
		T2:     InftarctionLevel.TimeForSwapping,
	})
	if err != nil {
		slog.Error("Marshal Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	AesKey, err := memguard.NewBufferFromReader(rand.Reader, 2048)
	if err != nil {
		slog.Error("Generate AesKey Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}
	defer AesKey.Destroy()

	plainText, err := s.EncryptionGrpc.GrpcAesEncryption(JsonData, AesKey.Bytes())
	if err != nil {
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}


	EncryptedAesKey,err := s.EncryptionGrpc.
}
