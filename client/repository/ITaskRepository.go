package repository

import (
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
)

type ITaskRepository interface {
	Save(task model.Task) (*model.Task, error)
}
