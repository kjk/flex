package flex

import "testing"

func TestStart_overrides(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, YGEdgeStart, 10)
	YGNodeStyleSetMargin(rootChild0, YGEdgeLeft, 20)
	YGNodeStyleSetMargin(rootChild0, YGEdgeRight, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetRight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(rootChild0))

}

func TestEnd_overrides(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, YGEdgeEnd, 10)
	YGNodeStyleSetMargin(rootChild0, YGEdgeLeft, 20)
	YGNodeStyleSetMargin(rootChild0, YGEdgeRight, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetRight(rootChild0))

}

func TestHorizontal_overridden(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, YGEdgeHorizontal, 10)
	YGNodeStyleSetMargin(rootChild0, YGEdgeLeft, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(rootChild0))

}

func TestVertical_overridden(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, YGEdgeVertical, 10)
	YGNodeStyleSetMargin(rootChild0, YGEdgeTop, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetBottom(rootChild0))

}

func TestHorizontal_overrides_all(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, YGEdgeHorizontal, 10)
	YGNodeStyleSetMargin(rootChild0, YGEdgeAll, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetBottom(rootChild0))

}

func TestVertical_overrides_all(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, YGEdgeVertical, 10)
	YGNodeStyleSetMargin(rootChild0, YGEdgeAll, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetRight(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetBottom(rootChild0))

}

func TestAll_overridden(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(rootChild0, YGEdgeTop, 10)
	YGNodeStyleSetMargin(rootChild0, YGEdgeRight, 10)
	YGNodeStyleSetMargin(rootChild0, YGEdgeBottom, 10)
	YGNodeStyleSetMargin(rootChild0, YGEdgeAll, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetBottom(rootChild0))

}
