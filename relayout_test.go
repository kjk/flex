package flex

import "testing"

func TestDont_cache_computed_flex_basis_between_layouts(t *testing.T) {
	config := NewConfig()
	ConfigSetExperimentalFeatureEnabled(config, ExperimentalFeatureWebFlexBasis, true)

	root := NewNodeWithConfig(config)
	NodeStyleSetHeightPercent(root, 100)
	NodeStyleSetWidthPercent(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, 100, Undefined, DirectionLTR)
	NodeCalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))
}

func TestRecalculate_resolvedDimonsion_onchange(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	YGNodeStyleSetMinHeight(rootChild0, 10)
	YGNodeStyleSetMaxHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	YGNodeStyleSetMinHeight(rootChild0, Undefined)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))
}
