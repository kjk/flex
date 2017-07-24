package flex

import "testing"

func TestComputed_layout_margin(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)
	root.StyleSetMarginPercent(EdgeStart, 10)

	CalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 10, root.LayoutGetMargin(EdgeLeft))
	assertFloatEqual(t, 0, root.LayoutGetMargin(EdgeRight))

	CalculateLayout(root, 100, 100, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetMargin(EdgeLeft))
	assertFloatEqual(t, 10, root.LayoutGetMargin(EdgeRight))
}
