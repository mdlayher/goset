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

// benchmarkCartesianProduct checks the performance of the set.CartesianProduct() method
func benchmarkCartesianProduct(n int, s *Set, t *Set) {
	// Run set.CartesianProduct() n times
	for i := 0; i < n; i++ {
		s.CartesianProduct(t)
	}
}

// BenchmarkCartesianProductSmall checks the performance of the set.CartesianProduct() method
// over a small data set
func BenchmarkCartesianProductSmall(b *testing.B) {
	benchmarkCartesianProduct(b.N, New(1, 2), New(2, 1))
}

// BenchmarkCartesianProductLarge checks the performance of the set.CartesianProduct() method
// over a large data set
func BenchmarkCartesianProductLarge(b *testing.B) {
	benchmarkCartesianProduct(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), New(9, 8, 7, 6, 5, 4, 3, 2, 1))
}

// benchmarkClone checks the performance of the set.Clone() method
func benchmarkClone(n int, s *Set) {
	// Run set.Clone() n times
	for i := 0; i < n; i++ {
		s.Clone()
	}
}

// BenchmarkCloneSmall checks the performance of the set.Clone() method
// over a small data set
func BenchmarkCloneSmall(b *testing.B) {
	benchmarkClone(b.N, New(1, 2))
}

// BenchmarkCloneLarge checks the performance of the set.Clone() method
// over a large data set
func BenchmarkCloneLarge(b *testing.B) {
	benchmarkClone(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

// benchmarkDifference checks the performance of the set.Difference() method
func benchmarkDifference(n int, s *Set, t *Set) {
	// Run set.Difference() n times
	for i := 0; i < n; i++ {
		s.Difference(t)
	}
}

// BenchmarkDifferenceSmall checks the performance of the set.Difference() method
// over a small data set
func BenchmarkDifferenceSmall(b *testing.B) {
	benchmarkDifference(b.N, New(1, 2), New(2, 1))
}

// BenchmarkDifferenceLarge checks the performance of the set.Difference() method
// over a large data set
func BenchmarkDifferenceLarge(b *testing.B) {
	benchmarkDifference(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), New(9, 8, 7, 6, 5, 4, 3, 2, 1))
}
