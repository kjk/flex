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
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_main_defined(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_both_dimensions_defined_row(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 100)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_both_dimensions_defined_column(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 100)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_align_stretch(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_flex_grow(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_flex_shrink(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 150)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_basis(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetFlexBasis(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_absolute_layout_width_defined(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeLeft, 0)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 0)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_absolute_layout_height_defined(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeLeft, 0)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 0)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_with_max_cross_defined(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetMaxWidth(rootChild0, 40)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_with_max_main_defined(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetMaxHeight(rootChild0, 40)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_with_min_cross_defined(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 30)
	YGNodeStyleSetMinWidth(rootChild0, 40)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_with_min_main_defined(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 30)
	YGNodeStyleSetMinHeight(rootChild0, 40)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_double_cross(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_half_cross(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 100)
	YGNodeStyleSetAspectRatio(rootChild0, 0.5)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_double_main(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 0.5)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_half_main(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 100)
	YGNodeStyleSetAspectRatio(rootChild0, 2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_with_measure_func(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeSetMeasureFunc(rootChild0, _measure)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_width_height_flex_grow_row(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_width_height_flex_grow_column(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_height_as_flex_basis(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	YGNodeStyleSetHeight(rootChild1, 100)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeStyleSetAspectRatio(rootChild1, 1)
	YGNodeInsertChild(root, rootChild1, 1)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 75, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 75, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 125, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 125, YGNodeLayoutGetHeight(rootChild1))

}

func TestAspect_ratio_width_as_flex_basis(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	YGNodeStyleSetWidth(rootChild1, 100)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeStyleSetAspectRatio(rootChild1, 1)
	YGNodeInsertChild(root, rootChild1, 1)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 75, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 75, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 125, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 125, YGNodeLayoutGetHeight(rootChild1))

}

func TestAspect_ratio_overrides_flex_grow_row(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetAspectRatio(rootChild0, 0.5)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_overrides_flex_grow_column(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetAspectRatio(rootChild0, 2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_left_right_absolute(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeLeft, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeRight, 10)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_top_bottom_absolute(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetPositionType(rootChild0, PositionTypeAbsolute)
	YGNodeStyleSetPosition(rootChild0, EdgeLeft, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 10)
	YGNodeStyleSetPosition(rootChild0, EdgeBottom, 10)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_width_overrides_align_stretch_row(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_height_overrides_align_stretch_column(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_allow_child_overflow_parent_size(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 4)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_defined_main_with_margin(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, EdgeLeft, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeRight, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestAspect_ratio_defined_cross_with_margin(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetAspectRatio(rootChild0, 1)
	YGNodeStyleSetMargin(rootChild0, EdgeLeft, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeRight, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}
