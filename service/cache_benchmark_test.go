package service

import (
	"fmt"
	"github.com/hive-ops/apiary/utils"
	"testing"
)

func BenchmarkSet(b *testing.B) {

	key := "key"
	value := "value"
	c := NewCache()

	for i := 0; i < b.N; i++ {
		c.Set(key, []byte(value))
	}

	utils.ReportOpsPerSec(b)

}

func BenchmarkGet(b *testing.B) {

	key := "key"
	value := "value"
	c := NewCache()
	c.Set(key, []byte(value))

	for i := 0; i < b.N; i++ {
		_, _ = c.Get(key)
	}

	utils.ReportOpsPerSec(b)

}

func BenchmarkDelete(b *testing.B) {

	c := NewCache()

	keys := make([]string, b.N)
	for i := range keys {
		key := fmt.Sprintf("key-%d", i)
		keys[i] = key
		c.Set(key, []byte("value"))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Delete(keys[i])
	}

	utils.ReportOpsPerSec(b)

}
