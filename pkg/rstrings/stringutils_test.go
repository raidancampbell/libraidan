package rstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""))
	assert.True(t, IsEmpty(" "))
	assert.True(t, IsEmpty("\t"))
	assert.True(t, IsEmpty("\r"))
	assert.True(t, IsEmpty("\n"))
	assert.True(t, IsEmpty(" \t\r\n\t"))
	assert.False(t, IsEmpty(" \t\ra\n\t"))
	assert.False(t, IsEmpty(" leading and trailing "))
	assert.False(t, IsEmpty("no leading ir trailing"))
}
