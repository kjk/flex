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
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 20, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 92, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 82, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 72, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild2.LayoutGetHeight())
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
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 72, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 82, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 92, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild2.LayoutGetHeight())
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
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 36, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 56, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 56, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 36, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild2.LayoutGetHeight())
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
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 92, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 92, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild2.LayoutGetHeight())
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
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 12, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 46, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 12, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 102, rootChild2.LayoutGetHeight())
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
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())
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
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 72, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 82, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 92, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 72, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 82, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 92, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())
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
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 36, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 56, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 36, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 56, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())
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
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 92, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 92, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())
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
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 12, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 102, root.LayoutGetWidth())
	assertFloatEqual(t, 102, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 12, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 46, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 102, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2.LayoutGetHeight())
}
