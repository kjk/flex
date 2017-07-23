package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy_style_same(t *testing.T) {
	node0 := YGNodeNew()
	node1 := YGNodeNew()
	assert.False(t, YGNodeIsDirty(node0))

	YGNodeCopyStyle(node0, node1)
	assert.False(t, YGNodeIsDirty(node0))
}

func TestCopy_style_modified(t *testing.T) {
	node0 := YGNodeNew()
	assert.False(t, YGNodeIsDirty(node0))
	assert.Equal(t, FlexDirectionColumn, YGNodeStyleGetFlexDirection(node0))
	assert.False(t, YGNodeStyleGetMaxHeight(node0).unit != YGUnitUndefined)

	node1 := YGNodeNew()
	YGNodeStyleSetFlexDirection(node1, FlexDirectionRow)
	YGNodeStyleSetMaxHeight(node1, 10)

	YGNodeCopyStyle(node0, node1)
	assert.True(t, YGNodeIsDirty(node0))
	assert.Equal(t, FlexDirectionRow, YGNodeStyleGetFlexDirection(node0))
	assertFloatEqual(t, 10, YGNodeStyleGetMaxHeight(node0).value)
}

func TestCopy_style_modified_same(t *testing.T) {
	node0 := YGNodeNew()
	YGNodeStyleSetFlexDirection(node0, FlexDirectionRow)
	YGNodeStyleSetMaxHeight(node0, 10)
	YGNodeCalculateLayout(node0, YGUndefined, YGUndefined, DirectionLTR)
	assert.False(t, YGNodeIsDirty(node0))

	node1 := YGNodeNew()
	YGNodeStyleSetFlexDirection(node1, FlexDirectionRow)
	YGNodeStyleSetMaxHeight(node1, 10)

	YGNodeCopyStyle(node0, node1)
	assert.False(t, YGNodeIsDirty(node0))
}
