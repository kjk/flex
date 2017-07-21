package flex

/*

TEST(YogaTest, absolute_layout_width_height_start_top) {
  const YGConfigRef config = YGConfigNew();

  const YGNodeRef root = YGNodeNewWithConfig(config);
  YGNodeStyleSetWidth(root, 100);
  YGNodeStyleSetHeight(root, 100);

  const YGNodeRef root_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute);
  YGNodeStyleSetPosition(root_child0, YGEdgeStart, 10);
  YGNodeStyleSetPosition(root_child0, YGEdgeTop, 10);
  YGNodeStyleSetWidth(root_child0, 10);
  YGNodeStyleSetHeight(root_child0, 10);
  YGNodeInsertChild(root, root_child0, 0);
  YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR);

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root));
  ASSERT_FLOAT_EQ(100, YGNodeLayoutGetWidth(root));
  ASSERT_FLOAT_EQ(100, YGNodeLayoutGetHeight(root));

  ASSERT_FLOAT_EQ(10, YGNodeLayoutGetLeft(root_child0));
  ASSERT_FLOAT_EQ(10, YGNodeLayoutGetTop(root_child0));
  ASSERT_FLOAT_EQ(10, YGNodeLayoutGetWidth(root_child0));
  ASSERT_FLOAT_EQ(10, YGNodeLayoutGetHeight(root_child0));

  YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL);

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root));
  ASSERT_FLOAT_EQ(100, YGNodeLayoutGetWidth(root));
  ASSERT_FLOAT_EQ(100, YGNodeLayoutGetHeight(root));

  ASSERT_FLOAT_EQ(80, YGNodeLayoutGetLeft(root_child0));
  ASSERT_FLOAT_EQ(10, YGNodeLayoutGetTop(root_child0));
  ASSERT_FLOAT_EQ(10, YGNodeLayoutGetWidth(root_child0));
  ASSERT_FLOAT_EQ(10, YGNodeLayoutGetHeight(root_child0));

  YGNodeFreeRecursive(root);

  YGConfigFree(config);
}
*/
