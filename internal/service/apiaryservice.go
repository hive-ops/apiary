package service

import (
	"context"
	"fmt"
	apiaryv1 "github.com/hive-ops/apiary/pb/apiary/v1"
)

type ApiaryService struct {
	apiaryv1.UnsafeApiaryServiceServer
	Config *Config
	caches map[string]*Cache
}

func (a *ApiaryService) GetEntries(ctx context.Context, req *apiaryv1.GetEntriesRequest) (*apiaryv1.GetEntriesResponse, error) {

	fmt.Println("GetEntries called")

	c, ok := a.caches[req.Keyspace]

	if !ok {
		return &apiaryv1.GetEntriesResponse{
			Entries:  nil,
			NotFound: req.Keys,
		}, nil

	}

	entries := make([]*apiaryv1.Entry, 0)
	notFound := make([]string, 0)

	for _, key := range req.Keys {
		val, err := c.Get(key)
		if err != nil {
			notFound = append(notFound, key)
			continue
		}

		entries = append(entries, &apiaryv1.Entry{
			Key:   key,
			Value: val,
		})
	}

	return &apiaryv1.GetEntriesResponse{
		Entries:  entries,
		NotFound: notFound,
	}, nil
}
func (a *ApiaryService) SetEntries(ctx context.Context, req *apiaryv1.SetEntriesRequest) (*apiaryv1.SetEntriesResponse, error) {
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

	return &apiaryv1.SetEntriesResponse{
		Successful: successful,
		Failed:     failed,
	}, nil
}
func (a *ApiaryService) DeleteEntries(ctx context.Context, req *apiaryv1.DeleteEntriesRequest) (*apiaryv1.DeleteEntriesResponse, error) {
	c, ok := a.caches[req.Keyspace]

	if !ok {
		return &apiaryv1.DeleteEntriesResponse{
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

	return &apiaryv1.DeleteEntriesResponse{
		Successful: successful,
		Failed:     failed,
	}, nil
}
func (a *ApiaryService) ClearEntries(ctx context.Context, req *apiaryv1.ClearEntriesRequest) (*apiaryv1.ClearEntriesResponse, error) {
	c, ok := a.caches[req.Keyspace]
	if !ok {
		return &apiaryv1.ClearEntriesResponse{
			Successful: false,
		}, nil
	}

	c.Clear()
	return &apiaryv1.ClearEntriesResponse{Successful: true}, nil
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
