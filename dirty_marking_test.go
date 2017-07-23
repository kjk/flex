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

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNew()
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	YGNodeStyleSetWidth(root_child0, 20)

	assert.True(t, YGNodeIsDirty(root_child0))
	assert.False(t, YGNodeIsDirty(root_child1))
	assert.True(t, YGNodeIsDirty(root))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assert.False(t, YGNodeIsDirty(root_child0))
	assert.False(t, YGNodeIsDirty(root_child1))
	assert.False(t, YGNodeIsDirty(root))

}

func TestDirty_propagation_only_if_prop_changed(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNew()
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	YGNodeStyleSetWidth(root_child0, 50)

	assert.False(t, YGNodeIsDirty(root_child0))
	assert.False(t, YGNodeIsDirty(root_child1))
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

	YGNodeStyleSetDisplay(child0, YGDisplayFlex)
	YGNodeStyleSetDisplay(child1, YGDisplayNone)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(child1_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(child1_child0_child0))

	YGNodeStyleSetDisplay(child0, YGDisplayNone)
	YGNodeStyleSetDisplay(child1, YGDisplayFlex)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 8, YGNodeLayoutGetWidth(child1_child0_child0))
	assertFloatEqual(t, 16, YGNodeLayoutGetHeight(child1_child0_child0))

	YGNodeStyleSetDisplay(child0, YGDisplayFlex)
	YGNodeStyleSetDisplay(child1, YGDisplayNone)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(child1_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(child1_child0_child0))

	YGNodeStyleSetDisplay(child0, YGDisplayNone)
	YGNodeStyleSetDisplay(child1, YGDisplayFlex)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
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

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	child1 := YGNodeNew()
	YGNodeRemoveChild(root, child1)
	assert.False(t, YGNodeIsDirty(root))

	YGNodeRemoveChild(root, child0)
	assert.True(t, YGNodeIsDirty(root))
}
