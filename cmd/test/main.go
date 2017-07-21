package main

import (
	"fmt"

	"github.com/kjk/flex"
)

func cmpFloat(exp, got float32) {
	if exp != got {
		fmt.Printf("got: %.2f, exp: %2.f\n", got, exp)
	}
}

func testAbsoluteLayoutWidthHeightStartTop() {
	config := flex.YGConfigNew()
	root := flex.YGNodeNewWithConfig(config)
	flex.YGNodeStyleSetWidth(root, 100)
	flex.YGNodeStyleSetHeight(root, 100)

	root_child0 := flex.YGNodeNewWithConfig(config)
	flex.YGNodeStyleSetPositionType(root_child0, flex.YGPositionTypeAbsolute)
	flex.YGNodeStyleSetPosition(root_child0, flex.YGEdgeStart, 10)
	flex.YGNodeStyleSetPosition(root_child0, flex.YGEdgeTop, 10)
	flex.YGNodeStyleSetWidth(root_child0, 10)
	flex.YGNodeStyleSetHeight(root_child0, 10)
	flex.YGNodeInsertChild(root, root_child0, 0)
	flex.YGNodeCalculateLayout(root, flex.YGUndefined, flex.YGUndefined, flex.YGDirectionLTR)

	cmpFloat(0, flex.YGNodeLayoutGetLeft(root))
	cmpFloat(0, flex.YGNodeLayoutGetTop(root))
	cmpFloat(100, flex.YGNodeLayoutGetWidth(root))
	cmpFloat(100, flex.YGNodeLayoutGetHeight(root))
}

func main() {
	testAbsoluteLayoutWidthHeightStartTop()
}
