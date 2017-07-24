package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert_default_values(t *testing.T) {
	root := NewNode()

	assert.Equal(t, 0, NodeGetChildCount(root))
	var nilNode *Node
	assert.Equal(t, nilNode, NodeGetChild(root, 1))
	assert.Equal(t, nilNode, NodeGetChild(root, 0))

	assert.Equal(t, DirectionInherit, NodeStyleGetDirection(root))
	assert.Equal(t, FlexDirectionColumn, NodeStyleGetFlexDirection(root))
	assert.Equal(t, JustifyFlexStart, NodeStyleGetJustifyContent(root))
	assert.Equal(t, AlignFlexStart, NodeStyleGetAlignContent(root))
	assert.Equal(t, AlignStretch, NodeStyleGetAlignItems(root))
	assert.Equal(t, AlignAuto, NodeStyleGetAlignSelf(root))
	assert.Equal(t, PositionTypeRelative, NodeStyleGetPositionType(root))
	assert.Equal(t, WrapNoWrap, NodeStyleGetFlexWrap(root))
	assert.Equal(t, OverflowVisible, NodeStyleGetOverflow(root))
	assertFloatEqual(t, 0, NodeStyleGetFlexGrow(root))
	assertFloatEqual(t, 0, NodeStyleGetFlexShrink(root))
	assert.Equal(t, NodeStyleGetFlexBasis(root).Unit, UnitAuto)

	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeLeft).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeTop).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeRight).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeBottom).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeStart).Unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeEnd).Unit, UnitUndefined)

	assert.Equal(t, NodeStyleGetMargin(root, EdgeLeft).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetMargin(root, EdgeTop).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetMargin(root, EdgeRight).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetMargin(root, EdgeBottom).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetMargin(root, EdgeStart).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetMargin(root, EdgeEnd).Unit, UnitUndefined)

	assert.Equal(t, NodeStyleGetPadding(root, EdgeLeft).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetPadding(root, EdgeTop).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetPadding(root, EdgeRight).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetPadding(root, EdgeBottom).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetPadding(root, EdgeStart).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetPadding(root, EdgeEnd).Unit, UnitUndefined)

	assert.True(t, FloatIsUndefined(NodeStyleGetBorder(root, EdgeLeft)))
	assert.True(t, FloatIsUndefined(NodeStyleGetBorder(root, EdgeTop)))
	assert.True(t, FloatIsUndefined(NodeStyleGetBorder(root, EdgeRight)))
	assert.True(t, FloatIsUndefined(NodeStyleGetBorder(root, EdgeBottom)))
	assert.True(t, FloatIsUndefined(NodeStyleGetBorder(root, EdgeStart)))
	assert.True(t, FloatIsUndefined(NodeStyleGetBorder(root, EdgeEnd)))

	assert.Equal(t, NodeStyleGetWidth(root).Unit, UnitAuto)
	assert.Equal(t, NodeStyleGetHeight(root).Unit, UnitAuto)
	assert.Equal(t, NodeStyleGetMinWidth(root).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetMinHeight(root).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetMaxWidth(root).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetMaxHeight(root).Unit, UnitUndefined)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 0, NodeLayoutGetRight(root))
	assertFloatEqual(t, 0, NodeLayoutGetBottom(root))

	assertFloatEqual(t, 0, NodeLayoutGetMargin(root, EdgeLeft))
	assertFloatEqual(t, 0, NodeLayoutGetMargin(root, EdgeTop))
	assertFloatEqual(t, 0, NodeLayoutGetMargin(root, EdgeRight))
	assertFloatEqual(t, 0, NodeLayoutGetMargin(root, EdgeBottom))

	assertFloatEqual(t, 0, NodeLayoutGetPadding(root, EdgeLeft))
	assertFloatEqual(t, 0, NodeLayoutGetPadding(root, EdgeTop))
	assertFloatEqual(t, 0, NodeLayoutGetPadding(root, EdgeRight))
	assertFloatEqual(t, 0, NodeLayoutGetPadding(root, EdgeBottom))

	assertFloatEqual(t, 0, NodeLayoutGetBorder(root, EdgeLeft))
	assertFloatEqual(t, 0, NodeLayoutGetBorder(root, EdgeTop))
	assertFloatEqual(t, 0, NodeLayoutGetBorder(root, EdgeRight))
	assertFloatEqual(t, 0, NodeLayoutGetBorder(root, EdgeBottom))

	assert.True(t, FloatIsUndefined(NodeLayoutGetWidth(root)))
	assert.True(t, FloatIsUndefined(NodeLayoutGetHeight(root)))
	assert.Equal(t, DirectionInherit, NodeLayoutGetDirection(root))

}

func TestAssert_webdefault_values(t *testing.T) {
	config := NewConfig()
	config.UseWebDefaults = true
	root := NewNodeWithConfig(config)

	assert.Equal(t, FlexDirectionRow, NodeStyleGetFlexDirection(root))
	assert.Equal(t, AlignStretch, NodeStyleGetAlignContent(root))
	assertFloatEqual(t, 1, NodeStyleGetFlexShrink(root))

}

func TestAssert_webdefault_values_reset(t *testing.T) {
	config := NewConfig()
	config.UseWebDefaults = true
	root := NewNodeWithConfig(config)
	NodeReset(root)

	assert.Equal(t, FlexDirectionRow, NodeStyleGetFlexDirection(root))
	assert.Equal(t, AlignStretch, NodeStyleGetAlignContent(root))
	assertFloatEqual(t, 1, NodeStyleGetFlexShrink(root))

}
