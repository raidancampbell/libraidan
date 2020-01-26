package roper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDefaultValue(t *testing.T) {
	assert.False(t, IsDefaultValue(42))
	assert.True(t, IsDefaultValue(0))
	assert.False(t, IsDefaultValue(-1))

	assert.True(t, IsDefaultValue(""))
	assert.False(t, IsDefaultValue("foo"))

	type f struct {
		x int
	}

	assert.True(t, IsDefaultValue(f{}))

	inst := f{x: 1}
	assert.False(t, IsDefaultValue(inst))
}
