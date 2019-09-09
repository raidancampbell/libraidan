package roper

import "libraidan/pkg/rstrings"

func NilOr(a, b interface{}) interface{} {
	if a != nil {
		return a
	} else {
		return b
	}
}

func EmptyOr(a, b string) string {
	if !rstrings.IsEmpty(a) {
		return a
	} else {
		return b
	}
}

func ZeroOr(a, b int) int {
	if a != 0 {
		return a
	} else {
		return b
	}
}