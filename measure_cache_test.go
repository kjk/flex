package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func _measureMax(node *YGNode, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize {
	measureCount := YGNodeGetContext(node).(int)
	measureCount++
	YGNodeSetContext(node, measureCount)

	if widthMode == YGMeasureModeUndefined {
		width = 10
	}
	if heightMode == YGMeasureModeUndefined {
		height = 10
	}
	return YGSize{
		width:  width,
		height: height,
	}
}

func _measureMin(node *YGNode, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize {
	measureCount := YGNodeGetContext(node).(int)
	measureCount++
	YGNodeSetContext(node, measureCount)

	if widthMode == YGMeasureModeUndefined || (widthMode == YGMeasureModeAtMost && width > 10) {
		width = 10
	}
	if heightMode == YGMeasureModeUndefined || (heightMode == YGMeasureModeAtMost && height > 10) {
		height = 10
	}
	return YGSize{
		width:  width,
		height: height,
	}
}

func _measure_84_49(node *YGNode, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize {
	measureCount, ok := YGNodeGetContext(node).(int)
	if ok {
		measureCount++
		YGNodeSetContext(node, measureCount)
	}

	return YGSize{
		width:  84,
		height: 49,
	}
}

func TestMeasure_once_single_flexible_child(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	measureCount := 0
	YGNodeSetContext(rootChild0, measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measureMax)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_same_exact_width_larger_than_needed_height(t *testing.T) {
	root := YGNodeNew()

	rootChild0 := YGNodeNew()
	measureCount := 0
	YGNodeSetContext(rootChild0, measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measureMin)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, 100, 100, DirectionLTR)
	YGNodeCalculateLayout(root, 100, 50, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_same_atmost_width_larger_than_needed_height(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	rootChild0 := YGNodeNew()
	measureCount := 0
	YGNodeSetContext(rootChild0, measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measureMin)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, 100, 100, DirectionLTR)
	YGNodeCalculateLayout(root, 100, 50, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_computed_width_larger_than_needed_height(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	rootChild0 := YGNodeNew()
	measureCount := 0
	YGNodeSetContext(rootChild0, measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measureMin)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, 100, 100, DirectionLTR)
	YGNodeStyleSetAlignItems(root, AlignStretch)
	YGNodeCalculateLayout(root, 10, 50, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_atmost_computed_width_undefined_height(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	rootChild0 := YGNodeNew()
	measureCount := 0
	YGNodeSetContext(rootChild0, measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measureMin)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, 100, YGUndefined, DirectionLTR)
	YGNodeCalculateLayout(root, 10, YGUndefined, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_already_measured_value_smaller_but_still_float_equal(t *testing.T) {
	measureCount := 0

	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 288)
	YGNodeStyleSetHeight(root, 288)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetPadding(rootChild0, YGEdgeAll, 2.88)
	YGNodeStyleSetFlexDirection(rootChild0, YGFlexDirectionRow)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNew()
	YGNodeSetContext(rootChild0Child0, measureCount)
	YGNodeSetMeasureFunc(rootChild0Child0, _measure_84_49)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	measureCount = YGNodeGetContext(rootChild0Child0).(int)
	assert.Equal(t, 1, measureCount)

}
