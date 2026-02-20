package keyInteration

import "MasterServer_/DomainLevel"

type ProcessContoller struct {
	Tokens DomainLevel.KeyInteracting
}

func NewProcessController(Token DomainLevel.KeyInteracting) *ProcessContoller {
	return &ProcessContoller{Tokens: Token}
}

func Connection() *ProcessContoller {

	apps := &KeyInterationController{}
	app := NewProcessController(apps)

	return app
}
