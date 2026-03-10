package GrpcImplementation

import (
	"MasterServer_/InfrastructureLevel/serverManagment"
	"crypto/sha256"
	"time"
)

func (s *GrpcDataManagement) FindHash(data []byte) bool {
	sha := sha256.Sum256(data)

	if Time, ok := serverManagment.MappingHash[sha]; ok {

		if time.Now().After(Time) {
			return true
		}
	}

	return false

}
