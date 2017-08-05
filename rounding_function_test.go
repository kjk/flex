package flex

import "testing"

func TestRounding_value(t *testing.T) {

	// Test that whole numbers are rounded to whole despite ceil/floor flags
	assertFloatEqual(t, 6.0, roundValueToPixelGrid(6.000001, 2.0, false, false))
	assertFloatEqual(t, 6.0, roundValueToPixelGrid(6.000001, 2.0, true, false))
	assertFloatEqual(t, 6.0, roundValueToPixelGrid(6.000001, 2.0, false, true))
	assertFloatEqual(t, 6.0, roundValueToPixelGrid(5.999999, 2.0, false, false))
	assertFloatEqual(t, 6.0, roundValueToPixelGrid(5.999999, 2.0, true, false))
	assertFloatEqual(t, 6.0, roundValueToPixelGrid(5.999999, 2.0, false, true))

	// Test that numbers with fraction are rounded correctly accounting for ceil/floor flags
	assertFloatEqual(t, 6.0, roundValueToPixelGrid(6.01, 2.0, false, false))
	assertFloatEqual(t, 6.5, roundValueToPixelGrid(6.01, 2.0, true, false))
	assertFloatEqual(t, 6.0, roundValueToPixelGrid(6.01, 2.0, false, true))
	assertFloatEqual(t, 6.0, roundValueToPixelGrid(5.99, 2.0, false, false))
	assertFloatEqual(t, 6.0, roundValueToPixelGrid(5.99, 2.0, true, false))
	assertFloatEqual(t, 5.5, roundValueToPixelGrid(5.99, 2.0, false, true))
}
