package popcount

import "testing"

func BenchmarkPopCountTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(4837294738)
	}
}

func BenchmarkPopCountLowestBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowestBit(4837294738)
	}
}
