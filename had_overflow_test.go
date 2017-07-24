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
	root.StyleSetFlexDirection(FlexDirectionColumn)
	root.StyleSetFlexWrap(WrapNoWrap)
	return config, root
}

func TestChildren_overflow_no_wrap_and_no_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	child0.StyleSetWidth(80)
	child0.StyleSetHeight(40)
	child0.StyleSetMargin(EdgeTop, 10)
	child0.StyleSetMargin(EdgeBottom, 15)
	root.InsertChild(child0, 0)
	child1 := NewNodeWithConfig(config)
	child1.StyleSetWidth(80)
	child1.StyleSetHeight(40)
	child1.StyleSetMargin(EdgeBottom, 5)
	root.InsertChild(child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, root.Layout.HadOverflow)
}

func TestSpacing_overflow_no_wrap_and_no_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	child0.StyleSetWidth(80)
	child0.StyleSetHeight(40)
	child0.StyleSetMargin(EdgeTop, 10)
	child0.StyleSetMargin(EdgeBottom, 10)
	root.InsertChild(child0, 0)
	child1 := NewNodeWithConfig(config)
	child1.StyleSetWidth(80)
	child1.StyleSetHeight(40)
	child1.StyleSetMargin(EdgeBottom, 5)
	root.InsertChild(child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, root.Layout.HadOverflow)
}

func TestNo_overflow_no_wrap_and_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	child0.StyleSetWidth(80)
	child0.StyleSetHeight(40)
	child0.StyleSetMargin(EdgeTop, 10)
	child0.StyleSetMargin(EdgeBottom, 10)
	root.InsertChild(child0, 0)
	child1 := NewNodeWithConfig(config)
	child1.StyleSetWidth(80)
	child1.StyleSetHeight(40)
	child1.StyleSetMargin(EdgeBottom, 5)
	child1.StyleSetFlexShrink(1)
	root.InsertChild(child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.False(t, root.Layout.HadOverflow)
}

func TestHadOverflow_gets_reset_if_not_logger_valid(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	child0.StyleSetWidth(80)
	child0.StyleSetHeight(40)
	child0.StyleSetMargin(EdgeTop, 10)
	child0.StyleSetMargin(EdgeBottom, 10)
	root.InsertChild(child0, 0)
	child1 := NewNodeWithConfig(config)
	child1.StyleSetWidth(80)
	child1.StyleSetHeight(40)
	child1.StyleSetMargin(EdgeBottom, 5)
	root.InsertChild(child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, root.Layout.HadOverflow)

	child1.StyleSetFlexShrink(1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.False(t, root.Layout.HadOverflow)
}

func TestSpacing_overflow_in_nested_nodes(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	child0.StyleSetWidth(80)
	child0.StyleSetHeight(40)
	child0.StyleSetMargin(EdgeTop, 10)
	child0.StyleSetMargin(EdgeBottom, 10)
	root.InsertChild(child0, 0)
	child1 := NewNodeWithConfig(config)
	child1.StyleSetWidth(80)
	child1.StyleSetHeight(40)
	root.InsertChild(child1, 1)
	child1_1 := NewNodeWithConfig(config)
	child1_1.StyleSetWidth(80)
	child1_1.StyleSetHeight(40)
	child1_1.StyleSetMargin(EdgeBottom, 5)
	child1.InsertChild(child1_1, 0)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, root.Layout.HadOverflow)
}
