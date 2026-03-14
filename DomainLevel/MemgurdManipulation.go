package DomainLevel

type MudguardManageKeys interface {
	SwapingOldKey()
	InstallingNewKey([]byte)
}
