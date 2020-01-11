package roper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilOr(t *testing.T) {
	var ni interface{}
	nn := "chosen"
	actual := NilOr(ni, nn)
	assert.Equal(t, "chosen", actual)
	actual = NilOr(nn, ni)
	assert.Equal(t, "chosen", actual)
}

func TestEmptyOr(t *testing.T) {
	ni := ""
	nn := "chosen"
	actual := EmptyOr(ni, nn)
	assert.Equal(t, "chosen", actual)
	actual = EmptyOr(nn, ni)
	assert.Equal(t, "chosen", actual)
}

func TestZeroOr(t *testing.T) {
	ni := 0
	nn := -1
	actual := ZeroOr(ni, nn)
	assert.Equal(t, -1, actual)
	actual = ZeroOr(nn, ni)
	assert.Equal(t, -1, actual)
}
