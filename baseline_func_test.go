package flex

import "testing"

func baselineFunc(node *Node, width float32, height float32) float32 {
	baseline := node.Context.(float32)
	return baseline
}

func TestAlign_baseline_customer_func(t *testing.T) {
	root := NewNode()
	root.StyleSetFlexDirection(FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNode()
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	NodeInsertChild(root, rootChild1, 1)

	var baselineValue float32 = 10
	rootChild1child0 := NewNode()
	rootChild1child0.Context = baselineValue
	rootChild1child0.StyleSetWidth(50)
	rootChild1child0.Baseline = baselineFunc
	rootChild1child0.StyleSetHeight(20)
	NodeInsertChild(rootChild1, rootChild1child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())
}
