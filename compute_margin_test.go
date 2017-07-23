package flex

import "testing"

func TestComputed_layout_margin(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)
	YGNodeStyleSetMarginPercent(root, EdgeStart, 10)

	YGNodeCalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 10, YGNodeLayoutGetMargin(root, EdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, EdgeRight))

	YGNodeCalculateLayout(root, 100, 100, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, EdgeLeft))
	assertFloatEqual(t, 10, YGNodeLayoutGetMargin(root, EdgeRight))
}
