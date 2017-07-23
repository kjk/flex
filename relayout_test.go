package flex

import "testing"

func TestDont_cache_computed_flex_basis_between_layouts(t *testing.T) {
	config := YGConfigNew()
	YGConfigSetExperimentalFeatureEnabled(config, YGExperimentalFeatureWebFlexBasis, true)

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeightPercent(root, 100)
	YGNodeStyleSetWidthPercent(root, 100)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexBasisPercent(root_child0, 100)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, 100, YGUndefined, YGDirectionLTR)
	YGNodeCalculateLayout(root, 100, 100, YGDirectionLTR)

	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root_child0))




}

func TestRecalculate_resolvedDimonsion_onchange(t *testing.T) {
	root := YGNodeNew()

	root_child0 := YGNodeNew()
	YGNodeStyleSetMinHeight(root_child0, 10)
	YGNodeStyleSetMaxHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	YGNodeStyleSetMinHeight(root_child0, YGUndefined)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(root_child0))


}
