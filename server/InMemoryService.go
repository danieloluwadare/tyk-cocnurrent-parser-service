package main

import (
	"context"
	pb "github.com/danieloluwadare/tyk-cocnurrent-parser-service/gen/proto"
	"log"
)

type inMemoryService struct {
	requests []*pb.TykServerRequest
	pb.UnimplementedInMemoryServiceServer
}

func NewInMemoryService() *inMemoryService {
	requests := make([]*pb.TykServerRequest, 0)
	return &inMemoryService{requests: requests}
	//&server{}
}

func (i *inMemoryService) SaveRequest(ctx context.Context, in *pb.TykServerRequest) (*pb.TykServerResponse, error) {
	log.Printf("Received: %v", in.GetName())
	i.requests = append(i.requests, in)
	return &pb.TykServerResponse{Name: "Successful " + in.GetName()}, nil
}
