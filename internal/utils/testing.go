package utils

import "testing"

func ReportOpsPerSec(b *testing.B) {
	opsPerSec := float64(b.N) / b.Elapsed().Seconds()
	b.ReportMetric(opsPerSec, "ops/s")
	b.ReportMetric(b.Elapsed().Seconds(), "s")
	b.ReportAllocs()
}
