package rruntime

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetCallerDetails(t *testing.T) {
	file, function, line := GetCallerDetails(0)
	assert.True(t, strings.HasSuffix(file, "callstack_test.go"))
	assert.True(t, strings.HasSuffix(function, "TestGetCallerDetails"))
	assert.Equal(t, 10, line)
}

func TestGetMyFileName(t *testing.T) {
	file := GetMyFileName()
	assert.True(t, strings.HasSuffix(file, "callstack_test.go"))
}

func TestGetMyFuncName(t *testing.T) {
	function := GetMyFuncName()
	assert.True(t, strings.HasSuffix(function, "TestGetMyFuncName"))
}

func TestGetMyLineNumber(t *testing.T) {
	line := GetMyLineNumber()
	assert.Equal(t, 27, line)
}