package PacketValidation

import (
	"log/slog"
	"time"
)

func (s *ValidatePacketData) CheckLifePacket(duration time.Time) bool {

	if ResultComparing := duration.Compare(time.Now()); ResultComparing == 1 {
		slog.Error("Packet's expired", "time packet's life", duration.Hour())
		return true
	}
	return false
}
