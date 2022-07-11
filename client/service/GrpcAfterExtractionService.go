package service

import (
	"context"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	pb "github.com/danieloluwadare/tyk-cocnurrent-parser-service/gen/proto"
	"github.com/kpango/glg"
	"log"
	"sync"
)

type grpcAfterExtractionService struct {
	ctx                context.Context
	inmemoryGrpcClient pb.InMemoryServiceClient
}

func NewGrpcAfterExtractionService(ctx context.Context, inmemoryGrpcClient pb.InMemoryServiceClient) grpcAfterExtractionService {
	return grpcAfterExtractionService{ctx: ctx, inmemoryGrpcClient: inmemoryGrpcClient}
}

func (g grpcAfterExtractionService) Execute(tykTaskConfigs []model.TykTaskConfig) {
	//ctx, _ := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	for i, tykTaskConfig := range tykTaskConfigs {
		tykServerRequest := &pb.TykServerRequest{Name: tykTaskConfig.Name, Description: tykTaskConfig.Description}
		glg.Log("calling grpc service for task ==>", tykTaskConfig, "with index =>", i)
		r, err := g.inmemoryGrpcClient.SaveRequest(g.ctx, tykServerRequest)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		glg.Log("Response from Grpc:", r.GetName())
	}

	//panic("implement me")
}

func (g grpcAfterExtractionService) ExecuteAsync(tykTaskConfigs []model.TykTaskConfig, wg *sync.WaitGroup) {
	wg.Add(1)
	g.Execute(tykTaskConfigs)
	wg.Done()
}
