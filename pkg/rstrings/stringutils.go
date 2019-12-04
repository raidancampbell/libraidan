package rstrings

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

// MapToString converts the given map to a string.
// The variadic bool is a flag indicating whether the result should be prettified with newlines and whitespace
// default is false, for compatibility
func MapToString(input map[string] interface{}, isPretty...bool) string {
	var (
		b []byte
		err error
	)

	if len(isPretty) > 0 && isPretty[0] {
		b, err = json.MarshalIndent(input, "", "  ")
	} else {
		b, err = json.Marshal(input)
	}
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
