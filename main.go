package main

import (
	"fmt"
	apiaryv1 "github.com/hive-ops/apiary/pb/apiary/v1"
	"github.com/hive-ops/apiary/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartApiaryServer() {

	config := service.LoadConfig("apiary.yaml")

	address := fmt.Sprintf("%s:%v", config.IP, config.Port)
	fmt.Println(fmt.Sprintf("Starting gRPC server on %s", address))

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := service.NewApiaryService(config)

	apiaryv1.RegisterApiaryServiceServer(grpcServer, s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	StartApiaryServer()
}
