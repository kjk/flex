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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 75, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 75, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 75, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 75, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetHeight(rootChild1))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 75, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 75, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 75, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 75, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 75, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild2))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0Child0))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 132, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 132, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 92, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 224, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 184, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 132, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 132, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 92, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 224, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 184, NodeLayoutGetHeight(rootChild2))
}
