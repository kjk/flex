package flex

import (
	"testing"
)

func TestAlignContentFlexStart(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 130)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeStyleSetHeight(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeStyleSetHeight(root_child3, 10)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeStyleSetHeight(root_child4, 10)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 130, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 130, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_flex_start_without_height_on_children(t *testing.T) {

	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeStyleSetHeight(root_child3, 10)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_flex_start_with_flex(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 120)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0, 0)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetFlexBasisPercent(root_child1, 0)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child3, 1)
	YGNodeStyleSetFlexShrink(root_child3, 1)
	YGNodeStyleSetFlexBasisPercent(root_child3, 0)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 120, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 120, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignContent(root, YGAlignFlexEnd)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeStyleSetHeight(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeStyleSetHeight(root_child3, 10)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeStyleSetHeight(root_child4, 10)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_spacebetween(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignSpaceBetween)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 130)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeStyleSetHeight(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeStyleSetHeight(root_child3, 10)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeStyleSetHeight(root_child4, 10)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 130, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 130, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_spacearound(t *testing.T) {

	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignSpaceAround)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 140)
	YGNodeStyleSetHeight(root, 120)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeStyleSetHeight(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeStyleSetHeight(root_child3, 10)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeStyleSetHeight(root_child4, 10)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 140, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 15, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 15, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 55, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 55, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 95, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 140, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 15, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 15, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 55, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 55, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 95, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_row(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_row_with_children(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child0, 1)
	YGNodeStyleSetFlexShrink(root_child0_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0_child0, 0)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_row_with_flex(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetFlexShrink(root_child1, 1)
	YGNodeStyleSetFlexBasisPercent(root_child1, 0)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child3, 1)
	YGNodeStyleSetFlexShrink(root_child3, 1)
	YGNodeStyleSetFlexBasisPercent(root_child3, 0)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_row_with_flex_no_shrink(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetFlexShrink(root_child1, 1)
	YGNodeStyleSetFlexBasisPercent(root_child1, 0)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child3, 1)
	YGNodeStyleSetFlexBasisPercent(root_child3, 0)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_row_with_margin(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root_child1, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(root_child1, YGEdgeTop, 10)
	YGNodeStyleSetMargin(root_child1, YGEdgeRight, 10)
	YGNodeStyleSetMargin(root_child1, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root_child3, YGEdgeLeft, 10)
	YGNodeStyleSetMargin(root_child3, YGEdgeTop, 10)
	YGNodeStyleSetMargin(root_child3, YGEdgeRight, 10)
	YGNodeStyleSetMargin(root_child3, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_row_with_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPadding(root_child1, YGEdgeLeft, 10)
	YGNodeStyleSetPadding(root_child1, YGEdgeTop, 10)
	YGNodeStyleSetPadding(root_child1, YGEdgeRight, 10)
	YGNodeStyleSetPadding(root_child1, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPadding(root_child3, YGEdgeLeft, 10)
	YGNodeStyleSetPadding(root_child3, YGEdgeTop, 10)
	YGNodeStyleSetPadding(root_child3, YGEdgeRight, 10)
	YGNodeStyleSetPadding(root_child3, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_row_with_single_row(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child1))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_row_with_fixed_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetHeight(root_child1, 60)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_row_with_max_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetMaxHeight(root_child1, 20)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_row_with_min_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 150)
	YGNodeStyleSetHeight(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 50)
	YGNodeStyleSetMinHeight(root_child1, 80)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_column(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)
	YGNodeStyleSetFlexWrap(root, YGWrapWrap)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 150)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child0, 50)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child0_child0, 1)
	YGNodeStyleSetFlexShrink(root_child0_child0, 1)
	YGNodeStyleSetFlexBasisPercent(root_child0_child0, 0)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexGrow(root_child1, 1)
	YGNodeStyleSetFlexShrink(root_child1, 1)
	YGNodeStyleSetFlexBasisPercent(root_child1, 0)
	YGNodeStyleSetHeight(root_child1, 50)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child2, 50)
	YGNodeInsertChild(root, root_child2, 2)

	root_child3 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child3, 50)
	YGNodeInsertChild(root, root_child3, 3)

	root_child4 := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeight(root_child4, 50)
	YGNodeInsertChild(root, root_child4, 4)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child2))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child3))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child3))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child3))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child4))

	YGNodeFreeRecursive(root)


}

func TestAlign_content_stretch_is_not_overriding_align_items(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignContent(root, YGAlignStretch)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionRow)
	YGNodeStyleSetAlignContent(root_child0, YGAlignStretch)
	YGNodeStyleSetAlignItems(root_child0, YGAlignCenter)
	YGNodeStyleSetWidth(root_child0, 100)
	YGNodeStyleSetHeight(root_child0, 100)
	YGNodeInsertChild(root, root_child0, 0)

	root_child0_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetAlignContent(root_child0_child0, YGAlignStretch)
	YGNodeStyleSetWidth(root_child0_child0, 10)
	YGNodeStyleSetHeight(root_child0_child0, 10)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child0_child0))
	assertFloatEqual(t, 45, YGNodeLayoutGetTop(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0_child0))

	YGNodeFreeRecursive(root)


}
