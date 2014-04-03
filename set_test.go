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

// TestClone verifies that the set.Clone() method is working properly
func TestClone(t *testing.T) {
	log.Println("TestClone()")

	// Create a table of tests and expected results of cloning
	var tests = []struct {
		source *Set
		target *Set
		result bool
	}{
		// Same items
		{New(1, 3, 5), New(1, 3, 5), true},
		// Re-ordered items
		{New(2, 4, 6), New(6, 4, 2), true},
		// Different items
		{New(1, 2, 3), New(1, 2, 3, 4), false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to clone the current set, verify result
		if clone := test.source.Clone(); clone.Equal(test.target) != test.result {
			t.Fatalf("set.Clone() - unexpected result: %t", test.result)
		}

		log.Println(test.source)
	}
}

// TestDifference verifies that the set.Difference() method is working properly
func TestDifference(t *testing.T) {
	log.Println("TestDifference()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results of Set differences
	var tests = []struct {
		source *Set
		target *Set
	}{
		// Same items
		{New(1, 3, 5), New()},
		// New items (no difference)
		{New(2, 4, 6), New(1, 3, 5)},
		// Combination of items
		{New(1, 2, 6), New(3, 5)},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to add an element to the set, verify result
		difference := set.Difference(test.source)
		if !difference.Equal(test.target) {
			t.Fatalf("set.Difference() - sets not equal: %s != %s", difference.String(), test.target.String())
		}

		log.Println(difference)
	}
}

// TestEnumerate verifies that the set.Enumerate() method is working properly
func TestEnumerate(t *testing.T) {
	log.Println("TestEnumerate()")

	// Create a slice of expected values upon set enumeration
	expected := []int{1, 3, 5, 7, 9}

	// Create a set
	set := New()

	// Add initial values
	for _, e := range expected {
		set.Add(e)
	}

	// Enumerate the values in the set
	for _, v := range set.Enumerate() {
		found := false

		// Check that the expected value was found upon set enumeration
		for _, e := range expected {
			if v == e {
				found = true
			}
		}

		// If value not found, test fails
		if !found {
			t.Fatalf("set.Enumerate() - element missing: %v", v)
		}
	}

	log.Println(set)
}

// TestEqual verifies that the set.Equal() method is working properly
func TestEqual(t *testing.T) {
	log.Println("TestEqual()")

	// Create a table of tests and expected results of cloning
	var tests = []struct {
		source *Set
		target *Set
		result bool
	}{
		// Same items
		{New(1, 3, 5), New(1, 3, 5), true},
		// Re-ordered items
		{New(2, 4, 6), New(6, 4, 2), true},
		// Repeated items
		{New(1, 2, 3), New(1, 2, 3, 1, 2), true},
		// Different items
		{New(1, 2, 3), New(1, 2, 4), false},
		// Different lengths
		{New(2, 4, 6), New(2, 4, 6, 8), false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Check set equality
		if test.source.Equal(test.target) != test.result {
			t.Fatalf("set.Equal() - unexpected result: %t", test.result)
		}

		log.Println(test.source)
	}
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

// TestUnion verifies that the set.Union() method is working properly
func TestUnion(t *testing.T) {
	log.Println("TestUnion()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results of Set unions
	var tests = []struct {
		source *Set
		target *Set
	}{
		// Same items
		{New(1, 3, 5), New(1, 3, 5)},
		// New items
		{New(2, 4, 6), New(1, 2, 3, 4, 5, 6)},
		// Combination of items
		{New(1, 2, 3), New(1, 2, 3, 5)},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to add an element to the set, verify result
		union := set.Union(test.source)
		if !union.Equal(test.target) {
			t.Fatalf("set.Union() - sets not equal: %s != %s", union.String(), test.target.String())
		}

		log.Println(union)
	}
}
