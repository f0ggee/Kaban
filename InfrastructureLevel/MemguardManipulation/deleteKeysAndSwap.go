package MemguardManipulation

import (
	"MasterServer_/Dto"

	"github.com/awnumar/memguard"
)

type MemgurdControl struct{}

func (m *MemgurdControl) SettingNewKey(NewRsaKey []byte) {

	Dto.Keys.NewPrivateKey = memguard.NewBuffer(len(NewRsaKey))
	Dto.Keys.NewPrivateKey.Copy(NewRsaKey)
}

func (m *MemgurdControl) DeleteKeysAndSwap() {

	Dto.Keys.OldPrivateKey.Destroy()

	Dto.Keys.OldPrivateKey = memguard.NewBuffer(Dto.Keys.NewPrivateKey.Size())
	Dto.Keys.OldPrivateKey.Copy(Dto.Keys.NewPrivateKey.Bytes())

	Dto.Keys.NewPrivateKey.Destroy()
}
