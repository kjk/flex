package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert_default_values(t *testing.T) {
	root := NewNode()

	assert.Equal(t, 0, YGNodeGetChildCount(root))
	var nilNode *Node
	assert.Equal(t, nilNode, YGNodeGetChild(root, 1))
	assert.Equal(t, nilNode, YGNodeGetChild(root, 0))

	assert.Equal(t, DirectionInherit, YGNodeStyleGetDirection(root))
	assert.Equal(t, FlexDirectionColumn, YGNodeStyleGetFlexDirection(root))
	assert.Equal(t, JustifyFlexStart, YGNodeStyleGetJustifyContent(root))
	assert.Equal(t, AlignFlexStart, YGNodeStyleGetAlignContent(root))
	assert.Equal(t, AlignStretch, YGNodeStyleGetAlignItems(root))
	assert.Equal(t, AlignAuto, YGNodeStyleGetAlignSelf(root))
	assert.Equal(t, PositionTypeRelative, YGNodeStyleGetPositionType(root))
	assert.Equal(t, WrapNoWrap, YGNodeStyleGetFlexWrap(root))
	assert.Equal(t, OverflowVisible, YGNodeStyleGetOverflow(root))
	assertFloatEqual(t, 0, NodeStyleGetFlexGrow(root))
	assertFloatEqual(t, 0, NodeStyleGetFlexShrink(root))
	assert.Equal(t, YGNodeStyleGetFlexBasis(root).Unit, UnitAuto)

	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeLeft).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeTop).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeRight).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeBottom).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeStart).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeEnd).Unit, UnitUndefined)

	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeLeft).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeTop).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeRight).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeBottom).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeStart).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeEnd).Unit, UnitUndefined)

	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeLeft).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeTop).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeRight).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeBottom).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeStart).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeEnd).Unit, UnitUndefined)

	assert.True(t, FloatIsUndefined(YGNodeStyleGetBorder(root, EdgeLeft)))
	assert.True(t, FloatIsUndefined(YGNodeStyleGetBorder(root, EdgeTop)))
	assert.True(t, FloatIsUndefined(YGNodeStyleGetBorder(root, EdgeRight)))
	assert.True(t, FloatIsUndefined(YGNodeStyleGetBorder(root, EdgeBottom)))
	assert.True(t, FloatIsUndefined(YGNodeStyleGetBorder(root, EdgeStart)))
	assert.True(t, FloatIsUndefined(YGNodeStyleGetBorder(root, EdgeEnd)))

	assert.Equal(t, YGNodeStyleGetWidth(root).Unit, UnitAuto)
	assert.Equal(t, YGNodeStyleGetHeight(root).Unit, UnitAuto)
	assert.Equal(t, YGNodeStyleGetMinWidth(root).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMinHeight(root).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMaxWidth(root).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMaxHeight(root).Unit, UnitUndefined)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetRight(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetBottom(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, EdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, EdgeTop))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, EdgeRight))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, EdgeBottom))

	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, EdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, EdgeTop))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, EdgeRight))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, EdgeBottom))

	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, EdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, EdgeTop))
	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, EdgeRight))
	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, EdgeBottom))

	assert.True(t, FloatIsUndefined(YGNodeLayoutGetWidth(root)))
	assert.True(t, FloatIsUndefined(YGNodeLayoutGetHeight(root)))
	assert.Equal(t, DirectionInherit, YGNodeLayoutGetDirection(root))

}

func TestAssert_webdefault_values(t *testing.T) {
	config := NewConfig()
	config.UseWebDefaults = true
	root := NewNodeWithConfig(config)

	assert.Equal(t, FlexDirectionRow, YGNodeStyleGetFlexDirection(root))
	assert.Equal(t, AlignStretch, YGNodeStyleGetAlignContent(root))
	assertFloatEqual(t, 1, NodeStyleGetFlexShrink(root))

}

func TestAssert_webdefault_values_reset(t *testing.T) {
	config := NewConfig()
	config.UseWebDefaults = true
	root := NewNodeWithConfig(config)
	NodeReset(root)

	assert.Equal(t, FlexDirectionRow, YGNodeStyleGetFlexDirection(root))
	assert.Equal(t, AlignStretch, YGNodeStyleGetAlignContent(root))
	assertFloatEqual(t, 1, NodeStyleGetFlexShrink(root))

}
