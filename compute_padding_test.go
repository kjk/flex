package flex

import "testing"

func TestComputed_layout_padding(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)
	NodeStyleSetPaddingPercent(root, EdgeStart, 10)

	NodeCalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 10, root.LayoutGetPadding(EdgeLeft))
	assertFloatEqual(t, 0, root.LayoutGetPadding(EdgeRight))

	NodeCalculateLayout(root, 100, 100, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetPadding(EdgeLeft))
	assertFloatEqual(t, 10, root.LayoutGetPadding(EdgeRight))
}
