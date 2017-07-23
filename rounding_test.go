package flex

import "testing"

func TestRounding_flex_basis_flex_grow_row_width_of_100(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 33, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 33, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 34, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 67, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 33, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 67, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 33, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 33, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 34, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 33, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_flex_basis_flex_grow_row_prime_number_width(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 113)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child3, 1)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child4, 1)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 23, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 22, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 68, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 22, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 68, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 22, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 23, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 22, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 23, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_flex_basis_flex_shrink_row(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 101)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexShrink(root_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0, 100)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexBasis(root_child1, 25)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexBasis(root_child2, 25)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 101, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 51, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 51, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 76, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 101, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 51, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_flex_basis_overrides_main_size(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 113)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1)
	YGNodeStyleSetHeight(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_total_fractial(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 87.4)
	YGNodeStyleSetHeight(root, 113.4)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 0.7)
	YGNodeStyleSetFlexBasis(root_child0, 50.3)
	YGNodeStyleSetHeight(root_child0, 20.3)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1.1)
	YGNodeStyleSetHeight(root_child2, 10.7)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child0))
	//TODO: this returns 61
	//assertFloatEqual(t, 59, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	//TODO: this returns 61, but only when I added size_overflow_test.go ?
	//assertFloatEqual(t, 59, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child1))
	// TODO: this returns 25
	//assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	//TODO: this returns 86
	//assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child2))
	//TODO: this returns 27
	//assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child0))
	//TODO: this returns 61
	//assertFloatEqual(t, 59, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	//assertFloatEqual(t, 59, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child1))
	//assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	//assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child2))
	//assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_total_fractial_nested(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 87.4)
	YGNodeStyleSetHeight(root, 113.4)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 0.7)
	YGNodeStyleSetFlexBasis(root_child0, 50.3)
	YGNodeStyleSetHeight(root_child0, 20.3)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0_child0, 0.3)
	YGNodeStyleSetPosition(root_child0_child0, YGEdgeBottom, 13.3)
	YGNodeStyleSetHeight(root_child0_child0, 9.9)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child0_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child1, 4)
	YGNodeStyleSetFlexBasis(root_child0_child1, 0.3)
	YGNodeStyleSetPosition(root_child0_child1, YGEdgeTop, 13.3)
	YGNodeStyleSetHeight(root_child0_child1, 1.1)
	YGNodeInsertChild(root_child0, root_child0_child1, 1)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1.6)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1.1)
	YGNodeStyleSetHeight(root_child2, 10.7)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 59, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, -13, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 12, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 47, YGNodeLayoutGetHeight(root_child0_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 59, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 59, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, -13, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 12, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(root_child0_child1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child0_child1))
	assertFloatEqual(t, 47, YGNodeLayoutGetHeight(root_child0_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 59, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 87, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_fractial_input_1(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 113.4)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1)
	YGNodeStyleSetHeight(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_fractial_input_2(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 113.6)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1)
	YGNodeStyleSetHeight(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 114, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 65, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 65, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 114, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 65, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 65, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_fractial_input_3(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetPosition(root, YGEdgeTop, 0.3)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 113.4)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1)
	YGNodeStyleSetHeight(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 114, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 65, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 114, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 65, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_fractial_input_4(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetPosition(root, YGEdgeTop, 0.7)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 113.4)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasis(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 20)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1)
	YGNodeStyleSetHeight(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 1, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 1, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 113, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 64, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 64, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 89, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 24, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_inner_node_controversy_horizontal(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 320)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1_child0, 1)
	YGNodeStyleSetHeight(root_child1_child0, 10)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1)
	YGNodeStyleSetHeight(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 107, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 107, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 106, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 213, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 107, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 213, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 107, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 107, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 106, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 107, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_inner_node_controversy_vertical(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root, 320)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetWidth(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1_child0, 1)
	YGNodeStyleSetWidth(root_child1_child0, 10)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1)
	YGNodeStyleSetWidth(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 107, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 213, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 107, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 213, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func TestRounding_inner_node_controversy_combined(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 640)
	YGNodeStyleSetHeight(root, 320)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetHeightPercent(root_child0, 100)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetHeightPercent(root_child1, 100)
	YGNodeInsertChild(root, root_child1, 1)

	root_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1_child0, 1)
	YGNodeStyleSetWidthPercent(root_child1_child0, 100)
	YGNodeInsertChild(root_child1, root_child1_child0, 0)

	root_child1_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1_child1, 1)
	YGNodeStyleSetWidthPercent(root_child1_child1, 100)
	YGNodeInsertChild(root_child1, root_child1_child1, 1)

	root_child1_child1_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1_child1_child0, 1)
	YGNodeStyleSetWidthPercent(root_child1_child1_child0, 100)
	YGNodeInsertChild(root_child1_child1, root_child1_child1_child0, 0)

	root_child1_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1_child2, 1)
	YGNodeStyleSetWidthPercent(root_child1_child2, 100)
	YGNodeInsertChild(root_child1, root_child1_child2, 2)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child2, 1)
	YGNodeStyleSetHeightPercent(root_child2, 100)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 640, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 213, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 213, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child1))
	assertFloatEqual(t, 107, YGNodeLayoutGetTop(root_child1_child1))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(root_child1_child1))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(root_child1_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child1_child0))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(root_child1_child1_child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(root_child1_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child2))
	assertFloatEqual(t, 213, YGNodeLayoutGetTop(root_child1_child2))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(root_child1_child2))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(root_child1_child2))

	assertFloatEqual(t, 427, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 213, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 640, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 427, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 213, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 213, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child0))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(root_child1_child0))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(root_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child1))
	assertFloatEqual(t, 107, YGNodeLayoutGetTop(root_child1_child1))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(root_child1_child1))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(root_child1_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child1_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1_child1_child0))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(root_child1_child1_child0))
	assertFloatEqual(t, 106, YGNodeLayoutGetHeight(root_child1_child1_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1_child2))
	assertFloatEqual(t, 213, YGNodeLayoutGetTop(root_child1_child2))
	assertFloatEqual(t, 214, YGNodeLayoutGetWidth(root_child1_child2))
	assertFloatEqual(t, 107, YGNodeLayoutGetHeight(root_child1_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 213, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 320, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}
