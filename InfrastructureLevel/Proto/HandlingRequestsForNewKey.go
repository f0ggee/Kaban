package Proto

import (
	"MasterServer_/DomainLevel"
	"MasterServer_/Dto"
	InftarctionLevel "MasterServer_/InfrastructureLevel"
	pb "MasterServer_/InfrastructureLevel/Proto/protoFiles"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/awnumar/memguard"
)

type HandlingRequestsForNewKey struct {
	S                DomainLevel.GrpcHandleData
	DescriptionGrpc  DomainLevel.Decryptor
	ServerManagement DomainLevel.GettingServersInfo
	EncryptionGrpc   DomainLevel.Encryption
}

type GrpcHandlerGettingNewKey struct {
	pb.UnimplementedSendingGettingServer
	S *HandlingRequestsForNewKey
}

func NewHandlingRequestsForNewKey(S *HandlingRequestsForNewKey) *GrpcHandlerGettingNewKey {
	return &GrpcHandlerGettingNewKey{S: S}
}

func (s *GrpcHandlerGettingNewKey) GettingNewKey(ctx context.Context, data *pb.InputSendData) (*pb.OutputSendData, error) {

	slog.Info("Start exchanging a key")
	if data == nil {
		slog.Error("Data was getting empty")
		return &pb.OutputSendData{}, errors.New("data is nil")
	}

	slog.Info("Here")
	if s.S.S.FindHash(data.SendData) {
		slog.Error("Hash has already been used")
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	s.S.S.SaveHash(data.SendData)

	DataIncomintLook := Dto.GrpcPacket{
		AesKeyData: nil,
		CipherData: nil,
	}

	err := json.Unmarshal(data.SendData, &DataIncomintLook)
	if err != nil {
		slog.Error("Unmarshal Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	DecryptedAesKey, err := s.S.DescriptionGrpc.DecryptAesKey(DataIncomintLook.AesKeyData)
	if err != nil {
		slog.Error("Decrypt Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	Data, err := s.S.DescriptionGrpc.DecryptCipherData(DecryptedAesKey, DataIncomintLook.CipherData)
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
	ResultComparingTime := s.S.S.CheckLifePacket(DataIntoPacket.Time)
	if ResultComparingTime {
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

	err = s.S.S.CheckSignature(DataIntoPacket.SignedServerName, ServerKey, DataIntoPacket.ServerName)
	if err != nil {
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	SignedKey, err := s.S.S.GenerateSignatureFromKey()
	if err != nil {
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	Dto.Keys.Mu.Lock()
	OutcomingDataJson, err := s.S.ServerManagement.ConvertDataToJsonType(Dto.GrpcOutcomingDataPacket{
		Sign:   SignedKey,
		RsaKey: Dto.Keys.NewPrivateKey.Bytes(),
		T1:     time.Now(),
		T2:     InftarctionLevel.TimeForSwapping,
	})
	if err != nil {
		slog.Error("Marshal Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}
	Dto.Keys.Mu.Unlock()

	AesKey, err := memguard.NewBufferFromReader(rand.Reader, 32)
	if err != nil {
		slog.Error("Generate AesKey Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}
	defer AesKey.Destroy()

	//var wg sync.WaitGroup
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//
	//}()
	slog.Info("Here")
	plainText, err := s.S.EncryptionGrpc.GrpcAesEncryption(OutcomingDataJson, AesKey.Bytes())
	if err != nil {
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	slog.Info("Here 2 ")
	EncryptedAesKey, err := s.S.EncryptionGrpc.EncryptAesKey(AesKey.Bytes(), ServerKey)
	if err != nil {
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	slog.Info("Here 3 ")
	OutcomingPacket, err := s.S.ServerManagement.ConvertDataToJsonType(Dto.GrpcPacket{
		AesKeyData: EncryptedAesKey,
		CipherData: plainText,
	})
	if err != nil {
		slog.Error("Marshal Error", "Error", err.Error())
		return &pb.OutputSendData{}, errors.New("something gone wrong")
	}

	slog.Info("Here 2 " + "")

	return &pb.OutputSendData{
		BytesOutput: OutcomingPacket,
		Error:       nil,
	}, nil
}
