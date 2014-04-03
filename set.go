package set

import (
	"fmt"
	"sync"
)

// Set represents an unordered collection of unique values
type Set struct {
	// Mutex to allow safe, concurrent access
	mutex sync.RWMutex
	// Empty struct consumes no memory, so we just use the map keys
	m map[interface{}]struct{}
}

// New creates a new Set, and initializes its internal map, optionally adding
// initial elements to the set
func New(values ...interface{}) *Set {
	// Initialize set
	s := Set{
		m: make(map[interface{}]struct{}),
	}

	// If items are specified in the initializer, immediately add them to the set
	for _, v := range values {
		s.Add(v)
	}

	return &s
}

// Add inserts a new element into the set, returning true if the element was newly
// added, or false if it already existed
func (s *Set) Add(value interface{}) bool {
	// Check existence
	found := s.Has(value)

	// Lock set for write
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Add value to set
	s.m[value] = struct{}{}

	// Return inverse, so true if element already existed
	return !found
}

// Has checks for membership of an element in the set
func (s *Set) Has(value interface{}) bool {
	// Lock set for read
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Check for value
	if _, ok := s.m[value]; ok {
		// Value found
		return true
	}

	// Value not found
	return false
}

// Size returns the size or cardinality of this set
func (s *Set) Size() int {
	// Lock set for read
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.m)
}

// String returns a string representation of this set
func (s *Set) String() string {
	// Print identifier
	str := fmt.Sprintf("Set(%d):[ ", s.Size())

	// Print all elements
	for k, _ := range s.m {
		str = str + fmt.Sprintf("%v ", k)
	}

	return str + "]"
}
