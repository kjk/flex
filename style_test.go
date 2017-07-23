package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy_style_same(t *testing.T) {
	node0 := NewNode()
	node1 := NewNode()
	assert.False(t, node0.IsDirty)

	NodeCopyStyle(node0, node1)
	assert.False(t, node0.IsDirty)
}

func TestCopy_style_modified(t *testing.T) {
	node0 := NewNode()
	assert.False(t, node0.IsDirty)
	assert.Equal(t, FlexDirectionColumn, YGNodeStyleGetFlexDirection(node0))
	assert.False(t, NodeStyleGetMaxHeight(node0).Unit != UnitUndefined)

	node1 := NewNode()
	YGNodeStyleSetFlexDirection(node1, FlexDirectionRow)
	YGNodeStyleSetMaxHeight(node1, 10)

	NodeCopyStyle(node0, node1)
	assert.True(t, node0.IsDirty)
	assert.Equal(t, FlexDirectionRow, YGNodeStyleGetFlexDirection(node0))
	assertFloatEqual(t, 10, NodeStyleGetMaxHeight(node0).Value)
}

func TestCopy_style_modified_same(t *testing.T) {
	node0 := NewNode()
	YGNodeStyleSetFlexDirection(node0, FlexDirectionRow)
	YGNodeStyleSetMaxHeight(node0, 10)
	NodeCalculateLayout(node0, Undefined, Undefined, DirectionLTR)
	assert.False(t, node0.IsDirty)

	node1 := NewNode()
	YGNodeStyleSetFlexDirection(node1, FlexDirectionRow)
	YGNodeStyleSetMaxHeight(node1, 10)

	NodeCopyStyle(node0, node1)
	assert.False(t, node0.IsDirty)
}
