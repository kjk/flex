package flex

import "testing"

func baselineFunc(node *Node, width float32, height float32) float32 {
	baseline := node.Context.(float32)
	return baseline
}

func TestAlign_baseline_customer_func(t *testing.T) {
	root := NewNode()
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	var baselineValue float32 = 10
	rootChild1child0 := NewNode()
	rootChild1child0.Context = baselineValue
	NodeStyleSetWidth(rootChild1child0, 50)
	rootChild1child0.Baseline = baselineFunc
	NodeStyleSetHeight(rootChild1child0, 20)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))
}
