package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirty_propagation(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNew()
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	YGNodeStyleSetWidth(rootChild0, 20)

	assert.True(t, YGNodeIsDirty(rootChild0))
	assert.False(t, YGNodeIsDirty(rootChild1))
	assert.True(t, YGNodeIsDirty(root))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.False(t, YGNodeIsDirty(rootChild0))
	assert.False(t, YGNodeIsDirty(rootChild1))
	assert.False(t, YGNodeIsDirty(root))

}

func TestDirty_propagation_only_if_prop_changed(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNew()
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	YGNodeStyleSetWidth(rootChild0, 50)

	assert.False(t, YGNodeIsDirty(rootChild0))
	assert.False(t, YGNodeIsDirty(rootChild1))
	assert.False(t, YGNodeIsDirty(root))

}

func TestDirty_mark_all_children_as_dirty_when_display_changes(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetHeight(root, 100)

	child0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(child0, 1)
	child1 := YGNodeNew()
	YGNodeStyleSetFlexGrow(child1, 1)

	child1_child0 := YGNodeNew()
	child1_child0_child0 := YGNodeNew()
	YGNodeStyleSetWidth(child1_child0_child0, 8)
	YGNodeStyleSetHeight(child1_child0_child0, 16)

	YGNodeInsertChild(child1_child0, child1_child0_child0, 0)

	YGNodeInsertChild(child1, child1_child0, 0)
	YGNodeInsertChild(root, child0, 0)
	YGNodeInsertChild(root, child1, 0)

	YGNodeStyleSetDisplay(child0, DisplayFlex)
	YGNodeStyleSetDisplay(child1, DisplayNone)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(child1_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(child1_child0_child0))

	YGNodeStyleSetDisplay(child0, DisplayNone)
	YGNodeStyleSetDisplay(child1, DisplayFlex)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 8, YGNodeLayoutGetWidth(child1_child0_child0))
	assertFloatEqual(t, 16, YGNodeLayoutGetHeight(child1_child0_child0))

	YGNodeStyleSetDisplay(child0, DisplayFlex)
	YGNodeStyleSetDisplay(child1, DisplayNone)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(child1_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(child1_child0_child0))

	YGNodeStyleSetDisplay(child0, DisplayNone)
	YGNodeStyleSetDisplay(child1, DisplayFlex)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 8, YGNodeLayoutGetWidth(child1_child0_child0))
	assertFloatEqual(t, 16, YGNodeLayoutGetHeight(child1_child0_child0))
}

func TestDirty_node_only_if_children_are_actually_removed(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	child0 := YGNodeNew()
	YGNodeStyleSetWidth(child0, 50)
	YGNodeStyleSetHeight(child0, 25)
	YGNodeInsertChild(root, child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	child1 := YGNodeNew()
	YGNodeRemoveChild(root, child1)
	assert.False(t, YGNodeIsDirty(root))

	YGNodeRemoveChild(root, child0)
	assert.True(t, YGNodeIsDirty(root))
}
