package flex

import "testing"

func TestDisplay_none(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetDisplay(rootChild1, DisplayNone)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))
}

func TestDisplay_none_fixed_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 20)
	NodeStyleSetHeight(rootChild1, 20)
	NodeStyleSetDisplay(rootChild1, DisplayNone)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))
}

func TestDisplay_none_with_margin(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild0, EdgeTop, 10)
	NodeStyleSetMargin(rootChild0, EdgeRight, 10)
	NodeStyleSetMargin(rootChild0, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild0, 20)
	NodeStyleSetHeight(rootChild0, 20)
	NodeStyleSetDisplay(rootChild0, DisplayNone)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))
}

func TestDisplay_none_with_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexShrink(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 0)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetFlexShrink(rootChild1, 1)
	NodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetDisplay(rootChild1, DisplayNone)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1child0, 1)
	NodeStyleSetFlexShrink(rootChild1child0, 1)
	NodeStyleSetFlexBasisPercent(rootChild1child0, 0)
	NodeStyleSetWidth(rootChild1child0, 20)
	NodeStyleSetMinWidth(rootChild1child0, 0)
	NodeStyleSetMinHeight(rootChild1child0, 0)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetFlexShrink(rootChild2, 1)
	NodeStyleSetFlexBasisPercent(rootChild2, 0)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}

func TestDisplay_none_with_position(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetPosition(rootChild1, EdgeTop, 10)
	NodeStyleSetDisplay(rootChild1, DisplayNone)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))
}
