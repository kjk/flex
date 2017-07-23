package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirty_propagation(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	YGNodeStyleSetWidth(rootChild0, 20)

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
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	YGNodeStyleSetWidth(rootChild0, 50)

	assert.False(t, rootChild0.IsDirty)
	assert.False(t, rootChild1.IsDirty)
	assert.False(t, root.IsDirty)

}

func TestDirty_mark_all_children_as_dirty_when_display_changes(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetHeight(root, 100)

	child0 := NewNode()
	YGNodeStyleSetFlexGrow(child0, 1)
	child1 := NewNode()
	YGNodeStyleSetFlexGrow(child1, 1)

	child1Child0 := NewNode()
	child1Child0Child0 := NewNode()
	YGNodeStyleSetWidth(child1Child0Child0, 8)
	YGNodeStyleSetHeight(child1Child0Child0, 16)

	YGNodeInsertChild(child1Child0, child1Child0Child0, 0)

	YGNodeInsertChild(child1, child1Child0, 0)
	YGNodeInsertChild(root, child0, 0)
	YGNodeInsertChild(root, child1, 0)

	YGNodeStyleSetDisplay(child0, DisplayFlex)
	YGNodeStyleSetDisplay(child1, DisplayNone)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(child1Child0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(child1Child0Child0))

	YGNodeStyleSetDisplay(child0, DisplayNone)
	YGNodeStyleSetDisplay(child1, DisplayFlex)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 8, YGNodeLayoutGetWidth(child1Child0Child0))
	assertFloatEqual(t, 16, YGNodeLayoutGetHeight(child1Child0Child0))

	YGNodeStyleSetDisplay(child0, DisplayFlex)
	YGNodeStyleSetDisplay(child1, DisplayNone)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(child1Child0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(child1Child0Child0))

	YGNodeStyleSetDisplay(child0, DisplayNone)
	YGNodeStyleSetDisplay(child1, DisplayFlex)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 8, YGNodeLayoutGetWidth(child1Child0Child0))
	assertFloatEqual(t, 16, YGNodeLayoutGetHeight(child1Child0Child0))
}

func TestDirty_node_only_if_children_are_actually_removed(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	child0 := NewNode()
	YGNodeStyleSetWidth(child0, 50)
	YGNodeStyleSetHeight(child0, 25)
	YGNodeInsertChild(root, child0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	child1 := NewNode()
	YGNodeRemoveChild(root, child1)
	assert.False(t, root.IsDirty)

	YGNodeRemoveChild(root, child0)
	assert.True(t, root.IsDirty)
}
