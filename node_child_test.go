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
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	NodeRemoveChild(root, rootChild0)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assert.True(t, FloatIsUndefined(NodeLayoutGetWidth(rootChild0)))
	assert.True(t, FloatIsUndefined(NodeLayoutGetHeight(rootChild0)))
}
