package DomainLevel

type MemgurdManipulation interface {
	DeleteKeysAndSwap()
	SettingNewKey([]byte)
	SayHi() string
}
