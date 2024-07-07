package server

import pb "github.com/hive-ops/apiary/pb/proto"

type Apiary struct {
	Caches map[string]*Cache
}

func NewApiary() *Apiary {
	return &Apiary{
		Caches: make(map[string]*Cache),
	}
}

func (a *Apiary) GetEntries(cmd *pb.GetEntriesCommand) *pb.GetEntriesResponse {

	cache, ok := a.Caches[cmd.Keyspace.GetKeyspaceRef()]
	if !ok {
		return &pb.GetEntriesResponse{
			Entries:  nil,
			NotFound: cmd.Keys,
		}
	}

	entries := make([]*pb.Entry, 0)
	notFound := make([]string, 0)

	for _, key := range cmd.Keys {
		val, err := cache.Get(key)
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
	}

}

func (a *Apiary) SetEntries(cmd *pb.SetEntriesCommand) *pb.SetEntriesResponse {

	cache, ok := a.Caches[cmd.Keyspace.GetKeyspaceRef()]
	if !ok {
		cache = NewCache()
		a.Caches[cmd.Keyspace.GetKeyspaceRef()] = cache
	}

	successful := make([]string, 0)
	failed := make([]string, 0)

	for _, entry := range cmd.Entries {
		cache.Set(entry.Key, entry.Value)
		successful = append(successful, entry.Key)
	}

	return &pb.SetEntriesResponse{
		Successful: successful,
		Failed:     failed,
	}

}

func (a *Apiary) DeleteEntries(cmd *pb.DeleteEntriesCommand) *pb.DeleteEntriesResponse {

	cache, ok := a.Caches[cmd.Keyspace.GetKeyspaceRef()]
	if !ok {
		return &pb.DeleteEntriesResponse{
			Successful: nil,
			NotFound:   cmd.Keys,
			Failed:     nil,
		}
	}

	successful := make([]string, 0)
	failed := make([]string, 0)

	for _, key := range cmd.Keys {
		cache.Delete(key)
		successful = append(successful, key)
	}

	return &pb.DeleteEntriesResponse{
		Successful: successful,
		Failed:     failed,
	}

}

func (a *Apiary) ClearEntries(cmd *pb.ClearEntriesCommand) *pb.ClearEntriesResponse {
	cache, ok := a.Caches[cmd.Keyspace.GetKeyspaceRef()]
	if !ok {
		return &pb.ClearEntriesResponse{
			Successful: false,
		}
	}

	cache.Clear()
	return &pb.ClearEntriesResponse{
		Successful: true,
	}
}
