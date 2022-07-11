package repository

import (
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	"github.com/jinzhu/gorm"
	"github.com/kpango/glg"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) taskRepository {
	return taskRepository{db: db}
}

func (t taskRepository) Save(task model.Task) (*model.Task, error) {
	// Create failed, do something e.g. return, panic etc.
	if dbc := t.db.Create(&task); dbc.Error != nil {
		return nil, dbc.Error
	}
	//u.db.Where("email = ?", user.Email).First(&newUser)
	glg.Log("task created =>", task)

	//return &user
	return &task, nil
}
