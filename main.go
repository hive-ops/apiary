package main

import (
	"fmt"
	pb "github.com/hive-ops/apiary/pb/proto"
	"github.com/hive-ops/apiary/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	port := "2468"

	fmt.Println(fmt.Sprintf("Starting gRPC server on port %s...", port))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := &server.ApiaryGRPCServer{}
	pb.RegisterApiaryServiceServer(grpcServer, s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
