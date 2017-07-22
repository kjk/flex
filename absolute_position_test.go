package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	eps = 0.000001
)

func init() {
	//gPrintTree = true
	//gPrintChanges = true
}

func assertFloatEqual(t *testing.T, got, exp float32) {
	assert.Equal(t, got, exp)
	diff := fabs(got - exp)
	if diff > eps {
		t.Fatalf("got: %.2f, exp: %.2f", got, exp)
	}
}

func TestAbsoluteLayoutWidthHeightStartTop(t *testing.T) {
	config := YGConfigNew()
	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeStart, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)
	YGConfigFree(config)
}

func testAbsoluteLayoutStartTopEndBottom(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeStart, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeEnd, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeBottom, 10)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func AbsoluteLayoutWidthHeightStartTopEndBottomTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeStart, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeEnd, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func DoNotClampHeightOfAbsoluteNodeToHeightOfItsOverflowHiddenParentTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetOverflow(root, YGOverflowHidden)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeStart, 0)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 0)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0, 100)
	YGNodeStyleSetHeight(root_child0_child0, 100)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func AbsoluteLayoutWithinBorderTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(root, YGEdgeTop, 10)
	YGNodeStyleSetMargin(root, YGEdgeRight, 10)
	YGNodeStyleSetMargin(root, YGEdgeBottom, 10)
	YGNodeStyleSetPadding(root, YGEdgeLeft, 10)
	YGNodeStyleSetPadding(root, YGEdgeTop, 10)
	YGNodeStyleSetPadding(root, YGEdgeRight, 10)
	YGNodeStyleSetPadding(root, YGEdgeBottom, 10)
	YGNodeStyleSetBorder(root, YGEdgeLeft, 10)
	YGNodeStyleSetBorder(root, YGEdgeTop, 10)
	YGNodeStyleSetBorder(root, YGEdgeRight, 10)
	YGNodeStyleSetBorder(root, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeLeft, 0)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 0)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child1, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child1, YGEdgeRight, 0)
	YGNodeStyleSetPosition(root_child1, YGEdgeBottom, 0)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child2, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child2, YGEdgeLeft, 0)
	YGNodeStyleSetPosition(root_child2, YGEdgeTop, 0)
	YGNodeStyleSetMargin(root_child2, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(root_child2, YGEdgeTop, 10)
	YGNodeStyleSetMargin(root_child2, YGEdgeRight, 10)
	YGNodeStyleSetMargin(root_child2, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeStyleSetHeight(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child3, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child3, YGEdgeRight, 0)
	YGNodeStyleSetPosition(root_child3, YGEdgeBottom, 0)
	YGNodeStyleSetMargin(root_child3, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(root_child3, YGEdgeTop, 10)
	YGNodeStyleSetMargin(root_child3, YGEdgeRight, 10)
	YGNodeStyleSetMargin(root_child3, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeStyleSetHeight(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func AbsoluteLayoutAlignItemsAndJustifyContentCenterTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 40)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func AbsoluteLayoutAlignItemsAndJustifyContentFlexEndText(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyFlexEnd)
	YGNodeStyleSetAlignItems(root, YGAlignFlexEnd)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 40)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func AbsoluteLayoutJustifyContentCenterTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 40)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func AbsoluteLayoutAlignItemsCenterTest(t *testing.T) {

	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 40)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func AbsoluteLayoutAlignItemsCenterOnChildOnlyTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child0, YGAlignCenter)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 40)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func AbsoluteLayoutAlignItemsAndJustifyContentCenterAndTopPositionTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 40)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func AbsoluteLayoutAlignItemsAndJustifyContentCenterAndBottomPositionTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 40)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func AbsoluteLayoutAlignItemsAndJustifyContentCenterAndLeftPositionTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeLeft, 5)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 40)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func absolute_layout_align_items_and_justify_content_center_and_right_positionTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeRight, 5)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 40)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func position_root_with_rtl_should_position_withoutdirectionTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetPosition(root, YGEdgeLeft, 72)
	YGNodeStyleSetWidth(root, 52)
	YGNodeStyleSetHeight(root, 52)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 72, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 72, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func absolute_layout_percentage_bottom_based_on_parent_heightTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(root_child0, YGEdgeTop, 50)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child1, YGPositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(root_child1, YGEdgeBottom, 50)
	YGNodeStyleSetWidth(root_child1, 10)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child2, YGPositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(root_child2, YGEdgeTop, 10)
	YGNodeStyleSetPositionPercent(root_child2, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func absolute_layout_in_wrap_reverse_column_containerTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetWidth(root_child0, 20)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func absolute_layout_in_wrap_reverse_row_containerTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetWidth(root_child0, 20)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func absolute_layout_in_wrap_reverse_column_container_flex_endTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child0, YGAlignFlexEnd)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetWidth(root_child0, 20)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func absolute_layout_in_wrap_reverse_row_container_flex_endTest(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child0, YGAlignFlexEnd)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetWidth(root_child0, 20)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}
