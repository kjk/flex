package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AbsoluteLayoutWidthHeightStartTopTest(t *testing.T) {
	config := YGConfigNew()
	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeStart, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assert.Equal(t, 0, YGNodeLayoutGetLeft(root))
	assert.Equal(t, 0, YGNodeLayoutGetTop(root))
	assert.Equal(t, 100, YGNodeLayoutGetWidth(root))
	assert.Equal(t, 100, YGNodeLayoutGetHeight(root))

	assert.Equal(t, 10, YGNodeLayoutGetLeft(root_child0))
	assert.Equal(t, 10, YGNodeLayoutGetTop(root_child0))
	assert.Equal(t, 10, YGNodeLayoutGetWidth(root_child0))
	assert.Equal(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assert.Equal(t, 0, YGNodeLayoutGetLeft(root))
	assert.Equal(t, 0, YGNodeLayoutGetTop(root))
	assert.Equal(t, 100, YGNodeLayoutGetWidth(root))
	assert.Equal(t, 100, YGNodeLayoutGetHeight(root))

	assert.Equal(t, 80, YGNodeLayoutGetLeft(root_child0))
	assert.Equal(t, 10, YGNodeLayoutGetTop(root_child0))
	assert.Equal(t, 10, YGNodeLayoutGetWidth(root_child0))
	assert.Equal(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)
	YGConfigFree(config)

}
