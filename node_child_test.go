package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset_layout_when_child_removed(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(100)
	rootChild0.StyleSetHeight(100)
	root.InsertChild(rootChild0, 0)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	root.RemoveChild(rootChild0)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assert.True(t, FloatIsUndefined(rootChild0.LayoutGetWidth()))
	assert.True(t, FloatIsUndefined(rootChild0.LayoutGetHeight()))
}
