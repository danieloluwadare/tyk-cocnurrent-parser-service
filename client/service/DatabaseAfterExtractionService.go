package service

import (
	model2 "github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/repository"
	"github.com/kpango/glg"
	"sync"
)

type databaseAfterExtractionService struct {
	taskRepository repository.ITaskRepository
}

func NewDatabaseAfterExtractionService(taskRepository repository.ITaskRepository) *databaseAfterExtractionService {
	return &databaseAfterExtractionService{taskRepository: taskRepository}
}

//Execute saves  an array of TykTaskConfig to the database
func (d databaseAfterExtractionService) Execute(tykTaskConfigs []model2.TykTaskConfig) {

	for i, tykTaskConfig := range tykTaskConfigs {
		task := model2.Task{TykTaskConfig: tykTaskConfig}
		glg.Log("saving ==>", tykTaskConfig, "with index =>", i, "into the database")
		_, _ = d.taskRepository.Save(task)
	}
}
func (d databaseAfterExtractionService) ExecuteAsync(tykTaskConfigs []model2.TykTaskConfig, wg *sync.WaitGroup) {
	d.Execute(tykTaskConfigs)
	wg.Done()
}
