package flex

import "testing"

func TestStart_overrides(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMargin(rootChild0, EdgeStart, 10)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 20)
	NodeStyleSetMargin(rootChild0, EdgeRight, 20)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetRight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)
	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetRight(rootChild0))
}

func TestEnd_overrides(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMargin(rootChild0, EdgeEnd, 10)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 20)
	NodeStyleSetMargin(rootChild0, EdgeRight, 20)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetRight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)
	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetRight(rootChild0))
}

func TestHorizontal_overridden(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMargin(rootChild0, EdgeHorizontal, 10)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 20)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetRight(rootChild0))
}

func TestVertical_overridden(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMargin(rootChild0, EdgeVertical, 10)
	NodeStyleSetMargin(rootChild0, EdgeTop, 20)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetBottom(rootChild0))
}

func TestHorizontal_overrides_all(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMargin(rootChild0, EdgeHorizontal, 10)
	NodeStyleSetMargin(rootChild0, EdgeAll, 20)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetRight(rootChild0))
	assertFloatEqual(t, 20, NodeLayoutGetBottom(rootChild0))
}

func TestVertical_overrides_all(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMargin(rootChild0, EdgeVertical, 10)
	NodeStyleSetMargin(rootChild0, EdgeAll, 20)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, NodeLayoutGetRight(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetBottom(rootChild0))
}

func TestAll_overridden(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild0, EdgeTop, 10)
	NodeStyleSetMargin(rootChild0, EdgeRight, 10)
	NodeStyleSetMargin(rootChild0, EdgeBottom, 10)
	NodeStyleSetMargin(rootChild0, EdgeAll, 20)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetRight(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetBottom(rootChild0))
}
