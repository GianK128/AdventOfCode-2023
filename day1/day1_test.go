package day1

import (
	"testing"
)

func BenchmarkDay1Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1("input.txt")
	}
}

func BenchmarkDay1Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2("input.txt")
	}
}
