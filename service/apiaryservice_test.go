package service

import (
	"github.com/hive-ops/apiary/pb"
	"github.com/hive-ops/apiary/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwarmDataServiceTestSuite(t *testing.T) {

	apiaryService := NewApiaryServiceWithDefaultConfig()

	keyspace := utils.RandomString(10, true, false, false)
	entries := []*pb.Entry{
		{
			Key:   "test1",
			Value: []byte("test1"),
		},
		{
			Key:   "test2",
			Value: []byte("test2"),
		},
	}
	keys := make([]string, 0)
	for _, entry := range entries {
		keys = append(keys, entry.Key)
	}

	values := make(map[string][]byte)
	for _, entry := range entries {
		values[entry.Key] = entry.Value
	}

	t.Run("GetEntrySuccessfully - Empty", func(t *testing.T) {
		res, err := apiaryService.GetEntries(nil, pb.NewGetEntriesRequest(keyspace, keys))

		assert.NoError(t, err)
		assert.Empty(t, res.Entries)
		assert.NotEmpty(t, res.NotFound)
	})

	t.Run("SetEntrySuccessfully", func(t *testing.T) {
		res, err := apiaryService.SetEntries(nil, pb.NewSetEntriesRequest(keyspace, entries))

		assert.NoError(t, err)
		assert.Equal(t, keys, res.Successful)
	})

	t.Run("GetEntrySuccessfully - Non-empty", func(t *testing.T) {
		res, err := apiaryService.GetEntries(nil, pb.NewGetEntriesRequest(keyspace, []string{"test1", "invalid-key"}))

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Entries)
		assert.NotEmpty(t, res.NotFound)
	})

	t.Run("DeleteEntrySuccessfully", func(t *testing.T) {
		keysToBeDeleted := []string{"test1"}
		res, err := apiaryService.DeleteEntries(nil, pb.NewDeleteEntriesRequest(keyspace, keysToBeDeleted))

		assert.NoError(t, err)
		assert.Equal(t, keysToBeDeleted, res.Successful)

		getRes, getErr := apiaryService.GetEntries(nil, pb.NewGetEntriesRequest(keyspace, keys))

		assert.NoError(t, getErr)
		assert.NotEmpty(t, getRes.Entries)
		assert.NotEmpty(t, getRes.NotFound)

	})

	t.Run("ClearEntriesSuccessfully", func(t *testing.T) {
		res, err := apiaryService.ClearEntries(nil, pb.NewClearEntriesRequest(keyspace))

		assert.NoError(t, err)
		assert.True(t, res.Successful)

		getRes, getErr := apiaryService.GetEntries(nil, pb.NewGetEntriesRequest(keyspace, keys))

		assert.NoError(t, getErr)
		assert.Empty(t, getRes.Entries)
		assert.NotEmpty(t, getRes.NotFound)
	})

}
