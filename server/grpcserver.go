package server

import (
	"context"
	pb "github.com/hive-ops/apiary/pb/proto"
)

var c = NewCache()

type ApiaryGRPCServer struct {
	pb.UnsafeApiaryServiceServer
}

func (s *ApiaryGRPCServer) GetEntries(ctx context.Context, cmd *pb.GetEntriesCommand) (*pb.GetEntriesResponse, error) {
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
			Value: val,
		})
	}

	return &pb.GetEntriesResponse{
		Entries:  entries,
		NotFound: notFound,
	}, nil
}

func (s *ApiaryGRPCServer) SetEntries(ctx context.Context, cmd *pb.SetEntriesCommand) (*pb.SetEntriesResponse, error) {

	successful := make([]string, 0)
	failed := make([]string, 0)

	for _, entry := range cmd.Entries {
		c.Set(entry.Key, entry.Value)
		successful = append(successful, entry.Key)
	}

	return &pb.SetEntriesResponse{
		Successful: successful,
		Failed:     failed,
	}, nil
}

func (s *ApiaryGRPCServer) DeleteEntries(ctx context.Context, cmd *pb.DeleteEntriesCommand) (*pb.DeleteEntriesResponse, error) {
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
