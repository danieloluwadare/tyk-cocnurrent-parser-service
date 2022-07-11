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
	"context"
	"fmt"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/repository"
	service2 "github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/service"
	utils2 "github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/utils"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"os"
	"time"
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

	//Create context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	inMemoryGrpcClient, conn := utils2.GetInMemoryGrpcClient()
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	taskRepository := repository.NewTaskRepository(dbConnection)

	databaseAfterExtractionService := service2.NewDatabaseAfterExtractionService(taskRepository)
	grpcAfterExtractionService := service2.NewGrpcAfterExtractionService(ctx, inMemoryGrpcClient)

	fileParser := service2.NewFileParser()
	initializer := service2.NewInitializer(fileParser, databaseAfterExtractionService, grpcAfterExtractionService)
	initializer.Execute()
}
