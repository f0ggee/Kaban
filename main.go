package main

import (
	"MasterServer_/DipendsInjective"
	"MasterServer_/Dto"
	InftarctionLevel "MasterServer_/InfrastructureLevel"
	"MasterServer_/InfrastructureLevel/CryptoImpl/CryprtoGenerator"
	"MasterServer_/InfrastructureLevel/CryptoImpl/Decryptor"
	"MasterServer_/InfrastructureLevel/CryptoImpl/Encrypter"
	"MasterServer_/InfrastructureLevel/GlobalProces"
	"MasterServer_/InfrastructureLevel/Grpc/GrpcHandleData"
	"MasterServer_/InfrastructureLevel/Grpc/PacketValidation"
	"MasterServer_/InfrastructureLevel/MemguardManipulation"
	pbRealization "MasterServer_/InfrastructureLevel/Proto"
	pbProtoFile "MasterServer_/InfrastructureLevel/Proto/protoFiles"
	"MasterServer_/InfrastructureLevel/RedisUse"
	"MasterServer_/InfrastructureLevel/rsaKeyManipulation"
	"MasterServer_/InfrastructureLevel/serveManage/ConverterData"
	"MasterServer_/InfrastructureLevel/serveManage/GettingInfo"

	"net"

	"crypto/rand"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/awnumar/memguard"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("cannot load env file", err.Error())
		return

	}

	Dto.Keys.NewPrivateKey, _ = memguard.NewBufferFromReader(rand.Reader, 2048)
	Dto.Keys.OldPrivateKey, _ = memguard.NewBufferFromReader(rand.Reader, 2048)
	Dto.Keys.MasterServerKey = os.Getenv("OurKey")

}

func main() {

	handler := slog.New(slog.NewTextHandler(os.Stdout, nil))
	child := handler.With(
		"Time", time.Now(),
		"ServersCount", InftarctionLevel.ServersCount,
	)
	slog.SetDefault(child)

	memguard.CatchInterrupt()
	defer memguard.Purge()
	ConnectRedis := RedisUse.RedisConnect()
	defer ConnectRedis.Close()

	CryptoGeneret := CryprtoGenerator.CryprtoGenerating{}
	Decryper := Decryptor.Decrypting{}
	Encrypting := Encrypter.Encryption{}
	AnotherProcessController := GlobalProces.ControllingExchange{}
	GrpcHandlingData := GrpcHandleData.GrpcDataManagement{}
	PacketValidating := PacketValidation.ValidatePacketData{}
	MemguardConrol := MemguardManipulation.MemgurdControl{}
	GrpcHandlerGettinNewKey := pbRealization.GrpcHandlerGettingNewKey{
		UnimplementedSendingGettingServer: pbProtoFile.UnimplementedSendingGettingServer{},
		S:                                 nil,
	}
	redisConn := RedisUse.RedisUsing{
		Connect: ConnectRedis}
	RsaKeyControl := rsaKeyManipulation.RsaKeyManipulation{}
	ConvertData := ConverterData.ConvertingData{}
	ServerInfo := GettingInfo.SeverManage{}

	RsaAndMemoryInteract := DipendsInjective.NewRsaKeyManipulationWithRsaAndMemory(&memguardManipulation, &rsaKeyInteraction)
	GrpcHandlingRealization := GrpcHandleData.GrpcDataManagement{ServerDataManagement: GrpcHandleData.PackForGrpcImplementation{S: &managment}}

	ServerManaging := GettingInfo.DataManipulation{}
	serverMangementPack := GettingInfo.Pack2{RsaKey: &rsaKeyInteraction}

	saz := GlobalProces.ProcessController{
		KeyInteracting:   &key,
		RedisInteracting: &RedisInteracting,
		ServerManagement: &ServerManaging,
	}

	NewProccesForGrpcHandling := pbRealization.NewHandlingRequestsForNewKey(&pbRealization.HandlingRequestsForNewKey{
		S:                &GrpcHandlingRealization,
		DescriptionGrpc:  GrcpDecryptor,
		ServerManagement: &ServerManaging,
		EncryptionGrpc:   GrpcEncryptor,
	})
	Sa := GlobalProces.NewAnotherProcessController(saz)
	SwapRsaKey(*RsaAndMemoryInteract)
	if StartHandling(&ServerManaging, Sa) {
		return
	}

	ticker := time.NewTicker(InftarctionLevel.TimeForSwapping)
	defer ticker.Stop()

	go func() {
		lis, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			slog.Error("failed to listen:", err.Error())
			return
		}

		grpcServer := grpc.NewServer()

		pbProtoFile.RegisterSendingGettingServer(grpcServer, NewProccesForGrpcHandling)
		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("failed to serve:", err.Error())
			return
		}

	}()

	for _ = range ticker.C {
		SwapRsaKey(*RsaAndMemoryInteract)
		slog.Info("Got a ticker")
		if StartHandling(GettingInfo.NewServerManipulation(serverMangementPack), Sa) {
			return
		}
	}
	slog.Info("Ep")

}

func StartHandling(serverMangementPack *GettingInfo.DataManipulation, Sa *GlobalProces.ControllingExchange) bool {
	for i := 1; i <= InftarctionLevel.ServersCount; i++ {
		ServerKey := serverMangementPack.GetServerKey(i)
		if ServerKey == nil {
			slog.Error("ServerKey is nil")
			continue
		}

		ServerName := serverMangementPack.GetServerName(i)
		if ServerName == "" {
			slog.Error("we can't find the server", "ServerNumber", i)
			continue
		}

		err := Sa.HandlingAndSendData(ServerKey, Dto.Keys.NewPrivateKey.Bytes(), ServerName)
		if err != nil {
			continue
		}

	}
	return false
}

func SwapRsaKey(RsaKey DipendsInjective.RsaKeyManipulationWithRsaAndMemory) {
	Dto.Keys.Mu.Lock()
	slog.Info("Swaping starts")

	TemporallySaving := memguard.NewBufferFromBytes(RsaKey.RsaKey.GenerateRsaKey())
	defer TemporallySaving.Destroy()

	RsaKey.KeyAndMemory.DeleteKeysAndSwap()

	RsaKey.KeyAndMemory.SettingNewKey(TemporallySaving.Bytes())
	log.Println("Swaping End")

	Dto.Keys.Mu.Unlock()
}
