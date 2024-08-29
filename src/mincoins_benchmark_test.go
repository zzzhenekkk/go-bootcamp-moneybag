package main

import "testing"

func BenchmarkMinCoins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCoins(13, []int{1, 5, 10, 50, 100, 500, 1000})
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCoins2(13, []int{1, 5, 10, 50, 100, 500, 1000})
	}
}
