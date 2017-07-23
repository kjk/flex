package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset_layout_when_child_removed(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 100)
	YGNodeStyleSetHeight(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	YGNodeRemoveChild(root, rootChild0)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assert.True(t, FloatIsUndefined(YGNodeLayoutGetWidth(rootChild0)))
	assert.True(t, FloatIsUndefined(YGNodeLayoutGetHeight(rootChild0)))
}
