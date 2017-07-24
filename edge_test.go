package flex

import "testing"

func TestStart_overrides(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetFlexGrow(1)
	NodeStyleSetMargin(rootChild0, EdgeStart, 10)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 20)
	NodeStyleSetMargin(rootChild0, EdgeRight, 20)
	NodeInsertChild(root, rootChild0, 0)

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
	NodeStyleSetMargin(rootChild0, EdgeEnd, 10)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 20)
	NodeStyleSetMargin(rootChild0, EdgeRight, 20)
	NodeInsertChild(root, rootChild0, 0)

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
	NodeStyleSetMargin(rootChild0, EdgeHorizontal, 10)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 20)
	NodeInsertChild(root, rootChild0, 0)

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
	NodeStyleSetMargin(rootChild0, EdgeVertical, 10)
	NodeStyleSetMargin(rootChild0, EdgeTop, 20)
	NodeInsertChild(root, rootChild0, 0)

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
	NodeStyleSetMargin(rootChild0, EdgeHorizontal, 10)
	NodeStyleSetMargin(rootChild0, EdgeAll, 20)
	NodeInsertChild(root, rootChild0, 0)

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
	NodeStyleSetMargin(rootChild0, EdgeVertical, 10)
	NodeStyleSetMargin(rootChild0, EdgeAll, 20)
	NodeInsertChild(root, rootChild0, 0)

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
	NodeStyleSetMargin(rootChild0, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild0, EdgeTop, 10)
	NodeStyleSetMargin(rootChild0, EdgeRight, 10)
	NodeStyleSetMargin(rootChild0, EdgeBottom, 10)
	NodeStyleSetMargin(rootChild0, EdgeAll, 20)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetRight())
	assertFloatEqual(t, 10, rootChild0.LayoutGetBottom())
}
