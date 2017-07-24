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
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_main_defined(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_both_dimensions_defined_row(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(100)
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_both_dimensions_defined_column(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(100)
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_align_stretch(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_flex_grow(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_flex_shrink(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(150)
	rootChild0.StyleSetFlexShrink(1)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_basis(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetFlexBasis(50)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_absolute_layout_width_defined(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeLeft, 0)
	rootChild0.StyleSetPosition(EdgeTop, 0)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_absolute_layout_height_defined(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeLeft, 0)
	rootChild0.StyleSetPosition(EdgeTop, 0)
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_with_max_cross_defined(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetMaxWidth(40)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_with_max_main_defined(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetMaxHeight(40)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_with_min_cross_defined(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(30)
	rootChild0.StyleSetMinWidth(40)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_with_min_main_defined(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(30)
	rootChild0.StyleSetMinHeight(40)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_double_cross(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetAspectRatio(2)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_half_cross(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(100)
	rootChild0.StyleSetAspectRatio(0.5)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_double_main(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetAspectRatio(0.5)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_half_main(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(100)
	rootChild0.StyleSetAspectRatio(2)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_with_measure_func(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.SetMeasureFunc(_measure)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_width_height_flex_grow_row(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(200)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_width_height_flex_grow_column(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(200)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_height_as_flex_basis(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNode()
	rootChild1.StyleSetHeight(100)
	rootChild1.StyleSetFlexGrow(1)
	rootChild1.StyleSetAspectRatio(1)
	root.InsertChild(rootChild1, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 75, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 75, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 75, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 125, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 125, rootChild1.LayoutGetHeight())
}

func TestAspect_ratio_width_as_flex_basis(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNode()
	rootChild1.StyleSetWidth(100)
	rootChild1.StyleSetFlexGrow(1)
	rootChild1.StyleSetAspectRatio(1)
	root.InsertChild(rootChild1, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 75, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 75, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 75, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 125, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 125, rootChild1.LayoutGetHeight())
}

func TestAspect_ratio_overrides_flex_grow_row(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetAspectRatio(0.5)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_overrides_flex_grow_column(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetAspectRatio(2)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_left_right_absolute(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeLeft, 10)
	rootChild0.StyleSetPosition(EdgeTop, 10)
	rootChild0.StyleSetPosition(EdgeRight, 10)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_top_bottom_absolute(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPosition(EdgeLeft, 10)
	rootChild0.StyleSetPosition(EdgeTop, 10)
	rootChild0.StyleSetPosition(EdgeBottom, 10)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_width_overrides_align_stretch_row(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_height_overrides_align_stretch_column(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetAspectRatio(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_allow_child_overflow_parent_size(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetAspectRatio(4)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_defined_main_with_margin(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetHeight(50)
	rootChild0.StyleSetAspectRatio(1)
	rootChild0.StyleSetMargin(EdgeLeft, 10)
	rootChild0.StyleSetMargin(EdgeRight, 10)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}

func TestAspect_ratio_defined_cross_with_margin(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetAspectRatio(1)
	rootChild0.StyleSetMargin(EdgeLeft, 10)
	rootChild0.StyleSetMargin(EdgeRight, 10)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())
}
