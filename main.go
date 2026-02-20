package main

import (
	"MasterServer_/InfrastructureLevel"
	"MasterServer_/InfrastructureLevel/GlobalProces"
	"MasterServer_/InfrastructureLevel/RedisUse"
	"MasterServer_/InfrastructureLevel/keyInteration"
	"log/slog"

	"MasterServer_/InfrastructureLevel/serverManagment"
	"github.com/awnumar/memguard"
	"github.com/joho/godotenv"
)

var Keys struct {
	NewPrivateKey *memguard.LockedBuffer
	OldPrivateKey *memguard.LockedBuffer
}

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("cannot load env file", err.Error())
		return

	}

	Keys.NewPrivateKey = memguard.NewBufferFromBytes([]byte("NewPrivateKey"))
	Keys.OldPrivateKey = memguard.NewBufferFromBytes([]byte("OldPrivateKey"))

}
func main() {
	memguard.CatchInterrupt()
	defer memguard.Purge()
	key := keyInteration.KeyInterationController{}
	RedisInteracting := RedisUse.RedisUseStruct{}
	//ServerManaging := serverManagment.ServerManagement{}
	globalProcess := GlobalProces.ProcessController{}
	//connect := InftarctionLevel.NewCollectPacks(&key, &RedisInteracting, ServerManaging, globalProcess)
	s := GlobalProces.ConnectProcessController(&RedisInteracting, &key, globalProcess)

	err := s.Process.HandlingAndSendData()

}
