package keyInteration

import "MasterServer_/DomainLevel"

type ProcessContoller struct {
	Tokens DomainLevel.KeyInteration
}

func NewProcessController(Token DomainLevel.KeyInteration) *ProcessContoller {
	return &ProcessContoller{Tokens: Token}
}

func Connection() *ProcessContoller {

	apps := &KeyInterationController{}
	app := NewProcessController(apps)

	return app
}
