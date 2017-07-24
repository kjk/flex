package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func measureMax(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	measureCount := node.Context.(int)
	measureCount++
	node.Context = measureCount

	if widthMode == MeasureModeUndefined {
		width = 10
	}
	if heightMode == MeasureModeUndefined {
		height = 10
	}
	return Size{
		Width:  width,
		Height: height,
	}
}

func measureMin(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	measureCount := node.Context.(int)
	measureCount++
	node.Context = measureCount

	if widthMode == MeasureModeUndefined || (widthMode == MeasureModeAtMost && width > 10) {
		width = 10
	}
	if heightMode == MeasureModeUndefined || (heightMode == MeasureModeAtMost && height > 10) {
		height = 10
	}
	return Size{
		Width:  width,
		Height: height,
	}
}

func measure8449(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	measureCount, ok := node.Context.(int)
	if ok {
		measureCount++
		node.Context = measureCount
	}

	return Size{
		Width:  84,
		Height: 49,
	}
}

func TestMeasure_once_single_flexible_child(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	measureCount := 0
	rootChild0.Context = measureCount
	rootChild0.SetMeasureFunc(measureMax)
	rootChild0.StyleSetFlexGrow(1)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	measureCount = rootChild0.Context.(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_same_exact_width_larger_than_needed_height(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	measureCount := 0
	rootChild0.Context = measureCount
	rootChild0.SetMeasureFunc(measureMin)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, 100, 100, DirectionLTR)
	NodeCalculateLayout(root, 100, 50, DirectionLTR)

	measureCount = rootChild0.Context.(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_same_atmost_width_larger_than_needed_height(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)

	rootChild0 := NewNode()
	measureCount := 0
	rootChild0.Context = measureCount
	rootChild0.SetMeasureFunc(measureMin)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, 100, 100, DirectionLTR)
	NodeCalculateLayout(root, 100, 50, DirectionLTR)

	measureCount = rootChild0.Context.(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_computed_width_larger_than_needed_height(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)

	rootChild0 := NewNode()
	measureCount := 0
	rootChild0.Context = measureCount
	rootChild0.SetMeasureFunc(measureMin)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, 100, 100, DirectionLTR)
	root.StyleSetAlignItems(AlignStretch)
	NodeCalculateLayout(root, 10, 50, DirectionLTR)

	measureCount = rootChild0.Context.(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_atmost_computed_width_undefined_height(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)

	rootChild0 := NewNode()
	measureCount := 0
	rootChild0.Context = measureCount
	rootChild0.SetMeasureFunc(measureMin)
	root.InsertChild(rootChild0, 0)

	NodeCalculateLayout(root, 100, Undefined, DirectionLTR)
	NodeCalculateLayout(root, 10, Undefined, DirectionLTR)

	measureCount = rootChild0.Context.(int)
	assert.Equal(t, 1, measureCount)
}

func TestRemeasure_with_already_measured_value_smaller_but_still_float_equal(t *testing.T) {
	measureCount := 0

	root := NewNode()
	root.StyleSetWidth(288)
	root.StyleSetHeight(288)
	root.StyleSetFlexDirection(FlexDirectionRow)

	rootChild0 := NewNode()
	rootChild0.StyleSetPadding(EdgeAll, 2.88)
	rootChild0.StyleSetFlexDirection(FlexDirectionRow)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNode()
	rootChild0Child0.Context = measureCount
	rootChild0Child0.SetMeasureFunc(measure8449)
	rootChild0.InsertChild(rootChild0Child0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	measureCount = rootChild0Child0.Context.(int)
	assert.Equal(t, 1, measureCount)
}
