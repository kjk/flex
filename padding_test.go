package flex

import "testing"

func TestPadding_no_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetPadding(root, EdgeLeft, 10)
	NodeStyleSetPadding(root, EdgeTop, 10)
	NodeStyleSetPadding(root, EdgeRight, 10)
	NodeStyleSetPadding(root, EdgeBottom, 10)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 20, root.LayoutGetWidth())
	assertFloatEqual(t, 20, root.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 20, root.LayoutGetWidth())
	assertFloatEqual(t, 20, root.LayoutGetHeight())
}

func TestPadding_container_match_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetPadding(root, EdgeLeft, 10)
	NodeStyleSetPadding(root, EdgeTop, 10)
	NodeStyleSetPadding(root, EdgeRight, 10)
	NodeStyleSetPadding(root, EdgeBottom, 10)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 30, root.LayoutGetWidth())
	assertFloatEqual(t, 30, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 30, root.LayoutGetWidth())
	assertFloatEqual(t, 30, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestPadding_flex_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetPadding(root, EdgeLeft, 10)
	NodeStyleSetPadding(root, EdgeTop, 10)
	NodeStyleSetPadding(root, EdgeRight, 10)
	NodeStyleSetPadding(root, EdgeBottom, 10)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetWidth(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0.LayoutGetHeight())
}

func TestPadding_stretch_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetPadding(root, EdgeLeft, 10)
	NodeStyleSetPadding(root, EdgeTop, 10)
	NodeStyleSetPadding(root, EdgeRight, 10)
	NodeStyleSetPadding(root, EdgeBottom, 10)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestPadding_center_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetPadding(root, EdgeStart, 10)
	NodeStyleSetPadding(root, EdgeEnd, 20)
	NodeStyleSetPadding(root, EdgeBottom, 20)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 35, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 35, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestChild_with_padding_align_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyFlexEnd)
	NodeStyleSetAlignItems(root, AlignFlexEnd)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetPadding(rootChild0, EdgeLeft, 20)
	NodeStyleSetPadding(rootChild0, EdgeTop, 20)
	NodeStyleSetPadding(rootChild0, EdgeRight, 20)
	NodeStyleSetPadding(rootChild0, EdgeBottom, 20)
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 100)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}
