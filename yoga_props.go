package flex

/*

Functions that get/set props, in C code generated from those C macros:

YG_NODE_PROPERTY_IMPL(void *, Context, context, context);
YG_NODE_PROPERTY_IMPL(YGPrintFunc, PrintFunc, printFunc, print);
YG_NODE_PROPERTY_IMPL(bool, HasNewLayout, hasNewLayout, hasNewLayout);
YG_NODE_PROPERTY_IMPL(YGNodeType, NodeType, nodeType, nodeType);

YG_NODE_STYLE_PROPERTY_IMPL(YGDirection, Direction, direction, direction);
YG_NODE_STYLE_PROPERTY_IMPL(YGFlexDirection, FlexDirection, flexDirection, flexDirection);
YG_NODE_STYLE_PROPERTY_IMPL(YGJustify, JustifyContent, justifyContent, justifyContent);
YG_NODE_STYLE_PROPERTY_IMPL(YGAlign, AlignContent, alignContent, alignContent);
YG_NODE_STYLE_PROPERTY_IMPL(YGAlign, AlignItems, alignItems, alignItems);
YG_NODE_STYLE_PROPERTY_IMPL(YGAlign, AlignSelf, alignSelf, alignSelf);
YG_NODE_STYLE_PROPERTY_IMPL(YGPositionType, PositionType, positionType, positionType);
YG_NODE_STYLE_PROPERTY_IMPL(YGWrap, FlexWrap, flexWrap, flexWrap);
YG_NODE_STYLE_PROPERTY_IMPL(YGOverflow, Overflow, overflow, overflow);
YG_NODE_STYLE_PROPERTY_IMPL(YGDisplay, Display, display, display);

YG_NODE_STYLE_PROPERTY_IMPL(float, Flex, flex, flex);
YG_NODE_STYLE_PROPERTY_SETTER_IMPL(float, FlexGrow, flexGrow, flexGrow);
YG_NODE_STYLE_PROPERTY_SETTER_IMPL(float, FlexShrink, flexShrink, flexShrink);
YG_NODE_STYLE_PROPERTY_UNIT_AUTO_IMPL(YGValue, FlexBasis, flexBasis, flexBasis);

YG_NODE_STYLE_EDGE_PROPERTY_UNIT_IMPL(YGValue, Position, position, position);
YG_NODE_STYLE_EDGE_PROPERTY_UNIT_IMPL(YGValue, Margin, margin, margin);
YG_NODE_STYLE_EDGE_PROPERTY_UNIT_AUTO_IMPL(YGValue, Margin, margin);
YG_NODE_STYLE_EDGE_PROPERTY_UNIT_IMPL(YGValue, Padding, padding, padding);
YG_NODE_STYLE_EDGE_PROPERTY_IMPL(float, Border, border, border);

YG_NODE_STYLE_PROPERTY_UNIT_AUTO_IMPL(YGValue, Width, width, dimensions[YGDimensionWidth]);
YG_NODE_STYLE_PROPERTY_UNIT_AUTO_IMPL(YGValue, Height, height, dimensions[YGDimensionHeight]);
YG_NODE_STYLE_PROPERTY_UNIT_IMPL(YGValue, MinWidth, minWidth, minDimensions[YGDimensionWidth]);
YG_NODE_STYLE_PROPERTY_UNIT_IMPL(YGValue, MinHeight, minHeight, minDimensions[YGDimensionHeight]);
YG_NODE_STYLE_PROPERTY_UNIT_IMPL(YGValue, MaxWidth, maxWidth, maxDimensions[YGDimensionWidth]);
YG_NODE_STYLE_PROPERTY_UNIT_IMPL(YGValue, MaxHeight, maxHeight, maxDimensions[YGDimensionHeight]);

// Yoga specific properties, not compatible with flexbox specification
YG_NODE_STYLE_PROPERTY_IMPL(float, AspectRatio, aspectRatio, aspectRatio);

YG_NODE_LAYOUT_PROPERTY_IMPL(float, Left, position[YGEdgeLeft]);
YG_NODE_LAYOUT_PROPERTY_IMPL(float, Top, position[YGEdgeTop]);
YG_NODE_LAYOUT_PROPERTY_IMPL(float, Right, position[YGEdgeRight]);
YG_NODE_LAYOUT_PROPERTY_IMPL(float, Bottom, position[YGEdgeBottom]);
YG_NODE_LAYOUT_PROPERTY_IMPL(float, Width, dimensions[YGDimensionWidth]);
YG_NODE_LAYOUT_PROPERTY_IMPL(float, Height, dimensions[YGDimensionHeight]);
YG_NODE_LAYOUT_PROPERTY_IMPL(YGDirection, Direction, direction);
YG_NODE_LAYOUT_PROPERTY_IMPL(bool, HadOverflow, hadOverflow);

YG_NODE_LAYOUT_RESOLVED_PROPERTY_IMPL(float, Margin, margin);
YG_NODE_LAYOUT_RESOLVED_PROPERTY_IMPL(float, Border, border);
YG_NODE_LAYOUT_RESOLVED_PROPERTY_IMPL(float, Padding, padding);
*/

// YGNodeStyleSetWidth sets width
func YGNodeStyleSetWidth(node *Node, width float32) {
	dim := &node.style.dimensions[DimensionWidth]
	if dim.Value != width || dim.Unit != UnitPoint {
		dim.Value = width
		dim.Unit = UnitPoint
		if YGFloatIsUndefined(width) {
			dim.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetWidthPercent sets width percent
func YGNodeStyleSetWidthPercent(node *Node, width float32) {
	dim := &node.style.dimensions[DimensionWidth]
	if dim.Value != width || dim.Unit != UnitPercent {
		dim.Value = width
		dim.Unit = UnitPercent
		if YGFloatIsUndefined(width) {
			dim.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetWidthAuto sets width auto
func YGNodeStyleSetWidthAuto(node *Node) {
	dim := &node.style.dimensions[DimensionWidth]
	if dim.Unit != UnitAuto {
		dim.Value = Undefined
		dim.Unit = UnitAuto
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetWidth gets width
func YGNodeStyleGetWidth(node *Node) Value {
	return node.style.dimensions[DimensionWidth]
}

// YGNodeStyleSetHeight sets height
func YGNodeStyleSetHeight(node *Node, height float32) {
	dim := &node.style.dimensions[DimensionHeight]
	if dim.Value != height || dim.Unit != UnitPoint {
		dim.Value = height
		dim.Unit = UnitPoint
		if YGFloatIsUndefined(height) {
			dim.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetHeightPercent sets height percent
func YGNodeStyleSetHeightPercent(node *Node, height float32) {
	dim := &node.style.dimensions[DimensionHeight]
	if dim.Value != height || dim.Unit != UnitPercent {
		dim.Value = height
		dim.Unit = UnitPercent
		if YGFloatIsUndefined(height) {
			dim.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetHeightAuto sets height auto
func YGNodeStyleSetHeightAuto(node *Node) {
	dim := &node.style.dimensions[DimensionHeight]
	if dim.Unit != UnitAuto {
		dim.Value = Undefined
		dim.Unit = UnitAuto
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetHeight gets height
func YGNodeStyleGetHeight(node *Node) Value {
	return node.style.dimensions[DimensionHeight]
}

// YGNodeStyleSetPositionType sets position type
func YGNodeStyleSetPositionType(node *Node, positionType PositionType) {
	if node.style.positionType != positionType {
		node.style.positionType = positionType
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetPositionType gets position type
func YGNodeStyleGetPositionType(node *Node) PositionType {
	return node.style.positionType
}

// YGNodeStyleSetPosition sets position
func YGNodeStyleSetPosition(node *Node, edge Edge, position float32) {
	pos := &node.style.position[edge]
	if pos.Value != position || pos.Unit != UnitPoint {
		pos.Value = position
		pos.Unit = UnitPoint
		if YGFloatIsUndefined(position) {
			pos.Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetPositionPercent sets position percent
func YGNodeStyleSetPositionPercent(node *Node, edge Edge, position float32) {
	pos := &node.style.position[edge]
	if pos.Value != position || pos.Unit != UnitPercent {
		pos.Value = position
		pos.Unit = UnitPercent
		if YGFloatIsUndefined(position) {
			pos.Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetPosition gets position
func YGNodeStyleGetPosition(node *Node, edge Edge) Value {
	return node.style.position[edge]
}

// YGNodeSetContext sets context
func YGNodeSetContext(node *Node, context interface{}) {
	node.context = context
}

// YGNodeGetContext gets context
func YGNodeGetContext(node *Node) interface{} {
	return node.context
}

// YGNodeSetPrintFunc sets print func
func YGNodeSetPrintFunc(node *Node, printFunc PrintFunc) {
	node.print = printFunc
}

// YGNodeGetPrintFunc gets print func
func YGNodeGetPrintFunc(node *Node) PrintFunc {
	return node.print
}

// YGNodeSetHasNewLayout sets has new layout
func YGNodeSetHasNewLayout(node *Node, hasNewLayout bool) {
	node.hasNewLayout = hasNewLayout
}

// YGNodeGetHasNewLayout gets has new layout
func YGNodeGetHasNewLayout(node *Node) bool {
	return node.hasNewLayout
}

// YGNodeSetNodeType sets node type
func YGNodeSetNodeType(node *Node, nodeType NodeType) {
	node.nodeType = nodeType
}

// YGNodeGetNodeType gets node type
func YGNodeGetNodeType(node *Node) NodeType {
	return node.nodeType
}

// YGNodeStyleSetDirection sets direction
func YGNodeStyleSetDirection(node *Node, direction Direction) {
	if node.style.direction != direction {
		node.style.direction = direction
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetDirection gets direction
func YGNodeStyleGetDirection(node *Node) Direction {
	return node.style.direction
}

// YGNodeStyleSetFlexDirection sets flex directions
func YGNodeStyleSetFlexDirection(node *Node, flexDirection FlexDirection) {
	if node.style.flexDirection != flexDirection {
		node.style.flexDirection = flexDirection
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetFlexDirection gets flex direction
func YGNodeStyleGetFlexDirection(node *Node) FlexDirection {
	return node.style.flexDirection
}

// YGNodeStyleSetJustifyContent sets justify content
func YGNodeStyleSetJustifyContent(node *Node, justifyContent Justify) {
	if node.style.justifyContent != justifyContent {
		node.style.justifyContent = justifyContent
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetJustifyContent gets justify content
func YGNodeStyleGetJustifyContent(node *Node) Justify {
	return node.style.justifyContent
}

// YGNodeStyleSetAlignContent sets align content
func YGNodeStyleSetAlignContent(node *Node, alignContent Align) {
	if node.style.alignContent != alignContent {
		node.style.alignContent = alignContent
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetAlignContent gets align content
func YGNodeStyleGetAlignContent(node *Node) Align {
	return node.style.alignContent
}

// YGNodeStyleSetAlignItems sets align content
func YGNodeStyleSetAlignItems(node *Node, alignItems Align) {
	if node.style.alignItems != alignItems {
		node.style.alignItems = alignItems
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetAlignItems gets align items
func YGNodeStyleGetAlignItems(node *Node) Align {
	return node.style.alignItems
}

// YGNodeStyleSetAlignSelf sets align self
func YGNodeStyleSetAlignSelf(node *Node, alignSelf Align) {
	if node.style.alignSelf != alignSelf {
		node.style.alignSelf = alignSelf
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetAlignSelf gets align self
func YGNodeStyleGetAlignSelf(node *Node) Align {
	return node.style.alignSelf
}

// YGNodeStyleSetFlexWrap sets flex wrap
func YGNodeStyleSetFlexWrap(node *Node, flexWrap Wrap) {
	if node.style.flexWrap != flexWrap {
		node.style.flexWrap = flexWrap
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetFlexWrap gets flex wrap
func YGNodeStyleGetFlexWrap(node *Node) Wrap {
	return node.style.flexWrap
}

// YGNodeStyleSetOverflow sets overflow
func YGNodeStyleSetOverflow(node *Node, overflow Overflow) {
	if node.style.overflow != overflow {
		node.style.overflow = overflow
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetOverflow gets overflow
func YGNodeStyleGetOverflow(node *Node) Overflow {
	return node.style.overflow
}

// YGNodeStyleSetDisplay sets display
func YGNodeStyleSetDisplay(node *Node, display Display) {
	if node.style.display != display {
		node.style.display = display
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetDisplay gets display
func YGNodeStyleGetDisplay(node *Node) Display {
	return node.style.display
}

// YGNodeStyleSetFlex sets flex
func YGNodeStyleSetFlex(node *Node, flex float32) {
	if node.style.flex != flex {
		node.style.flex = flex
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetFlex gets flex
func YGNodeStyleGetFlex(node *Node) float32 {
	return node.style.flex
}

// YGNodeStyleSetFlexGrow sets flex grow
func YGNodeStyleSetFlexGrow(node *Node, flexGrow float32) {
	if node.style.flexGrow != flexGrow {
		node.style.flexGrow = flexGrow
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetFlexShrink sets flex shrink
func YGNodeStyleSetFlexShrink(node *Node, flexShrink float32) {
	if node.style.flexShrink != flexShrink {
		node.style.flexShrink = flexShrink
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetFlexBasis sets flex basis
func YGNodeStyleSetFlexBasis(node *Node, flexBasis float32) {
	if node.style.flexBasis.Value != flexBasis ||
		node.style.flexBasis.Unit != UnitPoint {
		node.style.flexBasis.Value = flexBasis
		node.style.flexBasis.Unit = UnitPoint
		if YGFloatIsUndefined(flexBasis) {
			node.style.flexBasis.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetFlexBasisPercent sets flex basis percent
func YGNodeStyleSetFlexBasisPercent(node *Node, flexBasis float32) {
	if node.style.flexBasis.Value != flexBasis ||
		node.style.flexBasis.Unit != UnitPercent {
		node.style.flexBasis.Value = flexBasis
		node.style.flexBasis.Unit = UnitPercent
		if YGFloatIsUndefined(flexBasis) {
			node.style.flexBasis.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetFlexBasisAuto sets flex basis auto
func YGNodeStyleSetFlexBasisAuto(node *Node) {
	if node.style.flexBasis.Unit != UnitAuto {
		node.style.flexBasis.Value = Undefined
		node.style.flexBasis.Unit = UnitAuto
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetFlexBasis gets flex basis
func YGNodeStyleGetFlexBasis(node *Node) Value {
	return node.style.flexBasis
}

// YGNodeStyleSetMargin sets margin
func YGNodeStyleSetMargin(node *Node, edge Edge, margin float32) {
	if node.style.margin[edge].Value != margin ||
		node.style.margin[edge].Unit != UnitPoint {
		node.style.margin[edge].Value = margin
		node.style.margin[edge].Unit = UnitPoint
		if YGFloatIsUndefined(margin) {
			node.style.margin[edge].Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMarginPercent sets margin percent
func YGNodeStyleSetMarginPercent(node *Node, edge Edge, margin float32) {
	if node.style.margin[edge].Value != margin ||
		node.style.margin[edge].Unit != UnitPercent {
		node.style.margin[edge].Value = margin
		node.style.margin[edge].Unit = UnitPercent
		if YGFloatIsUndefined(margin) {
			node.style.margin[edge].Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetMargin gets margin
func YGNodeStyleGetMargin(node *Node, edge Edge) Value {
	return node.style.margin[edge]
}

// YGNodeStyleSetMarginAuto sets margin auto
func YGNodeStyleSetMarginAuto(node *Node, edge Edge) {
	if node.style.margin[edge].Unit != UnitAuto {
		node.style.margin[edge].Value = Undefined
		node.style.margin[edge].Unit = UnitAuto
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetPadding sets padding
func YGNodeStyleSetPadding(node *Node, edge Edge, padding float32) {
	if node.style.padding[edge].Value != padding ||
		node.style.padding[edge].Unit != UnitPoint {
		node.style.padding[edge].Value = padding
		node.style.padding[edge].Unit = UnitPoint
		if YGFloatIsUndefined(padding) {
			node.style.padding[edge].Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetPaddingPercent sets padding percent
func YGNodeStyleSetPaddingPercent(node *Node, edge Edge, padding float32) {
	if node.style.padding[edge].Value != padding ||
		node.style.padding[edge].Unit != UnitPercent {
		node.style.padding[edge].Value = padding
		node.style.padding[edge].Unit = UnitPercent
		if YGFloatIsUndefined(padding) {
			node.style.padding[edge].Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetPadding gets padding
func YGNodeStyleGetPadding(node *Node, edge Edge) Value {
	return node.style.padding[edge]
}

// YGNodeStyleSetBorder sets border
func YGNodeStyleSetBorder(node *Node, edge Edge, border float32) {
	if node.style.border[edge].Value != border ||
		node.style.border[edge].Unit != UnitPoint {
		node.style.border[edge].Value = border
		node.style.border[edge].Unit = UnitPoint
		if YGFloatIsUndefined(border) {
			node.style.border[edge].Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetBorder gets border
func YGNodeStyleGetBorder(node *Node, edge Edge) float32 {
	return node.style.border[edge].Value
}

// YGNodeStyleSetMinWidth sets min width
func YGNodeStyleSetMinWidth(node *Node, minWidth float32) {
	if node.style.minDimensions[DimensionWidth].Value != minWidth ||
		node.style.minDimensions[DimensionWidth].Unit != UnitPoint {
		node.style.minDimensions[DimensionWidth].Value = minWidth
		node.style.minDimensions[DimensionWidth].Unit = UnitPoint
		if YGFloatIsUndefined(minWidth) {
			node.style.minDimensions[DimensionWidth].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMinWidthPercent sets width percent
func YGNodeStyleSetMinWidthPercent(node *Node, minWidth float32) {
	if node.style.minDimensions[DimensionWidth].Value != minWidth ||
		node.style.minDimensions[DimensionWidth].Unit != UnitPercent {
		node.style.minDimensions[DimensionWidth].Value = minWidth
		node.style.minDimensions[DimensionWidth].Unit = UnitPercent
		if YGFloatIsUndefined(minWidth) {
			node.style.minDimensions[DimensionWidth].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetMinWidth gets min width
func YGNodeStyleGetMinWidth(node *Node) Value {
	return node.style.minDimensions[DimensionWidth]
}

// YGNodeStyleSetMinHeight sets min width
func YGNodeStyleSetMinHeight(node *Node, minHeight float32) {
	if node.style.minDimensions[DimensionHeight].Value != minHeight ||
		node.style.minDimensions[DimensionHeight].Unit != UnitPoint {
		node.style.minDimensions[DimensionHeight].Value = minHeight
		node.style.minDimensions[DimensionHeight].Unit = UnitPoint
		if YGFloatIsUndefined(minHeight) {
			node.style.minDimensions[DimensionHeight].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMinHeightPercent sets min height percent
func YGNodeStyleSetMinHeightPercent(node *Node, minHeight float32) {
	if node.style.minDimensions[DimensionHeight].Value != minHeight ||
		node.style.minDimensions[DimensionHeight].Unit != UnitPercent {
		node.style.minDimensions[DimensionHeight].Value = minHeight
		node.style.minDimensions[DimensionHeight].Unit = UnitPercent
		if YGFloatIsUndefined(minHeight) {
			node.style.minDimensions[DimensionHeight].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetMinHeight gets min height
func YGNodeStyleGetMinHeight(node *Node) Value {
	return node.style.minDimensions[DimensionHeight]
}

// YGNodeStyleSetMaxWidth sets max width
func YGNodeStyleSetMaxWidth(node *Node, maxWidth float32) {
	if node.style.maxDimensions[DimensionWidth].Value != maxWidth ||
		node.style.maxDimensions[DimensionWidth].Unit != UnitPoint {
		node.style.maxDimensions[DimensionWidth].Value = maxWidth
		node.style.maxDimensions[DimensionWidth].Unit = UnitPoint
		if YGFloatIsUndefined(maxWidth) {
			node.style.maxDimensions[DimensionWidth].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMaxWidthPercent sets max width percent
func YGNodeStyleSetMaxWidthPercent(node *Node, maxWidth float32) {
	if node.style.maxDimensions[DimensionWidth].Value != maxWidth ||
		node.style.maxDimensions[DimensionWidth].Unit != UnitPercent {
		node.style.maxDimensions[DimensionWidth].Value = maxWidth
		node.style.maxDimensions[DimensionWidth].Unit = UnitPercent
		if YGFloatIsUndefined(maxWidth) {
			node.style.maxDimensions[DimensionWidth].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetMaxWidth gets max width
func YGNodeStyleGetMaxWidth(node *Node) Value {
	return node.style.maxDimensions[DimensionWidth]
}

// YGNodeStyleSetMaxHeight sets max width
func YGNodeStyleSetMaxHeight(node *Node, maxHeight float32) {
	if node.style.maxDimensions[DimensionHeight].Value != maxHeight ||
		node.style.maxDimensions[DimensionHeight].Unit != UnitPoint {
		node.style.maxDimensions[DimensionHeight].Value = maxHeight
		node.style.maxDimensions[DimensionHeight].Unit = UnitPoint
		if YGFloatIsUndefined(maxHeight) {
			node.style.maxDimensions[DimensionHeight].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMaxHeightPercent sets max height percent
func YGNodeStyleSetMaxHeightPercent(node *Node, maxHeight float32) {
	if node.style.maxDimensions[DimensionHeight].Value != maxHeight ||
		node.style.maxDimensions[DimensionHeight].Unit != UnitPercent {
		node.style.maxDimensions[DimensionHeight].Value = maxHeight
		node.style.maxDimensions[DimensionHeight].Unit = UnitPercent
		if YGFloatIsUndefined(maxHeight) {
			node.style.maxDimensions[DimensionHeight].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetMaxHeight gets max height
func YGNodeStyleGetMaxHeight(node *Node) Value {
	return node.style.maxDimensions[DimensionHeight]
}

// YGNodeStyleSetAspectRatio sets axpect ratio
func YGNodeStyleSetAspectRatio(node *Node, aspectRatio float32) {
	if node.style.aspectRatio != aspectRatio {
		node.style.aspectRatio = aspectRatio
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetAspectRatio gets aspect ratio
func YGNodeStyleGetAspectRatio(node *Node) float32 {
	return node.style.aspectRatio
}

// YGNodeLayoutGetLeft gets left
func YGNodeLayoutGetLeft(node *Node) float32 {
	return node.layout.position[EdgeLeft]
}

// YGNodeLayoutGetTop gets top
func YGNodeLayoutGetTop(node *Node) float32 {
	return node.layout.position[EdgeTop]
}

// YGNodeLayoutGetRight gets right
func YGNodeLayoutGetRight(node *Node) float32 {
	return node.layout.position[EdgeRight]
}

// YGNodeLayoutGetBottom gets bottom
func YGNodeLayoutGetBottom(node *Node) float32 {
	return node.layout.position[EdgeBottom]
}

// YGNodeLayoutGetWidth gets width
func YGNodeLayoutGetWidth(node *Node) float32 {
	return node.layout.dimensions[DimensionWidth]
}

// YGNodeLayoutGetHeight gets height
func YGNodeLayoutGetHeight(node *Node) float32 {
	return node.layout.dimensions[DimensionHeight]
}

// YGNodeLayoutGetDirection gets direction
func YGNodeLayoutGetDirection(node *Node) Direction {
	return node.layout.direction
}

// YGNodeLayoutGetHadOverflow gets hadOverflow
func YGNodeLayoutGetHadOverflow(node *Node) bool {
	return node.layout.hadOverflow
}

// YGNodeLayoutGetMargin gets margin
func YGNodeLayoutGetMargin(node *Node, edge Edge) float32 {
	assertWithNode(node, edge < EdgeEnd, "Cannot get layout properties of multi-edge shorthands")
	if edge == EdgeLeft {
		if node.layout.direction == DirectionRTL {
			return node.layout.margin[EdgeEnd]
		}
		return node.layout.margin[EdgeStart]
	}
	if edge == EdgeRight {
		if node.layout.direction == DirectionRTL {
			return node.layout.margin[EdgeStart]
		}
		return node.layout.margin[EdgeEnd]
	}
	return node.layout.margin[edge]
}

// YGNodeLayoutGetBorder gets border
func YGNodeLayoutGetBorder(node *Node, edge Edge) float32 {
	assertWithNode(node, edge < EdgeEnd,
		"Cannot get layout properties of multi-edge shorthands")
	if edge == EdgeLeft {
		if node.layout.direction == DirectionRTL {
			return node.layout.border[EdgeEnd]
		}
		return node.layout.border[EdgeStart]
	}
	if edge == EdgeRight {
		if node.layout.direction == DirectionRTL {
			return node.layout.border[EdgeStart]
		}
		return node.layout.border[EdgeEnd]
	}
	return node.layout.border[edge]
}

// YGNodeLayoutGetPadding gets padding
func YGNodeLayoutGetPadding(node *Node, edge Edge) float32 {
	assertWithNode(node, edge < EdgeEnd,
		"Cannot get layout properties of multi-edge shorthands")
	if edge == EdgeLeft {
		if node.layout.direction == DirectionRTL {
			return node.layout.padding[EdgeEnd]
		}
		return node.layout.padding[EdgeStart]
	}
	if edge == EdgeRight {
		if node.layout.direction == DirectionRTL {
			return node.layout.padding[EdgeStart]
		}
		return node.layout.padding[EdgeEnd]
	}
	return node.layout.padding[edge]
}
