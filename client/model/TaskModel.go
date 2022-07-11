package model

//Task Entity  containing basic fields
//swagger:model user-model
type Task struct {
	ID uint `gorm:"primary_key"`
	TykTaskConfig
}
