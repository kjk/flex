package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset_layout_when_child_removed(t *testing.T) {
	root := YGNodeNew()

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 100)
	YGNodeStyleSetHeight(root_child0, 100)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	YGNodeRemoveChild(root, root_child0)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assert.True(t, YGFloatIsUndefined(YGNodeLayoutGetWidth(root_child0)))
	assert.True(t, YGFloatIsUndefined(YGNodeLayoutGetHeight(root_child0)))


}
