package server

import (
	"fmt"
	pb "github.com/hive-ops/apiary/pb/proto"
	"github.com/hive-ops/apiary/utils"
	"testing"
)

var namespace = pb.NewNamespace("hive-ops", "apiary")
var keyspace = pb.NewKeyspace(namespace, "benchmark")

func BenchmarkApiarySet(b *testing.B) {

	apiaryInstance := NewApiary()

	cmd := pb.NewSetEntriesCommand(keyspace, []*pb.Entry{
		{
			Key:   "foo",
			Value: "bar",
		},
	})

	for i := 0; i < b.N; i++ {
		apiaryInstance.SetEntries(cmd)
	}

	utils.ReportOpsPerSec(b)

}

func BenchmarkApiaryGet(b *testing.B) {

	entry := &pb.Entry{
		Key:   "foo",
		Value: "bar",
	}

	apiaryInstance := NewApiary()

	apiaryInstance.SetEntries(pb.NewSetEntriesCommand(keyspace, []*pb.Entry{entry}))

	getCmd := pb.NewGetEntriesCommand(keyspace, []string{entry.Key})

	for i := 0; i < b.N; i++ {
		apiaryInstance.GetEntries(getCmd)
	}

	utils.ReportOpsPerSec(b)

}

func BenchmarkApiaryDelete(b *testing.B) {

	apiaryInstance := NewApiary()

	cmds := make([]*pb.DeleteEntriesCommand, b.N)
	for i := range cmds {
		key := fmt.Sprintf("foo-%d", i)
		value := fmt.Sprintf("bar-%d", i)
		cmds[i] = pb.NewDeleteEntriesCommand(keyspace, []string{key})
		apiaryInstance.SetEntries(pb.NewSetEntriesCommand(keyspace, []*pb.Entry{
			{
				Key:   key,
				Value: value,
			},
		}))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		apiaryInstance.DeleteEntries(cmds[i])
	}

	utils.ReportOpsPerSec(b)

}
