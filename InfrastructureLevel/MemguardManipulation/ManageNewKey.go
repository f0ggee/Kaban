package MemguardManipulation

import (
	"MasterServer_/Dto"

	"github.com/awnumar/memguard"
)

func (m *MemgurdControl) InstallingNewKey(NewRsaKey []byte) {

	Dto.Keys.NewPrivateKey = memguard.NewBuffer(len(NewRsaKey))
	Dto.Keys.NewPrivateKey.Copy(NewRsaKey)
}
