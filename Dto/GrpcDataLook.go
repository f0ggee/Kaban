package Dto

import "time"

type GrpcDataIncomingPacket struct {
	Time             time.Time
	ServerName       []byte
	SignedServerName []byte
}
