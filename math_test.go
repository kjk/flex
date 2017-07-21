package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaN(t *testing.T) {
	isNan := IsNaN(NAN)
	assert.True(t, isNan)
	isNan = IsNaN(5)
	assert.False(t, isNan)
}
