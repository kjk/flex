package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert_default_values(t *testing.T) {
	root := NewNode()

	assert.Equal(t, 0, len(root.Children))
	var nilNode *Node
	assert.Equal(t, nilNode, root.GetChild(1))
	assert.Equal(t, nilNode, root.GetChild(0))

	assert.Equal(t, DirectionInherit, root.Style.Direction)
	assert.Equal(t, FlexDirectionColumn, root.Style.FlexDirection)
	assert.Equal(t, JustifyFlexStart, root.Style.JustifyContent)
	assert.Equal(t, AlignFlexStart, root.Style.AlignContent)
	assert.Equal(t, AlignStretch, root.Style.AlignItems)
	assert.Equal(t, AlignAuto, root.Style.AlignSelf)
	assert.Equal(t, PositionTypeRelative, root.Style.PositionType)
	assert.Equal(t, WrapNoWrap, root.Style.FlexWrap)
	assert.Equal(t, OverflowVisible, root.Style.Overflow)
	assertFloatEqual(t, 0, NodeStyleGetFlexGrow(root))
	assertFloatEqual(t, 0, NodeStyleGetFlexShrink(root))
	assert.Equal(t, root.Style.FlexBasis.Unit, UnitAuto)

	assert.Equal(t, NodeStyleGetPosition(root, EdgeLeft).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetPosition(root, EdgeTop).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetPosition(root, EdgeRight).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetPosition(root, EdgeBottom).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetPosition(root, EdgeStart).Unit, UnitUndefined)
	assert.Equal(t, NodeStyleGetPosition(root, EdgeEnd).Unit, UnitUndefined)

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
	assert.Equal(t, DirectionInherit, root.Layout.Direction)

}

func TestAssert_webdefault_values(t *testing.T) {
	config := NewConfig()
	config.UseWebDefaults = true
	root := NewNodeWithConfig(config)

	assert.Equal(t, FlexDirectionRow, root.Style.FlexDirection)
	assert.Equal(t, AlignStretch, root.Style.AlignContent)
	assertFloatEqual(t, 1, NodeStyleGetFlexShrink(root))

}

func TestAssert_webdefault_values_reset(t *testing.T) {
	config := NewConfig()
	config.UseWebDefaults = true
	root := NewNodeWithConfig(config)
	NodeReset(root)

	assert.Equal(t, FlexDirectionRow, root.Style.FlexDirection)
	assert.Equal(t, AlignStretch, root.Style.AlignContent)
	assertFloatEqual(t, 1, NodeStyleGetFlexShrink(root))

}
