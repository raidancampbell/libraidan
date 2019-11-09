package rstrings

import "strings"

// IsEmpty returns whether the given string has a length of 0 or is only whitespace
// examples:
// IsEmpty("") = true
// IsEmpty("    ") = True
// IsEmpty("b") = false
// IsEmpty("   b") = false
// IsEmpty("b   ") = false
// IsEmpty("   b   ") = false
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
