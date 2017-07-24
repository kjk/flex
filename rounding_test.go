package flex

import "testing"

func TestRounding_flex_basis_flex_grow_row_width_of_100(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 33, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 33, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 34, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 67, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 33, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 67, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 33, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 33, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 34, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 33, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())
}

func TestRounding_flex_basis_flex_grow_row_prime_number_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 113)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild3, 1)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild4, 1)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 113, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 23, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 23, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 22, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 23, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 68, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 22, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 23, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild4.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 113, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 23, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 68, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 22, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 23, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 23, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 22, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild3.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild4.LayoutGetTop())
	assertFloatEqual(t, 23, rootChild4.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild4.LayoutGetHeight())
}

func TestRounding_flex_basis_flex_shrink_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 101)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexShrink(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 100)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexBasis(rootChild1, 25)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexBasis(rootChild2, 25)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 101, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 51, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 51, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 76, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 101, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 51, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())
}

func TestRounding_flex_basis_overrides_main_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 113)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 113, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 64, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 64, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 113, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 64, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 64, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild2.LayoutGetHeight())
}

func TestRounding_total_fractial(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 87.4)
	NodeStyleSetHeight(root, 113.4)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 0.7)
	NodeStyleSetFlexBasis(rootChild0, 50.3)
	NodeStyleSetHeight(rootChild0, 20.3)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1.6)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1.1)
	NodeStyleSetHeight(rootChild2, 10.7)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 87, root.LayoutGetWidth())
	assertFloatEqual(t, 113, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 59, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 59, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 87, root.LayoutGetWidth())
	assertFloatEqual(t, 113, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 59, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 59, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild2.LayoutGetHeight())
}

func TestRounding_total_fractial_nested(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 87.4)
	NodeStyleSetHeight(root, 113.4)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 0.7)
	NodeStyleSetFlexBasis(rootChild0, 50.3)
	NodeStyleSetHeight(rootChild0, 20.3)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetFlexBasis(rootChild0Child0, 0.3)
	NodeStyleSetPosition(rootChild0Child0, EdgeBottom, 13.3)
	NodeStyleSetHeight(rootChild0Child0, 9.9)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0child1, 4)
	NodeStyleSetFlexBasis(rootChild0child1, 0.3)
	NodeStyleSetPosition(rootChild0child1, EdgeTop, 13.3)
	NodeStyleSetHeight(rootChild0child1, 1.1)
	NodeInsertChild(rootChild0, rootChild0child1, 1)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1.6)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1.1)
	NodeStyleSetHeight(rootChild2, 10.7)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 87, root.LayoutGetWidth())
	assertFloatEqual(t, 113, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 59, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, -13, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 12, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 25, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 47, rootChild0child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 59, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 87, root.LayoutGetWidth())
	assertFloatEqual(t, 113, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 59, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, -13, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 12, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 25, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 47, rootChild0child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 59, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 30, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 87, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild2.LayoutGetHeight())
}

func TestRounding_fractial_input_1(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 113.4)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 113, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 64, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 64, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 113, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 64, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 64, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild2.LayoutGetHeight())
}

func TestRounding_fractial_input_2(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 113.6)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 114, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 65, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 65, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 114, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 65, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 65, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild2.LayoutGetHeight())
}

func TestRounding_fractial_input_3(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetPosition(root, EdgeTop, 0.3)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 113.4)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 114, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 65, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 64, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 114, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 65, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 64, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild2.LayoutGetHeight())
}

func TestRounding_fractial_input_4(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetPosition(root, EdgeTop, 0.7)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 113.4)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 20)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 1, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 113, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 64, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 64, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 1, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 113, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 64, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 64, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 89, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 24, rootChild2.LayoutGetHeight())
}

func TestRounding_inner_node_controversy_horizontal(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 320)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1child0, 1)
	NodeStyleSetHeight(rootChild1child0, 10)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 320, root.LayoutGetWidth())
	assertFloatEqual(t, 10, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 107, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 107, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 106, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 106, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 213, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 107, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 320, root.LayoutGetWidth())
	assertFloatEqual(t, 10, root.LayoutGetHeight())

	assertFloatEqual(t, 213, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 107, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 107, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 106, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 106, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 107, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())
}

func TestRounding_inner_node_controversy_vertical(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetHeight(root, 320)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetWidth(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetWidth(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1child0, 1)
	NodeStyleSetWidth(rootChild1child0, 10)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetWidth(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 10, root.LayoutGetWidth())
	assertFloatEqual(t, 320, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 107, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 107, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 106, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 106, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 213, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 107, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 10, root.LayoutGetWidth())
	assertFloatEqual(t, 320, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 107, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 107, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 106, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 106, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 213, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 107, rootChild2.LayoutGetHeight())
}

func TestRounding_inner_node_controversy_combined(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 640)
	NodeStyleSetHeight(root, 320)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetHeightPercent(rootChild0, 100)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetHeightPercent(rootChild1, 100)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1child0, 1)
	NodeStyleSetWidthPercent(rootChild1child0, 100)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild1_child1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1_child1, 1)
	NodeStyleSetWidthPercent(rootChild1_child1, 100)
	NodeInsertChild(rootChild1, rootChild1_child1, 1)

	rootChild1_child1Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1_child1Child0, 1)
	NodeStyleSetWidthPercent(rootChild1_child1Child0, 100)
	NodeInsertChild(rootChild1_child1, rootChild1_child1Child0, 0)

	rootChild1_child2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1_child2, 1)
	NodeStyleSetWidthPercent(rootChild1_child2, 100)
	NodeInsertChild(rootChild1, rootChild1_child2, 2)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 1)
	NodeStyleSetHeightPercent(rootChild2, 100)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 640, root.LayoutGetWidth())
	assertFloatEqual(t, 320, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 213, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 320, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 213, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 214, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 320, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 214, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 107, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 107, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 214, rootChild1_child1.LayoutGetWidth())
	assertFloatEqual(t, 106, rootChild1_child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child1Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1Child0.LayoutGetTop())
	assertFloatEqual(t, 214, rootChild1_child1Child0.LayoutGetWidth())
	assertFloatEqual(t, 106, rootChild1_child1Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 213, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 214, rootChild1_child2.LayoutGetWidth())
	assertFloatEqual(t, 107, rootChild1_child2.LayoutGetHeight())

	assertFloatEqual(t, 427, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 213, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 320, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 640, root.LayoutGetWidth())
	assertFloatEqual(t, 320, root.LayoutGetHeight())

	assertFloatEqual(t, 427, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 213, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 320, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 213, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 214, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 320, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 214, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 107, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 107, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 214, rootChild1_child1.LayoutGetWidth())
	assertFloatEqual(t, 106, rootChild1_child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child1Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1Child0.LayoutGetTop())
	assertFloatEqual(t, 214, rootChild1_child1Child0.LayoutGetWidth())
	assertFloatEqual(t, 106, rootChild1_child1Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 213, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 214, rootChild1_child2.LayoutGetWidth())
	assertFloatEqual(t, 107, rootChild1_child2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 213, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 320, rootChild2.LayoutGetHeight())
}
