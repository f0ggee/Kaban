package PacketValidation

import (
	"MasterServer_/InfrastructureLevel/serveManage/GettingInfo"
	"time"
)

func (s *ValidatePacketData) FindHash(data [32]byte) bool {

	if Time, ok := GettingInfo.MappingHash[data]; ok {

		if time.Now().After(Time) {
			return true
		}
	}

	return false

}
