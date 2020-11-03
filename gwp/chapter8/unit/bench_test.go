package main

import (
	"testing"
)

// BenchmarkDecode ...
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("post.json")
	}
}
