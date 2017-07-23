package flex

import "testing"

func _baseline(node *Node, width float32, height float32) float32 {
	baseline := YGNodeGetContext(node).(float32)
	return baseline
}

func TestAlign_baseline_customer_func(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	var baselineValue float32 = 10
	rootChild1child0 := NewNode()
	YGNodeSetContext(rootChild1child0, baselineValue)
	YGNodeStyleSetWidth(rootChild1child0, 50)
	rootChild1child0.Baseline = _baseline
	YGNodeStyleSetHeight(rootChild1child0, 20)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

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
