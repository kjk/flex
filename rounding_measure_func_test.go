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
	rootChild0.SetMeasureFunc(_measureFloor)
	root.InsertChild(rootChild0, 0)

	config.SetPointScaleFactor(0)

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 10.2, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10.2, rootChild0.LayoutGetHeight())

	config.SetPointScaleFactor(1)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 11, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 11, rootChild0.LayoutGetHeight())

	config.SetPointScaleFactor(2)

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 10.5, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10.5, rootChild0.LayoutGetHeight())

	config.SetPointScaleFactor(4)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10.25, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10.25, rootChild0.LayoutGetHeight())

	config.SetPointScaleFactor(float32(1) / float32(3))

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 12.0, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 12.0, rootChild0.LayoutGetHeight())
}

func TestRounding_feature_with_custom_measure_func_ceil(t *testing.T) {
	config := NewConfig()
	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.SetMeasureFunc(_measureCeil)
	root.InsertChild(rootChild0, 0)

	config.SetPointScaleFactor(1)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 11, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 11, rootChild0.LayoutGetHeight())
}

func TestRounding_feature_with_custom_measure_and_fractial_matching_scale(t *testing.T) {
	config := NewConfig()
	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPosition(EdgeLeft, 73.625)
	rootChild0.SetMeasureFunc(_measureFractial)
	root.InsertChild(rootChild0, 0)

	config.SetPointScaleFactor(2)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0.5, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 0.5, rootChild0.LayoutGetHeight())
	assertFloatEqual(t, 73.5, rootChild0.LayoutGetLeft())
}
