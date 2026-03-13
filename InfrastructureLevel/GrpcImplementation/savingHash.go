package GrpcImplementation

import (
	"MasterServer_/DomainLevel"
	"MasterServer_/InfrastructureLevel/serverManagment"
	"crypto/sha256"
	"time"
)

type PackForGrpcImplementation struct {
	S DomainLevel.ServerDataManagement
}
type GrpcDataManagement struct {
	ServerDataManagement PackForGrpcImplementation
}

func (s *GrpcDataManagement) SaveHash(hash []byte) {
	hashing := sha256.Sum256(hash)
	serverManagment.MappingHash[hashing] = time.Now()
}
