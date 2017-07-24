package flex

import (
	"testing"
)

func TestWrap_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 30)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 30)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 30)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 60, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 60, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 30, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 30, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild3))
}

func TestWrap_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 30)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 30)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 30)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild3))
}

func TestWrap_row_align_items_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignFlexEnd)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 30)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild3))
}

func TestWrap_row_align_items_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 30)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild3))
}

func TestFlex_wrap_children_with_min_main_overriding_flex_basis(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetMinWidth(rootChild0, 55)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexBasis(rootChild1, 50)
	NodeStyleSetMinWidth(rootChild1, 55)
	NodeStyleSetHeight(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 55, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 55, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 55, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 45, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 55, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))
}

func TestFlex_wrap_wrap_to_child_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetAlignItems(rootChild0, AlignFlexStart)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 100)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0_child0, 100)
	NodeStyleSetHeight(rootChild0Child0_child0, 100)
	NodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 100)
	NodeStyleSetHeight(rootChild1, 100)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))
}

func TestFlex_wrap_align_stretch_fits_one_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))
}

func TestWrap_reverse_row_align_content_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_align_content_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignCenter)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_single_line_different_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 300)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 300, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 90, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 120, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 300, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 270, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 240, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 210, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 180, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 150, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_align_content_stretch(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_align_content_space_around(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignSpaceAround)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_column_fixed_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 170, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 170, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 170, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 170, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 140, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestWrapped_row_within_align_items_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 150)
	NodeStyleSetHeight(rootChild0Child0, 80)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1, 80)
	NodeStyleSetHeight(rootChild0child1, 80)
	NodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 120, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0child1))
}

func TestWrapped_row_within_align_items_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 150)
	NodeStyleSetHeight(rootChild0Child0, 80)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1, 80)
	NodeStyleSetHeight(rootChild0child1, 80)
	NodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 120, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0child1))
}

func TestWrapped_row_within_align_items_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignFlexEnd)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 150)
	NodeStyleSetHeight(rootChild0Child0, 80)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1, 80)
	NodeStyleSetHeight(rootChild0child1, 80)
	NodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 150, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 120, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0child1))
}

func TestWrapped_column_max_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignContent(root, AlignCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 700)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 500)
	NodeStyleSetMaxHeight(rootChild0, 200)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild1, EdgeLeft, 20)
	NodeStyleSetMargin(rootChild1, EdgeTop, 20)
	NodeStyleSetMargin(rootChild1, EdgeRight, 20)
	NodeStyleSetMargin(rootChild1, EdgeBottom, 20)
	NodeStyleSetWidth(rootChild1, 200)
	NodeStyleSetHeight(rootChild1, 200)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 100)
	NodeStyleSetHeight(rootChild2, 100)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 700, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 250, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 200, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 250, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 420, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 200, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 700, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 350, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 300, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 250, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 180, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 200, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))
}

func TestWrapped_column_max_height_flex(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignContent(root, AlignCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 700)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexShrink(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 0)
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 500)
	NodeStyleSetMaxHeight(rootChild0, 200)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetFlexShrink(rootChild1, 1)
	NodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetMargin(rootChild1, EdgeLeft, 20)
	NodeStyleSetMargin(rootChild1, EdgeTop, 20)
	NodeStyleSetMargin(rootChild1, EdgeRight, 20)
	NodeStyleSetMargin(rootChild1, EdgeBottom, 20)
	NodeStyleSetWidth(rootChild1, 200)
	NodeStyleSetHeight(rootChild1, 200)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 100)
	NodeStyleSetHeight(rootChild2, 100)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 700, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 300, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 180, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 250, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 200, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 180, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 300, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 400, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 700, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 300, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 180, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 250, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 200, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 180, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 300, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 400, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))
}

func TestWrap_nodes_with_content_sizing_overflowing_margin(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 500)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	NodeStyleSetWidth(rootChild0, 85)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0_child0, 40)
	NodeStyleSetHeight(rootChild0Child0_child0, 40)
	NodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild0child1, EdgeRight, 10)
	NodeInsertChild(rootChild0, rootChild0child1, 1)

	rootChild0child1Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1Child0, 40)
	NodeStyleSetHeight(rootChild0child1Child0, 40)
	NodeInsertChild(rootChild0child1, rootChild0child1Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 500, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 85, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0child1Child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0child1Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 500, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 415, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 85, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 45, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 35, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0child1Child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0child1Child0))
}

func TestWrap_nodes_with_content_sizing_margin_cross(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 500)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	NodeStyleSetWidth(rootChild0, 70)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0_child0, 40)
	NodeStyleSetHeight(rootChild0Child0_child0, 40)
	NodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild0child1, EdgeTop, 10)
	NodeInsertChild(rootChild0, rootChild0child1, 1)

	rootChild0child1Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1Child0, 40)
	NodeStyleSetHeight(rootChild0child1Child0, 40)
	NodeInsertChild(rootChild0child1, rootChild0child1Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 500, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 70, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 90, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0child1Child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0child1Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 500, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 430, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 70, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 90, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 30, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0child1Child0.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0child1Child0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0child1Child0))
}
