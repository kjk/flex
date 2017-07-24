package flex

import "testing"

func TestFlex_direction_column_no_height(t *testing.T) {
	config := NewConfig()
	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))
}

func TestFlex_direction_row_no_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 30, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 20, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 30, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))
}

func TestFlex_direction_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))
}

func TestFlex_direction_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 20, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 80, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 70, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))
}

func TestFlex_direction_column_reverse(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionColumnReverse)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 90, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 70, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 90, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 70, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))
}

func TestFlex_direction_row_reverse(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRowReverse)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 80, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 70, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 20, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))
}
