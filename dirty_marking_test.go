package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirty_propagation(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(20)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	rootChild0.StyleSetWidth(20)

	assert.True(t, rootChild0.IsDirty)
	assert.False(t, rootChild1.IsDirty)
	assert.True(t, root.IsDirty)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.False(t, rootChild0.IsDirty)
	assert.False(t, rootChild1.IsDirty)
	assert.False(t, root.IsDirty)

}

func TestDirty_propagation_only_if_prop_changed(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(20)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	rootChild0.StyleSetWidth(50)

	assert.False(t, rootChild0.IsDirty)
	assert.False(t, rootChild1.IsDirty)
	assert.False(t, root.IsDirty)

}

func TestDirty_mark_all_children_as_dirty_when_display_changes(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetHeight(100)

	child0 := NewNode()
	child0.StyleSetFlexGrow(1)
	child1 := NewNode()
	child1.StyleSetFlexGrow(1)

	child1Child0 := NewNode()
	child1Child0Child0 := NewNode()
	child1Child0Child0.StyleSetWidth(8)
	child1Child0Child0.StyleSetHeight(16)

	NodeInsertChild(child1Child0, child1Child0Child0, 0)

	NodeInsertChild(child1, child1Child0, 0)
	NodeInsertChild(root, child0, 0)
	NodeInsertChild(root, child1, 0)

	child0.StyleSetDisplay(DisplayFlex)
	child1.StyleSetDisplay(DisplayNone)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 0, child1Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 0, child1Child0Child0.LayoutGetHeight())

	child0.StyleSetDisplay(DisplayNone)
	child1.StyleSetDisplay(DisplayFlex)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 8, child1Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 16, child1Child0Child0.LayoutGetHeight())

	child0.StyleSetDisplay(DisplayFlex)
	child1.StyleSetDisplay(DisplayNone)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 0, child1Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 0, child1Child0Child0.LayoutGetHeight())

	child0.StyleSetDisplay(DisplayNone)
	child1.StyleSetDisplay(DisplayFlex)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 8, child1Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 16, child1Child0Child0.LayoutGetHeight())
}

func TestDirty_node_only_if_children_are_actually_removed(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(50)
	root.StyleSetHeight(50)

	child0 := NewNode()
	child0.StyleSetWidth(50)
	child0.StyleSetHeight(25)
	NodeInsertChild(root, child0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	child1 := NewNode()
	NodeRemoveChild(root, child1)
	assert.False(t, root.IsDirty)

	NodeRemoveChild(root, child0)
	assert.True(t, root.IsDirty)
}
