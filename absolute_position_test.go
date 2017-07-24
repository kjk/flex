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
	config := NewConfig()
	root := NewNodeWithConfig(config)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeStart, 10)
	rootChild0.StyleSetPosition(EdgeTop, 10)
	rootChild0.StyleSetWidth(10)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutStartTopEndBottom(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeStart, 10)
	rootChild0.StyleSetPosition(EdgeTop, 10)
	rootChild0.StyleSetPosition(EdgeEnd, 10)
	rootChild0.StyleSetPosition(EdgeBottom, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutWidthHeightStartTopEndBottom(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeStart, 10)
	rootChild0.StyleSetPosition(EdgeTop, 10)
	rootChild0.StyleSetPosition(EdgeEnd, 10)
	rootChild0.StyleSetPosition(EdgeBottom, 10)
	rootChild0.StyleSetWidth(10)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestDoNotClampHeightOfAbsoluteNodeToHeightOfItsOverflowHiddenParent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetOverflow(OverflowHidden)
	root.StyleSetWidth(50)
	root.StyleSetHeight(50)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeStart, 0)
	rootChild0.StyleSetPosition(EdgeTop, 0)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidth(100)
	rootChild0Child0.StyleSetHeight(100)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, -50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetHeight())
}

func TestAbsoluteLayoutWithinBorder(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetMargin(EdgeLeft, 10)
	root.StyleSetMargin(EdgeTop, 10)
	root.StyleSetMargin(EdgeRight, 10)
	root.StyleSetMargin(EdgeBottom, 10)
	root.StyleSetPadding(EdgeLeft, 10)
	root.StyleSetPadding(EdgeTop, 10)
	root.StyleSetPadding(EdgeRight, 10)
	root.StyleSetPadding(EdgeBottom, 10)
	NodeStyleSetBorder(root, EdgeLeft, 10)
	NodeStyleSetBorder(root, EdgeTop, 10)
	NodeStyleSetBorder(root, EdgeRight, 10)
	NodeStyleSetBorder(root, EdgeBottom, 10)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeLeft, 0)
	rootChild0.StyleSetPosition(EdgeTop, 0)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetPositionType(PositionTypeAbsolute)
	rootChild1.StyleSetPosition(EdgeRight, 0)
	rootChild1.StyleSetPosition(EdgeBottom, 0)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetPositionType(PositionTypeAbsolute)
	rootChild2.StyleSetPosition(EdgeLeft, 0)
	rootChild2.StyleSetPosition(EdgeTop, 0)
	rootChild2.StyleSetMargin(EdgeLeft, 10)
	rootChild2.StyleSetMargin(EdgeTop, 10)
	rootChild2.StyleSetMargin(EdgeRight, 10)
	rootChild2.StyleSetMargin(EdgeBottom, 10)
	rootChild2.StyleSetWidth(50)
	rootChild2.StyleSetHeight(50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetPositionType(PositionTypeAbsolute)
	rootChild3.StyleSetPosition(EdgeRight, 0)
	rootChild3.StyleSetPosition(EdgeBottom, 0)
	rootChild3.StyleSetMargin(EdgeLeft, 10)
	rootChild3.StyleSetMargin(EdgeTop, 10)
	rootChild3.StyleSetMargin(EdgeRight, 10)
	rootChild3.StyleSetMargin(EdgeBottom, 10)
	rootChild3.StyleSetWidth(50)
	rootChild3.StyleSetHeight(50)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10, root.LayoutGetLeft())
	assertFloatEqual(t, 10, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 20, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild3.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 10, root.LayoutGetLeft())
	assertFloatEqual(t, 10, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 20, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild3.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenter(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetFlexGrow(1)
	root.StyleSetWidth(110)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetWidth(60)
	rootChild0.StyleSetHeight(40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentFlexEnd(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyFlexEnd)
	root.StyleSetAlignItems(AlignFlexEnd)
	root.StyleSetFlexGrow(1)
	root.StyleSetWidth(110)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetWidth(60)
	rootChild0.StyleSetHeight(40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutJustifyContentCenter(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetFlexGrow(1)
	root.StyleSetWidth(110)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetWidth(60)
	rootChild0.StyleSetHeight(40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsCenter(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetFlexGrow(1)
	root.StyleSetWidth(110)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetWidth(60)
	rootChild0.StyleSetHeight(40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsCenterOnChildOnly(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexGrow(1)
	root.StyleSetWidth(110)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetAlignSelf(AlignCenter)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetWidth(60)
	rootChild0.StyleSetHeight(40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenterAndTopPosition(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetFlexGrow(1)
	root.StyleSetWidth(110)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeTop, 10)
	rootChild0.StyleSetWidth(60)
	rootChild0.StyleSetHeight(40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenterAndBottomPosition(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetFlexGrow(1)
	root.StyleSetWidth(110)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeBottom, 10)
	rootChild0.StyleSetWidth(60)
	rootChild0.StyleSetHeight(40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenterAndLeftPosition(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetFlexGrow(1)
	root.StyleSetWidth(110)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeLeft, 5)
	rootChild0.StyleSetWidth(60)
	rootChild0.StyleSetHeight(40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsolute_layout_align_items_and_justify_content_center_and_right_position(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetFlexGrow(1)
	root.StyleSetWidth(110)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeRight, 5)
	rootChild0.StyleSetWidth(60)
	rootChild0.StyleSetHeight(40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestPosition_root_with_rtl_should_position_withoutdirection(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetPosition(EdgeLeft, 72)
	root.StyleSetWidth(52)
	root.StyleSetHeight(52)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 72, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 72, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())
}

func TestAbsolute_layout_percentage_bottom_based_on_parent_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(100)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPositionPercent(EdgeTop, 50)
	rootChild0.StyleSetWidth(10)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetPositionType(PositionTypeAbsolute)
	rootChild1.StyleSetPositionPercent(EdgeBottom, 50)
	rootChild1.StyleSetWidth(10)
	rootChild1.StyleSetHeight(10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetPositionType(PositionTypeAbsolute)
	rootChild2.StyleSetPositionPercent(EdgeTop, 10)
	rootChild2.StyleSetPositionPercent(EdgeBottom, 10)
	rootChild2.StyleSetWidth(10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 90, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 160, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 90, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 160, rootChild2.LayoutGetHeight())
}

func TestAbsolute_layout_in_wrap_reverse_column_container(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexWrap(WrapWrapReverse)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetWidth(20)
	rootChild0.StyleSetHeight(20)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())
}

func TestAbsolute_layout_in_wrap_reverse_row_container(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetFlexWrap(WrapWrapReverse)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetWidth(20)
	rootChild0.StyleSetHeight(20)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())
}

func TestAbsolute_layout_in_wrap_reverse_column_container_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexWrap(WrapWrapReverse)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetAlignSelf(AlignFlexEnd)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetWidth(20)
	rootChild0.StyleSetHeight(20)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())
}

func TestAbsolute_layout_in_wrap_reverse_row_container_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetFlexWrap(WrapWrapReverse)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetAlignSelf(AlignFlexEnd)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetWidth(20)
	rootChild0.StyleSetHeight(20)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())
}
