package rruntime

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func BenchmarkGetCallerDetails_minus1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = GetCallerDetails(-1)
	}
}

func BenchmarkGetCallerDetails_zero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = GetCallerDetails(0)
	}
}

func BenchmarkGetCallerDetails_one(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = GetCallerDetails(1)
	}
}

func TestGetCallerDetails(t *testing.T) {
	file, function, line := GetCallerDetails(0)
	assert.True(t, strings.HasSuffix(file, "callstack_test.go"))
	assert.True(t, strings.HasSuffix(function, "TestGetCallerDetails"))
	assert.Equal(t, 28, line)
}

func BenchmarkGetMyFileName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetMyFileName()
	}
}

func TestGetMyFileName(t *testing.T) {
	file := GetMyFileName()
	assert.True(t, strings.HasSuffix(file, "callstack_test.go"))
}

func BenchmarkGetMyFuncName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetMyFuncName()
	}
}

func TestGetMyFuncName(t *testing.T) {
	function := GetMyFuncName()
	assert.Equal(t, "TestGetMyFuncName", function)
}

func BenchmarkGetMyLineNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetMyLineNumber()
	}
}

func TestGetMyLineNumber(t *testing.T) {
	line := GetMyLineNumber()
	assert.Equal(t, 63, line)
}
