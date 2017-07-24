package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newHadOverflowTests() (*Config, *Node) {
	config := NewConfig()
	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 100)
	NodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetFlexWrap(root, WrapNoWrap)
	return config, root
}

func TestChildren_overflow_no_wrap_and_no_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child0, 80)
	NodeStyleSetHeight(child0, 40)
	NodeStyleSetMargin(child0, EdgeTop, 10)
	NodeStyleSetMargin(child0, EdgeBottom, 15)
	NodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1, 80)
	NodeStyleSetHeight(child1, 40)
	NodeStyleSetMargin(child1, EdgeBottom, 5)
	NodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, NodeLayoutGetHadOverflow(root))
}

func TestSpacing_overflow_no_wrap_and_no_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child0, 80)
	NodeStyleSetHeight(child0, 40)
	NodeStyleSetMargin(child0, EdgeTop, 10)
	NodeStyleSetMargin(child0, EdgeBottom, 10)
	NodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1, 80)
	NodeStyleSetHeight(child1, 40)
	NodeStyleSetMargin(child1, EdgeBottom, 5)
	NodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, NodeLayoutGetHadOverflow(root))
}

func TestNo_overflow_no_wrap_and_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child0, 80)
	NodeStyleSetHeight(child0, 40)
	NodeStyleSetMargin(child0, EdgeTop, 10)
	NodeStyleSetMargin(child0, EdgeBottom, 10)
	NodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1, 80)
	NodeStyleSetHeight(child1, 40)
	NodeStyleSetMargin(child1, EdgeBottom, 5)
	NodeStyleSetFlexShrink(child1, 1)
	NodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.False(t, NodeLayoutGetHadOverflow(root))
}

func TestHadOverflow_gets_reset_if_not_logger_valid(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child0, 80)
	NodeStyleSetHeight(child0, 40)
	NodeStyleSetMargin(child0, EdgeTop, 10)
	NodeStyleSetMargin(child0, EdgeBottom, 10)
	NodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1, 80)
	NodeStyleSetHeight(child1, 40)
	NodeStyleSetMargin(child1, EdgeBottom, 5)
	NodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, NodeLayoutGetHadOverflow(root))

	NodeStyleSetFlexShrink(child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.False(t, NodeLayoutGetHadOverflow(root))
}

func TestSpacing_overflow_in_nested_nodes(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child0, 80)
	NodeStyleSetHeight(child0, 40)
	NodeStyleSetMargin(child0, EdgeTop, 10)
	NodeStyleSetMargin(child0, EdgeBottom, 10)
	NodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1, 80)
	NodeStyleSetHeight(child1, 40)
	NodeInsertChild(root, child1, 1)
	child1_1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1_1, 80)
	NodeStyleSetHeight(child1_1, 40)
	NodeStyleSetMargin(child1_1, EdgeBottom, 5)
	NodeInsertChild(child1, child1_1, 0)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, NodeLayoutGetHadOverflow(root))
}
