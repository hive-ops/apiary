package server

import (
	"context"
	"fmt"
	"github.com/hive-ops/apiary/pb"
	"github.com/hive-ops/apiary/utils"
	"testing"
)

var keyspace = "benchmark"
var config = LoadConfig("../apiary.yaml")

func BenchmarkApiarySet(b *testing.B) {

	server := NewApiaryService(config)
	ctx := context.Background()

	req := &pb.SetEntriesRequest{
		Keyspace: keyspace,
		Entries: []*pb.Entry{
			{
				Key:   "foo",
				Value: "bar",
			},
		},
	}

	for i := 0; i < b.N; i++ {
		_, _ = server.SetEntries(ctx, req)
	}

	utils.ReportOpsPerSec(b)

}

func BenchmarkApiaryGet(b *testing.B) {

	entry := &pb.Entry{
		Key:   "foo",
		Value: "bar",
	}

	server := NewApiaryService(config)
	ctx := context.Background()

	_, _ = server.SetEntries(ctx, &pb.SetEntriesRequest{
		Keyspace: keyspace,
		Entries:  []*pb.Entry{entry},
	})

	getCmd := &pb.GetEntriesRequest{
		Keyspace: keyspace,
		Keys:     []string{entry.Key},
	}

	for i := 0; i < b.N; i++ {
		_, _ = server.GetEntries(ctx, getCmd)
	}

	utils.ReportOpsPerSec(b)

}

func BenchmarkApiaryDelete(b *testing.B) {

	server := NewApiaryService(config)
	ctx := context.Background()

	cmds := make([]*pb.DeleteEntriesRequest, b.N)
	for i := range cmds {
		key := fmt.Sprintf("foo-%d", i)
		value := fmt.Sprintf("bar-%d", i)
		cmds[i] = &pb.DeleteEntriesRequest{Keyspace: keyspace, Keys: []string{key}}
		_, _ = server.SetEntries(ctx, &pb.SetEntriesRequest{Keyspace: keyspace, Entries: []*pb.Entry{
			{
				Key:   key,
				Value: value,
			},
		}})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = server.DeleteEntries(ctx, cmds[i])
	}

	utils.ReportOpsPerSec(b)

}
