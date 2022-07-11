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
	pb "github.com/danieloluwadare/tyk-cocnurrent-parser-service/gen/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	inMemoryGrpcClient, conn := utils2.GetInMemoryGrpcClient()
	taskRepository := repository.NewTaskRepository(dbConnection)

	databaseAfterExtractionService := service2.NewDatabaseAfterExtractionService(taskRepository)
	grpcAfterExtractionService := service2.NewGrpcAfterExtractionService(ctx, inMemoryGrpcClient)
	fileParser := service2.NewFileParser()
	//tykTaskConfigs, _ := fileParser.Parse()

	initializer := service2.NewInitializer(fileParser, databaseAfterExtractionService, grpcAfterExtractionService)
	initializer.Execute()

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	//getCLient()

}
func getCLient() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewInMemoryServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SaveRequest(ctx, &pb.TykServerRequest{Name: "Test Name", Description: "Alright"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetName())
}
