package server

import (
	"context"
	"fmt"
	pb "github.com/hive-ops/apiary/pb/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var caches = make(map[string]*Cache)

type ApiaryGRPCServer struct {
	pb.UnsafeApiaryServiceServer
	Config *Config
}

func (s *ApiaryGRPCServer) getCache(ctx context.Context, keyspace *pb.Keyspace) (*Cache, bool) {
	c, ok := caches[keyspace.GetKeyspaceRef()]
	return c, ok
}

func (s *ApiaryGRPCServer) GetEntries(ctx context.Context, cmd *pb.GetEntriesCommand) (*pb.GetEntriesResponse, error) {
	c, ok := s.getCache(ctx, cmd.Keyspace)
	if !ok {
		return &pb.GetEntriesResponse{
			Entries:  nil,
			NotFound: cmd.Keys,
		}, nil
	}

	entries := make([]*pb.Entry, 0)
	notFound := make([]string, 0)

	for _, key := range cmd.Keys {
		val, err := c.Get(key)
		if err != nil {
			notFound = append(notFound, key)
			continue
		}
		entries = append(entries, &pb.Entry{
			Key:   key,
			Value: string(val),
		})
	}

	return &pb.GetEntriesResponse{
		Entries:  entries,
		NotFound: notFound,
	}, nil
}

func (s *ApiaryGRPCServer) SetEntries(ctx context.Context, cmd *pb.SetEntriesCommand) (*pb.SetEntriesResponse, error) {
	c, ok := s.getCache(ctx, cmd.Keyspace)
	if !ok {
		c = NewCache()
		caches[cmd.Keyspace.GetKeyspaceRef()] = c
	}

	successful := make([]string, 0)
	failed := make([]string, 0)

	for _, entry := range cmd.Entries {
		c.Set(entry.Key, []byte(entry.Value))
		successful = append(successful, entry.Key)
	}

	return &pb.SetEntriesResponse{
		Successful: successful,
		Failed:     failed,
	}, nil
}

func (s *ApiaryGRPCServer) DeleteEntries(ctx context.Context, cmd *pb.DeleteEntriesCommand) (*pb.DeleteEntriesResponse, error) {
	c, ok := s.getCache(ctx, cmd.Keyspace)
	if !ok {
		return &pb.DeleteEntriesResponse{
			Successful: nil,
			NotFound:   cmd.Keys,
			Failed:     nil,
		}, nil
	}

	successful := make([]string, 0)
	failed := make([]string, 0)

	for _, key := range cmd.Keys {
		c.Delete(key)
		successful = append(successful, key)
	}

	return &pb.DeleteEntriesResponse{
		Successful: successful,
		Failed:     failed,
	}, nil
}

func (s *ApiaryGRPCServer) ClearEntries(ctx context.Context, cmd *pb.ClearEntriesCommand) (*pb.ClearEntriesResponse, error) {
	c, ok := s.getCache(ctx, cmd.Keyspace)
	if !ok {
		return &pb.ClearEntriesResponse{
			Successful: false,
		}, nil
	}

	c.Clear()
	return &pb.ClearEntriesResponse{Successful: true}, nil
}

func NewApiaryServer(config *Config) *ApiaryGRPCServer {
	return &ApiaryGRPCServer{
		Config: config,
	}
}

func StartApiaryServer() {

	config := LoadConfig("apiary.yaml")

	address := fmt.Sprintf("%s:%v", config.IP, config.Port)
	fmt.Println(fmt.Sprintf("Starting gRPC server on %s", address))

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := &ApiaryGRPCServer{
		Config: config,
	}

	pb.RegisterApiaryServiceServer(grpcServer, s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
