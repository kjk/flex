package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type measureConstraint struct {
	width      float32
	widthMode  MeasureMode
	height     float32
	heightMode MeasureMode
}

type measureConstraintList struct {
	length      int
	constraints []measureConstraint
}

func _measure2(node *Node,
	width float32,
	widthMode MeasureMode,
	height float32,
	heightMode MeasureMode) Size {
	constraintList := node.Context.(*measureConstraintList)
	constraints := constraintList.constraints
	currentIndex := constraintList.length
	(&constraints[currentIndex]).width = width
	(&constraints[currentIndex]).widthMode = widthMode
	(&constraints[currentIndex]).height = height
	(&constraints[currentIndex]).heightMode = heightMode
	constraintList.length = currentIndex + 1

	if widthMode == MeasureModeUndefined {
		width = 10
	}

	if heightMode == MeasureModeUndefined {
		height = 10
	} else {
		height = width // TODO:: is it a bug in tests ?
	}
	return Size{
		Width:  width,
		Height: height,
	}
}

func TestExactly_measure_stretched_child_column(t *testing.T) {
	constraintList := measureConstraintList{
		length:      0,
		constraints: make([]measureConstraint, 10, 10),
	}

	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	rootChild0.Context = &constraintList
	NodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].width)
	assert.Equal(t, MeasureModeExactly, constraintList.constraints[0].widthMode)

}

func TestExactly_measure_stretched_child_row(t *testing.T) {
	constraintList := measureConstraintList{
		length:      0,
		constraints: make([]measureConstraint, 10, 10),
	}

	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	rootChild0.Context = &constraintList
	NodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, MeasureModeExactly, constraintList.constraints[0].heightMode)

}

func TestAt_most_main_axis_column(t *testing.T) {
	constraintList := measureConstraintList{
		length:      0,
		constraints: make([]measureConstraint, 10, 10),
	}

	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	rootChild0.Context = &constraintList
	NodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, MeasureModeAtMost, constraintList.constraints[0].heightMode)

}

func TestAt_most_cross_axis_column(t *testing.T) {
	constraintList := measureConstraintList{
		length:      0,
		constraints: make([]measureConstraint, 10, 10),
	}

	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	rootChild0.Context = &constraintList
	NodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].width)
	assert.Equal(t, MeasureModeAtMost, constraintList.constraints[0].widthMode)

}

func TestAt_most_main_axis_row(t *testing.T) {
	constraintList := measureConstraintList{
		length:      0,
		constraints: make([]measureConstraint, 10, 10),
	}

	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	rootChild0.Context = &constraintList
	NodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].width)
	assert.Equal(t, MeasureModeAtMost, constraintList.constraints[0].widthMode)

}

func TestAt_most_cross_axis_row(t *testing.T) {
	constraintList := measureConstraintList{
		length:      0,
		constraints: make([]measureConstraint, 10, 10),
	}

	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	rootChild0.Context = &constraintList
	NodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, MeasureModeAtMost, constraintList.constraints[0].heightMode)

}

func TestFlex_child(t *testing.T) {
	constraintList := measureConstraintList{
		length:      0,
		constraints: make([]measureConstraint, 10, 10),
	}

	root := NewNode()
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetFlexGrow(rootChild0, 1)
	rootChild0.Context = &constraintList
	NodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 2, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, MeasureModeAtMost, constraintList.constraints[0].heightMode)

	assertFloatEqual(t, 100, constraintList.constraints[1].height)
	assert.Equal(t, MeasureModeExactly, constraintList.constraints[1].heightMode)

}

func TestFlex_child_with_flex_basis(t *testing.T) {
	constraintList := measureConstraintList{
		length:      0,
		constraints: make([]measureConstraint, 10, 10),
	}

	root := NewNode()
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 0)
	rootChild0.Context = &constraintList
	NodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, MeasureModeExactly, constraintList.constraints[0].heightMode)

}

func TestOverflow_scroll_column(t *testing.T) {
	constraintList := measureConstraintList{
		length:      0,
		constraints: make([]measureConstraint, 10, 10),
	}

	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetOverflow(root, OverflowScroll)
	NodeStyleSetHeight(root, 100)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNode()
	rootChild0.Context = &constraintList
	NodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].width)
	assert.Equal(t, MeasureModeAtMost, constraintList.constraints[0].widthMode)

	assert.True(t, FloatIsUndefined(constraintList.constraints[0].height))
	assert.Equal(t, MeasureModeUndefined, constraintList.constraints[0].heightMode)

}

func TestOverflow_scroll_row(t *testing.T) {
	constraintList := measureConstraintList{
		length:      0,
		constraints: make([]measureConstraint, 10, 10),
	}

	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetOverflow(root, OverflowScroll)
	NodeStyleSetHeight(root, 100)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNode()
	rootChild0.Context = &constraintList
	NodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assert.True(t, FloatIsUndefined(constraintList.constraints[0].width))
	assert.Equal(t, MeasureModeUndefined, constraintList.constraints[0].widthMode)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, MeasureModeAtMost, constraintList.constraints[0].heightMode)

}
