package flex

import "testing"

func _measureFloor(node *YGNode, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize {
	return YGSize{
		width: 10.2, height: 10.2,
	}
}

func _measureCeil(node *YGNode, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize {
	return YGSize{
		width: 10.5, height: 10.5,
	}
}

func _measureFractial(node *YGNode, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize {
	return YGSize{
		width: 0.5, height: 0.5,
	}
}

func TestRounding_feature_with_custom_measure_func_floor(t *testing.T) {
	config := YGConfigNew()
	root := YGNodeNewWithConfig(config)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _measureFloor)
	YGNodeInsertChild(root, root_child0, 0)

	YGConfigSetPointScaleFactor(config, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 10.2, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10.2, YGNodeLayoutGetHeight(root_child0))

	YGConfigSetPointScaleFactor(config, 1)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 11, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 11, YGNodeLayoutGetHeight(root_child0))

	YGConfigSetPointScaleFactor(config, 2)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 10.5, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10.5, YGNodeLayoutGetHeight(root_child0))

	YGConfigSetPointScaleFactor(config, 4)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 10.25, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10.25, YGNodeLayoutGetHeight(root_child0))

	YGConfigSetPointScaleFactor(config, float32(1)/float32(3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 12.0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 12.0, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_feature_with_custom_measure_func_ceil(t *testing.T) {
	config := YGConfigNew()
	root := YGNodeNewWithConfig(config)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _measureCeil)
	YGNodeInsertChild(root, root_child0, 0)

	YGConfigSetPointScaleFactor(config, 1)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 11, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 11, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_feature_with_custom_measure_and_fractial_matching_scale(t *testing.T) {
	config := YGConfigNew()
	root := YGNodeNewWithConfig(config)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPosition(root_child0, YGEdgeLeft, 73.625)
	YGNodeSetMeasureFunc(root_child0, _measureFractial)
	YGNodeInsertChild(root, root_child0, 0)

	YGConfigSetPointScaleFactor(config, 2)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0.5, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0.5, YGNodeLayoutGetHeight(root_child0))
	assertFloatEqual(t, 73.5, YGNodeLayoutGetLeft(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}
