package flex

import "testing"

func TestComputed_layout_padding(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)
	root.StyleSetPaddingPercent(EdgeStart, 10)

	CalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 10, root.LayoutGetPadding(EdgeLeft))
	assertFloatEqual(t, 0, root.LayoutGetPadding(EdgeRight))

	CalculateLayout(root, 100, 100, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetPadding(EdgeLeft))
	assertFloatEqual(t, 10, root.LayoutGetPadding(EdgeRight))
}
