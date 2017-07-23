package flex

import "testing"

func TestMax_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMaxWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

}

func TestMax_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetMaxHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

}

func TestMin_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMinHeight(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

}

func TestMin_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMinWidth(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

}

func TestJustify_content_min_max(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetMinHeight(root, 100)
	YGNodeStyleSetMaxHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

}

func TestAlign_items_min_max(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetMinWidth(root, 100)
	YGNodeStyleSetMaxWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 60)
	YGNodeStyleSetHeight(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

}

func TestJustify_content_overflow_min_max(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetMinHeight(root, 100)
	YGNodeStyleSetMaxHeight(root, 110)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeStyleSetHeight(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, -20, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, -20, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

}

func TestFlex_grow_to_min(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetMinHeight(root, 100)
	YGNodeStyleSetMaxHeight(root, 500)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexShrink(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

}

func TestFlex_grow_in_at_most_container(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0_child0, 0)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0_child0))

}

func TestFlex_grow_child(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0, 0)
	YGNodeStyleSetHeight(root_child0, 100)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

}

func TestFlex_grow_within_constrained_min_max_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMinHeight(root, 100)
	YGNodeStyleSetMaxHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

}

func TestFlex_grow_within_max_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetMaxWidth(root_child0, 100)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child0, 1)
	YGNodeStyleSetHeight(root_child0_child0, 20)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0))

}

func TestFlex_grow_within_constrained_max_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetMaxWidth(root_child0, 300)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child0, 1)
	YGNodeStyleSetHeight(root_child0_child0, 20)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child0_child0))

}

func TestFlex_root_ignored(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root, 1)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetMinHeight(root, 100)
	YGNodeStyleSetMaxHeight(root, 500)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0, 200)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child1, 100)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

}

func TestFlex_grow_root_minimized(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetMinHeight(root, 100)
	YGNodeStyleSetMaxHeight(root, 500)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMinHeight(root_child0, 100)
	YGNodeStyleSetMaxHeight(root_child0, 500)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0_child0, 200)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child0_child1, 100)
	YGNodeInsertChild(root_child0, root_child0_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child1))

}

func TestFlex_grow_height_maximized(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 500)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMinHeight(root_child0, 100)
	YGNodeStyleSetMaxHeight(root_child0, 500)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0_child0, 200)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child0_child1, 100)
	YGNodeInsertChild(root_child0, root_child0_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 400, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 400, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 400, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 400, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child1))

}

func TestFlex_grow_within_constrained_min_row(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetMinWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

}

func TestFlex_grow_within_constrained_min_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMinHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

}

func TestFlex_grow_within_constrained_max_row(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetMaxWidth(root_child0, 100)
	YGNodeStyleSetHeight(root_child0, 100)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexShrink(root_child0_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0_child0, 100)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child1, 50)
	YGNodeInsertChild(root_child0, root_child0_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0_child1))

}

func TestFlex_grow_within_constrained_max_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetMaxHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexShrink(root_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0, 100)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

}

func TestChild_min_max_width_flexing(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 120)
	YGNodeStyleSetHeight(root, 50)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0, 0)
	YGNodeStyleSetMinWidth(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetFlexBasisPercent(root_child1, 50)
	YGNodeStyleSetMaxWidth(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

}

func TestMin_width_overrides_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetMinWidth(root, 100)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

}

func TestMax_width_overrides_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetMaxWidth(root, 100)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

}

func TestMin_height_overrides_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root, 50)
	YGNodeStyleSetMinHeight(root, 100)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

}

func TestMax_height_overrides_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root, 200)
	YGNodeStyleSetMaxHeight(root, 100)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

}

func TestMin_max_percent_no_width_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMinWidthPercent(root_child0, 10)
	YGNodeStyleSetMaxWidthPercent(root_child0, 10)
	YGNodeStyleSetMinHeightPercent(root_child0, 10)
	YGNodeStyleSetMaxHeightPercent(root_child0, 10)
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
