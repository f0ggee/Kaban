package DomainLevel

import "MasterServer_/Dto"

type RedisUse interface {
	SendData(*Dto.RedisDataLooksLike, string) error
}
