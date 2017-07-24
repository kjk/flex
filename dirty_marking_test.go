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
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNode()
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	rootChild0.StyleSetWidth(20)

	assert.True(t, rootChild0.IsDirty)
	assert.False(t, rootChild1.IsDirty)
	assert.True(t, root.IsDirty)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

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
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNode()
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

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

	child1Child0.InsertChild(child1Child0Child0, 0)

	child1.InsertChild(child1Child0, 0)
	root.InsertChild(child0, 0)
	root.InsertChild(child1, 0)

	child0.StyleSetDisplay(DisplayFlex)
	child1.StyleSetDisplay(DisplayNone)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 0, child1Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 0, child1Child0Child0.LayoutGetHeight())

	child0.StyleSetDisplay(DisplayNone)
	child1.StyleSetDisplay(DisplayFlex)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 8, child1Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 16, child1Child0Child0.LayoutGetHeight())

	child0.StyleSetDisplay(DisplayFlex)
	child1.StyleSetDisplay(DisplayNone)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 0, child1Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 0, child1Child0Child0.LayoutGetHeight())

	child0.StyleSetDisplay(DisplayNone)
	child1.StyleSetDisplay(DisplayFlex)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)
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
	root.InsertChild(child0, 0)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	child1 := NewNode()
	root.RemoveChild(child1)
	assert.False(t, root.IsDirty)

	root.RemoveChild(child0)
	assert.True(t, root.IsDirty)
}
