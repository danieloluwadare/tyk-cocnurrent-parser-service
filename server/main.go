package main

import (
	pb "github.com/danieloluwadare/tyk-cocnurrent-parser-service/gen/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	serverInMemoryService := NewInMemoryService()
	pb.RegisterInMemoryServiceServer(grpcServer, serverInMemoryService)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
