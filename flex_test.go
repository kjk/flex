package flex

import "testing"

func TestFlex_basis_flex_grow_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 75, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 75, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 75, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 75, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())
}

func TestFlex_basis_flex_grow_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 75, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 75, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 75, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())
}

func TestFlex_basis_flex_shrink_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexShrink(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 100)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexBasis(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())
}

func TestFlex_basis_flex_shrink_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexShrink(rootChild0, 1)
	NodeStyleSetFlexBasis(rootChild0, 100)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexBasis(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())
}

func TestFlex_shrink_to_zero(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetHeight(root, 75)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexShrink(rootChild1, 1)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 75, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 75, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild2.LayoutGetHeight())
}

func TestFlex_basis_overrides_main_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

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
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild2.LayoutGetHeight())
}

func TestFlex_grow_shrink_at_most(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetFlexShrink(rootChild0Child0, 1)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetHeight())
}

func TestFlex_grow_less_than_factor_one(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 0.2)
	NodeStyleSetFlexBasis(rootChild0, 40)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 0.2)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild2, 0.4)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 132, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 132, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 92, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 224, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 184, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 132, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 132, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 92, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 224, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 184, rootChild2.LayoutGetHeight())
}
