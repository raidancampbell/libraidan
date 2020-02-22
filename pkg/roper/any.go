package roper

import "github.com/raidancampbell/libraidan/pkg/rstrings"

// AnyNil returns a bool whether any of the variadic inputs is nil
// if no inputs are passed, it defaults to false
func AnyNil(a ...interface{}) bool {
	for _, v := range a {
		if v == nil {
			return true
		}
	}
	return false
}

// AnyEmpty returns a bool whether any of the variadic inputs is an "empty" string.
// see rstrings::IsEmpty for empty-checking logic
// if no inputs are passed, it defaults to false
func AnyEmpty(a ...string) bool {
	for _, v := range a {
		if rstrings.IsEmpty(v) {
			return true
		}
	}
	return false
}

// AnyZero returns a bool whether any of the variadic inputs is zero
// if no inputs are passed, it defaults to false
func AnyZero(a ...int) bool {
	for _, v := range a {
		if v == 0 {
			return true
		}
	}
	return false
}

// Any returns a bool indicating whether any of the given bools are true
func Any(a ...bool) bool {
	flag := false
	for _, v := range a {
		flag = flag || v
	}
	return flag
}
