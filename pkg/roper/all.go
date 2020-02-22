// Package roper implements operator-style functions, largely around coalescence or variadic operations
package roper

import "github.com/raidancampbell/libraidan/pkg/rstrings"

// AllNil returns a bool whether All of the variadic inputs is nil
// if no inputs are passed, it defaults to true
func AllNil(a ...interface{}) bool {
	acc := true
	for _, v := range a {
		acc = acc && v == nil
	}
	return acc
}

// AllEmpty returns a bool whether All of the variadic inputs are "empty" strings.
// see rstrings::IsEmpty for empty-checking logic
// if no inputs are passed, it defaults to true
func AllEmpty(a ...string) bool {
	acc := true
	for _, v := range a {
		acc = acc && rstrings.IsEmpty(v)
	}
	return acc
}

// AllZero returns a bool whether All of the variadic inputs is zero
// if no inputs are passed, it defaults to true
func AllZero(a ...int) bool {
	acc := true
	for _, v := range a {
		acc = acc && v == 0
	}
	return acc
}

// All returns a bool indicating that all of the given booleans are true
func All(a ...bool) bool {
	flag := true
	for _, v := range a {
		flag = flag && v
	}
	return flag
}
