package flex

import "testing"

func TestPercentage_width_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidthPercent(root_child0, 30)
	YGNodeStyleSetHeightPercent(root_child0, 30)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 140, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

}

func TestPercentage_position_left_top(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 400)
	YGNodeStyleSetHeight(root, 400)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionPercent(root_child0, YGEdgeLeft, 10)
	YGNodeStyleSetPositionPercent(root_child0, YGEdgeTop, 20)
	YGNodeStyleSetWidthPercent(root_child0, 45)
	YGNodeStyleSetHeightPercent(root_child0, 55)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 400, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 400, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 180, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 220, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 400, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 400, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 260, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 180, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 220, YGNodeLayoutGetHeight(root_child0))

}

func TestPercentage_position_bottom_right(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 500)
	YGNodeStyleSetHeight(root, 500)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionPercent(root_child0, YGEdgeRight, 20)
	YGNodeStyleSetPositionPercent(root_child0, YGEdgeBottom, 10)
	YGNodeStyleSetWidthPercent(root_child0, 55)
	YGNodeStyleSetHeightPercent(root_child0, 15)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, -50, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 275, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 125, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, -50, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 275, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(root_child0))

}

func TestPercentage_flex_basis(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetFlexBasisPercent(root_child1, 25)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 125, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 125, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 75, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 75, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 125, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 75, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child1))

}

func TestPercentage_flex_basis_cross(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetFlexBasisPercent(root_child1, 25)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 125, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 125, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 125, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 125, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(root_child1))

}

func TestPercentage_flex_basis_cross_min_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMinHeightPercent(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 2)
	YGNodeStyleSetMinHeightPercent(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 140, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 140, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 140, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 140, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child1))

}

func TestPercentage_flex_basis_main_max_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 10)
	YGNodeStyleSetMaxHeightPercent(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 4)
	YGNodeStyleSetFlexBasisPercent(root_child1, 10)
	YGNodeStyleSetMaxHeightPercent(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 52, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 148, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 148, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 148, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child1))

}

func TestPercentage_flex_basis_cross_max_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 10)
	YGNodeStyleSetMaxHeightPercent(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 4)
	YGNodeStyleSetFlexBasisPercent(root_child1, 10)
	YGNodeStyleSetMaxHeightPercent(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 120, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 120, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child1))

}

func TestPercentage_flex_basis_main_max_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 15)
	YGNodeStyleSetMaxWidthPercent(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 4)
	YGNodeStyleSetFlexBasisPercent(root_child1, 10)
	YGNodeStyleSetMaxWidthPercent(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child1))

}

func TestPercentage_flex_basis_cross_max_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 10)
	YGNodeStyleSetMaxWidthPercent(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 4)
	YGNodeStyleSetFlexBasisPercent(root_child1, 15)
	YGNodeStyleSetMaxWidthPercent(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 160, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(root_child1))

}

func TestPercentage_flex_basis_main_min_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 15)
	YGNodeStyleSetMinWidthPercent(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 4)
	YGNodeStyleSetFlexBasisPercent(root_child1, 10)
	YGNodeStyleSetMinWidthPercent(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root_child1))

}

func TestPercentage_flex_basis_cross_min_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 10)
	YGNodeStyleSetMinWidthPercent(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 4)
	YGNodeStyleSetFlexBasisPercent(root_child1, 15)
	YGNodeStyleSetMinWidthPercent(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(root_child1))

}

func TestPercentage_multiple_nested_with_padding_margin_and_percentage_values(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 10)
	YGNodeStyleSetMargin(root_child0, YGEdgeLeft, 5)
	YGNodeStyleSetMargin(root_child0, YGEdgeTop, 5)
	YGNodeStyleSetMargin(root_child0, YGEdgeRight, 5)
	YGNodeStyleSetMargin(root_child0, YGEdgeBottom, 5)
	YGNodeStyleSetPadding(root_child0, YGEdgeLeft, 3)
	YGNodeStyleSetPadding(root_child0, YGEdgeTop, 3)
	YGNodeStyleSetPadding(root_child0, YGEdgeRight, 3)
	YGNodeStyleSetPadding(root_child0, YGEdgeBottom, 3)
	YGNodeStyleSetMinWidthPercent(root_child0, 60)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root_child0_child0, YGEdgeLeft, 5)
	YGNodeStyleSetMargin(root_child0_child0, YGEdgeTop, 5)
	YGNodeStyleSetMargin(root_child0_child0, YGEdgeRight, 5)
	YGNodeStyleSetMargin(root_child0_child0, YGEdgeBottom, 5)
	YGNodeStyleSetPaddingPercent(root_child0_child0, YGEdgeLeft, 3)
	YGNodeStyleSetPaddingPercent(root_child0_child0, YGEdgeTop, 3)
	YGNodeStyleSetPaddingPercent(root_child0_child0, YGEdgeRight, 3)
	YGNodeStyleSetPaddingPercent(root_child0_child0, YGEdgeBottom, 3)
	YGNodeStyleSetWidthPercent(root_child0_child0, 50)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMarginPercent(root_child0_child0_child0, YGEdgeLeft, 5)
	YGNodeStyleSetMarginPercent(root_child0_child0_child0, YGEdgeTop, 5)
	YGNodeStyleSetMarginPercent(root_child0_child0_child0, YGEdgeRight, 5)
	YGNodeStyleSetMarginPercent(root_child0_child0_child0, YGEdgeBottom, 5)
	YGNodeStyleSetPadding(root_child0_child0_child0, YGEdgeLeft, 3)
	YGNodeStyleSetPadding(root_child0_child0_child0, YGEdgeTop, 3)
	YGNodeStyleSetPadding(root_child0_child0_child0, YGEdgeRight, 3)
	YGNodeStyleSetPadding(root_child0_child0_child0, YGEdgeBottom, 3)
	YGNodeStyleSetWidthPercent(root_child0_child0_child0, 45)
	YGNodeInsertChild(root_child0_child0, root_child0_child0_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 4)
	YGNodeStyleSetFlexBasisPercent(root_child1, 15)
	YGNodeStyleSetMinWidthPercent(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 190, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 48, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 8, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 8, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 92, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 36, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 6, YGNodeLayoutGetHeight(root_child0_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 58, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 142, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 190, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 48, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 8, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 92, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 46, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 36, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 6, YGNodeLayoutGetHeight(root_child0_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 58, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 142, YGNodeLayoutGetHeight(root_child1))

}

func TestPercentage_margin_should_calculate_based_only_on_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetMarginPercent(root_child0, YGEdgeLeft, 10)
	YGNodeStyleSetMarginPercent(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetMarginPercent(root_child0, YGEdgeRight, 10)
	YGNodeStyleSetMarginPercent(root_child0, YGEdgeBottom, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0, 10)
	YGNodeStyleSetHeight(root_child0_child0, 10)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 160, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 160, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 150, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0_child0))

}

func TestPercentage_padding_should_calculate_based_only_on_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetPaddingPercent(root_child0, YGEdgeLeft, 10)
	YGNodeStyleSetPaddingPercent(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetPaddingPercent(root_child0, YGEdgeRight, 10)
	YGNodeStyleSetPaddingPercent(root_child0, YGEdgeBottom, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0, 10)
	YGNodeStyleSetHeight(root_child0_child0, 10)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0_child0))

}

func TestPercentage_absolute_position(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(root_child0, YGEdgeLeft, 30)
	YGNodeStyleSetPositionPercent(root_child0, YGEdgeTop, 10)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

}

func TestPercentage_width_height_undefined_parent_size(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidthPercent(root_child0, 50)
	YGNodeStyleSetHeightPercent(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))

}

func TestPercent_within_flex_grow(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 350)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 100)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidthPercent(root_child1_child0, 100)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 100)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 350, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 350, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

}

func TestPercentage_container_in_wrapping_container(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0_child0, YGFlexDirectionRow)
	YGNodeStyleSetJustifyContent(root_child0_child0, YGJustifyCenter)
	YGNodeStyleSetWidthPercent(root_child0_child0, 100)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0_child0, 50)
	YGNodeStyleSetHeight(root_child0_child0_child0, 50)
	YGNodeInsertChild(root_child0_child0, root_child0_child0_child0, 0)

	root_child0_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0_child0_child1, 50)
	YGNodeStyleSetHeight(root_child0_child0_child1, 50)
	YGNodeInsertChild(root_child0_child0, root_child0_child0_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 75, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0_child0_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 75, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0_child1))

}

func TestPercent_absolute_position(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 60)
	YGNodeStyleSetHeight(root, 50)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(root_child0, YGEdgeLeft, 50)
	YGNodeStyleSetWidthPercent(root_child0, 100)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidthPercent(root_child0_child0, 100)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidthPercent(root_child0_child1, 100)
	YGNodeInsertChild(root_child0, root_child0_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, -60, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child1))

}
