package flex

import "testing"

func TestAlign_self_center(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child0, AlignCenter)
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

}

func TestAlign_self_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child0, AlignFlexEnd)
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

}

func TestAlign_self_flex_start(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child0, AlignFlexStart)
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

}

func TestAlign_self_flex_end_override_flex_start(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child0, AlignFlexEnd)
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

}

func TestAlign_self_baseline(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child0, AlignBaseline)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignSelf(root_child1, AlignBaseline)
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

}
