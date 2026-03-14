package GrpcHandleData

import (
	"MasterServer_/InfrastructureLevel/serveManage/GettingInfo"
	"crypto/sha256"
	"time"
)

type GrpcDataManagement struct {
}

func (s *GrpcDataManagement) SaveHash(hash []byte) {
	hashing := sha256.Sum256(hash)
	GettingInfo.MappingHash[hashing] = time.Now()
}
