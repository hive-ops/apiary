package service

import (
	"context"
	"fmt"
	apiaryv1 "github.com/hive-ops/apiary/pb/apiary/v1"
	"github.com/hive-ops/apiary/utils"
	"testing"
)

var keyspace = "benchmark"
var config = LoadConfig("../apiary.yaml")

func BenchmarkApiarySet(b *testing.B) {

	server := NewApiaryService(config)
	ctx := context.Background()

	req := &apiaryv1.SetEntriesRequest{
		Keyspace: keyspace,
		Entries: []*apiaryv1.Entry{
			{
				Key:   "foo",
				Value: []byte("bar"),
			},
		},
	}

	for i := 0; i < b.N; i++ {
		_, _ = server.SetEntries(ctx, req)
	}

	utils.ReportOpsPerSec(b)

}

func BenchmarkApiaryGet(b *testing.B) {

	entry := &apiaryv1.Entry{
		Key:   "foo",
		Value: []byte("bar"),
	}

	server := NewApiaryService(config)
	ctx := context.Background()

	_, _ = server.SetEntries(ctx, &apiaryv1.SetEntriesRequest{
		Keyspace: keyspace,
		Entries:  []*apiaryv1.Entry{entry},
	})

	getCmd := &apiaryv1.GetEntriesRequest{
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

	cmds := make([]*apiaryv1.DeleteEntriesRequest, b.N)
	for i := range cmds {
		key := fmt.Sprintf("foo-%d", i)
		value := []byte(fmt.Sprintf("bar-%d", i))
		cmds[i] = &apiaryv1.DeleteEntriesRequest{Keyspace: keyspace, Keys: []string{key}}
		_, _ = server.SetEntries(ctx, &apiaryv1.SetEntriesRequest{Keyspace: keyspace, Entries: []*apiaryv1.Entry{
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
