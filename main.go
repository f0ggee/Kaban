package main

import (
	"MasterServer_/DipendsInjective"
	"MasterServer_/Dto"
	"MasterServer_/InfrastructureLevel/GlobalProces"
	"MasterServer_/InfrastructureLevel/MemguardManipulation"
	"MasterServer_/InfrastructureLevel/RedisUse"
	"MasterServer_/InfrastructureLevel/keyInteration"
	"MasterServer_/InfrastructureLevel/rsaKeyManipulation"
	"crypto/rand"
	"log"
	"log/slog"
	"os"
	"time"

	"MasterServer_/InfrastructureLevel/serverManagment"

	"github.com/awnumar/memguard"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("cannot load env file", err.Error())
		return

	}

	Dto.Keys.NewPrivateKey, _ = memguard.NewBufferFromReader(rand.Reader, 2048)
	Dto.Keys.OldPrivateKey, _ = memguard.NewBufferFromReader(rand.Reader, 2048)

}

const ServersCount = 1

func main() {

	handler := slog.New(slog.NewTextHandler(os.Stdout, nil))
	child := handler.With(
		"Time", time.Now(),
		"ServersCount", ServersCount,
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

	for i := 1; i <= ServersCount; i++ {
		ServerKey := serverMangementPack.GetServerKey(i)
		if ServerKey == nil {
			return
		}

		ServerName := serverMangementPack.GetServerName(i)

		//ServerName := s.ServerManagement.GetServerName(i)
		if ServerName == "" {
			slog.Info("we can't find the server", "ServerNumber", i)
			continue
		}
		slog.Info("we found the server", "ServerNumber", i, "ServerName", ServerName)

		slog.Info("DtoData", "Data", Dto.Keys.NewPrivateKey.String())

		err := Sa.HandlingAndSendData(ServerKey, Dto.Keys.NewPrivateKey.Bytes(), ServerName)
		if err != nil {
			return
		}
	}

	log.Println("Server has done work", "Bool", true)

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
