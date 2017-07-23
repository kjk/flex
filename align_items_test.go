package flex

import "testing"

func TestAlign_items_stretch(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_items_center(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_items_flex_start(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, YGAlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_items_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, YGAlignFlexEnd)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)
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
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_child(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 50)
	YGNodeStyleSetHeight(root_child1_child0, 10)
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
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_child_multiline(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child1, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root_child1, YGWrapWrap)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 25)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 25)
	YGNodeStyleSetHeight(root_child1_child0, 20)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child1_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child1, 25)
	YGNodeStyleSetHeight(root_child1_child1, 10)
	YGNodeInsertChild(root_child1, root_child1_child1, 1)

	root_child1_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child2, 25)
	YGNodeStyleSetHeight(root_child1_child2, 20)
	YGNodeInsertChild(root_child1, root_child1_child2, 2)

	root_child1_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child3, 25)
	YGNodeStyleSetHeight(root_child1_child3, 10)
	YGNodeInsertChild(root_child1, root_child1_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child2))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child3))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_child_multiline_override(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child1, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root_child1, YGWrapWrap)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 25)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 25)
	YGNodeStyleSetHeight(root_child1_child0, 20)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child1_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child1_child1, YGAlignBaseline)
	YGNodeStyleSetWidth(root_child1_child1, 25)
	YGNodeStyleSetHeight(root_child1_child1, 10)
	YGNodeInsertChild(root_child1, root_child1_child1, 1)

	root_child1_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child2, 25)
	YGNodeStyleSetHeight(root_child1_child2, 20)
	YGNodeInsertChild(root_child1, root_child1_child2, 2)

	root_child1_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child1_child3, YGAlignBaseline)
	YGNodeStyleSetWidth(root_child1_child3, 25)
	YGNodeStyleSetHeight(root_child1_child3, 10)
	YGNodeInsertChild(root_child1, root_child1_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child2))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child3))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_child_multiline_no_override_on_secondline(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child1, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root_child1, YGWrapWrap)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 25)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 25)
	YGNodeStyleSetHeight(root_child1_child0, 20)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child1_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child1, 25)
	YGNodeStyleSetHeight(root_child1_child1, 10)
	YGNodeInsertChild(root_child1, root_child1_child1, 1)

	root_child1_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child2, 25)
	YGNodeStyleSetHeight(root_child1_child2, 20)
	YGNodeInsertChild(root_child1, root_child1_child2, 2)

	root_child1_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child1_child3, YGAlignBaseline)
	YGNodeStyleSetWidth(root_child1_child3, 25)
	YGNodeStyleSetHeight(root_child1_child3, 10)
	YGNodeInsertChild(root_child1, root_child1_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child2))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child1))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child1_child3))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child3))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_child_top(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPosition(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 50)
	YGNodeStyleSetHeight(root_child1_child0, 10)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_child_top2(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPosition(root_child1, YGEdgeTop, 5)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 50)
	YGNodeStyleSetHeight(root_child1_child0, 10)
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
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_double_nested_child(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0, 50)
	YGNodeStyleSetHeight(root_child0_child0, 20)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 50)
	YGNodeStyleSetHeight(root_child1_child0, 15)
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

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 15, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 15, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_child_margin(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root_child0, YGEdgeLeft, 5)
	YGNodeStyleSetMargin(root_child0, YGEdgeTop, 5)
	YGNodeStyleSetMargin(root_child0, YGEdgeRight, 5)
	YGNodeStyleSetMargin(root_child0, YGEdgeBottom, 5)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root_child1_child0, YGEdgeLeft, 1)
	YGNodeStyleSetMargin(root_child1_child0, YGEdgeTop, 1)
	YGNodeStyleSetMargin(root_child1_child0, YGEdgeRight, 1)
	YGNodeStyleSetMargin(root_child1_child0, YGEdgeBottom, 1)
	YGNodeStyleSetWidth(root_child1_child0, 50)
	YGNodeStyleSetHeight(root_child1_child0, 10)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 44, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 1, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 1, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 44, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, -1, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 1, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_child_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetPadding(root, YGEdgeLeft, 5)
	YGNodeStyleSetPadding(root, YGEdgeTop, 5)
	YGNodeStyleSetPadding(root, YGEdgeRight, 5)
	YGNodeStyleSetPadding(root, YGEdgeBottom, 5)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPadding(root_child1, YGEdgeLeft, 5)
	YGNodeStyleSetPadding(root_child1, YGEdgeTop, 5)
	YGNodeStyleSetPadding(root_child1, YGEdgeRight, 5)
	YGNodeStyleSetPadding(root_child1, YGEdgeBottom, 5)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 50)
	YGNodeStyleSetHeight(root_child1_child0, 10)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 55, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, -5, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, -5, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_multiline(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 50)
	YGNodeStyleSetHeight(root_child1_child0, 10)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeStyleSetHeight(root_child2, 20)
	YGNodeInsertChild(root, root_child2, 2)

	root_child2_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2_child0, 50)
	YGNodeStyleSetHeight(root_child2_child0, 10)
	YGNodeInsertChild(root_child2, root_child2_child0, 0)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeStyleSetHeight(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)
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
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_multiline_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 20)
	YGNodeStyleSetHeight(root_child1_child0, 20)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 40)
	YGNodeStyleSetHeight(root_child2, 70)
	YGNodeInsertChild(root, root_child2, 2)

	root_child2_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2_child0, 10)
	YGNodeStyleSetHeight(root_child2_child0, 10)
	YGNodeInsertChild(root_child2, root_child2_child0, 0)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeStyleSetHeight(root_child3, 20)
	YGNodeInsertChild(root, root_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 70, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 70, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child3))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_multiline_column2(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 30)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 20)
	YGNodeStyleSetHeight(root_child1_child0, 20)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 40)
	YGNodeStyleSetHeight(root_child2, 70)
	YGNodeInsertChild(root, root_child2, 2)

	root_child2_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2_child0, 10)
	YGNodeStyleSetHeight(root_child2_child0, 10)
	YGNodeInsertChild(root_child2, root_child2_child0, 0)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeStyleSetHeight(root_child3, 20)
	YGNodeInsertChild(root, root_child3, 3)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 70, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 70, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child3))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_baseline_multiline_row_and_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, YGAlignBaseline)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1_child0, 50)
	YGNodeStyleSetHeight(root_child1_child0, 10)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeStyleSetHeight(root_child2, 20)
	YGNodeInsertChild(root, root_child2, 2)

	root_child2_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2_child0, 50)
	YGNodeStyleSetHeight(root_child2_child0, 10)
	YGNodeInsertChild(root_child2, root_child2_child0, 0)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeStyleSetHeight(root_child3, 20)
	YGNodeInsertChild(root, root_child3, 3)
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
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child3))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child3))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_items_center_child_with_margin_bigger_than_parent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetWidth(root, 52)
	YGNodeStyleSetHeight(root, 52)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root_child0, YGAlignCenter)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root_child0_child0, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(root_child0_child0, YGEdgeRight, 10)
	YGNodeStyleSetWidth(root_child0_child0, 52)
	YGNodeStyleSetHeight(root_child0_child0, 52)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_items_flex_end_child_with_margin_bigger_than_parent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetWidth(root, 52)
	YGNodeStyleSetHeight(root, 52)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root_child0, YGAlignFlexEnd)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root_child0_child0, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(root_child0_child0, YGEdgeRight, 10)
	YGNodeStyleSetWidth(root_child0_child0, 52)
	YGNodeStyleSetHeight(root_child0_child0, 52)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_items_center_child_without_margin_bigger_than_parent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetWidth(root, 52)
	YGNodeStyleSetHeight(root, 52)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root_child0, YGAlignCenter)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0, 72)
	YGNodeStyleSetHeight(root_child0_child0, 72)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, -10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, -10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_items_flex_end_child_without_margin_bigger_than_parent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetWidth(root, 52)
	YGNodeStyleSetHeight(root, 52)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root_child0, YGAlignFlexEnd)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0, 72)
	YGNodeStyleSetHeight(root_child0_child0, 72)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, -10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 52, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -10, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, -10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 72, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_center_should_size_based_on_content(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, YGAlignCenter)
	YGNodeStyleSetMargin(root, YGEdgeTop, 20)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root_child0, YGJustifyCenter)
	YGNodeStyleSetFlexShrink(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child0, 1)
	YGNodeStyleSetFlexShrink(root_child0_child0, 1)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0_child0, 20)
	YGNodeStyleSetHeight(root_child0_child0_child0, 20)
	YGNodeInsertChild(root_child0_child0, root_child0_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestAlign_strech_should_size_based_on_parent(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, YGEdgeTop, 20)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root_child0, YGJustifyCenter)
	YGNodeStyleSetFlexShrink(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child0, 1)
	YGNodeStyleSetFlexShrink(root_child0_child0, 1)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0_child0, 20)
	YGNodeStyleSetHeight(root_child0_child0_child0, 20)
	YGNodeInsertChild(root_child0_child0, root_child0_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0_child0))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}
