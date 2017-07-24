package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirty_propagation(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	NodeStyleSetWidth(rootChild0, 20)

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
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	NodeStyleSetWidth(rootChild0, 50)

	assert.False(t, rootChild0.IsDirty)
	assert.False(t, rootChild1.IsDirty)
	assert.False(t, root.IsDirty)

}

func TestDirty_mark_all_children_as_dirty_when_display_changes(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetHeight(root, 100)

	child0 := NewNode()
	NodeStyleSetFlexGrow(child0, 1)
	child1 := NewNode()
	NodeStyleSetFlexGrow(child1, 1)

	child1Child0 := NewNode()
	child1Child0Child0 := NewNode()
	NodeStyleSetWidth(child1Child0Child0, 8)
	NodeStyleSetHeight(child1Child0Child0, 16)

	YGNodeInsertChild(child1Child0, child1Child0Child0, 0)

	YGNodeInsertChild(child1, child1Child0, 0)
	YGNodeInsertChild(root, child0, 0)
	YGNodeInsertChild(root, child1, 0)

	NodeStyleSetDisplay(child0, DisplayFlex)
	NodeStyleSetDisplay(child1, DisplayNone)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 0, NodeLayoutGetWidth(child1Child0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(child1Child0Child0))

	NodeStyleSetDisplay(child0, DisplayNone)
	NodeStyleSetDisplay(child1, DisplayFlex)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 8, NodeLayoutGetWidth(child1Child0Child0))
	assertFloatEqual(t, 16, NodeLayoutGetHeight(child1Child0Child0))

	NodeStyleSetDisplay(child0, DisplayFlex)
	NodeStyleSetDisplay(child1, DisplayNone)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 0, NodeLayoutGetWidth(child1Child0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(child1Child0Child0))

	NodeStyleSetDisplay(child0, DisplayNone)
	NodeStyleSetDisplay(child1, DisplayFlex)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 8, NodeLayoutGetWidth(child1Child0Child0))
	assertFloatEqual(t, 16, NodeLayoutGetHeight(child1Child0Child0))
}

func TestDirty_node_only_if_children_are_actually_removed(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	child0 := NewNode()
	NodeStyleSetWidth(child0, 50)
	NodeStyleSetHeight(child0, 25)
	YGNodeInsertChild(root, child0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	child1 := NewNode()
	YGNodeRemoveChild(root, child1)
	assert.False(t, root.IsDirty)

	YGNodeRemoveChild(root, child0)
	assert.True(t, root.IsDirty)
}
