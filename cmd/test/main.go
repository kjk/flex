package main

import (
	"fmt"
	"testing"

	. "github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func cmpFloat(exp, got float32) {

}
func assertFloatEqual(t *testing.T, got, exp float32) {
	if exp != got {
		fmt.Printf("got: %.2f, exp: %2.f\n", got, exp)
	}
	assert.Equal(t, got, exp)
}

func testAbsolute_layout_percentage_bottom_based_on_parent_height(t *testing.T) {
	config := YGConfigNew()

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 200)

	root_child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child0, YGPositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(root_child0, YGEdgeTop, 50)
	YGNodeStyleSetWidth(root_child0, 10)
	YGNodeStyleSetHeight(root_child0, 10)
	YGNodeInsertChild(root, root_child0, 0)

	root_child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child1, YGPositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(root_child1, YGEdgeBottom, 50)
	YGNodeStyleSetWidth(root_child1, 10)
	YGNodeStyleSetHeight(root_child1, 10)
	YGNodeInsertChild(root, root_child1, 1)

	root_child2 := YGNodeNewWithConfig(config)
	YGNodeStyleSetPositionType(root_child2, YGPositionTypeAbsolute)
	YGNodeStyleSetPositionPercent(root_child2, YGEdgeTop, 10)
	YGNodeStyleSetPositionPercent(root_child2, YGEdgeBottom, 10)
	YGNodeStyleSetWidth(root_child2, 10)
	YGNodeInsertChild(root, root_child2, 2)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(root_child2))

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child0))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child1))
	assertFloatEqual(t, 90, YGNodeLayoutGetTop(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(root_child1))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(root_child2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root_child2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(root_child2))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(root_child2))

	YGNodeFreeRecursive(root)

	YGConfigFree(config)
}

func main() {
	t := &testing.T{}
	testAbsolute_layout_percentage_bottom_based_on_parent_height(t)
}
