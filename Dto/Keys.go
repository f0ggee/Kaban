package Dto

import (
	"sync"

	"github.com/awnumar/memguard"
)

var Keys struct {
	Mu            sync.RWMutex
	NewPrivateKey *memguard.LockedBuffer
	OldPrivateKey *memguard.LockedBuffer
}
