package client

import (
	"context"
	pb "github.com/hive-ops/apiary/pb/proto"
	"github.com/samber/lo"
	"testing"
	"time"
)

var namespace = pb.NewNamespace("hive-ops", "apiary")
var keyspace = pb.NewKeyspace(namespace, "benchmark")

func TestClient(t *testing.T) {

	client := NewClient("127.0.0.1:2468")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := "testKey"

	cmd := pb.NewGetEntriesCommand(keyspace, []string{key})

	res, _ := client.GetEntries(ctx, cmd)

	if !lo.Contains(res.NotFound, key) {
		t.Fatalf("GetEntries failed: %v", res)
	}

	client.close()
}
