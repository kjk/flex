package flex

import "testing"

func TestRounding_flex_basis_flex_grow_row_width_of_100(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 33, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 33, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 34, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 67, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 33, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 67, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 33, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 33, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 34, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 33, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_flex_basis_flex_grow_row_prime_number_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 113)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild3, 1)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild4, 1)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 23, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 22, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 68, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 22, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 68, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 22, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 23, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 22, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild4))
}

func TestRounding_flex_basis_flex_shrink_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 101)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeStyleSetFlexBasis(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexBasis(rootChild1, 25)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexBasis(rootChild2, 25)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 101, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 51, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 51, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 76, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 101, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 51, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_flex_basis_overrides_main_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 113)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_total_fractial(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 87.4)
	NodeStyleSetHeight(root, 113.4)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 0.7)
	YGNodeStyleSetFlexBasis(rootChild0, 50.3)
	NodeStyleSetHeight(rootChild0, 20.3)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1.6)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1.1)
	NodeStyleSetHeight(rootChild2, 10.7)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 59, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 59, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 59, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 59, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_total_fractial_nested(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 87.4)
	NodeStyleSetHeight(root, 113.4)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 0.7)
	YGNodeStyleSetFlexBasis(rootChild0, 50.3)
	NodeStyleSetHeight(rootChild0, 20.3)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0Child0, 1)
	YGNodeStyleSetFlexBasis(rootChild0Child0, 0.3)
	NodeStyleSetPosition(rootChild0Child0, EdgeBottom, 13.3)
	NodeStyleSetHeight(rootChild0Child0, 9.9)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0child1, 4)
	YGNodeStyleSetFlexBasis(rootChild0child1, 0.3)
	NodeStyleSetPosition(rootChild0child1, EdgeTop, 13.3)
	NodeStyleSetHeight(rootChild0child1, 1.1)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1.6)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1.1)
	NodeStyleSetHeight(rootChild2, 10.7)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 59, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, -13, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 12, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 47, YGNodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 59, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 59, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, -13, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 12, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 47, YGNodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 59, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_fractial_input_1(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 113.4)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_fractial_input_2(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 113.6)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 114, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 65, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 65, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 114, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 65, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 65, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_fractial_input_3(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetPosition(root, EdgeTop, 0.3)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 113.4)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 114, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 65, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 114, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 65, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_fractial_input_4(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetPosition(root, EdgeTop, 0.7)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 113.4)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 1, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 1, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_inner_node_controversy_horizontal(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 320)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1child0, 1)
	NodeStyleSetHeight(rootChild1child0, 10)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 107, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 107, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 106, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 213, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 107, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 213, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 107, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 107, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 106, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 107, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_inner_node_controversy_vertical(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetHeight(root, 320)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetWidth(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetWidth(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1child0, 1)
	NodeStyleSetWidth(rootChild1child0, 10)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetWidth(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 107, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 213, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 107, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 213, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(rootChild2))
}

func TestRounding_inner_node_controversy_combined(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 640)
	NodeStyleSetHeight(root, 320)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetHeightPercent(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeightPercent(rootChild1, 100)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1child0, 1)
	NodeStyleSetWidthPercent(rootChild1child0, 100)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild1_child1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1_child1, 1)
	NodeStyleSetWidthPercent(rootChild1_child1, 100)
	YGNodeInsertChild(rootChild1, rootChild1_child1, 1)

	rootChild1_child1Child0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1_child1Child0, 1)
	NodeStyleSetWidthPercent(rootChild1_child1Child0, 100)
	YGNodeInsertChild(rootChild1_child1, rootChild1_child1Child0, 0)

	rootChild1_child2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1_child2, 1)
	NodeStyleSetWidthPercent(rootChild1_child2, 100)
	YGNodeInsertChild(rootChild1, rootChild1_child2, 2)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeightPercent(rootChild2, 100)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 640, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 213, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 213, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child1))
	assertFloatEqual(t, 107, YGNodeLayoutGetTop(rootChild1_child1))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child1Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1_child1Child0))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(rootChild1_child1Child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(rootChild1_child1Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child2))
	assertFloatEqual(t, 213, YGNodeLayoutGetTop(rootChild1_child2))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 427, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 213, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 640, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 427, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 213, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 213, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child1))
	assertFloatEqual(t, 107, YGNodeLayoutGetTop(rootChild1_child1))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(rootChild1_child1))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(rootChild1_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child1Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1_child1Child0))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(rootChild1_child1Child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(rootChild1_child1Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1_child2))
	assertFloatEqual(t, 213, YGNodeLayoutGetTop(rootChild1_child2))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(rootChild1_child2))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(rootChild1_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 213, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(rootChild2))
}
