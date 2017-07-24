package flex

import "testing"

func _measure(node *Node, width float32, widthMode MeasureMode,
	height float32, heightMode MeasureMode) Size {

	if widthMode != MeasureModeExactly {
		width = 50
	}
	if heightMode != MeasureModeExactly {
		height = 50
	}
	return Size{
		Width:  width,
		Height: height,
	}
}

func TestAspect_ratio_cross_defined(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_main_defined(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_both_dimensions_defined_row(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_both_dimensions_defined_column(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_align_stretch(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_flex_grow(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_flex_shrink(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 150)
	NodeStyleSetFlexShrink(rootChild0, 1)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_basis(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_absolute_layout_width_defined(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeLeft, 0)
	NodeStyleSetPosition(rootChild0, EdgeTop, 0)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_absolute_layout_height_defined(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeLeft, 0)
	NodeStyleSetPosition(rootChild0, EdgeTop, 0)
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_with_max_cross_defined(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetMaxWidth(rootChild0, 40)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_with_max_main_defined(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetMaxHeight(rootChild0, 40)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_with_min_cross_defined(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 30)
	NodeStyleSetMinWidth(rootChild0, 40)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_with_min_main_defined(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetMinHeight(rootChild0, 40)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_double_cross(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_half_cross(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 100)
	NodeStyleSetAspectRatio(rootChild0, 0.5)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_double_main(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 0.5)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_half_main(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetAspectRatio(rootChild0, 2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_with_measure_func(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeSetMeasureFunc(rootChild0, _measure)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_width_height_flex_grow_row(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_width_height_flex_grow_column(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_height_as_flex_basis(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	NodeStyleSetHeight(rootChild1, 100)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetAspectRatio(rootChild1, 1)
	YGNodeInsertChild(root, rootChild1, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 75, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 75, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 75, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 125, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 125, NodeLayoutGetHeight(rootChild1))
}

func TestAspect_ratio_width_as_flex_basis(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	NodeStyleSetWidth(rootChild1, 100)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetAspectRatio(rootChild1, 1)
	YGNodeInsertChild(root, rootChild1, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 75, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 75, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 75, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 125, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 125, NodeLayoutGetHeight(rootChild1))
}

func TestAspect_ratio_overrides_flex_grow_row(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetAspectRatio(rootChild0, 0.5)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_overrides_flex_grow_column(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetAspectRatio(rootChild0, 2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_left_right_absolute(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeLeft, 10)
	NodeStyleSetPosition(rootChild0, EdgeTop, 10)
	NodeStyleSetPosition(rootChild0, EdgeRight, 10)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_top_bottom_absolute(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeLeft, 10)
	NodeStyleSetPosition(rootChild0, EdgeTop, 10)
	NodeStyleSetPosition(rootChild0, EdgeBottom, 10)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_width_overrides_align_stretch_row(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_height_overrides_align_stretch_column(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_allow_child_overflow_parent_size(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 4)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_defined_main_with_margin(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetHeight(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild0, EdgeRight, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}

func TestAspect_ratio_defined_cross_with_margin(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetAspectRatio(rootChild0, 1)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild0, EdgeRight, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))
}
