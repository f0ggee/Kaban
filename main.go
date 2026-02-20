package main

import (
	"MasterServer_/DipendsInjective"
	"MasterServer_/Dto"
	"MasterServer_/InfrastructureLevel"
	"MasterServer_/InfrastructureLevel/GlobalProces"
	"MasterServer_/InfrastructureLevel/MemguardManipulation"
	"MasterServer_/InfrastructureLevel/RedisUse"
	"MasterServer_/InfrastructureLevel/keyInteration"
	"MasterServer_/InfrastructureLevel/rsaKeyManipulation"
	"crypto/rand"
	"log"
	"log/slog"
	"os"
	"sync"
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
	ServerManaging := serverManagment.Pack2{}
	globalProcess := GlobalProces.ProcessController{}
	memguardManipulation := MemguardManipulation.MemgurdControl{}
	rsaKeyInteraction := rsaKeyManipulation.RsaKeyManipulation{}

	RsaAndMemoryInteract := DipendsInjective.NewRsaKeyManipulationWithRsaAndMemory(memguardManipulation, rsaKeyInteraction)

	//connect := InftarctionLevel.NewCollectPacks(&key, &RedisInteracting, ServerManaging, globalProcess)
	s := GlobalProces.NewProcessController(&key, &RedisInteracting, globalProcess, ServerManaging)
	SwapRsaKey(*RsaAndMemoryInteract)

	for i := 1; i <= ServersCount; i++ {
		ServerKey := s.ServerManagement.GetServerKey(i)

	}
}

func SwapRsaKey(RsaKey DipendsInjective.RsaKeyManipulationWithRsaAndMemory) {
	Dto.Keys.Mu.Lock()
	log.Println("Swaping starts")

	TemporallySaving := memguard.NewBufferFromBytes(RsaKey.RsaKey.GenerateRsaKey())
	defer TemporallySaving.Destroy()

	RsaKey.KeyAndMemory.DeleteKeysAndSwap()

	RsaKey.KeyAndMemory.SettingNewKey(TemporallySaving.Bytes())
	log.Println("Swaping starts")

	Dto.Keys.Mu.Unlock()
}
