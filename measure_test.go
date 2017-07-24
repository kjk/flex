package flex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func _measure3(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	measureCount, ok := node.Context.(*int)
	if ok {
		(*measureCount)++
	}

	return Size{Width: 10, Height: 10}
}

func _simulate_wrapping_text(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	if widthMode == MeasureModeUndefined || width >= 68 {
		return Size{Width: 68, Height: 16}
	}

	return Size{Width: 50, Height: 32}
}

func _measure_assert_negative(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	if width < 0 {
		panic(fmt.Sprintf("width is %.2f and should be >= 0", width))
	}
	if height < 0 {
		panic(fmt.Sprintf("height is %.2f should be >= 0, height", height))
	}
	// EXPECT_GE(width, 0);
	//EXPECT_GE(height, 0);

	return Size{
		Width: 0, Height: 0,
	}
}

func TestDont_measure_single_grow_shrink_child(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexShrink(1)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
}

func TestMeasure_absolute_child_with_no_constraints(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	NodeInsertChild(root, rootChild0, 0)

	measureCount := 0

	rootChild0Child0 := NewNode()
	rootChild0Child0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0Child0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0Child0, _measure3)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, measureCount)
}

func TestDont_measure_when_min_equals_max(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	rootChild0.StyleSetMinWidth(10)
	rootChild0.StyleSetMaxWidth(10)
	rootChild0.StyleSetMinHeight(10)
	rootChild0.StyleSetMaxHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestDont_measure_when_min_equals_max_percentages(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	rootChild0.StyleSetMinWidthPercent(10)
	rootChild0.StyleSetMaxWidthPercent(10)
	rootChild0.StyleSetMinHeightPercent(10)
	rootChild0.StyleSetMaxHeightPercent(10)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestDont_measure_when_min_equals_max_mixed_width_percent(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	rootChild0.StyleSetMinWidthPercent(10)
	rootChild0.StyleSetMaxWidthPercent(10)
	rootChild0.StyleSetMinHeight(10)
	rootChild0.StyleSetMaxHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestDont_measure_when_min_equals_max_mixed_height_percent(t *testing.T) {
	root := NewNode()
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	rootChild0.StyleSetMinWidth(10)
	rootChild0.StyleSetMaxWidth(10)
	rootChild0.StyleSetMinHeightPercent(10)
	rootChild0.StyleSetMaxHeightPercent(10)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestMeasure_enough_size_should_be_in_single_line(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(100)

	rootChild0 := NewNode()
	rootChild0.StyleSetAlignSelf(AlignFlexStart)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)

	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 68, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 16, rootChild0.LayoutGetHeight())
}

func TestMeasure_not_enough_size_should_wrap(t *testing.T) {
	root := NewNode()
	root.StyleSetWidth(55)

	rootChild0 := NewNode()
	rootChild0.StyleSetAlignSelf(AlignFlexStart)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)

	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 32, rootChild0.LayoutGetHeight())
}

func TestMeasure_zero_space_should_grow(t *testing.T) {
	root := NewNode()
	root.StyleSetHeight(200)
	root.StyleSetFlexDirection(FlexDirectionColumn)
	root.StyleSetFlexGrow(0)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.StyleSetFlexDirection(FlexDirectionColumn)
	rootChild0.StyleSetPadding(EdgeAll, 100)
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)

	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, 282, Undefined, DirectionLTR)

	assertFloatEqual(t, 282, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
}

func TestMeasure_flex_direction_row_and_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetPadding(EdgeLeft, 25)
	root.StyleSetPadding(EdgeTop, 25)
	root.StyleSetPadding(EdgeRight, 25)
	root.StyleSetPadding(EdgeBottom, 25)
	root.StyleSetWidth(50)
	root.StyleSetHeight(50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(5)
	rootChild1.StyleSetHeight(5)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 25, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 75, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 25, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 5, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 5, rootChild1.LayoutGetHeight())
}

func TestMeasure_flex_direction_column_and_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetMargin(EdgeTop, 20)
	root.StyleSetPadding(EdgeAll, 25)
	root.StyleSetWidth(50)
	root.StyleSetHeight(50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(5)
	rootChild1.StyleSetHeight(5)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 25, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 0, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 32, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 57, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 5, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 5, rootChild1.LayoutGetHeight())
}

func TestMeasure_flex_direction_row_no_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetMargin(EdgeTop, 20)
	root.StyleSetWidth(50)
	root.StyleSetHeight(50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(5)
	rootChild1.StyleSetHeight(5)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 5, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 5, rootChild1.LayoutGetHeight())
}

func TestMeasure_flex_direction_row_no_padding_align_items_flexstart(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetMargin(EdgeTop, 20)
	root.StyleSetWidth(50)
	root.StyleSetHeight(50)
	root.StyleSetAlignItems(AlignFlexStart)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(5)
	rootChild1.StyleSetHeight(5)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 32, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 5, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 5, rootChild1.LayoutGetHeight())
}

func TestMeasure_with_fixed_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetMargin(EdgeTop, 20)
	root.StyleSetPadding(EdgeAll, 25)
	root.StyleSetWidth(50)
	root.StyleSetHeight(50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	rootChild0.StyleSetWidth(10)
	rootChild0.StyleSetHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(5)
	rootChild1.StyleSetHeight(5)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 25, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 35, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 5, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 5, rootChild1.LayoutGetHeight())
}

func TestMeasure_with_flex_shrink(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetMargin(EdgeTop, 20)
	root.StyleSetPadding(EdgeAll, 25)
	root.StyleSetWidth(50)
	root.StyleSetHeight(50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	rootChild0.StyleSetFlexShrink(1)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(5)
	rootChild1.StyleSetHeight(5)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 25, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 0, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 25, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 5, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 5, rootChild1.LayoutGetHeight())
}

func TestMeasure_no_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetMargin(EdgeTop, 20)
	root.StyleSetWidth(50)
	root.StyleSetHeight(50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	rootChild0.StyleSetFlexShrink(1)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(5)
	rootChild1.StyleSetHeight(5)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 32, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 32, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 5, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 5, rootChild1.LayoutGetHeight())
}

/*
#if GTEST_HAS_DEATH_TEST
TEST(YogaDeathTest, cannot_add_child_to_node_with_measure_func) {
  root := YGNodeNew();
  YGNodeSetMeasureFunc(root, _measure3);

  rootChild0 := YGNodeNew();
  ASSERT_DEATH(YGNodeInsertChild(root, rootChild0, 0), "Cannot add child.*");
  YGNodeFree(rootChild0);
  ;
}

TEST(YogaDeathTest, cannot_add_nonnull_measure_func_to_non_leaf_node) {
  root := YGNodeNew();
  rootChild0 := YGNodeNew();
  YGNodeInsertChild(root, rootChild0, 0);

  ASSERT_DEATH(YGNodeSetMeasureFunc(root, _measure3), "Cannot set measure function.*");
  ;
}
#endif
*/

func TestCan_nullify_measure_func_on_any_node(t *testing.T) {
	root := NewNode()
	NodeInsertChild(root, NewNode(), 0)

	NodeSetMeasureFunc(root, nil)
	assert.True(t, root.Measure == nil)
}

func TestCant_call_negative_measure(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionColumn)
	root.StyleSetWidth(50)
	root.StyleSetHeight(10)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _measure_assert_negative)
	rootChild0.StyleSetMargin(EdgeTop, 20)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
}

func TestCant_call_negative_measure_horizontal(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(10)
	root.StyleSetHeight(20)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _measure_assert_negative)
	rootChild0.StyleSetMargin(EdgeStart, 20)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
}
