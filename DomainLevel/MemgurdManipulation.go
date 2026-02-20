package DomainLevel

type MemgurdManipulation interface {
	DeleteKeysAndSwap()
	SettingNewKey([]byte)
}
