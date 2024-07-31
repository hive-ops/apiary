package client

import (
	"context"
	"github.com/hive-ops/apiary/pb"
	"github.com/samber/lo"
	"testing"
	"time"
)

var keyspace = "benchmark"

func TestClient(t *testing.T) {

	client := NewClient("127.0.0.1:2468")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := "testKey"

	req := &pb.GetEntriesRequest{Keyspace: keyspace, Keys: []string{key}}

	res, err := client.GetEntries(ctx, req)

	if err != nil {
		t.Fatalf("GetEntries failed: %v", err)
	}

	if !lo.Contains(res.NotFound, key) {
		t.Fatalf("GetEntries failed: %v", res)
	}

	client.close()
}
