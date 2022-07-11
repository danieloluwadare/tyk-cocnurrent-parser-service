package model

import (
	"github.com/jinzhu/gorm"
)

//Task Entity  containing basic fields
//swagger:model user-model
type Task struct {
	gorm.Model
	TykTaskConfig
}
