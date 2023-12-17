package day3

import (
	"testing"
)

func BenchmarkDay2Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1("input.txt")
	}
}

func BenchmarkDay2Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2("input.txt")
	}
}
