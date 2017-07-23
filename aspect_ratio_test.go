package flex

import "testing"

func _measure(node *YGNode, width float32, widthMode YGMeasureMode,
	height float32, heightMode YGMeasureMode) YGSize {

	if widthMode != YGMeasureModeExactly {
		width = 50
	}
	if heightMode != YGMeasureModeExactly {
		height = 50
	}
	return YGSize{
		width:  width,
		height: height,
	}
}

func TestAspect_ratio_cross_defined(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_main_defined(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_both_dimensions_defined_row(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 100)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_both_dimensions_defined_column(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 100)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_align_stretch(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_flex_grow(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_flex_shrink(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 150)
	YGNodeStyleSetFlexShrink(root_child0, 1)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_basis(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetFlexBasis(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_absolute_layout_width_defined(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeLeft, 0)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 0)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_absolute_layout_height_defined(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeLeft, 0)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 0)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_with_max_cross_defined(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetMaxWidth(root_child0, 40)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_with_max_main_defined(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetMaxHeight(root_child0, 40)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_with_min_cross_defined(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 30)
	YGNodeStyleSetMinWidth(root_child0, 40)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_with_min_main_defined(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 30)
	YGNodeStyleSetMinHeight(root_child0, 40)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_double_cross(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 2)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_half_cross(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 100)
	YGNodeStyleSetAspectRatio(root_child0, 0.5)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_double_main(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 0.5)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_half_main(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 100)
	YGNodeStyleSetAspectRatio(root_child0, 2)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_with_measure_func(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeSetMeasureFunc(root_child0, _measure)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_width_height_flex_grow_row(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_width_height_flex_grow_column(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_height_as_flex_basis(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNew()
	YGNodeStyleSetHeight(root_child1, 100)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetAspectRatio(root_child1, 1)
	YGNodeInsertChild(root, root_child1, 1)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 75, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 75, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 125, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 125, YGNodeLayoutGetHeight(root_child1))

}

func TestAspect_ratio_width_as_flex_basis(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNew()
	YGNodeStyleSetWidth(root_child1, 100)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetAspectRatio(root_child1, 1)
	YGNodeInsertChild(root, root_child1, 1)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 75, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 75, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 125, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 125, YGNodeLayoutGetHeight(root_child1))

}

func TestAspect_ratio_overrides_flex_grow_row(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetAspectRatio(root_child0, 0.5)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_overrides_flex_grow_column(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetAspectRatio(root_child0, 2)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_left_right_absolute(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeLeft, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeRight, 10)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_top_bottom_absolute(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPosition(root_child0, YGEdgeLeft, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetPosition(root_child0, YGEdgeBottom, 10)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_width_overrides_align_stretch_row(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_height_overrides_align_stretch_column(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_allow_child_overflow_parent_size(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 4)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_defined_main_with_margin(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeStyleSetMargin(root_child0, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeRight, 10)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestAspect_ratio_defined_cross_with_margin(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetAspectRatio(root_child0, 1)
	YGNodeStyleSetMargin(root_child0, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeRight, 10)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}
