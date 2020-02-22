// Package rcollections contains a (deprecated) interface and implementation for collections
// Deprecated: use GoDS instead
package rcollections

import (
	"fmt"
)

// StrMap is an interface for a String-based map
type StrMap interface {
	// Add adds the given key, value pair to the map.
	// return indicates whether a value was overwritten
	Add(key string, value interface{}) StrMap
	// AddAll creates a new map of the combination of both.
	// For intersections, the value from the other map is chosen
	AddAll(other StrMap) StrMap
	// Contains indicates whether the given key is inside the map
	Contains(needle string) bool
	// ContainsValue indicates whether the given value is inside the map
	ContainsValue(needle interface{}) bool
	// Get retrieves the value at the given key.  Returns a non-nil error If the value is not found
	Get(needle string) (interface{}, error)
	// GetWithDefault retrieves the value at the given key.
	// If the key is found, the associated value is returned
	// Otherwise, the given default value is returned
	GetWithDefault(needle string, deflt interface{}) interface{}
	// Remove removes the given key from the map
	// If the key doesn't exist in the map, it's considered a no-op
	Remove(key string) StrMap
	// RemoveAll removes all the keys from the given keyslice from the map, and returns a new resulting map
	// If the key doesn't exist in the map, it's considered a no-op
	RemoveAll(keys []string) StrMap
	// RemoveIntersections calls Remove for all keys in the other map, and returns a new resulting map
	RemoveIntersections(other StrMap) StrMap
	// Intersection is the resulting map of all keys in both maps.  The value is taken from the other map
	Intersection(other StrMap) StrMap
	// Disjunction is the resulting map of all keys NOT in both maps.
	Disjunction(other StrMap) StrMap
	// ContainsAll indicates whether this map contains all the keys of the other map
	// i.e. is the set of this map's keys a superset of the other map's keys
	// if other is empty, true is returned
	ContainsAll(other StrMap) bool
	// ContainsAny indicates whether this map contains any of the keys of the other map
	// i.e. if ContainsAny returns true, the intersection would be non-empty
	// if other is empty, true is returned
	ContainsAny(other StrMap) bool
	// Len returns the number of elements in the map
	Len() int
	// IsEmpty indicates whether the map is empty. i.e. its Len() is 0
	IsEmpty() bool
	// Equals indicates whether both maps have the same keysets, and each key's value is equal
	Equals(other StrMap) bool
	// ValueFrequency counts the number of times a value occurs in the map
	ValueFrequency(needle interface{}) int
	// Filter applies the given function to all values in the map.
	// The function must receive a key:value keypair and return an indicator whether the element should be retained
	// i.e. if the function returns true, it will appear in the resulting map
	Filter(func(key string, val interface{}) bool) StrMap
	// Map applies the given function to all values in the map.
	// The function must receive a key:value keypair and return the mutated value
	Map(func(key string, val interface{}) interface{}) StrMap
	// Reduce applies the given function to all values in the map.
	// The function must receive a key:value keypair along with an accumulator and return the mutated accumulator
	// Since the accumulator is an interface, the function must handle an uninitialized accumulator and type it appropriately
	Reduce(func(key string, val interface{}, acc interface{}) interface{}) interface{}
	// innerMap is an accessor for this library
	innerMap() map[string]interface{}
}

type strMap struct {
	s map[string]interface{}
}

// AddAll overwrites the exiting keys
func (s strMap) AddAll(other StrMap) StrMap {
	newMap := NewStrMap()

	for k, v := range s.s {
		newMap.innerMap()[k] = v
	}

	for ok, ov := range other.innerMap() {
		newMap.innerMap()[ok] = ov
	}
	return newMap
}

// NewStrMap returns an implementation of the StrMap interface
func NewStrMap() StrMap {
	s := make(map[string]interface{})
	return &strMap{s: s}
}

// NewStrMapFrom returns an implementation of the StrMap interface, pre-populated with the given map
func NewStrMapFrom(orig map[string]interface{}) StrMap {
	return &strMap{s: orig}
}

func (s strMap) Contains(needle string) bool {
	for k := range s.s {
		if needle == k {
			return true
		}
	}
	return false
}

func (s strMap) Add(key string, value interface{}) StrMap {
	newStrMap := NewStrMap()
	for k, v := range s.innerMap() {
		newStrMap.innerMap()[k] = v
	}
	newStrMap.innerMap()[key] = value
	return newStrMap
}

func (s strMap) Get(needle string) (interface{}, error) {
	for k, v := range s.s {
		if needle == k {
			return v, nil
		}
	}
	return nil, fmt.Errorf("key: %s not found", needle)
}

func (s strMap) GetWithDefault(needle string, deflt interface{}) interface{} {
	val, err := s.Get(needle)
	if err != nil {
		return deflt
	}
	return val
}

func (s strMap) Remove(key string) StrMap {
	newStrMap := NewStrMap()
	for k, v := range s.innerMap() {
		if k != key {
			newStrMap.innerMap()[k] = v
		}
	}
	return newStrMap
}

func (s strMap) RemoveAll(keys []string) StrMap {
	otherKeySet := NewStrMap()
	for _, key := range keys {
		otherKeySet.innerMap()[key] = nil
	}
	return s.RemoveIntersections(otherKeySet)
}

func (s strMap) ContainsValue(needle interface{}) bool {
	for _, v := range s.s {
		if v == needle {
			return true
		}
	}
	return false
}

func (s strMap) ValueFrequency(needle interface{}) int {
	freq := 0
	for _, v := range s.s {
		if v == needle {
			freq++
		}
	}
	return freq

}

func (s strMap) innerMap() map[string]interface{} {
	return s.s
}

func (s strMap) Filter(f func(key string, val interface{}) bool) StrMap {
	newMap := NewStrMap()
	for k, v := range s.s {
		shouldRetain := f(k, v)
		if shouldRetain {
			newMap.innerMap()[k] = v
		}
	}
	return newMap
}

func (s strMap) Map(f func(key string, val interface{}) interface{}) StrMap {
	newMap := NewStrMap()
	for k, v := range s.s {
		newMap.innerMap()[k] = f(k, v)
	}
	return newMap
}

func (s strMap) Reduce(f func(key string, val interface{}, acc interface{}) interface{}) interface{} {
	var accumulator interface{}
	for k, v := range s.s {
		accumulator = f(k, v, accumulator)
	}
	return accumulator
}

func (s strMap) Intersection(other StrMap) StrMap {
	newMap := NewStrMap()
	for k, v := range other.innerMap() {
		if s.Contains(k) {
			newMap.innerMap()[k] = v
		}
	}
	return newMap
}

func (s strMap) Disjunction(other StrMap) StrMap {
	newMap := NewStrMap()
	for k, v := range s.s {
		if !other.Contains(k) {
			newMap.innerMap()[k] = v
		}
	}

	for k, v := range other.innerMap() {
		if !s.Contains(k) {
			newMap.innerMap()[k] = v
		}
	}

	return newMap

}

func (s strMap) RemoveIntersections(other StrMap) StrMap {
	newMap := NewStrMap()
	for k, v := range s.s {
		if !other.Contains(k) {
			newMap.innerMap()[k] = v
		}
	}
	return newMap
}

func (s strMap) ContainsAll(other StrMap) bool {
	if other.IsEmpty() {
		return true
	}
	for k := range other.innerMap() {
		if !s.Contains(k) {
			return false
		}
	}
	return true
}

func (s strMap) ContainsAny(other StrMap) bool {
	for k := range other.innerMap() {
		if s.Contains(k) {
			return true
		}
	}
	return other.IsEmpty()

}

func (s strMap) Len() int {
	return len(s.innerMap())
}

func (s strMap) IsEmpty() bool {
	return s.Len() == 0
}

// are these equal
func (s strMap) Equals(other StrMap) bool {
	equality := true
	for k, v := range s.s {
		ov, err := other.Get(k)
		equality = equality && err == nil && v == ov
	}
	return equality
}
