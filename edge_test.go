package flex

import "testing"

func TestStart_overrides(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, EdgeStart, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeLeft, 20)
	YGNodeStyleSetMargin(rootChild0, EdgeRight, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetRight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(rootChild0))
}

func TestEnd_overrides(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, EdgeEnd, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeLeft, 20)
	YGNodeStyleSetMargin(rootChild0, EdgeRight, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetRight(rootChild0))
}

func TestHorizontal_overridden(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, EdgeHorizontal, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeLeft, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(rootChild0))
}

func TestVertical_overridden(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, EdgeVertical, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeTop, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetBottom(rootChild0))
}

func TestHorizontal_overrides_all(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, EdgeHorizontal, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeAll, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetBottom(rootChild0))
}

func TestVertical_overrides_all(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, EdgeVertical, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeAll, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetRight(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetBottom(rootChild0))
}

func TestAll_overridden(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, EdgeLeft, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeTop, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeRight, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeBottom, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeAll, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetBottom(rootChild0))
}
