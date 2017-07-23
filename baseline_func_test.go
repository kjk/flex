package flex

import "testing"

func _baseline(node *YGNode, width float32, height float32) float32 {
	baseline := YGNodeGetContext(node).(float32)
	return baseline
}

func TestAlign_baseline_customer_func(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNew()
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	var baselineValue float32 = 10
	root_child1_child0 := YGNodeNew()
	YGNodeSetContext(root_child1_child0, baselineValue)
	YGNodeStyleSetWidth(root_child1_child0, 50)
	YGNodeSetBaselineFunc(root_child1_child0, _baseline)
	YGNodeStyleSetHeight(root_child1_child0, 20)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))
}
