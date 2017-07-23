package flex

import "testing"

func TestFlex_direction_column_no_height(t *testing.T) {
	config := NewConfig()
	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))
}

func TestFlex_direction_row_no_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}

func TestFlex_direction_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))
}

func TestFlex_direction_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}

func TestFlex_direction_column_reverse(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionColumnReverse)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))
}

func TestFlex_direction_row_reverse(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRowReverse)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}
