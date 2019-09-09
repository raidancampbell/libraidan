package rstrings

import "strings"

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}