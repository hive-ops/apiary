package service

import (
	"context"
	"github.com/hive-ops/apiary/pb"
)

type ApiaryService struct {
	pb.UnsafeApiaryServiceServer
	Config *Config
	caches map[string]*Cache
}

func (a *ApiaryService) GetEntries(ctx context.Context, req *pb.GetEntriesRequest) (*pb.GetEntriesResponse, error) {
	c, ok := a.caches[req.Keyspace]

	if !ok {
		return &pb.GetEntriesResponse{
			Entries:  nil,
			NotFound: req.Keys,
		}, nil

	}

	entries := make([]*pb.Entry, 0)
	notFound := make([]string, 0)

	for _, key := range req.Keys {
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
func (a *ApiaryService) SetEntries(ctx context.Context, req *pb.SetEntriesRequest) (*pb.SetEntriesResponse, error) {
	c, ok := a.caches[req.Keyspace]
	if !ok {
		c = NewCache()
		a.caches[req.Keyspace] = c
	}

	successful := make([]string, 0)
	failed := make([]string, 0)

	for _, entry := range req.Entries {
		c.Set(entry.Key, []byte(entry.Value))
		successful = append(successful, entry.Key)
	}

	return &pb.SetEntriesResponse{
		Successful: successful,
		Failed:     failed,
	}, nil
}
func (a *ApiaryService) DeleteEntries(ctx context.Context, req *pb.DeleteEntriesRequest) (*pb.DeleteEntriesResponse, error) {
	c, ok := a.caches[req.Keyspace]

	if !ok {
		return &pb.DeleteEntriesResponse{
			Successful: nil,
			NotFound:   req.Keys,
			Failed:     nil,
		}, nil
	}

	successful := make([]string, 0)
	failed := make([]string, 0)

	for _, key := range req.Keys {
		c.Delete(key)
		successful = append(successful, key)
	}

	return &pb.DeleteEntriesResponse{
		Successful: successful,
		Failed:     failed,
	}, nil
}
func (a *ApiaryService) ClearEntries(ctx context.Context, req *pb.ClearEntriesRequest) (*pb.ClearEntriesResponse, error) {
	c, ok := a.caches[req.Keyspace]
	if !ok {
		return &pb.ClearEntriesResponse{
			Successful: false,
		}, nil
	}

	c.Clear()
	return &pb.ClearEntriesResponse{Successful: true}, nil
}

func NewApiaryService(config *Config) *ApiaryService {
	return &ApiaryService{
		Config: config,
		caches: make(map[string]*Cache),
	}
}

func NewApiaryServiceWithDefaultConfig() *ApiaryService {
	return NewApiaryService(NewDefaultConfig())
}
