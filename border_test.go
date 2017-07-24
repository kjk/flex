package flex

import "testing"

func TestBorder_no_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetBorder(root, EdgeLeft, 10)
	NodeStyleSetBorder(root, EdgeTop, 10)
	NodeStyleSetBorder(root, EdgeRight, 10)
	NodeStyleSetBorder(root, EdgeBottom, 10)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 20, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(root))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 20, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(root))
}

func TestBorder_container_match_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetBorder(root, EdgeLeft, 10)
	NodeStyleSetBorder(root, EdgeTop, 10)
	NodeStyleSetBorder(root, EdgeRight, 10)
	NodeStyleSetBorder(root, EdgeBottom, 10)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 30, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 30, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 30, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
}

func TestBorder_flex_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetBorder(root, EdgeLeft, 10)
	NodeStyleSetBorder(root, EdgeTop, 10)
	NodeStyleSetBorder(root, EdgeRight, 10)
	NodeStyleSetBorder(root, EdgeBottom, 10)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetWidth(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0))
}

func TestBorder_stretch_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetBorder(root, EdgeLeft, 10)
	NodeStyleSetBorder(root, EdgeTop, 10)
	NodeStyleSetBorder(root, EdgeRight, 10)
	NodeStyleSetBorder(root, EdgeBottom, 10)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
}

func TestBorder_center_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetBorder(root, EdgeStart, 10)
	NodeStyleSetBorder(root, EdgeEnd, 20)
	NodeStyleSetBorder(root, EdgeBottom, 20)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 40, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 35, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 35, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
}
