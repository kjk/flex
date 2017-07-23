package flex

import "testing"

func TestPadding_no_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetPadding(root, EdgeLeft, 10)
	YGNodeStyleSetPadding(root, EdgeTop, 10)
	YGNodeStyleSetPadding(root, EdgeRight, 10)
	YGNodeStyleSetPadding(root, EdgeBottom, 10)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root))
}

func TestPadding_container_match_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetPadding(root, EdgeLeft, 10)
	YGNodeStyleSetPadding(root, EdgeTop, 10)
	YGNodeStyleSetPadding(root, EdgeRight, 10)
	YGNodeStyleSetPadding(root, EdgeBottom, 10)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestPadding_flex_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetPadding(root, EdgeLeft, 10)
	YGNodeStyleSetPadding(root, EdgeTop, 10)
	YGNodeStyleSetPadding(root, EdgeRight, 10)
	YGNodeStyleSetPadding(root, EdgeBottom, 10)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetWidth(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))
}

func TestPadding_stretch_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetPadding(root, EdgeLeft, 10)
	YGNodeStyleSetPadding(root, EdgeTop, 10)
	YGNodeStyleSetPadding(root, EdgeRight, 10)
	YGNodeStyleSetPadding(root, EdgeBottom, 10)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestPadding_center_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetAlignItems(root, AlignCenter)
	YGNodeStyleSetPadding(root, EdgeStart, 10)
	YGNodeStyleSetPadding(root, EdgeEnd, 20)
	YGNodeStyleSetPadding(root, EdgeBottom, 20)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 35, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 35, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestChild_with_padding_align_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyFlexEnd)
	YGNodeStyleSetAlignItems(root, AlignFlexEnd)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetPadding(rootChild0, EdgeLeft, 20)
	YGNodeStyleSetPadding(rootChild0, EdgeTop, 20)
	YGNodeStyleSetPadding(rootChild0, EdgeRight, 20)
	YGNodeStyleSetPadding(rootChild0, EdgeBottom, 20)
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))
}
