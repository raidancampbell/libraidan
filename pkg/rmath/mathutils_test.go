package rmath

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 0, Abs(0))
	assert.Equal(t, 1, Abs(1))
	assert.Equal(t, 1, Abs(1))
	assert.Equal(t, math.MaxInt32, Abs(math.MaxInt32))

	// expected difference due to one's complement
	assert.Equal(t, math.MaxInt32+1, Abs(math.MinInt32))
}
