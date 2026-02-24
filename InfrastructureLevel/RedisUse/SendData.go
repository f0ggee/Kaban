package RedisUse

import (
	"MasterServer_/Dto"
	"log"
)

type RedisUseStruct struct{}

//func (*RedisUseStruct) SendData(data *Dto.RedisDataLooksLike, serverName string) error {
//
//	connect := RedisConnect()
//	defer connect.Close()
//
//	err := connect.HSet(context.Background(), serverName, data).Err()
//	if err != nil {
//		slog.Error("RedisUseStruct.SendData()", "Error", err)
//		return err
//	}
//	return nil
//}

func (*RedisUseStruct) SendData(data *Dto.RedisDataLooksLike, serverName string) error {

	log.Println("We emulate working redis")
	return nil
}
