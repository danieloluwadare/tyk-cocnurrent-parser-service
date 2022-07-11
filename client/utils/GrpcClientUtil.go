package utils

import (
	pb "github.com/danieloluwadare/tyk-cocnurrent-parser-service/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GetInMemoryGrpcClient() (pb.InMemoryServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewInMemoryServiceClient(conn)
	return c, conn
}
