package RedisUse

import "MasterServer_/DomainLevel"

type RedisConn struct {
	Client DomainLevel.RedisUse
}

func NewRedisConn(Tokes DomainLevel.RedisUse) *RedisConn {
	return &RedisConn{Client: Tokes}
}

func ConnectionRedis() *RedisConn {
	apps := &RedisUseStruct{}

	app := NewRedisConn(apps)
	return app
}
