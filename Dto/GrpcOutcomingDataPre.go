package Dto

import "time"

type GrpcOutcomingDataPacket struct {
	Sign   []byte `json:"Sign"`
	RsaKey []byte `json:"RsaKey"`
	T1     time.Time
	T2     time.Duration
}
