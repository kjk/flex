package flex

import "testing"

func TestAlign_items_stretch(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

}

func TestAlign_items_center(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

}

func TestAlign_items_flex_start(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

}

func TestAlign_items_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignFlexEnd)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

}

func TestAlign_baseline(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)
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
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

}

func TestAlign_baseline_child(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 50)
	YGNodeStyleSetHeight(rootChild1child0, 10)
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
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_child_multiline(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild1, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(rootChild1, YGWrapWrap)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 25)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 25)
	YGNodeStyleSetHeight(rootChild1child0, 20)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild1_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1_child1, 25)
	YGNodeStyleSetHeight(rootChild1_child1, 10)
	YGNodeInsertChild(rootChild1, rootChild1_child1, 1)

	rootChild1_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1_child2, 25)
	YGNodeStyleSetHeight(rootChild1_child2, 20)
	YGNodeInsertChild(rootChild1, rootChild1_child2, 2)

	rootChild1_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1_child3, 25)
	YGNodeStyleSetHeight(rootChild1_child3, 10)
	YGNodeInsertChild(rootChild1, rootChild1_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child3))

}

func TestAlign_baseline_child_multiline_override(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild1, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(rootChild1, YGWrapWrap)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 25)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 25)
	YGNodeStyleSetHeight(rootChild1child0, 20)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild1_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(rootChild1_child1, AlignBaseline)
	YGNodeStyleSetWidth(rootChild1_child1, 25)
	YGNodeStyleSetHeight(rootChild1_child1, 10)
	YGNodeInsertChild(rootChild1, rootChild1_child1, 1)

	rootChild1_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1_child2, 25)
	YGNodeStyleSetHeight(rootChild1_child2, 20)
	YGNodeInsertChild(rootChild1, rootChild1_child2, 2)

	rootChild1_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(rootChild1_child3, AlignBaseline)
	YGNodeStyleSetWidth(rootChild1_child3, 25)
	YGNodeStyleSetHeight(rootChild1_child3, 10)
	YGNodeInsertChild(rootChild1, rootChild1_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child3))

}

func TestAlign_baseline_child_multiline_no_override_on_secondline(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild1, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(rootChild1, YGWrapWrap)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 25)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 25)
	YGNodeStyleSetHeight(rootChild1child0, 20)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild1_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1_child1, 25)
	YGNodeStyleSetHeight(rootChild1_child1, 10)
	YGNodeInsertChild(rootChild1, rootChild1_child1, 1)

	rootChild1_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1_child2, 25)
	YGNodeStyleSetHeight(rootChild1_child2, 20)
	YGNodeInsertChild(rootChild1, rootChild1_child2, 2)

	rootChild1_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(rootChild1_child3, AlignBaseline)
	YGNodeStyleSetWidth(rootChild1_child3, 25)
	YGNodeStyleSetHeight(rootChild1_child3, 10)
	YGNodeInsertChild(rootChild1, rootChild1_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1_child3))

}

func TestAlign_baseline_child_top(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPosition(rootChild0, EdgeTop, 10)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 50)
	YGNodeStyleSetHeight(rootChild1child0, 10)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_child_top2(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPosition(rootChild1, EdgeTop, 5)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 50)
	YGNodeStyleSetHeight(rootChild1child0, 10)
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
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_double_nested_child(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0, 50)
	YGNodeStyleSetHeight(rootChild0Child0, 20)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 50)
	YGNodeStyleSetHeight(rootChild1child0, 15)
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

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 15, YGNodeLayoutGetHeight(rootChild1child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 15, YGNodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

}

func TestAlign_baseline_child_margin(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(rootChild0, EdgeLeft, 5)
	YGNodeStyleSetMargin(rootChild0, EdgeTop, 5)
	YGNodeStyleSetMargin(rootChild0, EdgeRight, 5)
	YGNodeStyleSetMargin(rootChild0, EdgeBottom, 5)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(rootChild1child0, EdgeLeft, 1)
	YGNodeStyleSetMargin(rootChild1child0, EdgeTop, 1)
	YGNodeStyleSetMargin(rootChild1child0, EdgeRight, 1)
	YGNodeStyleSetMargin(rootChild1child0, EdgeBottom, 1)
	YGNodeStyleSetWidth(rootChild1child0, 50)
	YGNodeStyleSetHeight(rootChild1child0, 10)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 44, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 1, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 1, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 44, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, -1, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 1, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_child_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetPadding(root, EdgeLeft, 5)
	YGNodeStyleSetPadding(root, EdgeTop, 5)
	YGNodeStyleSetPadding(root, EdgeRight, 5)
	YGNodeStyleSetPadding(root, EdgeBottom, 5)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPadding(rootChild1, EdgeLeft, 5)
	YGNodeStyleSetPadding(rootChild1, EdgeTop, 5)
	YGNodeStyleSetPadding(rootChild1, EdgeRight, 5)
	YGNodeStyleSetPadding(rootChild1, EdgeBottom, 5)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 50)
	YGNodeStyleSetHeight(rootChild1child0, 10)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 55, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, -5, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, -5, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

}

func TestAlign_baseline_multiline(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 50)
	YGNodeStyleSetHeight(rootChild1child0, 10)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 50)
	YGNodeStyleSetHeight(rootChild2, 20)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild2_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild2_child0, 50)
	YGNodeStyleSetHeight(rootChild2_child0, 10)
	YGNodeInsertChild(rootChild2, rootChild2_child0, 0)

	rootChild3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 50)
	YGNodeStyleSetHeight(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)
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
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

}

func TestAlign_baseline_multiline_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 20)
	YGNodeStyleSetHeight(rootChild1child0, 20)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 40)
	YGNodeStyleSetHeight(rootChild2, 70)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild2_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild2_child0, 10)
	YGNodeStyleSetHeight(rootChild2_child0, 10)
	YGNodeInsertChild(rootChild2, rootChild2_child0, 0)

	rootChild3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 50)
	YGNodeStyleSetHeight(rootChild3, 20)
	YGNodeInsertChild(root, rootChild3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 70, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 70, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild3))

}

func TestAlign_baseline_multiline_column2(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 30)
	YGNodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 20)
	YGNodeStyleSetHeight(rootChild1child0, 20)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 40)
	YGNodeStyleSetHeight(rootChild2, 70)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild2_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild2_child0, 10)
	YGNodeStyleSetHeight(rootChild2_child0, 10)
	YGNodeInsertChild(rootChild2, rootChild2_child0, 0)

	rootChild3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 50)
	YGNodeStyleSetHeight(rootChild3, 20)
	YGNodeInsertChild(root, rootChild3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 70, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 70, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild3))

}

func TestAlign_baseline_multiline_row_and_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignBaseline)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 50)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1child0, 50)
	YGNodeStyleSetHeight(rootChild1child0, 10)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 50)
	YGNodeStyleSetHeight(rootChild2, 20)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild2_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild2_child0, 50)
	YGNodeStyleSetHeight(rootChild2_child0, 10)
	YGNodeInsertChild(rootChild2, rootChild2_child0, 0)

	rootChild3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild3, 50)
	YGNodeStyleSetHeight(rootChild3, 20)
	YGNodeInsertChild(root, rootChild3, 3)
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
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild3))

}

func TestAlign_items_center_child_with_margin_bigger_than_parent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetWidth(root, 52)
	YGNodeStyleSetHeight(root, 52)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(rootChild0, AlignCenter)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(rootChild0Child0, EdgeLeft, 10)
	YGNodeStyleSetMargin(rootChild0Child0, EdgeRight, 10)
	YGNodeStyleSetWidth(rootChild0Child0, 52)
	YGNodeStyleSetHeight(rootChild0Child0, 52)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(rootChild0Child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(rootChild0Child0))

}

func TestAlign_items_flex_end_child_with_margin_bigger_than_parent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetWidth(root, 52)
	YGNodeStyleSetHeight(root, 52)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(rootChild0, AlignFlexEnd)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(rootChild0Child0, EdgeLeft, 10)
	YGNodeStyleSetMargin(rootChild0Child0, EdgeRight, 10)
	YGNodeStyleSetWidth(rootChild0Child0, 52)
	YGNodeStyleSetHeight(rootChild0Child0, 52)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(rootChild0Child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(rootChild0Child0))

}

func TestAlign_items_center_child_without_margin_bigger_than_parent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetWidth(root, 52)
	YGNodeStyleSetHeight(root, 52)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(rootChild0, AlignCenter)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0, 72)
	YGNodeStyleSetHeight(rootChild0Child0, 72)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, -10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(rootChild0Child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, -10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(rootChild0Child0))

}

func TestAlign_items_flex_end_child_without_margin_bigger_than_parent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetWidth(root, 52)
	YGNodeStyleSetHeight(root, 52)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(rootChild0, AlignFlexEnd)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0, 72)
	YGNodeStyleSetHeight(rootChild0Child0, 72)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, -10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(rootChild0Child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, -10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(rootChild0Child0))

}

func TestAlign_center_should_size_based_on_content(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(rootChild0, YGJustifyCenter)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0Child0, 1)
	YGNodeStyleSetFlexShrink(rootChild0Child0, 1)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0_child0, 20)
	YGNodeStyleSetHeight(rootChild0Child0_child0, 20)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0_child0))

}

func TestAlign_strech_should_size_based_on_parent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(rootChild0, YGJustifyCenter)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0Child0, 1)
	YGNodeStyleSetFlexShrink(rootChild0Child0, 1)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0_child0, 20)
	YGNodeStyleSetHeight(rootChild0Child0_child0, 20)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0_child0))

}
