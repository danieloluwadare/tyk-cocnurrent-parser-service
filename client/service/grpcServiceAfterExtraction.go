package service

import (
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	"sync"
)

type grpcServiceAfterExtraction struct {
}

func (g grpcServiceAfterExtraction) Execute(tykJsonConfigArray []model.TykTaskConfig) {
	//TODO implement me
	panic("implement me")
}

func (g grpcServiceAfterExtraction) ExecuteAsync(tykTaskConfigs []model.TykTaskConfig, wg *sync.WaitGroup) {
	wg.Add(1)
	g.Execute(tykTaskConfigs)
	wg.Done()
}
