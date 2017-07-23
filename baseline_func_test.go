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

	rootChild0 := YGNodeNew()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNew()
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	var baselineValue float32 = 10
	rootChild1child0 := YGNodeNew()
	YGNodeSetContext(rootChild1child0, baselineValue)
	YGNodeStyleSetWidth(rootChild1child0, 50)
	YGNodeSetBaselineFunc(rootChild1child0, _baseline)
	YGNodeStyleSetHeight(rootChild1child0, 20)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))
}
