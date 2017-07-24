package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newHadOverflowTests() (*Config, *Node) {
	config := NewConfig()
	root := NewNodeWithConfig(config)
	root.StyleSetWidth(200)
	root.StyleSetHeight(100)
	NodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetFlexWrap(root, WrapNoWrap)
	return config, root
}

func TestChildren_overflow_no_wrap_and_no_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	child0.StyleSetWidth(80)
	child0.StyleSetHeight(40)
	NodeStyleSetMargin(child0, EdgeTop, 10)
	NodeStyleSetMargin(child0, EdgeBottom, 15)
	NodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	child1.StyleSetWidth(80)
	child1.StyleSetHeight(40)
	NodeStyleSetMargin(child1, EdgeBottom, 5)
	NodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, root.Layout.HadOverflow)
}

func TestSpacing_overflow_no_wrap_and_no_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	child0.StyleSetWidth(80)
	child0.StyleSetHeight(40)
	NodeStyleSetMargin(child0, EdgeTop, 10)
	NodeStyleSetMargin(child0, EdgeBottom, 10)
	NodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	child1.StyleSetWidth(80)
	child1.StyleSetHeight(40)
	NodeStyleSetMargin(child1, EdgeBottom, 5)
	NodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, root.Layout.HadOverflow)
}

func TestNo_overflow_no_wrap_and_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	child0.StyleSetWidth(80)
	child0.StyleSetHeight(40)
	NodeStyleSetMargin(child0, EdgeTop, 10)
	NodeStyleSetMargin(child0, EdgeBottom, 10)
	NodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	child1.StyleSetWidth(80)
	child1.StyleSetHeight(40)
	NodeStyleSetMargin(child1, EdgeBottom, 5)
	NodeStyleSetFlexShrink(child1, 1)
	NodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.False(t, root.Layout.HadOverflow)
}

func TestHadOverflow_gets_reset_if_not_logger_valid(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	child0.StyleSetWidth(80)
	child0.StyleSetHeight(40)
	NodeStyleSetMargin(child0, EdgeTop, 10)
	NodeStyleSetMargin(child0, EdgeBottom, 10)
	NodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	child1.StyleSetWidth(80)
	child1.StyleSetHeight(40)
	NodeStyleSetMargin(child1, EdgeBottom, 5)
	NodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, root.Layout.HadOverflow)

	NodeStyleSetFlexShrink(child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.False(t, root.Layout.HadOverflow)
}

func TestSpacing_overflow_in_nested_nodes(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	child0.StyleSetWidth(80)
	child0.StyleSetHeight(40)
	NodeStyleSetMargin(child0, EdgeTop, 10)
	NodeStyleSetMargin(child0, EdgeBottom, 10)
	NodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	child1.StyleSetWidth(80)
	child1.StyleSetHeight(40)
	NodeInsertChild(root, child1, 1)
	child1_1 := NewNodeWithConfig(config)
	child1_1.StyleSetWidth(80)
	child1_1.StyleSetHeight(40)
	NodeStyleSetMargin(child1_1, EdgeBottom, 5)
	NodeInsertChild(child1, child1_1, 0)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, root.Layout.HadOverflow)
}
