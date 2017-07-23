package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	eps = 0.000001
)

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

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeStart, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 10)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsoluteLayoutStartTopEndBottom(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeStart, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeEnd, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeBottom, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsoluteLayoutWidthHeightStartTopEndBottom(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeStart, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeEnd, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeBottom, 10)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestDoNotClampHeightOfAbsoluteNodeToHeightOfItsOverflowHiddenParent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetOverflow(root, OverflowHidden)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeStart, 0)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 0)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0, 100)
	YGNodeStyleSetHeight(rootChild0Child0, 100)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0))
}

func TestAbsoluteLayoutWithinBorder(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, EdgeLeft, 10)
	YGNodeStyleSetMargin(root, EdgeTop, 10)
	YGNodeStyleSetMargin(root, EdgeRight, 10)
	YGNodeStyleSetMargin(root, EdgeBottom, 10)
	YGNodeStyleSetPadding(root, EdgeLeft, 10)
	YGNodeStyleSetPadding(root, EdgeTop, 10)
	YGNodeStyleSetPadding(root, EdgeRight, 10)
	YGNodeStyleSetPadding(root, EdgeBottom, 10)
	YGNodeStyleSetBorder(root, EdgeLeft, 10)
	YGNodeStyleSetBorder(root, EdgeTop, 10)
	YGNodeStyleSetBorder(root, EdgeRight, 10)
	YGNodeStyleSetBorder(root, EdgeBottom, 10)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeLeft, 0)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 0)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild1, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild1, EdgeRight, 0)
	YGNodeStyleSetPosition(rootChild1, EdgeBottom, 0)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild2, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild2, EdgeLeft, 0)
	YGNodeStyleSetPosition(rootChild2, EdgeTop, 0)
	YGNodeStyleSetMargin(rootChild2, EdgeLeft, 10)
	YGNodeStyleSetMargin(rootChild2, EdgeTop, 10)
	YGNodeStyleSetMargin(rootChild2, EdgeRight, 10)
	YGNodeStyleSetMargin(rootChild2, EdgeBottom, 10)
	YGNodeStyleSetWidth(rootChild2, 50)
	YGNodeStyleSetHeight(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild3, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild3, EdgeRight, 0)
	YGNodeStyleSetPosition(rootChild3, EdgeBottom, 0)
	YGNodeStyleSetMargin(rootChild3, EdgeLeft, 10)
	YGNodeStyleSetMargin(rootChild3, EdgeTop, 10)
	YGNodeStyleSetMargin(rootChild3, EdgeRight, 10)
	YGNodeStyleSetMargin(rootChild3, EdgeBottom, 10)
	YGNodeStyleSetWidth(rootChild3, 50)
	YGNodeStyleSetHeight(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenter(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetWidth(rootChild0, 60)
	YGNodeStyleSetHeight(rootChild0, 40)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentFlexEnd(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyFlexEnd)
	YGNodeStyleSetAlignItems(root, AlignFlexEnd)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetWidth(rootChild0, 60)
	YGNodeStyleSetHeight(rootChild0, 40)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsoluteLayoutJustifyContentCenter(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetWidth(rootChild0, 60)
	YGNodeStyleSetHeight(rootChild0, 40)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsoluteLayoutAlignItemsCenter(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetWidth(rootChild0, 60)
	YGNodeStyleSetHeight(rootChild0, 40)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsoluteLayoutAlignItemsCenterOnChildOnly(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(rootChild0, AlignCenter)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetWidth(rootChild0, 60)
	YGNodeStyleSetHeight(rootChild0, 40)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenterAndTopPosition(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 10)
	YGNodeStyleSetWidth(rootChild0, 60)
	YGNodeStyleSetHeight(rootChild0, 40)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenterAndBottomPosition(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeBottom, 10)
	YGNodeStyleSetWidth(rootChild0, 60)
	YGNodeStyleSetHeight(rootChild0, 40)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenterAndLeftPosition(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeLeft, 5)
	YGNodeStyleSetWidth(rootChild0, 60)
	YGNodeStyleSetHeight(rootChild0, 40)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsolute_layout_align_items_and_justify_content_center_and_right_position(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 110)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeRight, 5)
	YGNodeStyleSetWidth(rootChild0, 60)
	YGNodeStyleSetHeight(rootChild0, 40)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))
}

func TestPosition_root_with_rtl_should_position_withoutdirection(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetPosition(root, EdgeLeft, 72)
	YGNodeStyleSetWidth(root, 52)
	YGNodeStyleSetHeight(root, 52)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 72, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 72, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))
}

func TestAbsolute_layout_percentage_bottom_based_on_parent_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(rootChild0, EdgeTop, 50)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild1, PositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(rootChild1, EdgeBottom, 50)
	YGNodeStyleSetWidth(rootChild1, 10)
	YGNodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild2, PositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(rootChild2, EdgeTop, 10)
	YGNodeStyleSetPositionPercent(rootChild2, EdgeBottom, 10)
	YGNodeStyleSetWidth(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild2))
}

func TestAbsolute_layout_in_wrap_reverse_column_container(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetWidth(rootChild0, 20)
	YGNodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsolute_layout_in_wrap_reverse_row_container(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetWidth(rootChild0, 20)
	YGNodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsolute_layout_in_wrap_reverse_column_container_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(rootChild0, AlignFlexEnd)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetWidth(rootChild0, 20)
	YGNodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))
}

func TestAbsolute_layout_in_wrap_reverse_row_container_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(rootChild0, AlignFlexEnd)
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetWidth(rootChild0, 20)
	YGNodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))
}
