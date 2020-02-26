// Package rruntime contains useful runtime or meta functions, including caller (function, filename, line) introspection
// and context.Context serialization/deserialization
package rruntime

import (
	"runtime"
	"strings"
)

// GetCallerDetails returns a (filename, function name, line number) tuple of the current callstack
// the function internally skips itself, so invoking GetCallerDetails(0) returns the invoker's location
// similarly, invoking GetCallerDetails(1) returns the invoker's caller's location,
// and GetCallerDetails(-1) will always return the location of GetCallerDetails
func GetCallerDetails(stackSkip int) (string, string, int) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2+stackSkip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fqfn := frame.Function
	s := strings.Split(fqfn, ".")
	return frame.File, s[len(s)-1], frame.Line
}

// GetMyFuncName returns the caller's function name
func GetMyFuncName() string {
	_, fName, _ := GetCallerDetails(1)
	return fName
}

// GetMyFileName returns the caller's absolute file path
func GetMyFileName() string {
	fName, _, _ := GetCallerDetails(1)
	return fName
}

// GetMyLineNumber returns the caller's line number
func GetMyLineNumber() int {
	_, _, line := GetCallerDetails(1)
	return line
}
