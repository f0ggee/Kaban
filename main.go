package main

import (
	"log/slog"

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

}
