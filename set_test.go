package set

import (
	"log"
	"testing"
)

// TestAdd verifies that the set.Add() method is working properly
func TestAdd(t *testing.T) {
	log.Println("TestAdd()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results for adding new elements
	var tests = []struct {
		element interface{}
		result  bool
	}{
		// New items
		{2, true},
		{4, true},
		{6, true},
		// Existing items
		{1, false},
		{3, false},
		{5, false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to add an element to the set, verify result
		if ok := set.Add(test.element); ok != test.result {
			t.Fatalf("set.Add(%d) - unexpected result: %t", test.element, ok)
		}
	}

	log.Println(set)
}

// TestHas verifies that the set.Has() method is working properly
func TestHas(t *testing.T) {
	log.Println("TestHas()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results for checking membership of elements
	var tests = []struct {
		element interface{}
		result  bool
	}{
		// Existing items
		{1, true},
		{3, true},
		{5, true},
		// Non-existant items
		{2, false},
		{4, false},
		{6, false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to check if the element is contained in the set, verify result
		if ok := set.Has(test.element); ok != test.result {
			t.Fatalf("set.Has(%d) - unexpected result: %t", test.element, ok)
		}
	}

	log.Println(set)
}

// TestRemove verifies that the set.Remove() method is working properly
func TestRemove(t *testing.T) {
	log.Println("TestRemove()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results for removing elements
	var tests = []struct {
		element interface{}
		result  bool
	}{
		// Existing items
		{1, true},
		{3, true},
		{5, true},
		// Non-existant items
		{2, false},
		{4, false},
		{6, false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to remove an element from the set, verify result
		if ok := set.Remove(test.element); ok != test.result {
			t.Fatalf("set.Remove(%d) - unexpected result: %t", test.element, ok)
		}
	}

	log.Println(set)
}

// TestSize verifies that the set.Size() method is working properly
func TestSize(t *testing.T) {
	log.Println("TestSize()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Verify initial size
	if set.Size() != 3 {
		t.Fatalf("set.Size() - unexpected result: %d", set.Size())
	}

	// Create a table of tests and expected size when adding new elements
	var tests = []struct {
		element interface{}
		size    int
	}{
		// New items
		{2, 4},
		{4, 5},
		{6, 6},
		// Existing items
		{1, 6},
		{3, 6},
		{5, 6},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Add an element to the set, check size
		set.Add(test.element)

		if set.Size() != test.size {
			t.Fatalf("set.Size()- unexpected result: %d", set.Size())
		}
	}

	log.Println(set)
}
