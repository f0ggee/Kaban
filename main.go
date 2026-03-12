package main

import (
	"MasterServer_/DipendsInjective"
	"MasterServer_/Dto"
	InftarctionLevel "MasterServer_/InfrastructureLevel"
	"MasterServer_/InfrastructureLevel/GlobalProces"
	"MasterServer_/InfrastructureLevel/MemguardManipulation"
	pbRealization "MasterServer_/InfrastructureLevel/Proto"
	pbProtoFile "MasterServer_/InfrastructureLevel/Proto/protoFiles"
	"MasterServer_/InfrastructureLevel/RedisUse"
	"MasterServer_/InfrastructureLevel/keyInteration"
	"MasterServer_/InfrastructureLevel/rsaKeyManipulation"
	"net"

	"crypto/rand"
	"log"
	"log/slog"
	"os"
	"time"

	"MasterServer_/InfrastructureLevel/serverManagment"

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

	key := keyInteration.KeyInterationController{}
	RedisInteracting := RedisUse.RedisUseStruct{}
	//globalProcess := GlobalProces.ProcessController{}
	memguardManipulation := MemguardManipulation.MemgurdControl{}
	rsaKeyInteraction := rsaKeyManipulation.RsaKeyManipulation{}

	RsaAndMemoryInteract := DipendsInjective.NewRsaKeyManipulationWithRsaAndMemory(&memguardManipulation, &rsaKeyInteraction)

	ServerManaging := serverManagment.Pack2{RsaKey: &rsaKeyInteraction}
	serverMangementPack := serverManagment.NewServerManagement(ServerManaging)

	saz := GlobalProces.ProcessController{
		KeyInteracting:   &key,
		RedisInteracting: &RedisInteracting,
		ServerManagement: serverMangementPack,
	}

	Sa := GlobalProces.NewAnotherProcessController(saz)
	SwapRsaKey(*RsaAndMemoryInteract)
	if StartHandling(serverMangementPack, Sa) {
		return
	}

	ticker := time.NewTicker(InftarctionLevel.TimeForSwapping)
	defer ticker.Stop()

	go func() {

	}()

	for _ = range ticker.C {
		SwapRsaKey(*RsaAndMemoryInteract)
		slog.Info("Got a ticker")
		if StartHandling(serverMangementPack, Sa) {
			return
		}
	}
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		slog.Error("failed to listen:", err.Error())
		return
	}

	grpcServer := grpc.NewServer()

	pbProtoFile.RegisterSendingGettingServer(grpcServer, &pbRealization.HandlingRequestsForNewKey{})
	if err := grpcServer.Serve(lis); err != nil {
		slog.Error("failed to serve:", err.Error())
		return
	}

}

func StartHandling(serverMangementPack *serverManagment.ServerManagement, Sa *GlobalProces.AnotherProcessController) bool {
	for i := 1; i <= InftarctionLevel.ServersCount; i++ {
		ServerKey := serverMangementPack.GetServerKey(i)

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
