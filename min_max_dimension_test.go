package flex

import "testing"

func TestMax_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetMaxWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

}

func TestMax_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetMaxHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

}

func TestMin_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMinHeight(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

}

func TestMin_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMinWidth(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

}

func TestJustify_content_min_max(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetMinHeight(root, 100)
	NodeStyleSetMaxHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

}

func TestAlign_items_min_max(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetMinWidth(root, 100)
	NodeStyleSetMaxWidth(root, 200)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

}

func TestJustify_content_overflow_min_max(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetMinHeight(root, 100)
	NodeStyleSetMaxHeight(root, 110)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, -20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 110, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, -20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

}

func TestFlex_grow_to_min(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetMinHeight(root, 100)
	NodeStyleSetMaxHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

}

func TestFlex_grow_in_at_most_container(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetFlexBasis(rootChild0Child0, 0)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0Child0))

}

func TestFlex_grow_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 0)
	NodeStyleSetHeight(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestFlex_grow_within_constrained_min_max_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetMinHeight(root, 100)
	NodeStyleSetMaxHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

}

func TestFlex_grow_within_max_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetMaxWidth(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetHeight(rootChild0Child0, 20)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0))

}

func TestFlex_grow_within_constrained_max_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetMaxWidth(rootChild0, 300)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetHeight(rootChild0Child0, 20)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild0Child0))

}

func TestFlex_root_ignored(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(root, 1)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetMinHeight(root, 100)
	NodeStyleSetMaxHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 200)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 100)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

}

func TestFlex_grow_root_minimized(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetMinHeight(root, 100)
	NodeStyleSetMaxHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMinHeight(rootChild0, 100)
	NodeStyleSetMaxHeight(rootChild0, 500)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetFlexBasis(rootChild0Child0, 200)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0child1, 100)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 300, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0child1))

}

func TestFlex_grow_height_maximized(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMinHeight(rootChild0, 100)
	NodeStyleSetMaxHeight(rootChild0, 500)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetFlexBasis(rootChild0Child0, 200)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0child1, 100)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 400, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 400, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 400, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 400, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0child1))

}

func TestFlex_grow_within_constrained_min_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetMinWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

}

func TestFlex_grow_within_constrained_min_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetMinHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

}

func TestFlex_grow_within_constrained_max_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetMaxWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexShrink(rootChild0Child0, 1)
	NodeStyleSetFlexBasis(rootChild0Child0, 100)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1, 50)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0child1))

}

func TestFlex_grow_within_constrained_max_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetMaxHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexShrink(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

}

func TestChild_min_max_width_flexing(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 120)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 0)
	NodeStyleSetMinWidth(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetFlexBasisPercent(rootChild1, 50)
	NodeStyleSetMaxWidth(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

}

func TestMin_width_overrides_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetMinWidth(root, 100)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

}

func TestMax_width_overrides_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetMaxWidth(root, 100)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

}

func TestMin_height_overrides_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetHeight(root, 50)
	NodeStyleSetMinHeight(root, 100)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

}

func TestMax_height_overrides_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetHeight(root, 200)
	NodeStyleSetMaxHeight(root, 100)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

}

func TestMin_max_percent_no_width_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetMinWidthPercent(rootChild0, 10)
	NodeStyleSetMaxWidthPercent(rootChild0, 10)
	NodeStyleSetMinHeightPercent(rootChild0, 10)
	NodeStyleSetMaxHeightPercent(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

}
