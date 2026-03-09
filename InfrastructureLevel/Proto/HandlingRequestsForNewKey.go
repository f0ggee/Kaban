package Proto

import (
	"MasterServer_/DomainLevel"
	pb "MasterServer_/InfrastructureLevel/Proto/protoFiles"
	"context"
	"errors"
)

type HandlingRequestsForNewKey struct {
	pb.UnimplementedSendingGettingServer
	S DomainLevel.ServerDataManagement
}

func (s *HandlingRequestsForNewKey) GettingNewKey(ctx context.Context, data *pb.InputSendData) (*pb.OutputSendData, error) {

	if data == nil {
		return &pb.OutputSendData{}, errors.New("data is wrong")
	}

	if s.S.FindHash(data.SendData) {
		return &pb.OutputSendData{}, errors.New("something went wrong")
	}

}
