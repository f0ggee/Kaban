package GlobalProces

import "MasterServer_/DomainLevel"

type ProcessController struct {
	Cryptos          DomainLevel.Encryption
	CryptoGen        DomainLevel.CryptoGenerator
	RedisInteracting DomainLevel.RedisUse
	ServerManagement DomainLevel.GettingServersInfo
}

type ControllingExchange struct {
	E ProcessController
}

func NewAnotherProcessController(e ProcessController) *ControllingExchange {
	return &ControllingExchange{E: e}
}
