package main

import "testing"

// Benchmark function for Sum
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1000)
	}
}

// Benchmark function for SumOptimized
func BenchmarkSumOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOptimized(1000)
	}
}
