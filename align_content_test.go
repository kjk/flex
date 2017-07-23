package flex

import (
	"testing"
)

func TestAlignContentFlexStart(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 130)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 10)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeStyleSetHeight(rootChild4, 10)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 130, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 130, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_flex_start_without_height_on_children(t *testing.T) {

	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 10)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_flex_start_with_flex(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 120)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 0)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild3, 1)
	YGNodeStyleSetFlexShrink(rootChild3, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild3, 0)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 120, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 120, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetAlignContent(root, AlignFlexEnd)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 10)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeStyleSetHeight(rootChild4, 10)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_spacebetween(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignSpaceBetween)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 130)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 10)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeStyleSetHeight(rootChild4, 10)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 130, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 130, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild4))

}

func TestAlign_content_spacearound(t *testing.T) {

	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignSpaceAround)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 140)
	NodeStyleSetHeight(root, 120)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 10)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeStyleSetHeight(rootChild4, 10)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 140, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 15, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 15, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 55, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 55, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 95, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 140, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 15, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 15, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 55, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 55, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 95, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_children(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0Child0, 1)
	YGNodeStyleSetFlexShrink(rootChild0Child0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0Child0, 0)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_flex(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeStyleSetFlexShrink(rootChild1, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild3, 1)
	YGNodeStyleSetFlexShrink(rootChild3, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild3, 0)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_flex_no_shrink(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeStyleSetFlexShrink(rootChild1, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild3, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild3, 0)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_margin(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetMargin(rootChild1, EdgeLeft, 10)
	YGNodeStyleSetMargin(rootChild1, EdgeTop, 10)
	YGNodeStyleSetMargin(rootChild1, EdgeRight, 10)
	YGNodeStyleSetMargin(rootChild1, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetMargin(rootChild3, EdgeLeft, 10)
	YGNodeStyleSetMargin(rootChild3, EdgeTop, 10)
	YGNodeStyleSetMargin(rootChild3, EdgeRight, 10)
	YGNodeStyleSetMargin(rootChild3, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetPadding(rootChild1, EdgeLeft, 10)
	YGNodeStyleSetPadding(rootChild1, EdgeTop, 10)
	YGNodeStyleSetPadding(rootChild1, EdgeRight, 10)
	YGNodeStyleSetPadding(rootChild1, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	YGNodeStyleSetPadding(rootChild3, EdgeLeft, 10)
	YGNodeStyleSetPadding(rootChild3, EdgeTop, 10)
	YGNodeStyleSetPadding(rootChild3, EdgeRight, 10)
	YGNodeStyleSetPadding(rootChild3, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_single_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))
}

func TestAlign_content_stretch_row_with_fixed_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 60)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_max_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetMaxHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_min_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeStyleSetMinHeight(rootChild1, 80)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetAlignContent(root, AlignStretch)
	YGNodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 150)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild0Child0, 1)
	YGNodeStyleSetFlexShrink(rootChild0Child0, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild0Child0, 0)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexGrow(rootChild1, 1)
	YGNodeStyleSetFlexShrink(rootChild1, 1)
	YGNodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild2, 50)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild3, 50)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_is_not_overriding_align_items(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetAlignContent(root, AlignStretch)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	YGNodeStyleSetAlignContent(rootChild0, AlignStretch)
	YGNodeStyleSetAlignItems(rootChild0, AlignCenter)
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeStyleSetAlignContent(rootChild0Child0, AlignStretch)
	NodeStyleSetWidth(rootChild0Child0, 10)
	NodeStyleSetHeight(rootChild0Child0, 10)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0Child0))
}
