package InftarctionLevel

import (
	"MasterServer_/DomainLevel"
)

type CollectPacks struct {
	KeyInteracting    DomainLevel.KeyInteracting
	RedisInteracting  DomainLevel.RedisUse
	ServerManaging    DomainLevel.ServerDataManagement
	StartTransferring DomainLevel.Process
}

func NewCollectPacks(keyInteracting DomainLevel.KeyInteracting, redisInteracting DomainLevel.RedisUse, serverManaging DomainLevel.ServerDataManagement, startTransferring DomainLevel.Process) *CollectPacks {
	return &CollectPacks{KeyInteracting: keyInteracting, RedisInteracting: redisInteracting, ServerManaging: serverManaging, StartTransferring: startTransferring}
}
