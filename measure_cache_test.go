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

	root_child0 := YGNodeNew()
	measureCount := 0
	YGNodeSetContext(root_child0, measureCount)
	YGNodeSetMeasureFunc(root_child0, _measureMax)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	measureCount = YGNodeGetContext(root_child0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_same_exact_width_larger_than_needed_height(t *testing.T) {
	root := YGNodeNew()

	root_child0 := YGNodeNew()
	measureCount := 0
	YGNodeSetContext(root_child0, measureCount)
	YGNodeSetMeasureFunc(root_child0, _measureMin)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, 100, 100, YGDirectionLTR)
	YGNodeCalculateLayout(root, 100, 50, YGDirectionLTR)

	measureCount = YGNodeGetContext(root_child0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_same_atmost_width_larger_than_needed_height(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	root_child0 := YGNodeNew()
	measureCount := 0
	YGNodeSetContext(root_child0, measureCount)
	YGNodeSetMeasureFunc(root_child0, _measureMin)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, 100, 100, YGDirectionLTR)
	YGNodeCalculateLayout(root, 100, 50, YGDirectionLTR)

	measureCount = YGNodeGetContext(root_child0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_computed_width_larger_than_needed_height(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	root_child0 := YGNodeNew()
	measureCount := 0
	YGNodeSetContext(root_child0, measureCount)
	YGNodeSetMeasureFunc(root_child0, _measureMin)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, 100, 100, YGDirectionLTR)
	YGNodeStyleSetAlignItems(root, AlignStretch)
	YGNodeCalculateLayout(root, 10, 50, YGDirectionLTR)

	measureCount = YGNodeGetContext(root_child0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_atmost_computed_width_undefined_height(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	root_child0 := YGNodeNew()
	measureCount := 0
	YGNodeSetContext(root_child0, measureCount)
	YGNodeSetMeasureFunc(root_child0, _measureMin)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, 100, YGUndefined, YGDirectionLTR)
	YGNodeCalculateLayout(root, 10, YGUndefined, YGDirectionLTR)

	measureCount = YGNodeGetContext(root_child0).(int)
	assert.Equal(t, 1, measureCount)

}

func TestRemeasure_with_already_measured_value_smaller_but_still_float_equal(t *testing.T) {
	measureCount := 0

	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 288)
	YGNodeStyleSetHeight(root, 288)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)

	root_child0 := YGNodeNew()
	YGNodeStyleSetPadding(root_child0, YGEdgeAll, 2.88)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNew()
	YGNodeSetContext(root_child0_child0, measureCount)
	YGNodeSetMeasureFunc(root_child0_child0, _measure_84_49)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	measureCount = YGNodeGetContext(root_child0_child0).(int)
	assert.Equal(t, 1, measureCount)

}
