package flex

import (
	"testing"
)

func TestWrap_column(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetHeight(rootChild0, 30)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 30)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 30)
	YGNodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 30)
	YGNodeStyleSetHeight(rootChild3, 30)
	YGNodeInsertChild(root, rootChild3, 3)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))
}

func TestWrap_row(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetHeight(rootChild0, 30)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 30)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 30)
	YGNodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 30)
	YGNodeStyleSetHeight(rootChild3, 30)
	YGNodeInsertChild(root, rootChild3, 3)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))
}

func TestWrap_row_align_items_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignFlexEnd)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 30)
	YGNodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 30)
	YGNodeStyleSetHeight(rootChild3, 30)
	YGNodeInsertChild(root, rootChild3, 3)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))
}

func TestWrap_row_align_items_center(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 30)
	YGNodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 30)
	YGNodeStyleSetHeight(rootChild3, 30)
	YGNodeInsertChild(root, rootChild3, 3)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))
}

func TestFlex_wrap_children_with_min_main_overriding_flex_basis(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexBasis(rootChild0, 50)
	YGNodeStyleSetMinWidth(rootChild0, 55)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexBasis(rootChild1, 50)
	YGNodeStyleSetMinWidth(rootChild1, 55)
	YGNodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))
}

func TestFlex_wrap_wrap_to_child_height(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	YGNodeStyleSetAlignItems(rootChild0, AlignFlexStart)
	YGNodeStyleSetFlexWrap(rootChild0, WrapWrap)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0, 100)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0_child0, 100)
	YGNodeStyleSetHeight(rootChild0Child0_child0, 100)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 100)
	YGNodeStyleSetHeight(rootChild1, 100)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))
}

func TestFlex_wrap_align_stretch_fits_one_row(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))
}

func TestWrap_reverse_row_align_content_flex_start(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, WrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 30)
	YGNodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 30)
	YGNodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild4, 30)
	YGNodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_align_content_center(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignCenter)
	YGNodeStyleSetFlexWrap(root, WrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 30)
	YGNodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 30)
	YGNodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild4, 30)
	YGNodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_single_line_different_size(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, WrapWrapReverse)
	YGNodeStyleSetWidth(root, 300)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 30)
	YGNodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 30)
	YGNodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild4, 30)
	YGNodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 270, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 240, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 210, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 180, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 150, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_align_content_stretch(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 30)
	YGNodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 30)
	YGNodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild4, 30)
	YGNodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_align_content_space_around(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignSpaceAround)
	YGNodeStyleSetFlexWrap(root, WrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 30)
	YGNodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 30)
	YGNodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild4, 30)
	YGNodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_column_fixed_size(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexWrap(root, WrapWrapReverse)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 30)
	YGNodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 30)
	YGNodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild4, 30)
	YGNodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 140, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrapped_row_within_align_items_center(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(rootChild0, WrapWrap)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0, 150)
	YGNodeStyleSetHeight(rootChild0Child0, 80)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0child1, 80)
	YGNodeStyleSetHeight(rootChild0child1, 80)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))
}

func TestWrapped_row_within_align_items_flex_start(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(rootChild0, WrapWrap)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0, 150)
	YGNodeStyleSetHeight(rootChild0Child0, 80)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0child1, 80)
	YGNodeStyleSetHeight(rootChild0child1, 80)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))
}

func TestWrapped_row_within_align_items_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignFlexEnd)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(rootChild0, WrapWrap)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0, 150)
	YGNodeStyleSetHeight(rootChild0Child0, 80)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0child1, 80)
	YGNodeStyleSetHeight(rootChild0child1, 80)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))
}

func TestWrapped_column_max_height(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetAlignContent(root, AlignCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	YGNodeStyleSetWidth(root, 700)
	YGNodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 100)
	YGNodeStyleSetHeight(rootChild0, 500)
	YGNodeStyleSetMaxHeight(rootChild0, 200)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetMargin(rootChild1, EdgeLeft, 20)
	YGNodeStyleSetMargin(rootChild1, EdgeTop, 20)
	YGNodeStyleSetMargin(rootChild1, EdgeRight, 20)
	YGNodeStyleSetMargin(rootChild1, EdgeBottom, 20)
	YGNodeStyleSetWidth(rootChild1, 200)
	YGNodeStyleSetHeight(rootChild1, 200)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 100)
	YGNodeStyleSetHeight(rootChild2, 100)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 200, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 250, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 420, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 350, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 250, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 180, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}

func TestWrapped_column_max_height_flex(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetAlignContent(root, AlignCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	YGNodeStyleSetWidth(root, 700)
	YGNodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 0)
	YGNodeStyleSetWidth(rootChild0, 100)
	YGNodeStyleSetHeight(rootChild0, 500)
	YGNodeStyleSetMaxHeight(rootChild0, 200)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeStyleSetFlexShrink(rootChild1, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 0)
	YGNodeStyleSetMargin(rootChild1, EdgeLeft, 20)
	YGNodeStyleSetMargin(rootChild1, EdgeTop, 20)
	YGNodeStyleSetMargin(rootChild1, EdgeRight, 20)
	YGNodeStyleSetMargin(rootChild1, EdgeBottom, 20)
	YGNodeStyleSetWidth(rootChild1, 200)
	YGNodeStyleSetHeight(rootChild1, 200)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 100)
	YGNodeStyleSetHeight(rootChild2, 100)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 400, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 400, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}

func TestWrap_nodes_with_content_sizing_overflowing_margin(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(root, 500)
	YGNodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(rootChild0, WrapWrap)
	YGNodeStyleSetWidth(rootChild0, 85)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0_child0, 40)
	YGNodeStyleSetHeight(rootChild0Child0_child0, 40)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	YGNodeStyleSetMargin(rootChild0child1, EdgeRight, 10)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)

	rootChild0child1_child0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0child1_child0, 40)
	YGNodeStyleSetHeight(rootChild0child1_child0, 40)
	YGNodeInsertChild(rootChild0child1, rootChild0child1_child0, 0)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 85, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1_child0))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 415, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 85, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 35, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1_child0))
}

func TestWrap_nodes_with_content_sizing_margin_cross(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(root, 500)
	YGNodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(rootChild0, WrapWrap)
	YGNodeStyleSetWidth(rootChild0, 70)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0_child0, 40)
	YGNodeStyleSetHeight(rootChild0Child0_child0, 40)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	YGNodeStyleSetMargin(rootChild0child1, EdgeTop, 10)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)

	rootChild0child1_child0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0child1_child0, 40)
	YGNodeStyleSetHeight(rootChild0child1_child0, 40)
	YGNodeInsertChild(rootChild0child1, rootChild0child1_child0, 0)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1_child0))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 430, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1_child0))
}
