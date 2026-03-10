package Dto

import "time"

type GrpcDataLookPacket struct {
	Time             time.Time
	ServerName       []byte
	SignedServerName []byte
}
