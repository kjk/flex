package flex

import "testing"

func TestStart_overrides(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetMargin(EdgeStart, 10)
	rootChild0.StyleSetMargin(EdgeLeft, 20)
	rootChild0.StyleSetMargin(EdgeRight, 20)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0.LayoutGetRight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)
	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetRight())
}

func TestEnd_overrides(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetMargin(EdgeEnd, 10)
	rootChild0.StyleSetMargin(EdgeLeft, 20)
	rootChild0.StyleSetMargin(EdgeRight, 20)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetRight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)
	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0.LayoutGetRight())
}

func TestHorizontal_overridden(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetMargin(EdgeHorizontal, 10)
	rootChild0.StyleSetMargin(EdgeLeft, 20)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetRight())
}

func TestVertical_overridden(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionColumn)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetMargin(EdgeVertical, 10)
	rootChild0.StyleSetMargin(EdgeTop, 20)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetBottom())
}

func TestHorizontal_overrides_all(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionColumn)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetMargin(EdgeHorizontal, 10)
	rootChild0.StyleSetMargin(EdgeAll, 20)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetRight())
	assertFloatEqual(t, 20, rootChild0.LayoutGetBottom())
}

func TestVertical_overrides_all(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionColumn)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetMargin(EdgeVertical, 10)
	rootChild0.StyleSetMargin(EdgeAll, 20)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetRight())
	assertFloatEqual(t, 10, rootChild0.LayoutGetBottom())
}

func TestAll_overridden(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionColumn)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetMargin(EdgeLeft, 10)
	rootChild0.StyleSetMargin(EdgeTop, 10)
	rootChild0.StyleSetMargin(EdgeRight, 10)
	rootChild0.StyleSetMargin(EdgeBottom, 10)
	rootChild0.StyleSetMargin(EdgeAll, 20)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetRight())
	assertFloatEqual(t, 10, rootChild0.LayoutGetBottom())
}
