package flex

import "testing"

func TestAlign_items_stretch(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

}

func TestAlign_items_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

}

func TestAlign_items_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

}

func TestAlign_items_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignFlexEnd)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

}

func TestAlign_baseline(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

}

func TestAlign_baseline_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 50)
	NodeStyleSetHeight(rootChild1child0, 10)
	NodeInsertChild(rootChild1, rootChild1child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_child_multiline(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild1, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild1, WrapWrap)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 25)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 25)
	NodeStyleSetHeight(rootChild1child0, 20)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild1_child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1_child1, 25)
	NodeStyleSetHeight(rootChild1_child1, 10)
	NodeInsertChild(rootChild1, rootChild1_child1, 1)

	rootChild1_child2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1_child2, 25)
	NodeStyleSetHeight(rootChild1_child2, 20)
	NodeInsertChild(rootChild1, rootChild1_child2, 2)

	rootChild1_child3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1_child3, 25)
	NodeStyleSetHeight(rootChild1_child3, 10)
	NodeInsertChild(rootChild1, rootChild1_child3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 0, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 25, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 0, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child3))

}

func TestAlign_baseline_child_multiline_override(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild1, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild1, WrapWrap)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 25)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 25)
	NodeStyleSetHeight(rootChild1child0, 20)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild1_child1 := NewNodeWithConfig(config)
	NodeStyleSetAlignSelf(rootChild1_child1, AlignBaseline)
	NodeStyleSetWidth(rootChild1_child1, 25)
	NodeStyleSetHeight(rootChild1_child1, 10)
	NodeInsertChild(rootChild1, rootChild1_child1, 1)

	rootChild1_child2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1_child2, 25)
	NodeStyleSetHeight(rootChild1_child2, 20)
	NodeInsertChild(rootChild1, rootChild1_child2, 2)

	rootChild1_child3 := NewNodeWithConfig(config)
	NodeStyleSetAlignSelf(rootChild1_child3, AlignBaseline)
	NodeStyleSetWidth(rootChild1_child3, 25)
	NodeStyleSetHeight(rootChild1_child3, 10)
	NodeInsertChild(rootChild1, rootChild1_child3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 0, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 25, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 0, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child3))

}

func TestAlign_baseline_child_multiline_no_override_on_secondline(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild1, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild1, WrapWrap)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 25)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 25)
	NodeStyleSetHeight(rootChild1child0, 20)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild1_child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1_child1, 25)
	NodeStyleSetHeight(rootChild1_child1, 10)
	NodeInsertChild(rootChild1, rootChild1_child1, 1)

	rootChild1_child2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1_child2, 25)
	NodeStyleSetHeight(rootChild1_child2, 20)
	NodeInsertChild(rootChild1, rootChild1_child2, 2)

	rootChild1_child3 := NewNodeWithConfig(config)
	NodeStyleSetAlignSelf(rootChild1_child3, AlignBaseline)
	NodeStyleSetWidth(rootChild1_child3, 25)
	NodeStyleSetHeight(rootChild1_child3, 10)
	NodeInsertChild(rootChild1, rootChild1_child3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 0, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 25, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 0, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1_child3))

}

func TestAlign_baseline_child_top(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetPosition(rootChild0, EdgeTop, 10)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 50)
	NodeStyleSetHeight(rootChild1child0, 10)
	NodeInsertChild(rootChild1, rootChild1child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_child_top2(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetPosition(rootChild1, EdgeTop, 5)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 50)
	NodeStyleSetHeight(rootChild1child0, 10)
	NodeInsertChild(rootChild1, rootChild1child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 45, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 45, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_double_nested_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 50)
	NodeStyleSetHeight(rootChild0Child0, 20)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 50)
	NodeStyleSetHeight(rootChild1child0, 15)
	NodeInsertChild(rootChild1, rootChild1child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 15, NodeLayoutGetHeight(rootChild1child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 15, NodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

}

func TestAlign_baseline_child_margin(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 5)
	NodeStyleSetMargin(rootChild0, EdgeTop, 5)
	NodeStyleSetMargin(rootChild0, EdgeRight, 5)
	NodeStyleSetMargin(rootChild0, EdgeBottom, 5)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild1child0, EdgeLeft, 1)
	NodeStyleSetMargin(rootChild1child0, EdgeTop, 1)
	NodeStyleSetMargin(rootChild1child0, EdgeRight, 1)
	NodeStyleSetMargin(rootChild1child0, EdgeBottom, 1)
	NodeStyleSetWidth(rootChild1child0, 50)
	NodeStyleSetHeight(rootChild1child0, 10)
	NodeInsertChild(rootChild1, rootChild1child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 60, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 44, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 1, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 1, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, -10, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 44, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, -1, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 1, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_child_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetPadding(root, EdgeLeft, 5)
	NodeStyleSetPadding(root, EdgeTop, 5)
	NodeStyleSetPadding(root, EdgeRight, 5)
	NodeStyleSetPadding(root, EdgeBottom, 5)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetPadding(rootChild1, EdgeLeft, 5)
	NodeStyleSetPadding(rootChild1, EdgeTop, 5)
	NodeStyleSetPadding(rootChild1, EdgeRight, 5)
	NodeStyleSetPadding(rootChild1, EdgeBottom, 5)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 50)
	NodeStyleSetHeight(rootChild1child0, 10)
	NodeInsertChild(rootChild1, rootChild1child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 55, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 5, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, -5, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, -5, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_multiline(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 50)
	NodeStyleSetHeight(rootChild1child0, 10)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 20)
	NodeInsertChild(root, rootChild2, 2)

	rootChild2_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2_child0, 50)
	NodeStyleSetHeight(rootChild2_child0, 10)
	NodeInsertChild(rootChild2, rootChild2_child0, 0)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

}

func TestAlign_baseline_multiline_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 20)
	NodeStyleSetHeight(rootChild1child0, 20)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 40)
	NodeStyleSetHeight(rootChild2, 70)
	NodeInsertChild(root, rootChild2, 2)

	rootChild2_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2_child0, 10)
	NodeStyleSetHeight(rootChild2_child0, 10)
	NodeInsertChild(rootChild2, rootChild2_child0, 0)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 20)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 70, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 70, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 70, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild3))

}

func TestAlign_baseline_multiline_column2(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 20)
	NodeStyleSetHeight(rootChild1child0, 20)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 40)
	NodeStyleSetHeight(rootChild2, 70)
	NodeInsertChild(root, rootChild2, 2)

	rootChild2_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2_child0, 10)
	NodeStyleSetHeight(rootChild2_child0, 10)
	NodeInsertChild(rootChild2, rootChild2_child0, 0)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 20)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 70, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 70, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 70, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild3))

}

func TestAlign_baseline_multiline_row_and_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignBaseline)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1child0, 50)
	NodeStyleSetHeight(rootChild1child0, 10)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 20)
	NodeInsertChild(root, rootChild2, 2)

	rootChild2_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2_child0, 50)
	NodeStyleSetHeight(rootChild2_child0, 10)
	NodeInsertChild(rootChild2, rootChild2_child0, 0)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 20)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 90, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 90, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild3))

}

func TestAlign_items_center_child_with_margin_bigger_than_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetWidth(root, 52)
	NodeStyleSetHeight(root, 52)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(rootChild0, AlignCenter)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild0Child0, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild0Child0, EdgeRight, 10)
	NodeStyleSetWidth(rootChild0Child0, 52)
	NodeStyleSetHeight(rootChild0Child0, 52)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(rootChild0Child0))

}

func TestAlign_items_flex_end_child_with_margin_bigger_than_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetWidth(root, 52)
	NodeStyleSetHeight(root, 52)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(rootChild0, AlignFlexEnd)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild0Child0, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild0Child0, EdgeRight, 10)
	NodeStyleSetWidth(rootChild0Child0, 52)
	NodeStyleSetHeight(rootChild0Child0, 52)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(rootChild0Child0))

}

func TestAlign_items_center_child_without_margin_bigger_than_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetWidth(root, 52)
	NodeStyleSetHeight(root, 52)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(rootChild0, AlignCenter)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 72)
	NodeStyleSetHeight(rootChild0Child0, 72)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 72, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 72, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 72, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 72, NodeLayoutGetHeight(rootChild0Child0))

}

func TestAlign_items_flex_end_child_without_margin_bigger_than_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetWidth(root, 52)
	NodeStyleSetHeight(root, 52)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(rootChild0, AlignFlexEnd)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 72)
	NodeStyleSetHeight(rootChild0Child0, 72)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 72, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 72, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, NodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 72, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 72, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 72, NodeLayoutGetHeight(rootChild0Child0))

}

func TestAlign_center_should_size_based_on_content(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(rootChild0, JustifyCenter)
	NodeStyleSetFlexShrink(rootChild0, 1)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetFlexShrink(rootChild0Child0, 1)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0_child0, 20)
	NodeStyleSetHeight(rootChild0Child0_child0, 20)
	NodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0Child0_child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0Child0_child0))

}

func TestAlign_strech_should_size_based_on_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(rootChild0, JustifyCenter)
	NodeStyleSetFlexShrink(rootChild0, 1)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetFlexShrink(rootChild0Child0, 1)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0_child0, 20)
	NodeStyleSetHeight(rootChild0Child0_child0, 20)
	NodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0Child0_child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 80, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetTop())
	assertFloatEqual(t, 20, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild0Child0_child0))

}
