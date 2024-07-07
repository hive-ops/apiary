package server

import (
	"context"
	"fmt"
	pb "github.com/hive-ops/apiary/pb/proto"
	"github.com/hive-ops/apiary/utils"
	"testing"
)

var namespace = pb.NewNamespace("hive-ops", "apiary")
var keyspace = pb.NewKeyspace(namespace, "benchmark")
var config = LoadConfig("../apiary.yaml")

func BenchmarkApiarySet(b *testing.B) {

	server := NewApiaryServer(config)
	ctx := context.Background()

	cmd := pb.NewSetEntriesCommand(keyspace, []*pb.Entry{
		{
			Key:   "foo",
			Value: "bar",
		},
	})

	for i := 0; i < b.N; i++ {
		_, _ = server.SetEntries(ctx, cmd)
	}

	utils.ReportOpsPerSec(b)

}

func BenchmarkApiaryGet(b *testing.B) {

	entry := &pb.Entry{
		Key:   "foo",
		Value: "bar",
	}

	server := NewApiaryServer(config)
	ctx := context.Background()

	_, _ = server.SetEntries(ctx, pb.NewSetEntriesCommand(keyspace, []*pb.Entry{entry}))

	getCmd := pb.NewGetEntriesCommand(keyspace, []string{entry.Key})

	for i := 0; i < b.N; i++ {
		_, _ = server.GetEntries(ctx, getCmd)
	}

	utils.ReportOpsPerSec(b)

}

func BenchmarkApiaryDelete(b *testing.B) {

	server := NewApiaryServer(config)
	ctx := context.Background()

	cmds := make([]*pb.DeleteEntriesCommand, b.N)
	for i := range cmds {
		key := fmt.Sprintf("foo-%d", i)
		value := fmt.Sprintf("bar-%d", i)
		cmds[i] = pb.NewDeleteEntriesCommand(keyspace, []string{key})
		_, _ = server.SetEntries(ctx, pb.NewSetEntriesCommand(keyspace, []*pb.Entry{
			{
				Key:   key,
				Value: value,
			},
		}))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = server.DeleteEntries(ctx, cmds[i])
	}

	utils.ReportOpsPerSec(b)

}
