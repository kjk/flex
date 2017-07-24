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
	assert.Equal(t, FlexDirectionColumn, node0.Style.FlexDirection)
	assert.False(t, node0.StyleGetMaxHeight().Unit != UnitUndefined)

	node1 := NewNode()
	NodeStyleSetFlexDirection(node1, FlexDirectionRow)
	NodeStyleSetMaxHeight(node1, 10)

	NodeCopyStyle(node0, node1)
	assert.True(t, node0.IsDirty)
	assert.Equal(t, FlexDirectionRow, node0.Style.FlexDirection)
	assertFloatEqual(t, 10, node0.StyleGetMaxHeight().Value)
}

func TestCopy_style_modified_same(t *testing.T) {
	node0 := NewNode()
	NodeStyleSetFlexDirection(node0, FlexDirectionRow)
	NodeStyleSetMaxHeight(node0, 10)
	NodeCalculateLayout(node0, Undefined, Undefined, DirectionLTR)
	assert.False(t, node0.IsDirty)

	node1 := NewNode()
	NodeStyleSetFlexDirection(node1, FlexDirectionRow)
	NodeStyleSetMaxHeight(node1, 10)

	NodeCopyStyle(node0, node1)
	assert.False(t, node0.IsDirty)
}
