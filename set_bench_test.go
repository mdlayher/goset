package set

import (
	"testing"
)

// BenchmarkAdd checks the performance of the set.Add() method
func BenchmarkAdd(b *testing.B) {
	// Create a new set
	set := New()

	// Run set.Add() b.N times
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}
