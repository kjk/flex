package flex

import "testing"

func TestDont_cache_computed_flex_basis_between_layouts(t *testing.T) {
	config := NewConfig()
	config.SetExperimentalFeatureEnabled(ExperimentalFeatureWebFlexBasis, true)

	root := NewNodeWithConfig(config)
	root.StyleSetHeightPercent(100)
	root.StyleSetWidthPercent(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexBasisPercent(100)
	root.InsertChild(rootChild0, 0)

	CalculateLayout(root, 100, Undefined, DirectionLTR)
	CalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestRecalculate_resolvedDimonsion_onchange(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	rootChild0.StyleSetMinHeight(10)
	rootChild0.StyleSetMaxHeight(10)
	root.InsertChild(rootChild0, 0)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	rootChild0.StyleSetMinHeight(Undefined)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetHeight())
}
