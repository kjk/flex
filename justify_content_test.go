package flex

import "testing"

func TestJustify_content_row_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	root.StyleSetWidth(102)
	root.StyleSetHeight(102)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(10)
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
	root.StyleSetWidth(102)
	root.StyleSetHeight(102)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(10)
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
	root.StyleSetWidth(102)
	root.StyleSetHeight(102)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(10)
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
	root.StyleSetWidth(102)
	root.StyleSetHeight(102)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(10)
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
	root.StyleSetWidth(102)
	root.StyleSetHeight(102)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(10)
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
	root.StyleSetWidth(102)
	root.StyleSetHeight(102)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetHeight(10)
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
	root.StyleSetWidth(102)
	root.StyleSetHeight(102)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetHeight(10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetHeight(10)
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
	root.StyleSetWidth(102)
	root.StyleSetHeight(102)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetHeight(10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetHeight(10)
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
	root.StyleSetWidth(102)
	root.StyleSetHeight(102)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetHeight(10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetHeight(10)
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
	root.StyleSetWidth(102)
	root.StyleSetHeight(102)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetHeight(10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetHeight(10)
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
