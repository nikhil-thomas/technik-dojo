package main

import "testing"

// Wrong benchmark test
// test will not give any output
// value of b.N increases within each bench mark run
// hence benchmarks will not converge to a stable value
func XBenchmarkFibo1(b *testing.B) {
	for i := 0; i < b.N; i++ {

		_ = fibo1(i)
	}
}

// Wrong benchmark test
// test will not give any output
// value of b.N increases within each bench mark run
// hence benchmarks will not converge to a stable value
func XBenchmarkFibo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibo2(b.N)
	}
}

// valid benchmark test
func BenchmarkFibo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibo2(10)
	}
}

// valid benchmark test
// but, this benchmark can take a long time to complete
func BenchmarkFibo2(b *testing.B) {
	_ = fibo2(b.N)
}
