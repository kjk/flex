package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset_layout_when_child_removed(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 100)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	NodeRemoveChild(root, rootChild0)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assert.True(t, FloatIsUndefined(rootChild0.LayoutGetWidth()))
	assert.True(t, FloatIsUndefined(rootChild0.LayoutGetHeight()))
}
