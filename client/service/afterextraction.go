package service

import (
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	"sync"
)

type AfterExtraction interface {
	Execute(tykTaskConfigs []model.TykTaskConfig)
	ExecuteAsync(tykTaskConfigs []model.TykTaskConfig, wg *sync.WaitGroup)
}
