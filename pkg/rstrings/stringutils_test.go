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

func TestDefaultIfEmpty(t *testing.T) {
	assert.Equal(t, "foo", DefaultIfEmpty("", "foo"))
	assert.Equal(t, "foo", DefaultIfEmpty("foo", ""))
	assert.Equal(t, "foo", DefaultIfEmpty("foo", "bar"))
	assert.Equal(t, "", DefaultIfEmpty("", ""))
	assert.Equal(t, "foo", DefaultIfEmpty("   ", "foo"))
	assert.Equal(t, "foo", DefaultIfEmpty("foo", "   "))
	assert.Equal(t, "   ", DefaultIfEmpty("", "   "))
}

func TestMapToString(t *testing.T) {
	v := make(map[string]interface{})
	v["foo"] = "bar"
	assert.Equal(t, "{\"foo\":\"bar\"}", MapToString(v))
	assert.Equal(t, "{\"foo\":\"bar\"}", MapToString(v, false))
	assert.Equal(t, "{\n  \"foo\": \"bar\"\n}", MapToString(v, true))

	v["baz"] = 2
	assert.Contains(t, MapToString(v), "\"foo\":\"bar\"")
	assert.Contains(t, MapToString(v), "\"baz\":2")

	v["qux"] = nil
	assert.Contains(t, MapToString(v), "\"qux\":null")

	assert.Equal(t, "{}", MapToString(make(map[string]interface{})))
}

func TestLeftPad(t *testing.T) {
	type args struct {
		input  string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"untouched", args{"foo", 3}, "foo"},
		{"too large", args{"foo", 1}, "foo"},
		{"empty", args{"", 3}, "   "},
		{"negative", args{"foo", -1}, "foo"},
		{"happy", args{"foo", 5}, "  foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeftPad(tt.args.input, tt.args.length); got != tt.want {
				t.Errorf("LeftPad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRightPad(t *testing.T) {
	type args struct {
		input  string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"untouched", args{"foo", 3}, "foo"},
		{"too large", args{"foo", 1}, "foo"},
		{"empty", args{"", 3}, "   "},
		{"negative", args{"foo", -1}, "foo"},
		{"happy", args{"foo", 5}, "foo  "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RightPad(tt.args.input, tt.args.length); got != tt.want {
				t.Errorf("LeftPad() = %v, want %v", got, tt.want)
			}
		})
	}
}
