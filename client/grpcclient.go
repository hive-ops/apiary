package client

import (
	"github.com/hive-ops/apiary/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ApiaryGRPCClient struct {
	pb.ApiaryServiceClient
	conn *grpc.ClientConn
}

func NewClient(address string) *ApiaryGRPCClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &ApiaryGRPCClient{conn: conn, ApiaryServiceClient: pb.NewApiaryServiceClient(conn)}
}

func (c *ApiaryGRPCClient) close() {
	err := c.conn.Close()
	if err != nil {
		return
	}
}
