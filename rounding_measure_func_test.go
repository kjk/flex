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
	config := NewConfig()
	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _measureFloor)
	YGNodeInsertChild(root, rootChild0, 0)

	ConfigSetPointScaleFactor(config, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 10.2, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10.2, NodeLayoutGetHeight(rootChild0))

	ConfigSetPointScaleFactor(config, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 11, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 11, NodeLayoutGetHeight(rootChild0))

	ConfigSetPointScaleFactor(config, 2)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 10.5, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10.5, NodeLayoutGetHeight(rootChild0))

	ConfigSetPointScaleFactor(config, 4)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10.25, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10.25, NodeLayoutGetHeight(rootChild0))

	ConfigSetPointScaleFactor(config, float32(1)/float32(3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 12.0, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 12.0, NodeLayoutGetHeight(rootChild0))
}

func TestRounding_feature_with_custom_measure_func_ceil(t *testing.T) {
	config := NewConfig()
	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _measureCeil)
	YGNodeInsertChild(root, rootChild0, 0)

	ConfigSetPointScaleFactor(config, 1)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 11, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 11, NodeLayoutGetHeight(rootChild0))
}

func TestRounding_feature_with_custom_measure_and_fractial_matching_scale(t *testing.T) {
	config := NewConfig()
	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetPosition(rootChild0, EdgeLeft, 73.625)
	NodeSetMeasureFunc(rootChild0, _measureFractial)
	YGNodeInsertChild(root, rootChild0, 0)

	ConfigSetPointScaleFactor(config, 2)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0.5, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0.5, NodeLayoutGetHeight(rootChild0))
	assertFloatEqual(t, 73.5, NodeLayoutGetLeft(rootChild0))
}
