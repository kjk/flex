package flex

import "testing"

func TestComputed_layout_margin(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)
	YGNodeStyleSetMarginPercent(root, YGEdgeStart, 10)

	YGNodeCalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 10, YGNodeLayoutGetMargin(root, YGEdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, YGEdgeRight))

	YGNodeCalculateLayout(root, 100, 100, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, YGEdgeLeft))
	assertFloatEqual(t, 10, YGNodeLayoutGetMargin(root, YGEdgeRight))

}
