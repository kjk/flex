package flex

import "testing"

func TestJustify_content_row_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 102)
	NodeStyleSetHeight(root, 102)

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
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 20, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 92, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 82, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 72, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_row_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetJustifyContent(root, JustifyFlexEnd)
	NodeStyleSetWidth(root, 102)
	NodeStyleSetHeight(root, 102)

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
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 72, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 82, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 92, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_row_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetWidth(root, 102)
	NodeStyleSetHeight(root, 102)

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
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 36, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 56, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 56, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 36, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_row_space_between(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetJustifyContent(root, JustifySpaceBetween)
	NodeStyleSetWidth(root, 102)
	NodeStyleSetHeight(root, 102)

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
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 92, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 92, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_row_space_around(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetJustifyContent(root, JustifySpaceAround)
	NodeStyleSetWidth(root, 102)
	NodeStyleSetHeight(root, 102)

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
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 12, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 80, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 12, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_column_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 102)
	NodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_column_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyFlexEnd)
	NodeStyleSetWidth(root, 102)
	NodeStyleSetHeight(root, 102)

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
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 72, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 82, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 92, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 72, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 82, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 92, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_column_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetWidth(root, 102)
	NodeStyleSetHeight(root, 102)

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
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 36, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 56, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 36, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 56, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_column_space_between(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifySpaceBetween)
	NodeStyleSetWidth(root, 102)
	NodeStyleSetHeight(root, 102)

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
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 92, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 92, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_column_space_around(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifySpaceAround)
	NodeStyleSetWidth(root, 102)
	NodeStyleSetHeight(root, 102)

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
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 12, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 12, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))
}
