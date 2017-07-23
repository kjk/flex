package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert_default_values(t *testing.T) {
	root := YGNodeNew()

	assert.Equal(t, 0, YGNodeGetChildCount(root))
	var nilNode *YGNode
	assert.Equal(t, nilNode, YGNodeGetChild(root, 1))
	assert.Equal(t, nilNode, YGNodeGetChild(root, 0))

	assert.Equal(t, YGDirectionInherit, YGNodeStyleGetDirection(root))
	assert.Equal(t, YGFlexDirectionColumn, YGNodeStyleGetFlexDirection(root))
	assert.Equal(t, YGJustifyFlexStart, YGNodeStyleGetJustifyContent(root))
	assert.Equal(t, YGAlignFlexStart, YGNodeStyleGetAlignContent(root))
	assert.Equal(t, YGAlignStretch, YGNodeStyleGetAlignItems(root))
	assert.Equal(t, YGAlignAuto, YGNodeStyleGetAlignSelf(root))
	assert.Equal(t, YGPositionTypeRelative, YGNodeStyleGetPositionType(root))
	assert.Equal(t, YGWrapNoWrap, YGNodeStyleGetFlexWrap(root))
	assert.Equal(t, YGOverflowVisible, YGNodeStyleGetOverflow(root))
	assertFloatEqual(t, 0, YGNodeStyleGetFlexGrow(root))
	assertFloatEqual(t, 0, YGNodeStyleGetFlexShrink(root))
	assert.Equal(t, YGNodeStyleGetFlexBasis(root).unit, YGUnitAuto)

	assert.Equal(t, YGNodeStyleGetPosition(root, YGEdgeLeft).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, YGEdgeTop).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, YGEdgeRight).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, YGEdgeBottom).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, YGEdgeStart).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, YGEdgeEnd).unit, YGUnitUndefined)

	assert.Equal(t, YGNodeStyleGetMargin(root, YGEdgeLeft).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, YGEdgeTop).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, YGEdgeRight).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, YGEdgeBottom).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, YGEdgeStart).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, YGEdgeEnd).unit, YGUnitUndefined)

	assert.Equal(t, YGNodeStyleGetPadding(root, YGEdgeLeft).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, YGEdgeTop).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, YGEdgeRight).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, YGEdgeBottom).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, YGEdgeStart).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, YGEdgeEnd).unit, YGUnitUndefined)

	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, YGEdgeLeft)))
	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, YGEdgeTop)))
	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, YGEdgeRight)))
	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, YGEdgeBottom)))
	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, YGEdgeStart)))
	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, YGEdgeEnd)))

	assert.Equal(t, YGNodeStyleGetWidth(root).unit, YGUnitAuto)
	assert.Equal(t, YGNodeStyleGetHeight(root).unit, YGUnitAuto)
	assert.Equal(t, YGNodeStyleGetMinWidth(root).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetMinHeight(root).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetMaxWidth(root).unit, YGUnitUndefined)
	assert.Equal(t, YGNodeStyleGetMaxHeight(root).unit, YGUnitUndefined)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetRight(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetBottom(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, YGEdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, YGEdgeTop))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, YGEdgeRight))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, YGEdgeBottom))

	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, YGEdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, YGEdgeTop))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, YGEdgeRight))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, YGEdgeBottom))

	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, YGEdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, YGEdgeTop))
	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, YGEdgeRight))
	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, YGEdgeBottom))

	assert.True(t, YGFloatIsUndefined(YGNodeLayoutGetWidth(root)))
	assert.True(t, YGFloatIsUndefined(YGNodeLayoutGetHeight(root)))
	assert.Equal(t, YGDirectionInherit, YGNodeLayoutGetDirection(root))

	YGNodeFreeRecursive(root)
}

func TestAssert_webdefault_values(t *testing.T) {
	config := YGConfigNew()
	YGConfigSetUseWebDefaults(config, true)
	root := YGNodeNewWithConfig(config)

	assert.Equal(t, YGFlexDirectionRow, YGNodeStyleGetFlexDirection(root))
	assert.Equal(t, YGAlignStretch, YGNodeStyleGetAlignContent(root))
	assertFloatEqual(t, 1, YGNodeStyleGetFlexShrink(root))

	YGNodeFreeRecursive(root)
	YGConfigFree(config)
}

func TestAssert_webdefault_values_reset(t *testing.T) {
	config := YGConfigNew()
	YGConfigSetUseWebDefaults(config, true)
	root := YGNodeNewWithConfig(config)
	YGNodeReset(root)

	assert.Equal(t, YGFlexDirectionRow, YGNodeStyleGetFlexDirection(root))
	assert.Equal(t, YGAlignStretch, YGNodeStyleGetAlignContent(root))
	assertFloatEqual(t, 1, YGNodeStyleGetFlexShrink(root))

	YGNodeFreeRecursive(root)
	YGConfigFree(config)
}
