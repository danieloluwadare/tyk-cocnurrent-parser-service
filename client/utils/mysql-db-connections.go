package utils

import (
	"fmt"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kpango/glg"
)

func GetDataBaseConnection(dialect string, username string, password string, dbName string, dbHost string, dbPort string) *gorm.DB {

	dbUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)
	glg.Log(dbUrl)

	conn, err := gorm.Open(dialect, dbUrl)
	if err != nil {
		glg.Error(err)
	}
	db := conn

	return db
}

func InitiateModelMigration(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(
		&model.Task{},
	)
}
