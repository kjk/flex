package flex

import "testing"

func TestPercentage_width_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidthPercent(rootChild0, 30)
	YGNodeStyleSetHeightPercent(rootChild0, 30)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 140, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))
}

func TestPercentage_position_left_top(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 400)
	YGNodeStyleSetHeight(root, 400)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionPercent(rootChild0, EdgeLeft, 10)
	YGNodeStyleSetPositionPercent(rootChild0, EdgeTop, 20)
	YGNodeStyleSetWidthPercent(rootChild0, 45)
	YGNodeStyleSetHeightPercent(rootChild0, 55)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 400, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 400, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 180, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 220, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 400, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 400, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 260, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 180, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 220, YGNodeLayoutGetHeight(rootChild0))
}

func TestPercentage_position_bottom_right(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 500)
	YGNodeStyleSetHeight(root, 500)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionPercent(rootChild0, EdgeRight, 20)
	YGNodeStyleSetPositionPercent(rootChild0, EdgeBottom, 10)
	YGNodeStyleSetWidthPercent(rootChild0, 55)
	YGNodeStyleSetHeightPercent(rootChild0, 15)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, -100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, -50, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 275, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 125, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, -50, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 275, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(rootChild0))
}

func TestPercentage_flex_basis(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 25)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 125, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 125, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 75, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 75, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 125, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 75, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_cross(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 25)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 125, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 125, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 125, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 125, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 75, YGNodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_cross_min_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMinHeightPercent(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 2)
	YGNodeStyleSetMinHeightPercent(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 140, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 140, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 140, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 140, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_main_max_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 10)
	YGNodeStyleSetMaxHeightPercent(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 4)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 10)
	YGNodeStyleSetMaxHeightPercent(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 52, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 148, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 148, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 52, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 148, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_cross_max_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 10)
	YGNodeStyleSetMaxHeightPercent(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 4)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 10)
	YGNodeStyleSetMaxHeightPercent(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 120, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 120, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_main_max_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 15)
	YGNodeStyleSetMaxWidthPercent(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 4)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 10)
	YGNodeStyleSetMaxWidthPercent(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_cross_max_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 10)
	YGNodeStyleSetMaxWidthPercent(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 4)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 15)
	YGNodeStyleSetMaxWidthPercent(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 160, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_main_min_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 15)
	YGNodeStyleSetMinWidthPercent(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 4)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 10)
	YGNodeStyleSetMinWidthPercent(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 120, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild1))
}

func TestPercentage_flex_basis_cross_min_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 10)
	YGNodeStyleSetMinWidthPercent(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 4)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 15)
	YGNodeStyleSetMinWidthPercent(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(rootChild1))
}

func TestPercentage_multiple_nested_with_padding_margin_and_percentage_values(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 10)
	YGNodeStyleSetMargin(rootChild0, EdgeLeft, 5)
	YGNodeStyleSetMargin(rootChild0, EdgeTop, 5)
	YGNodeStyleSetMargin(rootChild0, EdgeRight, 5)
	YGNodeStyleSetMargin(rootChild0, EdgeBottom, 5)
	YGNodeStyleSetPadding(rootChild0, EdgeLeft, 3)
	YGNodeStyleSetPadding(rootChild0, EdgeTop, 3)
	YGNodeStyleSetPadding(rootChild0, EdgeRight, 3)
	YGNodeStyleSetPadding(rootChild0, EdgeBottom, 3)
	YGNodeStyleSetMinWidthPercent(rootChild0, 60)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(rootChild0Child0, EdgeLeft, 5)
	YGNodeStyleSetMargin(rootChild0Child0, EdgeTop, 5)
	YGNodeStyleSetMargin(rootChild0Child0, EdgeRight, 5)
	YGNodeStyleSetMargin(rootChild0Child0, EdgeBottom, 5)
	YGNodeStyleSetPaddingPercent(rootChild0Child0, EdgeLeft, 3)
	YGNodeStyleSetPaddingPercent(rootChild0Child0, EdgeTop, 3)
	YGNodeStyleSetPaddingPercent(rootChild0Child0, EdgeRight, 3)
	YGNodeStyleSetPaddingPercent(rootChild0Child0, EdgeBottom, 3)
	YGNodeStyleSetWidthPercent(rootChild0Child0, 50)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMarginPercent(rootChild0Child0_child0, EdgeLeft, 5)
	YGNodeStyleSetMarginPercent(rootChild0Child0_child0, EdgeTop, 5)
	YGNodeStyleSetMarginPercent(rootChild0Child0_child0, EdgeRight, 5)
	YGNodeStyleSetMarginPercent(rootChild0Child0_child0, EdgeBottom, 5)
	YGNodeStyleSetPadding(rootChild0Child0_child0, EdgeLeft, 3)
	YGNodeStyleSetPadding(rootChild0Child0_child0, EdgeTop, 3)
	YGNodeStyleSetPadding(rootChild0Child0_child0, EdgeRight, 3)
	YGNodeStyleSetPadding(rootChild0Child0_child0, EdgeBottom, 3)
	YGNodeStyleSetWidthPercent(rootChild0Child0_child0, 45)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 4)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 15)
	YGNodeStyleSetMinWidthPercent(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 190, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 48, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 8, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 8, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 92, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 36, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 6, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 58, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 142, YGNodeLayoutGetHeight(rootChild1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 5, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 190, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 48, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 8, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 92, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 46, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 36, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 6, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 58, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 142, YGNodeLayoutGetHeight(rootChild1))
}

func TestPercentage_margin_should_calculate_based_only_on_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetMarginPercent(rootChild0, EdgeLeft, 10)
	YGNodeStyleSetMarginPercent(rootChild0, EdgeTop, 10)
	YGNodeStyleSetMarginPercent(rootChild0, EdgeRight, 10)
	YGNodeStyleSetMarginPercent(rootChild0, EdgeBottom, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0, 10)
	YGNodeStyleSetHeight(rootChild0Child0, 10)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0Child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 150, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0Child0))
}

func TestPercentage_padding_should_calculate_based_only_on_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetPaddingPercent(rootChild0, EdgeLeft, 10)
	YGNodeStyleSetPaddingPercent(rootChild0, EdgeTop, 10)
	YGNodeStyleSetPaddingPercent(rootChild0, EdgeRight, 10)
	YGNodeStyleSetPaddingPercent(rootChild0, EdgeBottom, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0, 10)
	YGNodeStyleSetHeight(rootChild0Child0, 10)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0Child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0Child0))
}

func TestPercentage_absolute_position(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(rootChild0, YGPositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(rootChild0, EdgeLeft, 30)
	YGNodeStyleSetPositionPercent(rootChild0, EdgeTop, 10)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestPercentage_width_height_undefined_parent_size(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidthPercent(rootChild0, 50)
	YGNodeStyleSetHeightPercent(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))
}

func TestPercent_within_flex_grow(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 350)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild1child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidthPercent(rootChild1child0, 100)
	YGNodeInsertChild(rootChild1, rootChild1child0, 0)

	rootChild2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 100)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 350, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 350, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild1child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}

func TestPercentage_container_in_wrapping_container(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetJustifyContent(root, YGJustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 200)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild0Child0, YGFlexDirectionRow)
	YGNodeStyleSetJustifyContent(rootChild0Child0, YGJustifyCenter)
	YGNodeStyleSetWidthPercent(rootChild0Child0, 100)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0_child0, 50)
	YGNodeStyleSetHeight(rootChild0Child0_child0, 50)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild0Child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild0Child0_child1, 50)
	YGNodeStyleSetHeight(rootChild0Child0_child1, 50)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 75, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0Child0_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 75, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0_child1))
}

func TestPercent_absolute_position(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 60)
	YGNodeStyleSetHeight(root, 50)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild0, YGFlexDirectionRow)
	YGNodeStyleSetPositionType(rootChild0, YGPositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(rootChild0, EdgeLeft, 50)
	YGNodeStyleSetWidthPercent(rootChild0, 100)
	YGNodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidthPercent(rootChild0Child0, 100)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidthPercent(rootChild0child1, 100)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, -60, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0child1))
}
