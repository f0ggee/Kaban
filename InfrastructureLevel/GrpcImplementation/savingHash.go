package GrpcImplementation

import (
	"MasterServer_/DomainLevel"
	"MasterServer_/InfrastructureLevel/serverManagment"
	"crypto/sha256"
	"log/slog"
	"time"
)

type PackGrpc struct {
	S DomainLevel.ServerDataManagement
}
type GrpcDataManagement struct {
	ServerDataManagement DomainLevel.ServerDataManagement
}

func (s *GrpcDataManagement) SaveHash(hash []byte) {
	hashing := sha256.Sum256(hash)
	serverManagment.MappingHash[hashing] = time.Now()
}
