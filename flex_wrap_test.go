package flex

import (
	"testing"
)

func TestWrap_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetHeight(root_child0, 30)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 30)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 30)
	YGNodeStyleSetHeight(root_child2, 30)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 30)
	YGNodeStyleSetHeight(root_child3, 30)
	YGNodeInsertChild(root, root_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child3))

}

func TestWrap_row(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetHeight(root_child0, 30)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 30)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 30)
	YGNodeStyleSetHeight(root_child2, 30)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 30)
	YGNodeStyleSetHeight(root_child3, 30)
	YGNodeInsertChild(root, root_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child3))

}

func TestWrap_row_align_items_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignFlexEnd)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 30)
	YGNodeStyleSetHeight(root_child2, 30)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 30)
	YGNodeStyleSetHeight(root_child3, 30)
	YGNodeInsertChild(root, root_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child3))

}

func TestWrap_row_align_items_center(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 30)
	YGNodeStyleSetHeight(root_child2, 30)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 30)
	YGNodeStyleSetHeight(root_child3, 30)
	YGNodeInsertChild(root, root_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child3))

}

func TestFlex_wrap_children_with_min_main_overriding_flex_basis(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexBasis(root_child0, 50)
	YGNodeStyleSetMinWidth(root_child0, 55)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexBasis(root_child1, 50)
	YGNodeStyleSetMinWidth(root_child1, 55)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

}

func TestFlex_wrap_wrap_to_child_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root_child0, AlignFlexStart)
	YGNodeStyleSetFlexWrap(root_child0, YGWrapWrap)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0, 100)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0_child0, 100)
	YGNodeStyleSetHeight(root_child0_child0_child0, 100)
	YGNodeInsertChild(root_child0_child0, root_child0_child0_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 100)
	YGNodeStyleSetHeight(root_child1, 100)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

}

func TestFlex_wrap_align_stretch_fits_one_row(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

}

func TestWrap_reverse_row_align_content_flex_start(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 30)
	YGNodeStyleSetHeight(root_child2, 30)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 30)
	YGNodeStyleSetHeight(root_child3, 40)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 30)
	YGNodeStyleSetHeight(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

}

func TestWrap_reverse_row_align_content_center(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignCenter)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 30)
	YGNodeStyleSetHeight(root_child2, 30)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 30)
	YGNodeStyleSetHeight(root_child3, 40)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 30)
	YGNodeStyleSetHeight(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

}

func TestWrap_reverse_row_single_line_different_size(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 300)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 30)
	YGNodeStyleSetHeight(root_child2, 30)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 30)
	YGNodeStyleSetHeight(root_child3, 40)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 30)
	YGNodeStyleSetHeight(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 270, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 240, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 210, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 180, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 150, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

}

func TestWrap_reverse_row_align_content_stretch(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 30)
	YGNodeStyleSetHeight(root_child2, 30)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 30)
	YGNodeStyleSetHeight(root_child3, 40)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 30)
	YGNodeStyleSetHeight(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

}

func TestWrap_reverse_row_align_content_space_around(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignSpaceAround)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 30)
	YGNodeStyleSetHeight(root_child2, 30)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 30)
	YGNodeStyleSetHeight(root_child3, 40)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 30)
	YGNodeStyleSetHeight(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

}

func TestWrap_reverse_column_fixed_size(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexWrap(root, YGWrapWrapReverse)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 30)
	YGNodeStyleSetHeight(root_child2, 30)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 30)
	YGNodeStyleSetHeight(root_child3, 40)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 30)
	YGNodeStyleSetHeight(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 140, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

}

func TestWrapped_row_within_align_items_center(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root_child0, YGWrapWrap)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0, 150)
	YGNodeStyleSetHeight(root_child0_child0, 80)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child1, 80)
	YGNodeStyleSetHeight(root_child0_child1, 80)
	YGNodeInsertChild(root_child0, root_child0_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child1))

}

func TestWrapped_row_within_align_items_flex_start(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root_child0, YGWrapWrap)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0, 150)
	YGNodeStyleSetHeight(root_child0_child0, 80)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child1, 80)
	YGNodeStyleSetHeight(root_child0_child1, 80)
	YGNodeInsertChild(root_child0, root_child0_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child1))

}

func TestWrapped_row_within_align_items_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignFlexEnd)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root_child0, YGWrapWrap)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0, 150)
	YGNodeStyleSetHeight(root_child0_child0, 80)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child1, 80)
	YGNodeStyleSetHeight(root_child0_child1, 80)
	YGNodeInsertChild(root_child0, root_child0_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0_child1))

}

func TestWrapped_column_max_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignContent(root, AlignCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 700)
	YGNodeStyleSetHeight(root, 500)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 100)
	YGNodeStyleSetHeight(root_child0, 500)
	YGNodeStyleSetMaxHeight(root_child0, 200)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root_child1, YGEdgeLeft, 20)
	YGNodeStyleSetMargin(root_child1, YGEdgeTop, 20)
	YGNodeStyleSetMargin(root_child1, YGEdgeRight, 20)
	YGNodeStyleSetMargin(root_child1, YGEdgeBottom, 20)
	YGNodeStyleSetWidth(root_child1, 200)
	YGNodeStyleSetHeight(root_child1, 200)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 100)
	YGNodeStyleSetHeight(root_child2, 100)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 200, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 250, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 420, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 350, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 250, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 180, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

}

func TestWrapped_column_max_height_flex(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignContent(root, AlignCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 700)
	YGNodeStyleSetHeight(root, 500)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexShrink(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 0)
	YGNodeStyleSetWidth(root_child0, 100)
	YGNodeStyleSetHeight(root_child0, 500)
	YGNodeStyleSetMaxHeight(root_child0, 200)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetFlexShrink(root_child1, 1)
	YGNodeStyleSetFlexBasisPercent(root_child1, 0)
	YGNodeStyleSetMargin(root_child1, YGEdgeLeft, 20)
	YGNodeStyleSetMargin(root_child1, YGEdgeTop, 20)
	YGNodeStyleSetMargin(root_child1, YGEdgeRight, 20)
	YGNodeStyleSetMargin(root_child1, YGEdgeBottom, 20)
	YGNodeStyleSetWidth(root_child1, 200)
	YGNodeStyleSetHeight(root_child1, 200)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 100)
	YGNodeStyleSetHeight(root_child2, 100)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 400, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 400, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

}

func TestWrap_nodes_with_content_sizing_overflowing_margin(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 500)
	YGNodeStyleSetHeight(root, 500)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root_child0, YGWrapWrap)
	YGNodeStyleSetWidth(root_child0, 85)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0_child0, 40)
	YGNodeStyleSetHeight(root_child0_child0_child0, 40)
	YGNodeInsertChild(root_child0_child0, root_child0_child0_child0, 0)

	root_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root_child0_child1, YGEdgeRight, 10)
	YGNodeInsertChild(root_child0, root_child0_child1, 1)

	root_child0_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child1_child0, 40)
	YGNodeStyleSetHeight(root_child0_child1_child0, 40)
	YGNodeInsertChild(root_child0_child1, root_child0_child1_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 85, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child1_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 415, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 85, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child0_child0))

	assertFloatEqual(t, 35, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child1_child0))

}

func TestWrap_nodes_with_content_sizing_margin_cross(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 500)
	YGNodeStyleSetHeight(root, 500)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root_child0, YGWrapWrap)
	YGNodeStyleSetWidth(root_child0, 70)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0_child0, 40)
	YGNodeStyleSetHeight(root_child0_child0_child0, 40)
	YGNodeInsertChild(root_child0_child0, root_child0_child0_child0, 0)

	root_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root_child0_child1, YGEdgeTop, 10)
	YGNodeInsertChild(root_child0, root_child0_child1, 1)

	root_child0_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child1_child0, 40)
	YGNodeStyleSetHeight(root_child0_child1_child0, 40)
	YGNodeInsertChild(root_child0_child1, root_child0_child1_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 70, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child1_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 430, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 70, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child0_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0_child1_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0_child1_child0))

}
