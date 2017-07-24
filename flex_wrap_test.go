package flex

import (
	"testing"
)

func TestWrap_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetHeight(30)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(30)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(30)
	rootChild2.StyleSetHeight(30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(30)
	rootChild3.StyleSetHeight(30)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 60, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild3.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 60, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild3.LayoutGetHeight())
}

func TestWrap_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetHeight(30)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(30)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(30)
	rootChild2.StyleSetHeight(30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(30)
	rootChild3.StyleSetHeight(30)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 60, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild3.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 60, root.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild3.LayoutGetHeight())
}

func TestWrap_row_align_items_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignFlexEnd)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(30)
	rootChild2.StyleSetHeight(30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(30)
	rootChild3.StyleSetHeight(30)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 60, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild3.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 60, root.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild3.LayoutGetHeight())
}

func TestWrap_row_align_items_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(30)
	rootChild2.StyleSetHeight(30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(30)
	rootChild3.StyleSetHeight(30)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 60, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild3.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 60, root.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild3.LayoutGetHeight())
}

func TestFlex_wrap_children_with_min_main_overriding_flex_basis(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexBasis(50)
	NodeStyleSetMinWidth(rootChild0, 55)
	rootChild0.StyleSetHeight(50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexBasis(50)
	NodeStyleSetMinWidth(rootChild1, 55)
	rootChild1.StyleSetHeight(50)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 55, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 55, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 55, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 55, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())
}

func TestFlex_wrap_wrap_to_child_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexDirection(FlexDirectionRow)
	rootChild0.StyleSetAlignItems(AlignFlexStart)
	rootChild0.StyleSetFlexWrap(WrapWrap)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidth(100)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0Child0 := NewNodeWithConfig(config)
	rootChild0Child0Child0.StyleSetWidth(100)
	rootChild0Child0Child0.StyleSetHeight(100)
	NodeInsertChild(rootChild0Child0, rootChild0Child0Child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(100)
	rootChild1.StyleSetHeight(100)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0Child0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0Child0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())
}

func TestFlex_wrap_align_stretch_fits_one_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(150)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(50)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 150, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 150, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())
}

func TestWrap_reverse_row_align_content_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetFlexWrap(WrapWrapReverse)
	root.StyleSetWidth(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(30)
	rootChild2.StyleSetHeight(30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(30)
	rootChild3.StyleSetHeight(40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	rootChild4.StyleSetWidth(30)
	rootChild4.StyleSetHeight(50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 80, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 80, root.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())
}

func TestWrap_reverse_row_align_content_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignContent(AlignCenter)
	root.StyleSetFlexWrap(WrapWrapReverse)
	root.StyleSetWidth(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(30)
	rootChild2.StyleSetHeight(30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(30)
	rootChild3.StyleSetHeight(40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	rootChild4.StyleSetWidth(30)
	rootChild4.StyleSetHeight(50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 80, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 80, root.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())
}

func TestWrap_reverse_row_single_line_different_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetFlexWrap(WrapWrapReverse)
	root.StyleSetWidth(300)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(30)
	rootChild2.StyleSetHeight(30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(30)
	rootChild3.StyleSetHeight(40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	rootChild4.StyleSetWidth(30)
	rootChild4.StyleSetHeight(50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 300, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 120, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 300, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 270, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 240, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 210, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 180, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 150, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())
}

func TestWrap_reverse_row_align_content_stretch(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignContent(AlignStretch)
	root.StyleSetFlexWrap(WrapWrapReverse)
	root.StyleSetWidth(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(30)
	rootChild2.StyleSetHeight(30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(30)
	rootChild3.StyleSetHeight(40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	rootChild4.StyleSetWidth(30)
	rootChild4.StyleSetHeight(50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 80, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 80, root.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())
}

func TestWrap_reverse_row_align_content_space_around(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignContent(AlignSpaceAround)
	root.StyleSetFlexWrap(WrapWrapReverse)
	root.StyleSetWidth(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(30)
	rootChild2.StyleSetHeight(30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(30)
	rootChild3.StyleSetHeight(40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	rootChild4.StyleSetWidth(30)
	rootChild4.StyleSetHeight(50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 80, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 80, root.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())
}

func TestWrap_reverse_column_fixed_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetFlexWrap(WrapWrapReverse)
	root.StyleSetWidth(200)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(30)
	rootChild2.StyleSetHeight(30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(30)
	rootChild3.StyleSetHeight(40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	rootChild4.StyleSetWidth(30)
	rootChild4.StyleSetHeight(50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 170, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 170, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 170, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 170, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 140, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild4.LayoutGetHeight())
}

func TestWrapped_row_within_align_items_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexDirection(FlexDirectionRow)
	rootChild0.StyleSetFlexWrap(WrapWrap)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidth(150)
	rootChild0Child0.StyleSetHeight(80)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	rootChild0child1.StyleSetWidth(80)
	rootChild0child1.StyleSetHeight(80)
	NodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 160, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 160, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 120, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetHeight())
}

func TestWrapped_row_within_align_items_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexDirection(FlexDirectionRow)
	rootChild0.StyleSetFlexWrap(WrapWrap)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidth(150)
	rootChild0Child0.StyleSetHeight(80)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	rootChild0child1.StyleSetWidth(80)
	rootChild0child1.StyleSetHeight(80)
	NodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 160, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 160, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 120, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetHeight())
}

func TestWrapped_row_within_align_items_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignFlexEnd)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexDirection(FlexDirectionRow)
	rootChild0.StyleSetFlexWrap(WrapWrap)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidth(150)
	rootChild0Child0.StyleSetHeight(80)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	rootChild0child1.StyleSetWidth(80)
	rootChild0child1.StyleSetHeight(80)
	NodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 160, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 160, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 120, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetHeight())
}

func TestWrapped_column_max_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignContent(AlignCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(700)
	root.StyleSetHeight(500)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(100)
	rootChild0.StyleSetHeight(500)
	NodeStyleSetMaxHeight(rootChild0, 200)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetMargin(EdgeLeft, 20)
	rootChild1.StyleSetMargin(EdgeTop, 20)
	rootChild1.StyleSetMargin(EdgeRight, 20)
	rootChild1.StyleSetMargin(EdgeBottom, 20)
	rootChild1.StyleSetWidth(200)
	rootChild1.StyleSetHeight(200)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(100)
	rootChild2.StyleSetHeight(100)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 700, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 250, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 200, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 250, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 420, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 200, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 700, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 350, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 300, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 250, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 180, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 200, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())
}

func TestWrapped_column_max_height_flex(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignContent(AlignCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(700)
	root.StyleSetHeight(500)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexShrink(1)
	rootChild0.StyleSetFlexBasisPercent(0)
	rootChild0.StyleSetWidth(100)
	rootChild0.StyleSetHeight(500)
	NodeStyleSetMaxHeight(rootChild0, 200)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(1)
	rootChild1.StyleSetFlexShrink(1)
	rootChild1.StyleSetFlexBasisPercent(0)
	rootChild1.StyleSetMargin(EdgeLeft, 20)
	rootChild1.StyleSetMargin(EdgeTop, 20)
	rootChild1.StyleSetMargin(EdgeRight, 20)
	rootChild1.StyleSetMargin(EdgeBottom, 20)
	rootChild1.StyleSetWidth(200)
	rootChild1.StyleSetHeight(200)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(100)
	rootChild2.StyleSetHeight(100)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 700, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 300, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 180, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 250, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 200, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 180, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 300, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 400, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 700, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 300, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 180, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 250, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 200, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 180, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 300, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 400, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())
}

func TestWrap_nodes_with_content_sizing_overflowing_margin(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(500)
	root.StyleSetHeight(500)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexDirection(FlexDirectionRow)
	rootChild0.StyleSetFlexWrap(WrapWrap)
	rootChild0.StyleSetWidth(85)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0Child0 := NewNodeWithConfig(config)
	rootChild0Child0Child0.StyleSetWidth(40)
	rootChild0Child0Child0.StyleSetHeight(40)
	NodeInsertChild(rootChild0Child0, rootChild0Child0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	rootChild0child1.StyleSetMargin(EdgeRight, 10)
	NodeInsertChild(rootChild0, rootChild0child1, 1)

	rootChild0child1Child0 := NewNodeWithConfig(config)
	rootChild0child1Child0.StyleSetWidth(40)
	rootChild0child1Child0.StyleSetHeight(40)
	NodeInsertChild(rootChild0child1, rootChild0child1Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 500, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 85, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0Child0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0child1Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0child1Child0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 500, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 415, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 85, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0Child0Child0.LayoutGetHeight())

	assertFloatEqual(t, 35, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0child1Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0child1Child0.LayoutGetHeight())
}

func TestWrap_nodes_with_content_sizing_margin_cross(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(500)
	root.StyleSetHeight(500)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexDirection(FlexDirectionRow)
	rootChild0.StyleSetFlexWrap(WrapWrap)
	rootChild0.StyleSetWidth(70)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0Child0 := NewNodeWithConfig(config)
	rootChild0Child0Child0.StyleSetWidth(40)
	rootChild0Child0Child0.StyleSetHeight(40)
	NodeInsertChild(rootChild0Child0, rootChild0Child0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	rootChild0child1.StyleSetMargin(EdgeTop, 10)
	NodeInsertChild(rootChild0, rootChild0child1, 1)

	rootChild0child1Child0 := NewNodeWithConfig(config)
	rootChild0child1Child0.StyleSetWidth(40)
	rootChild0child1Child0.StyleSetHeight(40)
	NodeInsertChild(rootChild0child1, rootChild0child1Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 500, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 70, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 90, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0Child0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0child1Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0child1Child0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 500, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 430, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 70, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 90, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0Child0Child0.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0child1Child0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0child1Child0.LayoutGetHeight())
}
