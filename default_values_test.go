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
	assertFloatEqual(t, 0, root.StyleGetFlexGrow())
	assertFloatEqual(t, 0, root.StyleGetFlexShrink())
	assert.Equal(t, root.Style.FlexBasis.Unit, UnitAuto)

	assert.Equal(t, root.StyleGetPosition(EdgeLeft).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetPosition(EdgeTop).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetPosition(EdgeRight).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetPosition(EdgeBottom).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetPosition(EdgeStart).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetPosition(EdgeEnd).Unit, UnitUndefined)

	assert.Equal(t, root.StyleGetMargin(EdgeLeft).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetMargin(EdgeTop).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetMargin(EdgeRight).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetMargin(EdgeBottom).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetMargin(EdgeStart).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetMargin(EdgeEnd).Unit, UnitUndefined)

	assert.Equal(t, root.StyleGetPadding(EdgeLeft).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetPadding(EdgeTop).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetPadding(EdgeRight).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetPadding(EdgeBottom).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetPadding(EdgeStart).Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetPadding(EdgeEnd).Unit, UnitUndefined)

	assert.True(t, FloatIsUndefined(root.StyleGetBorder(EdgeLeft)))
	assert.True(t, FloatIsUndefined(root.StyleGetBorder(EdgeTop)))
	assert.True(t, FloatIsUndefined(root.StyleGetBorder(EdgeRight)))
	assert.True(t, FloatIsUndefined(root.StyleGetBorder(EdgeBottom)))
	assert.True(t, FloatIsUndefined(root.StyleGetBorder(EdgeStart)))
	assert.True(t, FloatIsUndefined(root.StyleGetBorder(EdgeEnd)))

	assert.Equal(t, root.StyleGetWidth().Unit, UnitAuto)
	assert.Equal(t, root.StyleGetHeight().Unit, UnitAuto)
	assert.Equal(t, root.StyleGetMinWidth().Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetMinHeight().Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetMaxWidth().Unit, UnitUndefined)
	assert.Equal(t, root.StyleGetMaxHeight().Unit, UnitUndefined)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 0, root.LayoutGetRight())
	assertFloatEqual(t, 0, root.LayoutGetBottom())

	assertFloatEqual(t, 0, root.LayoutGetMargin(EdgeLeft))
	assertFloatEqual(t, 0, root.LayoutGetMargin(EdgeTop))
	assertFloatEqual(t, 0, root.LayoutGetMargin(EdgeRight))
	assertFloatEqual(t, 0, root.LayoutGetMargin(EdgeBottom))

	assertFloatEqual(t, 0, NodeLayoutGetPadding(root, EdgeLeft))
	assertFloatEqual(t, 0, NodeLayoutGetPadding(root, EdgeTop))
	assertFloatEqual(t, 0, NodeLayoutGetPadding(root, EdgeRight))
	assertFloatEqual(t, 0, NodeLayoutGetPadding(root, EdgeBottom))

	assertFloatEqual(t, 0, NodeLayoutGetBorder(root, EdgeLeft))
	assertFloatEqual(t, 0, NodeLayoutGetBorder(root, EdgeTop))
	assertFloatEqual(t, 0, NodeLayoutGetBorder(root, EdgeRight))
	assertFloatEqual(t, 0, NodeLayoutGetBorder(root, EdgeBottom))

	assert.True(t, FloatIsUndefined(root.LayoutGetWidth()))
	assert.True(t, FloatIsUndefined(root.LayoutGetHeight()))
	assert.Equal(t, DirectionInherit, root.Layout.Direction)

}

func TestAssert_webdefault_values(t *testing.T) {
	config := NewConfig()
	config.UseWebDefaults = true
	root := NewNodeWithConfig(config)

	assert.Equal(t, FlexDirectionRow, root.Style.FlexDirection)
	assert.Equal(t, AlignStretch, root.Style.AlignContent)
	assertFloatEqual(t, 1, root.StyleGetFlexShrink())

}

func TestAssert_webdefault_values_reset(t *testing.T) {
	config := NewConfig()
	config.UseWebDefaults = true
	root := NewNodeWithConfig(config)
	root.Reset()

	assert.Equal(t, FlexDirectionRow, root.Style.FlexDirection)
	assert.Equal(t, AlignStretch, root.Style.AlignContent)
	assertFloatEqual(t, 1, root.StyleGetFlexShrink())

}
