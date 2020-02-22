package roper

import "github.com/raidancampbell/libraidan/pkg/rstrings"

// NilOr is a coalescence operator similar to the python style "x = x or y", used to choose the first non-nil variable
func NilOr(a, b interface{}) interface{} {
	if a != nil {
		return a
	}
	return b
}

// EmptyOr is a coalescence operator similar to the python style "x = x or y", used to choose the first non-empty string
func EmptyOr(a, b string) string {
	if !rstrings.IsEmpty(a) {
		return a
	}
	return b
}

// ZeroOr is a coalescence operator similar to the python style "x = x or y", used to choose the first nonzero integer
func ZeroOr(a, b int) int {
	if a != 0 {
		return a
	}
	return b
}

// First is a coalescing function to return the first non-type-default argument in the variadic arg list
// see IsDefaultValue for the default checking implementation
func First(i ...interface{}) interface{} {
	for _, v := range i {
		if !IsDefaultValue(v) {
			return v
		}
	}
	return nil
}

// Coalesce is an alias for First
func Coalesce(i ...interface{}) interface{} {
	return First(i...)
}

// FirstStr is a string typed wrapper for First
func FirstStr(s ...string) string {
	// conversion is required, see https://golang.org/doc/faq#convert_slice_of_interface for details
	i := make([]interface{}, len(s))
	for idx, v := range s {
		i[idx] = v
	}
	result := First(i...)
	if result == nil {
		return ""
	}
	return result.(string)
}

// FirstInt is an int typed wrapper for First
func FirstInt(s ...int) int {
	// conversion is required, see https://golang.org/doc/faq#convert_slice_of_interface for details
	i := make([]interface{}, len(s))
	for idx, v := range s {
		i[idx] = v
	}
	result := First(i...)
	if result == nil {
		return 0
	}
	return result.(int)
}
