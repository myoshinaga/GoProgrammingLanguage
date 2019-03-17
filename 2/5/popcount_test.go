package popcount

import "testing"

func BenchmarkPopCountTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(4837294738)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(4837294738)
	}
}

func BenchmarkPopCountLowestBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowestBit(4837294738)
	}
}

func BenchmarkPopCountLowestBitClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLowestBitClear(4837294738)
	}
}
