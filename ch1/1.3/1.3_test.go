package main

import (
	"testing"
)

// go test -bench=.
var args = []string{"a", "fast", "dog", "jumped", "over", "my", "balls"}

func BenchmarkWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithJoin(args[0:])
	}
}

func BenchmarkWithoutJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithoutJoin(args[0:])
	}
}
