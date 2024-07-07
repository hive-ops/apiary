package server

import (
	"fmt"
	"testing"
)

func BenchmarkSet(b *testing.B) {

	key := "key"
	value := "value"
	c := NewCache()

	for i := 0; i < b.N; i++ {
		c.Set(key, value)
	}

	reportOpsPerSec(b)

}

func BenchmarkGet(b *testing.B) {

	key := "key"
	value := "value"
	c := NewCache()
	c.Set(key, value)

	for i := 0; i < b.N; i++ {
		_, _ = c.Get(key)
	}

	reportOpsPerSec(b)

}

func BenchmarkDelete(b *testing.B) {

	keys := make([]string, b.N)
	for i := range keys {
		key := fmt.Sprintf("key-%d", i)
		keys[i] = key
		c.Set(key, "value")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Delete(keys[i])
	}

	reportOpsPerSec(b)

}

func reportOpsPerSec(b *testing.B) {
	opsPerSec := float64(b.N) / b.Elapsed().Seconds()
	b.ReportMetric(opsPerSec, "ops/s")
	b.ReportMetric(b.Elapsed().Seconds(), "s")
	b.ReportAllocs()
}
