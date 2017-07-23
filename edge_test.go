package flex

import "testing"

func TestStart_overrides(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMargin(root_child0, YGEdgeStart, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeLeft, 20)
	YGNodeStyleSetMargin(root_child0, YGEdgeRight, 20)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetRight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(root_child0))

	YGNodeFreeRecursive(root)
}

func TestEnd_overrides(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMargin(root_child0, YGEdgeEnd, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeLeft, 20)
	YGNodeStyleSetMargin(root_child0, YGEdgeRight, 20)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetRight(root_child0))

	YGNodeFreeRecursive(root)
}

func TestHorizontal_overridden(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMargin(root_child0, YGEdgeHorizontal, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeLeft, 20)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(root_child0))

	YGNodeFreeRecursive(root)
}

func TestVertical_overridden(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMargin(root_child0, YGEdgeVertical, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeTop, 20)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetBottom(root_child0))

	YGNodeFreeRecursive(root)
}

func TestHorizontal_overrides_all(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMargin(root_child0, YGEdgeHorizontal, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeAll, 20)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetBottom(root_child0))

	YGNodeFreeRecursive(root)
}

func TestVertical_overrides_all(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMargin(root_child0, YGEdgeVertical, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeAll, 20)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetRight(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetBottom(root_child0))

	YGNodeFreeRecursive(root)
}

func TestAll_overridden(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMargin(root_child0, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeRight, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeBottom, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeAll, 20)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetRight(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetBottom(root_child0))

	YGNodeFreeRecursive(root)
}
