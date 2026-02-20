package DipendsInjective

import "MasterServer_/DomainLevel"

type RsaKeyManipulationWithRsaAndMemory struct {
	KeyAndMemory DomainLevel.MemgurdManipulation
	RsaKey       DomainLevel.RsaKeyManipulation
}

func NewRsaKeyManipulationWithRsaAndMemory(keyAndMemory DomainLevel.MemgurdManipulation, rsaKey DomainLevel.RsaKeyManipulation) *RsaKeyManipulationWithRsaAndMemory {
	return &RsaKeyManipulationWithRsaAndMemory{KeyAndMemory: keyAndMemory, RsaKey: rsaKey}
}
