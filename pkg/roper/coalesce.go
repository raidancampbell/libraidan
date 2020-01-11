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
