package MemguardManipulation

import (
	"MasterServer_/Dto"

	"github.com/awnumar/memguard"
)

func (m *MemgurdControl) SwapingOldKey() {

	Dto.Keys.OldPrivateKey.Destroy()

	Dto.Keys.OldPrivateKey = memguard.NewBuffer(Dto.Keys.NewPrivateKey.Size())
	Dto.Keys.OldPrivateKey.Copy(Dto.Keys.NewPrivateKey.Bytes())

	Dto.Keys.NewPrivateKey.Destroy()
}
