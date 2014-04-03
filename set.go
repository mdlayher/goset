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

// Clone copies the current set into a new, identical set
func (s Set) Clone() *Set {
	// Copy set into a new set
	outSet := New()
	for _, v := range s.Enumerate() {
		outSet.Add(v)
	}

	return outSet
}

// Difference returns a set containing all elements present in this set, but
// without any elements present in the parameter set
func (s Set) Difference(paramSet *Set) *Set {
	// Create a set of differences between the sets
	diffSet := New()

	// Enumerate and check all elements in the current set
	for _, e := range s.Enumerate() {
		found := false

		// Check if element is present in parameter set
		for _, p := range paramSet.Enumerate() {
			// Element found
			if e == p {
				found = true
			}
		}

		// If element was not found, add it to diff set
		if !found {
			diffSet.Add(e)
		}
	}

	return diffSet
}

// Enumerate returns an unordered slice of all elements in the set
func (s *Set) Enumerate() []interface{} {
	// Lock set for read
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Gather all values into a slice
	values := make([]interface{}, 0)
	for k, _ := range s.m {
		values = append(values, k)
	}

	return values
}

// Equal returns whether or not two sets have the same length and
// no differences, meaning they are equal
func (s Set) Equal(paramSet *Set) bool {
	return s.Size() == paramSet.Size() && s.Difference(paramSet).Size() == 0
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

// Remove destroys an element in the set, returning true if the element was
// destroyed, or false if it did not exist
func (s *Set) Remove(value interface{}) bool {
	// Check existence
	found := s.Has(value)

	// Lock set for write
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Remove value from set
	delete(s.m, value)

	return found
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

// Subset determines if a parameter set is a subset of elements within this
// set, returning true if it is a subset, or false if it is not
func (s Set) Subset(paramSet *Set) bool {
	// Check if all elements in the parameter set are contained within the set
	for _, v := range paramSet.Enumerate() {
		// Check if element is contained, if not, return false
		if !s.Has(v) {
			return false
		}
	}

	return true
}

// Union returns a set containing all elements present in this set, as well
// as all elements present in the parameter set
func (s Set) Union(paramSet *Set) *Set {
	// Clone the current set into a new set
	outSet := s.Clone()

	// Enumerate and add all elements from the parameter set
	for _, e := range paramSet.Enumerate() {
		outSet.Add(e)
	}

	return outSet
}
