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

	CryptoGenerate := CryprtoGenerator.CryprtoGenerating{}
	Decryper := Decryptor.Decrypting{}
	Encrypting := Encrypter.Encryption{}
	AnotherProcessController := GlobalProces.ControllingExchange{}
	GrpcHandlingData := GrpcHandleData.GrpcDataManagement{}
	PacketValidating := PacketValidation.ValidatePacketData{}
	MemguardControl := MemguardManipulation.MemgurdControl{}
	ServerInfo := GettingInfo.SeverManage{}
	redisConn := RedisUse.RedisUsing{Connect: ConnectRedis}
	RsaKeyControl := rsaKeyManipulation.RsaKeyManipulation{}
	ConvertData := ConverterData.ConvertingData{}

	//GrpcHandlerGettinNewKey := pbRealization.GrpcHandlerGettingNewKey{
	//	UnimplementedSendingGettingServer: pbProtoFile.UnimplementedSendingGettingServer{},
	//	S: pbRealization.HandlingRequestsForNewKey{
	//		Grpc:             &GrpcHandlingData,
	//		ServerManagement: &ServerInfo,
	//		Encryption:       &Encrypting,
	//		Checker:          &PacketValidating,
	//		CryptoGenerating: &CryptoGenerate,
	//		ConverterJson:    &ConvertData,
	//	},
	//}

	Injective1 := DipendsInjective.NewRsaKeyManipulationWithRsaAndMemory(&MemguardControl, &RsaKeyControl)
	Injective2 := GlobalProces.NewAnotherProcessController(GlobalProces.ProcessController{
		Cryptos:          &Encrypting,
		CryptoGen:        &CryptoGenerate,
		RedisInteracting: &redisConn,
		ServerManagement: &ServerInfo,
	})

	//injective3Grpc := pbRealization.NewGrpcHandlerGettingNewKey(&pbProtoFile.UnimplementedSendingGettingServer{}, &pbRealization.HandlingRequestsForNewKey{
	//
	//})
	SwapRsaKey(*Injective1)
	if StartHandling(&ServerInfo, Injective2) {
		return
	}

	ticker := time.NewTicker(InftarctionLevel.TimeForSwapping)
	defer ticker.Stop()

	go func() {
		lis, err := net.Listen("tcp", ":8081")
		if err != nil {
			slog.Error("failed to listen:", err.Error())
			return
		}

		grpcServer := grpc.NewServer()

		pbProtoFile.RegisterSendingGettingServer(grpcServer, &pbRealization.GrpcHandlerGettingNewKey{
			UnimplementedSendingGettingServer: pbProtoFile.UnimplementedSendingGettingServer{},
			S: pbRealization.HandlingRequestsForNewKey{
				Grpc:             &GrpcHandlingData,
				ServerManagement: &ServerInfo,
				Encryption:       &Encrypting,
				Checker:          &PacketValidating,
				CryptoGenerating: &CryptoGenerate,
				ConverterJson:    &ConvertData,
				Decrypting:       &Decryper,
			},
		})
		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("failed to serve:", err.Error())
			return
		}

	}()

	for _ = range ticker.C {
		slog.Info("We got the tick")
		SwapRsaKey(*Injective1)
		slog.Info("Got a ticker")
		if StartHandling(&ServerInfo, &AnotherProcessController) {
			return
		}
		slog.Info("Finished the exchange")
	}

}

func StartHandling(serverManagementPack *GettingInfo.SeverManage, Sa *GlobalProces.ControllingExchange) bool {
	for i := 1; i <= InftarctionLevel.ServersCount; i++ {
		ServerKey := serverManagementPack.GetServerKey(i)
		if ServerKey == nil {
			slog.Error("ServerKey is nil")
			continue
		}

		ServerName := serverManagementPack.GetServerName(i)
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

	slog.Info("Swaping starts")

	TemporallySaving := memguard.NewBufferFromBytes(RsaKey.RsaKey.GenerateRsaKey())
	defer TemporallySaving.Destroy()

	Dto.Keys.Mu.Lock()
	RsaKey.KeyAndMemory.SwapingOldKey()
	RsaKey.KeyAndMemory.InstallingNewKey(TemporallySaving.Bytes())
	log.Println("Swaping End")
	Dto.Keys.Mu.Unlock()
}
