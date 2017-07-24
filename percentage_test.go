package flex

import "testing"

func TestPercentage_width_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidthPercent(rootChild0, 30)
	NodeStyleSetHeightPercent(rootChild0, 30)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 140, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))
}

func TestPercentage_position_left_top(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 400)
	NodeStyleSetHeight(root, 400)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetPositionPercent(rootChild0, EdgeLeft, 10)
	NodeStyleSetPositionPercent(rootChild0, EdgeTop, 20)
	NodeStyleSetWidthPercent(rootChild0, 45)
	NodeStyleSetHeightPercent(rootChild0, 55)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 400, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 400, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 180, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 220, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 400, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 400, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 260, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 180, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 220, NodeLayoutGetHeight(rootChild0))
}

func TestPercentage_position_bottom_right(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 500)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetPositionPercent(rootChild0, EdgeRight, 20)
	NodeStyleSetPositionPercent(rootChild0, EdgeBottom, 10)
	NodeStyleSetWidthPercent(rootChild0, 55)
	NodeStyleSetHeightPercent(rootChild0, 15)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 500, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, -100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -50, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 275, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 75, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 500, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 125, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -50, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 275, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 75, NodeLayoutGetHeight(rootChild0))
}

func TestPercentage_flex_basis(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetFlexBasisPercent(rootChild1, 25)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 125, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 125, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 75, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 75, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 125, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 75, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_cross(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetFlexBasisPercent(rootChild1, 25)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 125, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 125, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 75, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 125, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 125, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 75, NodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_cross_min_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetMinHeightPercent(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 2)
	NodeStyleSetMinHeightPercent(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 140, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 140, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 140, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 140, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_main_max_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 10)
	NodeStyleSetMaxHeightPercent(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 4)
	NodeStyleSetFlexBasisPercent(rootChild1, 10)
	NodeStyleSetMaxHeightPercent(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 52, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 120, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 52, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 148, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 148, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 52, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 120, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 148, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_cross_max_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 10)
	NodeStyleSetMaxHeightPercent(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 4)
	NodeStyleSetFlexBasisPercent(rootChild1, 10)
	NodeStyleSetMaxHeightPercent(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 120, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 120, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 120, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 120, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_main_max_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 15)
	NodeStyleSetMaxWidthPercent(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 4)
	NodeStyleSetFlexBasisPercent(rootChild1, 10)
	NodeStyleSetMaxWidthPercent(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 120, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_cross_max_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 10)
	NodeStyleSetMaxWidthPercent(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 4)
	NodeStyleSetFlexBasisPercent(rootChild1, 15)
	NodeStyleSetMaxWidthPercent(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 150, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 160, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 150, NodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_main_min_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 15)
	NodeStyleSetMinWidthPercent(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 4)
	NodeStyleSetFlexBasisPercent(rootChild1, 10)
	NodeStyleSetMinWidthPercent(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 120, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_cross_min_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 10)
	NodeStyleSetMinWidthPercent(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 4)
	NodeStyleSetFlexBasisPercent(rootChild1, 15)
	NodeStyleSetMinWidthPercent(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 150, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 150, NodeLayoutGetHeight(rootChild1))
}

func TestPercentage_multiple_nested_with_padding_margin_and_percentage_values(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 10)
	NodeStyleSetMargin(rootChild0, EdgeLeft, 5)
	NodeStyleSetMargin(rootChild0, EdgeTop, 5)
	NodeStyleSetMargin(rootChild0, EdgeRight, 5)
	NodeStyleSetMargin(rootChild0, EdgeBottom, 5)
	NodeStyleSetPadding(rootChild0, EdgeLeft, 3)
	NodeStyleSetPadding(rootChild0, EdgeTop, 3)
	NodeStyleSetPadding(rootChild0, EdgeRight, 3)
	NodeStyleSetPadding(rootChild0, EdgeBottom, 3)
	NodeStyleSetMinWidthPercent(rootChild0, 60)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild0Child0, EdgeLeft, 5)
	NodeStyleSetMargin(rootChild0Child0, EdgeTop, 5)
	NodeStyleSetMargin(rootChild0Child0, EdgeRight, 5)
	NodeStyleSetMargin(rootChild0Child0, EdgeBottom, 5)
	NodeStyleSetPaddingPercent(rootChild0Child0, EdgeLeft, 3)
	NodeStyleSetPaddingPercent(rootChild0Child0, EdgeTop, 3)
	NodeStyleSetPaddingPercent(rootChild0Child0, EdgeRight, 3)
	NodeStyleSetPaddingPercent(rootChild0Child0, EdgeBottom, 3)
	NodeStyleSetWidthPercent(rootChild0Child0, 50)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	YGNodeStyleSetMarginPercent(rootChild0Child0_child0, EdgeLeft, 5)
	YGNodeStyleSetMarginPercent(rootChild0Child0_child0, EdgeTop, 5)
	YGNodeStyleSetMarginPercent(rootChild0Child0_child0, EdgeRight, 5)
	YGNodeStyleSetMarginPercent(rootChild0Child0_child0, EdgeBottom, 5)
	NodeStyleSetPadding(rootChild0Child0_child0, EdgeLeft, 3)
	NodeStyleSetPadding(rootChild0Child0_child0, EdgeTop, 3)
	NodeStyleSetPadding(rootChild0Child0_child0, EdgeRight, 3)
	NodeStyleSetPadding(rootChild0Child0_child0, EdgeBottom, 3)
	NodeStyleSetWidthPercent(rootChild0Child0_child0, 45)
	NodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 4)
	NodeStyleSetFlexBasisPercent(rootChild1, 15)
	NodeStyleSetMinWidthPercent(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 190, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 48, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 8, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 8, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 92, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 25, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 10, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 36, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 6, NodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 58, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 142, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 190, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 48, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 90, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 8, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 92, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 25, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 46, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 36, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 6, NodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 58, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 142, NodeLayoutGetHeight(rootChild1))
}

func TestPercentage_margin_should_calculate_based_only_on_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMarginPercent(rootChild0, EdgeLeft, 10)
	YGNodeStyleSetMarginPercent(rootChild0, EdgeTop, 10)
	YGNodeStyleSetMarginPercent(rootChild0, EdgeRight, 10)
	YGNodeStyleSetMarginPercent(rootChild0, EdgeBottom, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 10)
	NodeStyleSetHeight(rootChild0Child0, 10)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 160, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 160, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 150, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0Child0))
}

func TestPercentage_padding_should_calculate_based_only_on_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetPaddingPercent(rootChild0, EdgeLeft, 10)
	NodeStyleSetPaddingPercent(rootChild0, EdgeTop, 10)
	NodeStyleSetPaddingPercent(rootChild0, EdgeRight, 10)
	NodeStyleSetPaddingPercent(rootChild0, EdgeBottom, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 10)
	NodeStyleSetHeight(rootChild0Child0, 10)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 170, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0Child0))
}

func TestPercentage_absolute_position(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPositionPercent(rootChild0, EdgeLeft, 30)
	NodeStyleSetPositionPercent(rootChild0, EdgeTop, 10)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 60, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 60, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
}

func TestPercentage_width_height_undefined_parent_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidthPercent(rootChild0, 50)
	NodeStyleSetHeightPercent(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0))
}

func TestPercent_within_flex_grow(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 350)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 100)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	NodeStyleSetWidthPercent(rootChild1child0, 100)
	NodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 100)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 350, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 250, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 350, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 250, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))
}

func TestPercentage_container_in_wrapping_container(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0Child0, FlexDirectionRow)
	NodeStyleSetJustifyContent(rootChild0Child0, JustifyCenter)
	NodeStyleSetWidthPercent(rootChild0Child0, 100)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0_child0, 50)
	NodeStyleSetHeight(rootChild0Child0_child0, 50)
	NodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild0Child0_child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0_child1, 50)
	NodeStyleSetHeight(rootChild0Child0_child1, 50)
	NodeInsertChild(rootChild0Child0, rootChild0Child0_child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 75, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 50, rootChild0Child0_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0_child1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0Child0_child1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0_child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 75, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, rootChild0Child0_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, rootChild0Child0_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0_child1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0Child0_child1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0_child1))
}

func TestPercent_absolute_position(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 60)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPositionPercent(rootChild0, EdgeLeft, 50)
	NodeStyleSetWidthPercent(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidthPercent(rootChild0Child0, 100)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetWidthPercent(rootChild0child1, 100)
	NodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 60, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 30, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 60, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 60, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 60, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 60, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 30, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 60, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, -60, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 60, NodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0child1))
}
