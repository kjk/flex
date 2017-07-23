package flex

import "testing"

func TestComputed_layout_padding(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)
	YGNodeStyleSetPaddingPercent(root, EdgeStart, 10)

	YGNodeCalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 10, YGNodeLayoutGetPadding(root, EdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, EdgeRight))

	YGNodeCalculateLayout(root, 100, 100, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, EdgeLeft))
	assertFloatEqual(t, 10, YGNodeLayoutGetPadding(root, EdgeRight))
}
