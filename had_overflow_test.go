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
	YGNodeStyleSetFlexDirection(root, FlexDirectionColumn)
	YGNodeStyleSetFlexWrap(root, WrapNoWrap)
	return config, root
}

func TestChildren_overflow_no_wrap_and_no_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child0, 80)
	NodeStyleSetHeight(child0, 40)
	YGNodeStyleSetMargin(child0, EdgeTop, 10)
	YGNodeStyleSetMargin(child0, EdgeBottom, 15)
	YGNodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1, 80)
	NodeStyleSetHeight(child1, 40)
	YGNodeStyleSetMargin(child1, EdgeBottom, 5)
	YGNodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, YGNodeLayoutGetHadOverflow(root))
}

func TestSpacing_overflow_no_wrap_and_no_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child0, 80)
	NodeStyleSetHeight(child0, 40)
	YGNodeStyleSetMargin(child0, EdgeTop, 10)
	YGNodeStyleSetMargin(child0, EdgeBottom, 10)
	YGNodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1, 80)
	NodeStyleSetHeight(child1, 40)
	YGNodeStyleSetMargin(child1, EdgeBottom, 5)
	YGNodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, YGNodeLayoutGetHadOverflow(root))
}

func TestNo_overflow_no_wrap_and_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child0, 80)
	NodeStyleSetHeight(child0, 40)
	YGNodeStyleSetMargin(child0, EdgeTop, 10)
	YGNodeStyleSetMargin(child0, EdgeBottom, 10)
	YGNodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1, 80)
	NodeStyleSetHeight(child1, 40)
	YGNodeStyleSetMargin(child1, EdgeBottom, 5)
	YGNodeStyleSetFlexShrink(child1, 1)
	YGNodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.False(t, YGNodeLayoutGetHadOverflow(root))
}

func TestHadOverflow_gets_reset_if_not_logger_valid(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child0, 80)
	NodeStyleSetHeight(child0, 40)
	YGNodeStyleSetMargin(child0, EdgeTop, 10)
	YGNodeStyleSetMargin(child0, EdgeBottom, 10)
	YGNodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1, 80)
	NodeStyleSetHeight(child1, 40)
	YGNodeStyleSetMargin(child1, EdgeBottom, 5)
	YGNodeInsertChild(root, child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, YGNodeLayoutGetHadOverflow(root))

	YGNodeStyleSetFlexShrink(child1, 1)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.False(t, YGNodeLayoutGetHadOverflow(root))
}

func TestSpacing_overflow_in_nested_nodes(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child0, 80)
	NodeStyleSetHeight(child0, 40)
	YGNodeStyleSetMargin(child0, EdgeTop, 10)
	YGNodeStyleSetMargin(child0, EdgeBottom, 10)
	YGNodeInsertChild(root, child0, 0)
	child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1, 80)
	NodeStyleSetHeight(child1, 40)
	YGNodeInsertChild(root, child1, 1)
	child1_1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(child1_1, 80)
	NodeStyleSetHeight(child1_1, 40)
	YGNodeStyleSetMargin(child1_1, EdgeBottom, 5)
	YGNodeInsertChild(child1, child1_1, 0)

	NodeCalculateLayout(root, 200, 100, DirectionLTR)

	assert.True(t, YGNodeLayoutGetHadOverflow(root))
}
