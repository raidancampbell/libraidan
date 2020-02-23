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

// DefaultIfEmpty returns the given string.
//If the given string is empty (see rstrings.IsEmpty), the given default is returned
func DefaultIfEmpty(s, deflt string) string {
	if IsEmpty(s) {
		return deflt
	}
	return s
}

// MapToString converts the given map to a string.
// The variadic bool is a flag indicating whether the result should be prettified with newlines and whitespace
// default is false, for compatibility
func MapToString(input map[string]interface{}, isPretty ...bool) string {
	var (
		b   []byte
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

// LeftPad prepend-pads the given string to the desired length with space characters.
// See LeftPadWith for implementation details
func LeftPad(input string, length int) string {
	return LeftPadWith(input, ' ', length)
}

// LeftPadWith prepend-pads the given string to the given length with the given rune.
// If the input string is already longer than the desired length, it is returned and NOT truncated
func LeftPadWith(input string, char rune, length int) string {
	l := length - len(input)
	if l <= 0 {
		return input
	}
	return strings.Repeat(string(char), l) + input
}

// RightPad append-pads the given string to the desired length with space characters.
// See RightPadWith for implementation details
func RightPad(input string, length int) string {
	return RightPadWith(input, ' ', length)
}

// RightPadWith append-pads the given string to the given length with the given rune.
// If the input string is already longer than the desired length, it is returned and NOT truncated
func RightPadWith(input string, char rune, length int) string {
	l := length - len(input)
	if l <= 0 {
		return input
	}
	return input + strings.Repeat(string(char), l)
}

// EnsurePrefix prefixes the given string with the given prefix if it doesn't have it already
// examples:
// EnsurePrefix("foo", "f") = "foo"
// EnsurePrefix("foo", "new-") = "new-foo"
func EnsurePrefix(input string, prefix string) string {
	if !strings.HasPrefix(input, prefix) {
		input = prefix + input
	}
	return input
}

// EnsureSuffix suffixes the given string with the given suffix if it doesn't have it already
// examples:
// EnsureSuffix("foo", "o") = "foo"
// EnsureSuffix("foo", "-new") = "foo-new"
func EnsureSuffix(input string, suffix string) string {
	if !strings.HasSuffix(input, suffix) {
		input = input + suffix
	}
	return input
}

// EnsureWrapped wraps the given string with the given prefix/suffix if it doesn't have it already
// examples:
// EnsureWrapped("foo", "'") = "'foo'"
// EnsureSuffix("fluff", "f") = "fluff"
func EnsureWrapped(input string, wrapper string) string {
	if len(input) == 0 {
		return wrapper + wrapper
	}
	return EnsureSuffix(EnsurePrefix(input, wrapper), wrapper)
}
