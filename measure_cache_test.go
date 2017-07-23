package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func measureMax(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	measureCount := YGNodeGetContext(node).(int)
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
	measureCount := YGNodeGetContext(node).(int)
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
	measureCount, ok := YGNodeGetContext(node).(int)
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
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	measureCount := 0
	rootChild0.Context = measureCount
	NodeSetMeasureFunc(rootChild0, measureMax)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_same_exact_width_larger_than_needed_height(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	measureCount := 0
	rootChild0.Context = measureCount
	NodeSetMeasureFunc(rootChild0, measureMin)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, 100, 100, DirectionLTR)
	NodeCalculateLayout(root, 100, 50, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_same_atmost_width_larger_than_needed_height(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	rootChild0 := NewNode()
	measureCount := 0
	rootChild0.Context = measureCount
	NodeSetMeasureFunc(rootChild0, measureMin)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, 100, 100, DirectionLTR)
	NodeCalculateLayout(root, 100, 50, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_computed_width_larger_than_needed_height(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	rootChild0 := NewNode()
	measureCount := 0
	rootChild0.Context = measureCount
	NodeSetMeasureFunc(rootChild0, measureMin)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, 100, 100, DirectionLTR)
	YGNodeStyleSetAlignItems(root, AlignStretch)
	NodeCalculateLayout(root, 10, 50, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_atmost_computed_width_undefined_height(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	rootChild0 := NewNode()
	measureCount := 0
	rootChild0.Context = measureCount
	NodeSetMeasureFunc(rootChild0, measureMin)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, 100, Undefined, DirectionLTR)
	NodeCalculateLayout(root, 10, Undefined, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0).(int)
	assert.Equal(t, 1, measureCount)
}

func TestRemeasure_with_already_measured_value_smaller_but_still_float_equal(t *testing.T) {
	measureCount := 0

	root := NewNode()
	NodeStyleSetWidth(root, 288)
	NodeStyleSetHeight(root, 288)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)

	rootChild0 := NewNode()
	YGNodeStyleSetPadding(rootChild0, EdgeAll, 2.88)
	YGNodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNode()
	rootChild0Child0.Context = measureCount
	NodeSetMeasureFunc(rootChild0Child0, measure8449)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0Child0).(int)
	assert.Equal(t, 1, measureCount)
}
