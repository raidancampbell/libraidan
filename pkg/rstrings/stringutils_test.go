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

func TestMapToString(t *testing.T) {
	v := make(map[string]interface{})
	v["foo"] = "bar"
	assert.Equal(t,"{\"foo\":\"bar\"}", MapToString(v))
	assert.Equal(t,"{\"foo\":\"bar\"}", MapToString(v, false))
	assert.Equal(t,"{\n  \"foo\": \"bar\"\n}", MapToString(v, true))

	v["baz"] = 2
	assert.Contains(t, MapToString(v), "\"foo\":\"bar\"")
	assert.Contains(t, MapToString(v), "\"baz\":2")

	v["qux"] = nil
	assert.Contains(t, MapToString(v), "\"qux\":null")

	assert.Equal(t, "{}", MapToString(make(map[string]interface{})))
}