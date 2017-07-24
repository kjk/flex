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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 20, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(root))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 20, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(root))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 30, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 30, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 35, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 35, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
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

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))
}
