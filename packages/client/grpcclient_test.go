package client

import (
	"context"
	apiaryv1 "github.com/hive-ops/apiary/pb/apiary/v1"
	"github.com/samber/lo"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
)

var keyspace = "benchmark"

func TestClient(t *testing.T) {

	client := NewClient("127.0.0.1:2468", insecure.NewCredentials())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := "testKey"

	req := &apiaryv1.GetEntriesRequest{Keyspace: keyspace, Keys: []string{key}}

	res, err := client.GetEntries(ctx, req)

	if err != nil {
		t.Fatalf("GetEntries failed: %v", err)
	}

	if !lo.Contains(res.NotFound, key) {
		t.Fatalf("GetEntries failed: %v", res)
	}

	client.close()
}
