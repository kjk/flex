package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset_layout_when_child_removed(t *testing.T) {
	root := YGNodeNew()

	rootChild0 := YGNodeNew()
	YGNodeStyleSetWidth(rootChild0, 100)
	YGNodeStyleSetHeight(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	YGNodeRemoveChild(root, rootChild0)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assert.True(t, YGFloatIsUndefined(YGNodeLayoutGetWidth(rootChild0)))
	assert.True(t, YGFloatIsUndefined(YGNodeLayoutGetHeight(rootChild0)))

}
