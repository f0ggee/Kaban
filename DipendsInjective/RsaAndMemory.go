package DipendsInjective

import "MasterServer_/DomainLevel"

type RsaKeyManipulationWithRsaAndMemory struct {
	KeyAndMemory DomainLevel.MudguardManageKeys
	RsaKey       DomainLevel.RsaKeyManipulation
}

func NewRsaKeyManipulationWithRsaAndMemory(keyAndMemory DomainLevel.MudguardManageKeys, rsaKey DomainLevel.RsaKeyManipulation) *RsaKeyManipulationWithRsaAndMemory {
	return &RsaKeyManipulationWithRsaAndMemory{KeyAndMemory: keyAndMemory, RsaKey: rsaKey}
}
