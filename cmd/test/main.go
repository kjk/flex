package main

import (
	"fmt"
	"testing"

	. "github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func assertFloatEqual(t *testing.T, exp, got float32) {
	if exp != got {
		fmt.Printf("got: %.2f, exp: %2.f\n", got, exp)
	}
	assert.Equal(t, got, exp)
}

func testCopy_style_same(t *testing.T) {
	node0 := NewNode()
	node1 := NewNode()
	assert.False(t, YGNodeIsDirty(node0))

	YGNodeCopyStyle(node0, node1)
	assert.False(t, YGNodeIsDirty(node0))

}

func main() {
	t := &testing.T{}
	testCopy_style_same(t)
}
