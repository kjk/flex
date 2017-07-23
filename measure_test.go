package flex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func _measure3(node *YGNode, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	measureCount, ok := YGNodeGetContext(node).(*int)
	if ok {
		(*measureCount)++
	}

	return Size{Width: 10, Height: 10}
}

func _simulate_wrapping_text(node *YGNode, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	if widthMode == MeasureModeUndefined || width >= 68 {
		return Size{Width: 68, Height: 16}
	}

	return Size{Width: 50, Height: 32}
}

func _measure_assert_negative(node *YGNode, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
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
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measure)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
}

func TestMeasure_absolute_child_with_no_constraints(t *testing.T) {
	root := YGNodeNew()

	rootChild0 := YGNodeNew()
	YGNodeInsertChild(root, rootChild0, 0)

	measureCount := 0

	rootChild0Child0 := YGNodeNew()
	YGNodeStyleSetPositionType(rootChild0Child0, PositionTypeAbsolute)
	YGNodeSetContext(rootChild0Child0, &measureCount)
	YGNodeSetMeasureFunc(rootChild0Child0, _measure3)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 1, measureCount)
}

func TestDont_measure_when_min_equals_max(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measure3)
	YGNodeStyleSetMinWidth(rootChild0, 10)
	YGNodeStyleSetMaxWidth(rootChild0, 10)
	YGNodeStyleSetMinHeight(rootChild0, 10)
	YGNodeStyleSetMaxHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestDont_measure_when_min_equals_max_percentages(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measure3)
	YGNodeStyleSetMinWidthPercent(rootChild0, 10)
	YGNodeStyleSetMaxWidthPercent(rootChild0, 10)
	YGNodeStyleSetMinHeightPercent(rootChild0, 10)
	YGNodeStyleSetMaxHeightPercent(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestDont_measure_when_min_equals_max_mixed_width_percent(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measure3)
	YGNodeStyleSetMinWidthPercent(rootChild0, 10)
	YGNodeStyleSetMaxWidthPercent(rootChild0, 10)
	YGNodeStyleSetMinHeight(rootChild0, 10)
	YGNodeStyleSetMaxHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestDont_measure_when_min_equals_max_mixed_height_percent(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measure3)
	YGNodeStyleSetMinWidth(rootChild0, 10)
	YGNodeStyleSetMaxWidth(rootChild0, 10)
	YGNodeStyleSetMinHeightPercent(rootChild0, 10)
	YGNodeStyleSetMaxHeightPercent(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestMeasure_enough_size_should_be_in_single_line(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetAlignSelf(rootChild0, AlignFlexStart)
	YGNodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)

	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 68, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 16, YGNodeLayoutGetHeight(rootChild0))
}

func TestMeasure_not_enough_size_should_wrap(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 55)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetAlignSelf(rootChild0, AlignFlexStart)
	YGNodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)

	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(rootChild0))
}

func TestMeasure_zero_space_should_grow(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetHeight(root, 200)
	YGNodeStyleSetFlexDirection(root, FlexDirectionColumn)
	YGNodeStyleSetFlexGrow(root, 0)

	measureCount := 0

	rootChild0 := YGNodeNew()
	YGNodeStyleSetFlexDirection(rootChild0, FlexDirectionColumn)
	YGNodeStyleSetPadding(rootChild0, EdgeAll, 100)
	YGNodeSetContext(rootChild0, &measureCount)
	YGNodeSetMeasureFunc(rootChild0, _measure3)

	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, 282, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 282, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
}

func TestMeasure_flex_direction_row_and_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetPadding(root, EdgeLeft, 25)
	YGNodeStyleSetPadding(root, EdgeTop, 25)
	YGNodeStyleSetPadding(root, EdgeRight, 25)
	YGNodeStyleSetPadding(root, EdgeBottom, 25)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 5)
	YGNodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 75, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_flex_direction_column_and_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetPadding(root, EdgeAll, 25)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 5)
	YGNodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 57, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_flex_direction_row_no_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 5)
	YGNodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_flex_direction_row_no_padding_align_items_flexstart(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 5)
	YGNodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_with_fixed_size(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetPadding(root, EdgeAll, 25)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 5)
	YGNodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 35, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_with_flex_shrink(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetPadding(root, EdgeAll, 25)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 5)
	YGNodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_no_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 5)
	YGNodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 32, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
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
	root := YGNodeNew()
	YGNodeInsertChild(root, YGNodeNew(), 0)

	YGNodeSetMeasureFunc(root, nil)
	assert.True(t, YGNodeGetMeasureFunc(root) == nil)
}

func TestCant_call_negative_measure(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionColumn)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 10)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _measure_assert_negative)
	YGNodeStyleSetMargin(rootChild0, EdgeTop, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
}

func TestCant_call_negative_measure_horizontal(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetWidth(root, 10)
	YGNodeStyleSetHeight(root, 20)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(rootChild0, _measure_assert_negative)
	YGNodeStyleSetMargin(rootChild0, EdgeStart, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
}
