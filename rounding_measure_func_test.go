package flex

import "testing"

func _measureFloor(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	return Size{
		Width: 10.2, Height: 10.2,
	}
}

func _measureCeil(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	return Size{
		Width: 10.5, Height: 10.5,
	}
}

func _measureFractial(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	return Size{
		Width: 0.5, Height: 0.5,
	}
}

func TestRounding_feature_with_custom_measure_func_floor(t *testing.T) {
	config := YGConfigNew()
	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _measureFloor)
	YGNodeInsertChild(root, rootChild0, 0)

	YGConfigSetPointScaleFactor(config, 0)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 10.2, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10.2, YGNodeLayoutGetHeight(rootChild0))

	YGConfigSetPointScaleFactor(config, 1)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 11, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 11, YGNodeLayoutGetHeight(rootChild0))

	YGConfigSetPointScaleFactor(config, 2)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 10.5, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10.5, YGNodeLayoutGetHeight(rootChild0))

	YGConfigSetPointScaleFactor(config, 4)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10.25, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10.25, YGNodeLayoutGetHeight(rootChild0))

	YGConfigSetPointScaleFactor(config, float32(1)/float32(3))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 12.0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 12.0, YGNodeLayoutGetHeight(rootChild0))
}

func TestRounding_feature_with_custom_measure_func_ceil(t *testing.T) {
	config := YGConfigNew()
	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _measureCeil)
	YGNodeInsertChild(root, rootChild0, 0)

	YGConfigSetPointScaleFactor(config, 1)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 11, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 11, YGNodeLayoutGetHeight(rootChild0))
}

func TestRounding_feature_with_custom_measure_and_fractial_matching_scale(t *testing.T) {
	config := YGConfigNew()
	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetPosition(rootChild0, EdgeLeft, 73.625)
	YGNodeSetMeasureFunc(rootChild0, _measureFractial)
	YGNodeInsertChild(root, rootChild0, 0)

	YGConfigSetPointScaleFactor(config, 2)

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0.5, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0.5, YGNodeLayoutGetHeight(rootChild0))
	assertFloatEqual(t, 73.5, YGNodeLayoutGetLeft(rootChild0))
}
