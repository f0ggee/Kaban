package RedisUse

import (
	"MasterServer_/Dto"
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type RedisUsing struct {
	Connect *redis.Client
}

func (s *RedisUsing) SendData(data *Dto.RedisDataLooksLike, serverName string) error {

	err := s.Connect.HSet(context.Background(), serverName, data).Err()
	if err != nil {
		slog.Error("RedisUsing.SendData()", "Error", err)
		return err
	}
	return nil
}

//func (*RedisUsing) SendData(data *Dto.RedisDataLooksLike, serverName string) error {
//
//	log.Println("We emulate working redis")
//	return nil
//}
