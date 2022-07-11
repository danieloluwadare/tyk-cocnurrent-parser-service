// De-messenger
//
// De-messenger
//
//     Schemes: http,https,127.0.0.1
//     Host: localhost:8900
//     Version: 0.0.1
//	   BasePath: /api
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/repository"
	service2 "github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/service"
	utils2 "github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/utils"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	utils2.InitializeLog()
	dialect := os.Getenv("DB_CONNECTION")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	//Get database connection
	dbConnection := utils2.GetDataBaseConnection(dialect, username, password, dbName, dbHost, dbPort)
	//Migrate all models
	utils2.InitiateModelMigration(dbConnection)

	taskRepository := repository.NewTaskRepository(dbConnection)
	databaseAfterExtractionService := service2.NewDatabaseAfterExtractionService(taskRepository)

	fileParser := service2.NewFileParser()
	//tykTaskConfigs, _ := fileParser.Parse()

	initializer := service2.NewInitializer(fileParser, databaseAfterExtractionService, databaseAfterExtractionService)
	initializer.Execute()

}
