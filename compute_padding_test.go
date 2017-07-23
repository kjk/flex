package flex

import "testing"

func TestComputed_layout_padding(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)
	YGNodeStyleSetPaddingPercent(root, YGEdgeStart, 10)

	YGNodeCalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 10, YGNodeLayoutGetPadding(root, YGEdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, YGEdgeRight))

	YGNodeCalculateLayout(root, 100, 100, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, YGEdgeLeft))
	assertFloatEqual(t, 10, YGNodeLayoutGetPadding(root, YGEdgeRight))

}
