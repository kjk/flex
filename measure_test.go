package flex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func _measure3(node *YGNode, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize {
	measureCount, ok := YGNodeGetContext(node).(*int)
	if ok {
		(*measureCount)++
	}

	return YGSize{
		width:  10,
		height: 10,
	}
}

func _simulate_wrapping_text(node *YGNode, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize {
	if widthMode == YGMeasureModeUndefined || width >= 68 {
		return YGSize{width: 68, height: 16}
	}

	return YGSize{
		width: 50, height: 32,
	}
}

func _measure_assert_negative(node *YGNode, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize {
	if width < 0 {
		panic(fmt.Sprintf("width is %.2f and should be >= 0", width))
	}
	if height < 0 {
		panic(fmt.Sprintf("height is %.2f should be >= 0, height", height))
	}
	// EXPECT_GE(width, 0);
	//EXPECT_GE(height, 0);

	return YGSize{
		width: 0, height: 0,
	}
}

func TestDont_measure_single_grow_shrink_child(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	measureCount := 0

	root_child0 := YGNodeNew()
	YGNodeSetContext(root_child0, &measureCount)
	YGNodeSetMeasureFunc(root_child0, _measure)
	YGNodeStyleSetFlexGrow(root_child0, 1)
	YGNodeStyleSetFlexShrink(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assert.Equal(t, 0, measureCount)

	YGNodeFreeRecursive(root)
}

func TestMeasure_absolute_child_with_no_constraints(t *testing.T) {
	root := YGNodeNew()

	root_child0 := YGNodeNew()
	YGNodeInsertChild(root, root_child0, 0)

	measureCount := 0

	root_child0_child0 := YGNodeNew()
	YGNodeStyleSetPositionType(root_child0_child0, YGPositionTypeAbsolute)
	YGNodeSetContext(root_child0_child0, &measureCount)
	YGNodeSetMeasureFunc(root_child0_child0, _measure3)
	YGNodeInsertChild(root_child0, root_child0_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assert.Equal(t, 1, measureCount)

	YGNodeFreeRecursive(root)
}

func TestDont_measure_when_min_equals_max(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, YGAlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	measureCount := 0

	root_child0 := YGNodeNew()
	YGNodeSetContext(root_child0, &measureCount)
	YGNodeSetMeasureFunc(root_child0, _measure3)
	YGNodeStyleSetMinWidth(root_child0, 10)
	YGNodeStyleSetMaxWidth(root_child0, 10)
	YGNodeStyleSetMinHeight(root_child0, 10)
	YGNodeStyleSetMaxHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)
}

func TestDont_measure_when_min_equals_max_percentages(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, YGAlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	measureCount := 0

	root_child0 := YGNodeNew()
	YGNodeSetContext(root_child0, &measureCount)
	YGNodeSetMeasureFunc(root_child0, _measure3)
	YGNodeStyleSetMinWidthPercent(root_child0, 10)
	YGNodeStyleSetMaxWidthPercent(root_child0, 10)
	YGNodeStyleSetMinHeightPercent(root_child0, 10)
	YGNodeStyleSetMaxHeightPercent(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)
}

func TestDont_measure_when_min_equals_max_mixed_width_percent(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, YGAlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	measureCount := 0

	root_child0 := YGNodeNew()
	YGNodeSetContext(root_child0, &measureCount)
	YGNodeSetMeasureFunc(root_child0, _measure3)
	YGNodeStyleSetMinWidthPercent(root_child0, 10)
	YGNodeStyleSetMaxWidthPercent(root_child0, 10)
	YGNodeStyleSetMinHeight(root_child0, 10)
	YGNodeStyleSetMaxHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)
}

func TestDont_measure_when_min_equals_max_mixed_height_percent(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, YGAlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	measureCount := 0

	root_child0 := YGNodeNew()
	YGNodeSetContext(root_child0, &measureCount)
	YGNodeSetMeasureFunc(root_child0, _measure3)
	YGNodeStyleSetMinWidth(root_child0, 10)
	YGNodeStyleSetMaxWidth(root_child0, 10)
	YGNodeStyleSetMinHeightPercent(root_child0, 10)
	YGNodeStyleSetMaxHeightPercent(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)
}

func TestMeasure_enough_size_should_be_in_single_line(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)

	root_child0 := YGNodeNew()
	YGNodeStyleSetAlignSelf(root_child0, YGAlignFlexStart)
	YGNodeSetMeasureFunc(root_child0, _simulate_wrapping_text)

	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 68, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 16, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)
}

func TestMeasure_not_enough_size_should_wrap(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 55)

	root_child0 := YGNodeNew()
	YGNodeStyleSetAlignSelf(root_child0, YGAlignFlexStart)
	YGNodeSetMeasureFunc(root_child0, _simulate_wrapping_text)

	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(root_child0))

	YGNodeFreeRecursive(root)
}

func TestMeasure_zero_space_should_grow(t *testing.T) {
	root := YGNodeNew()
	YGNodeStyleSetHeight(root, 200)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetFlexGrow(root, 0)

	measureCount := 0

	root_child0 := YGNodeNew()
	YGNodeStyleSetFlexDirection(root_child0, YGFlexDirectionColumn)
	YGNodeStyleSetPadding(root_child0, YGEdgeAll, 100)
	YGNodeSetContext(root_child0, &measureCount)
	YGNodeSetMeasureFunc(root_child0, _measure3)

	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, 282, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 282, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))

	YGNodeFreeRecursive(root)
}

func TestMeasure_flex_direction_row_and_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetPadding(root, YGEdgeLeft, 25)
	YGNodeStyleSetPadding(root, YGEdgeTop, 25)
	YGNodeStyleSetPadding(root, YGEdgeRight, 25)
	YGNodeStyleSetPadding(root, YGEdgeBottom, 25)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _simulate_wrapping_text)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 5)
	YGNodeStyleSetHeight(root_child1, 5)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 75, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(root_child1))

	YGNodeFreeRecursive(root)


}

func TestMeasure_flex_direction_column_and_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, YGEdgeTop, 20)
	YGNodeStyleSetPadding(root, YGEdgeAll, 25)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _simulate_wrapping_text)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 5)
	YGNodeStyleSetHeight(root_child1, 5)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 57, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(root_child1))

	YGNodeFreeRecursive(root)


}

func TestMeasure_flex_direction_row_no_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetMargin(root, YGEdgeTop, 20)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _simulate_wrapping_text)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 5)
	YGNodeStyleSetHeight(root_child1, 5)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(root_child1))

	YGNodeFreeRecursive(root)


}

func TestMeasure_flex_direction_row_no_padding_align_items_flexstart(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetMargin(root, YGEdgeTop, 20)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)
	YGNodeStyleSetAlignItems(root, YGAlignFlexStart)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _simulate_wrapping_text)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 5)
	YGNodeStyleSetHeight(root_child1, 5)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(root_child1))

	YGNodeFreeRecursive(root)


}

func TestMeasure_with_fixed_size(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, YGEdgeTop, 20)
	YGNodeStyleSetPadding(root, YGEdgeAll, 25)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _simulate_wrapping_text)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 5)
	YGNodeStyleSetHeight(root_child1, 5)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 35, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(root_child1))

	YGNodeFreeRecursive(root)


}

func TestMeasure_with_flex_shrink(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, YGEdgeTop, 20)
	YGNodeStyleSetPadding(root, YGEdgeAll, 25)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _simulate_wrapping_text)
	YGNodeStyleSetFlexShrink(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 5)
	YGNodeStyleSetHeight(root_child1, 5)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(root_child1))

	YGNodeFreeRecursive(root)


}

func TestMeasure_no_padding(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetMargin(root, YGEdgeTop, 20)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 50)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _simulate_wrapping_text)
	YGNodeStyleSetFlexShrink(root_child0, 1)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root_child1, 5)
	YGNodeStyleSetHeight(root_child1, 5)
	YGNodeInsertChild(root, root_child1, 1)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 32, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(root_child1))

	YGNodeFreeRecursive(root)


}

/*
#if GTEST_HAS_DEATH_TEST
TEST(YogaDeathTest, cannot_add_child_to_node_with_measure_func) {
  root := YGNodeNew();
  YGNodeSetMeasureFunc(root, _measure3);

  root_child0 := YGNodeNew();
  ASSERT_DEATH(YGNodeInsertChild(root, root_child0, 0), "Cannot add child.*");
  YGNodeFree(root_child0);
  YGNodeFreeRecursive(root);
}

TEST(YogaDeathTest, cannot_add_nonnull_measure_func_to_non_leaf_node) {
  root := YGNodeNew();
  root_child0 := YGNodeNew();
  YGNodeInsertChild(root, root_child0, 0);

  ASSERT_DEATH(YGNodeSetMeasureFunc(root, _measure3), "Cannot set measure function.*");
  YGNodeFreeRecursive(root);
}
#endif
*/

func TestCan_nullify_measure_func_on_any_node(t *testing.T) {
	root := YGNodeNew()
	YGNodeInsertChild(root, YGNodeNew(), 0)

	YGNodeSetMeasureFunc(root, nil)
	assert.True(t, YGNodeGetMeasureFunc(root) == nil)
	YGNodeFreeRecursive(root)
}

func TestCant_call_negative_measure(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetWidth(root, 50)
	YGNodeStyleSetHeight(root, 10)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _measure_assert_negative)
	YGNodeStyleSetMargin(root_child0, YGEdgeTop, 20)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	YGNodeFreeRecursive(root)

}

func TestCant_call_negative_measure_horizontal(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 10)
	YGNodeStyleSetHeight(root, 20)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeSetMeasureFunc(root_child0, _measure_assert_negative)
	YGNodeStyleSetMargin(root_child0, YGEdgeStart, 20)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	YGNodeFreeRecursive(root)

}
