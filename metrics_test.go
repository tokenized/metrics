package metrics

import (
	"testing"
	"time"
)

func BenchmarkTiming_time_Since(b *testing.B) {
	// benchmark calculating elapsed time using time.Since
	start := time.Now()

	for n := 0; n < b.N; n++ {
		_ = float64(time.Since(start).Nanoseconds()) / float64(time.Millisecond)
	}
}

func BenchmarkTiming_nano_nano(b *testing.B) {
	// benchmarks calculating elapsed time using time.Nanosecond difference.
	start := time.Now()

	for n := 0; n < b.N; n++ {
		_ = float64(time.Now().Nanosecond()-start.Nanosecond()) / float64(time.Millisecond)
	}
}
