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
func YGNodeStyleSetWidth(node *YGNode, width float32) {
	dim := &node.style.dimensions[DimensionWidth]
	if dim.value != width || dim.unit != UnitPoint {
		dim.value = width
		dim.unit = UnitPoint
		if YGFloatIsUndefined(width) {
			dim.unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetWidthPercent sets width percent
func YGNodeStyleSetWidthPercent(node *YGNode, width float32) {
	dim := &node.style.dimensions[DimensionWidth]
	if dim.value != width || dim.unit != UnitPercent {
		dim.value = width
		dim.unit = UnitPercent
		if YGFloatIsUndefined(width) {
			dim.unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetWidthAuto sets width auto
func YGNodeStyleSetWidthAuto(node *YGNode) {
	dim := &node.style.dimensions[DimensionWidth]
	if dim.unit != UnitAuto {
		dim.value = YGUndefined
		dim.unit = UnitAuto
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetWidth gets width
func YGNodeStyleGetWidth(node *YGNode) YGValue {
	return node.style.dimensions[DimensionWidth]
}

// YGNodeStyleSetHeight sets height
func YGNodeStyleSetHeight(node *YGNode, height float32) {
	dim := &node.style.dimensions[DimensionHeight]
	if dim.value != height || dim.unit != UnitPoint {
		dim.value = height
		dim.unit = UnitPoint
		if YGFloatIsUndefined(height) {
			dim.unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetHeightPercent sets height percent
func YGNodeStyleSetHeightPercent(node *YGNode, height float32) {
	dim := &node.style.dimensions[DimensionHeight]
	if dim.value != height || dim.unit != UnitPercent {
		dim.value = height
		dim.unit = UnitPercent
		if YGFloatIsUndefined(height) {
			dim.unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetHeightAuto sets height auto
func YGNodeStyleSetHeightAuto(node *YGNode) {
	dim := &node.style.dimensions[DimensionHeight]
	if dim.unit != UnitAuto {
		dim.value = YGUndefined
		dim.unit = UnitAuto
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetHeight gets height
func YGNodeStyleGetHeight(node *YGNode) YGValue {
	return node.style.dimensions[DimensionHeight]
}

// YGNodeStyleSetPositionType sets position type
func YGNodeStyleSetPositionType(node *YGNode, positionType PositionType) {
	if node.style.positionType != positionType {
		node.style.positionType = positionType
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetPositionType gets position type
func YGNodeStyleGetPositionType(node *YGNode) PositionType {
	return node.style.positionType
}

// YGNodeStyleSetPosition sets position
func YGNodeStyleSetPosition(node *YGNode, edge Edge, position float32) {
	pos := &node.style.position[edge]
	if pos.value != position || pos.unit != UnitPoint {
		pos.value = position
		pos.unit = UnitPoint
		if YGFloatIsUndefined(position) {
			pos.unit = UnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetPositionPercent sets position percent
func YGNodeStyleSetPositionPercent(node *YGNode, edge Edge, position float32) {
	pos := &node.style.position[edge]
	if pos.value != position || pos.unit != UnitPercent {
		pos.value = position
		pos.unit = UnitPercent
		if YGFloatIsUndefined(position) {
			pos.unit = UnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetPosition gets position
func YGNodeStyleGetPosition(node *YGNode, edge Edge) YGValue {
	return node.style.position[edge]
}

// YGNodeSetContext sets context
func YGNodeSetContext(node *YGNode, context interface{}) {
	node.context = context
}

// YGNodeGetContext gets context
func YGNodeGetContext(node *YGNode) interface{} {
	return node.context
}

// YGNodeSetPrintFunc sets print func
func YGNodeSetPrintFunc(node *YGNode, printFunc YGPrintFunc) {
	node.print = printFunc
}

// YGNodeGetPrintFunc gets print func
func YGNodeGetPrintFunc(node *YGNode) YGPrintFunc {
	return node.print
}

// YGNodeSetHasNewLayout sets has new layout
func YGNodeSetHasNewLayout(node *YGNode, hasNewLayout bool) {
	node.hasNewLayout = hasNewLayout
}

// YGNodeGetHasNewLayout gets has new layout
func YGNodeGetHasNewLayout(node *YGNode) bool {
	return node.hasNewLayout
}

// YGNodeSetNodeType sets node type
func YGNodeSetNodeType(node *YGNode, nodeType NodeType) {
	node.nodeType = nodeType
}

// YGNodeGetNodeType gets node type
func YGNodeGetNodeType(node *YGNode) NodeType {
	return node.nodeType
}

// YGNodeStyleSetDirection sets direction
func YGNodeStyleSetDirection(node *YGNode, direction Direction) {
	if node.style.direction != direction {
		node.style.direction = direction
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetDirection gets direction
func YGNodeStyleGetDirection(node *YGNode) Direction {
	return node.style.direction
}

// YGNodeStyleSetFlexDirection sets flex directions
func YGNodeStyleSetFlexDirection(node *YGNode, flexDirection FlexDirection) {
	if node.style.flexDirection != flexDirection {
		node.style.flexDirection = flexDirection
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetFlexDirection gets flex direction
func YGNodeStyleGetFlexDirection(node *YGNode) FlexDirection {
	return node.style.flexDirection
}

// YGNodeStyleSetJustifyContent sets justify content
func YGNodeStyleSetJustifyContent(node *YGNode, justifyContent Justify) {
	if node.style.justifyContent != justifyContent {
		node.style.justifyContent = justifyContent
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetJustifyContent gets justify content
func YGNodeStyleGetJustifyContent(node *YGNode) Justify {
	return node.style.justifyContent
}

// YGNodeStyleSetAlignContent sets align content
func YGNodeStyleSetAlignContent(node *YGNode, alignContent Align) {
	if node.style.alignContent != alignContent {
		node.style.alignContent = alignContent
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetAlignContent gets align content
func YGNodeStyleGetAlignContent(node *YGNode) Align {
	return node.style.alignContent
}

// YGNodeStyleSetAlignItems sets align content
func YGNodeStyleSetAlignItems(node *YGNode, alignItems Align) {
	if node.style.alignItems != alignItems {
		node.style.alignItems = alignItems
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetAlignItems gets align items
func YGNodeStyleGetAlignItems(node *YGNode) Align {
	return node.style.alignItems
}

// YGNodeStyleSetAlignSelf sets align self
func YGNodeStyleSetAlignSelf(node *YGNode, alignSelf Align) {
	if node.style.alignSelf != alignSelf {
		node.style.alignSelf = alignSelf
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetAlignSelf gets align self
func YGNodeStyleGetAlignSelf(node *YGNode) Align {
	return node.style.alignSelf
}

// YGNodeStyleSetFlexWrap sets flex wrap
func YGNodeStyleSetFlexWrap(node *YGNode, flexWrap Wrap) {
	if node.style.flexWrap != flexWrap {
		node.style.flexWrap = flexWrap
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetFlexWrap gets flex wrap
func YGNodeStyleGetFlexWrap(node *YGNode) Wrap {
	return node.style.flexWrap
}

// YGNodeStyleSetOverflow sets overflow
func YGNodeStyleSetOverflow(node *YGNode, overflow Overflow) {
	if node.style.overflow != overflow {
		node.style.overflow = overflow
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetOverflow gets overflow
func YGNodeStyleGetOverflow(node *YGNode) Overflow {
	return node.style.overflow
}

// YGNodeStyleSetDisplay sets display
func YGNodeStyleSetDisplay(node *YGNode, display Display) {
	if node.style.display != display {
		node.style.display = display
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetDisplay gets display
func YGNodeStyleGetDisplay(node *YGNode) Display {
	return node.style.display
}

// YGNodeStyleSetFlex sets flex
func YGNodeStyleSetFlex(node *YGNode, flex float32) {
	if node.style.flex != flex {
		node.style.flex = flex
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetFlex gets flex
func YGNodeStyleGetFlex(node *YGNode) float32 {
	return node.style.flex
}

// YGNodeStyleSetFlexGrow sets flex grow
func YGNodeStyleSetFlexGrow(node *YGNode, flexGrow float32) {
	if node.style.flexGrow != flexGrow {
		node.style.flexGrow = flexGrow
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetFlexShrink sets flex shrink
func YGNodeStyleSetFlexShrink(node *YGNode, flexShrink float32) {
	if node.style.flexShrink != flexShrink {
		node.style.flexShrink = flexShrink
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetFlexBasis sets flex basis
func YGNodeStyleSetFlexBasis(node *YGNode, flexBasis float32) {
	if node.style.flexBasis.value != flexBasis ||
		node.style.flexBasis.unit != UnitPoint {
		node.style.flexBasis.value = flexBasis
		node.style.flexBasis.unit = UnitPoint
		if YGFloatIsUndefined(flexBasis) {
			node.style.flexBasis.unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetFlexBasisPercent sets flex basis percent
func YGNodeStyleSetFlexBasisPercent(node *YGNode, flexBasis float32) {
	if node.style.flexBasis.value != flexBasis ||
		node.style.flexBasis.unit != UnitPercent {
		node.style.flexBasis.value = flexBasis
		node.style.flexBasis.unit = UnitPercent
		if YGFloatIsUndefined(flexBasis) {
			node.style.flexBasis.unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetFlexBasisAuto sets flex basis auto
func YGNodeStyleSetFlexBasisAuto(node *YGNode) {
	if node.style.flexBasis.unit != UnitAuto {
		node.style.flexBasis.value = YGUndefined
		node.style.flexBasis.unit = UnitAuto
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetFlexBasis gets flex basis
func YGNodeStyleGetFlexBasis(node *YGNode) YGValue {
	return node.style.flexBasis
}

// YGNodeStyleSetMargin sets margin
func YGNodeStyleSetMargin(node *YGNode, edge Edge, margin float32) {
	if node.style.margin[edge].value != margin ||
		node.style.margin[edge].unit != UnitPoint {
		node.style.margin[edge].value = margin
		node.style.margin[edge].unit = UnitPoint
		if YGFloatIsUndefined(margin) {
			node.style.margin[edge].unit = UnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMarginPercent sets margin percent
func YGNodeStyleSetMarginPercent(node *YGNode, edge Edge, margin float32) {
	if node.style.margin[edge].value != margin ||
		node.style.margin[edge].unit != UnitPercent {
		node.style.margin[edge].value = margin
		node.style.margin[edge].unit = UnitPercent
		if YGFloatIsUndefined(margin) {
			node.style.margin[edge].unit = UnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetMargin gets margin
func YGNodeStyleGetMargin(node *YGNode, edge Edge) YGValue {
	return node.style.margin[edge]
}

// YGNodeStyleSetMarginAuto sets margin auto
func YGNodeStyleSetMarginAuto(node *YGNode, edge Edge) {
	if node.style.margin[edge].unit != UnitAuto {
		node.style.margin[edge].value = YGUndefined
		node.style.margin[edge].unit = UnitAuto
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetPadding sets padding
func YGNodeStyleSetPadding(node *YGNode, edge Edge, padding float32) {
	if node.style.padding[edge].value != padding ||
		node.style.padding[edge].unit != UnitPoint {
		node.style.padding[edge].value = padding
		node.style.padding[edge].unit = UnitPoint
		if YGFloatIsUndefined(padding) {
			node.style.padding[edge].unit = UnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetPaddingPercent sets padding percent
func YGNodeStyleSetPaddingPercent(node *YGNode, edge Edge, padding float32) {
	if node.style.padding[edge].value != padding ||
		node.style.padding[edge].unit != UnitPercent {
		node.style.padding[edge].value = padding
		node.style.padding[edge].unit = UnitPercent
		if YGFloatIsUndefined(padding) {
			node.style.padding[edge].unit = UnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetPadding gets padding
func YGNodeStyleGetPadding(node *YGNode, edge Edge) YGValue {
	return node.style.padding[edge]
}

// YGNodeStyleSetBorder sets border
func YGNodeStyleSetBorder(node *YGNode, edge Edge, border float32) {
	if node.style.border[edge].value != border ||
		node.style.border[edge].unit != UnitPoint {
		node.style.border[edge].value = border
		node.style.border[edge].unit = UnitPoint
		if YGFloatIsUndefined(border) {
			node.style.border[edge].unit = UnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetBorder gets border
func YGNodeStyleGetBorder(node *YGNode, edge Edge) float32 {
	return node.style.border[edge].value
}

// YGNodeStyleSetMinWidth sets min width
func YGNodeStyleSetMinWidth(node *YGNode, minWidth float32) {
	if node.style.minDimensions[DimensionWidth].value != minWidth ||
		node.style.minDimensions[DimensionWidth].unit != UnitPoint {
		node.style.minDimensions[DimensionWidth].value = minWidth
		node.style.minDimensions[DimensionWidth].unit = UnitPoint
		if YGFloatIsUndefined(minWidth) {
			node.style.minDimensions[DimensionWidth].unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMinWidthPercent sets width percent
func YGNodeStyleSetMinWidthPercent(node *YGNode, minWidth float32) {
	if node.style.minDimensions[DimensionWidth].value != minWidth ||
		node.style.minDimensions[DimensionWidth].unit != UnitPercent {
		node.style.minDimensions[DimensionWidth].value = minWidth
		node.style.minDimensions[DimensionWidth].unit = UnitPercent
		if YGFloatIsUndefined(minWidth) {
			node.style.minDimensions[DimensionWidth].unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetMinWidth gets min width
func YGNodeStyleGetMinWidth(node *YGNode) YGValue {
	return node.style.minDimensions[DimensionWidth]
}

// YGNodeStyleSetMinHeight sets min width
func YGNodeStyleSetMinHeight(node *YGNode, minHeight float32) {
	if node.style.minDimensions[DimensionHeight].value != minHeight ||
		node.style.minDimensions[DimensionHeight].unit != UnitPoint {
		node.style.minDimensions[DimensionHeight].value = minHeight
		node.style.minDimensions[DimensionHeight].unit = UnitPoint
		if YGFloatIsUndefined(minHeight) {
			node.style.minDimensions[DimensionHeight].unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMinHeightPercent sets min height percent
func YGNodeStyleSetMinHeightPercent(node *YGNode, minHeight float32) {
	if node.style.minDimensions[DimensionHeight].value != minHeight ||
		node.style.minDimensions[DimensionHeight].unit != UnitPercent {
		node.style.minDimensions[DimensionHeight].value = minHeight
		node.style.minDimensions[DimensionHeight].unit = UnitPercent
		if YGFloatIsUndefined(minHeight) {
			node.style.minDimensions[DimensionHeight].unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetMinHeight gets min height
func YGNodeStyleGetMinHeight(node *YGNode) YGValue {
	return node.style.minDimensions[DimensionHeight]
}

// YGNodeStyleSetMaxWidth sets max width
func YGNodeStyleSetMaxWidth(node *YGNode, maxWidth float32) {
	if node.style.maxDimensions[DimensionWidth].value != maxWidth ||
		node.style.maxDimensions[DimensionWidth].unit != UnitPoint {
		node.style.maxDimensions[DimensionWidth].value = maxWidth
		node.style.maxDimensions[DimensionWidth].unit = UnitPoint
		if YGFloatIsUndefined(maxWidth) {
			node.style.maxDimensions[DimensionWidth].unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMaxWidthPercent sets max width percent
func YGNodeStyleSetMaxWidthPercent(node *YGNode, maxWidth float32) {
	if node.style.maxDimensions[DimensionWidth].value != maxWidth ||
		node.style.maxDimensions[DimensionWidth].unit != UnitPercent {
		node.style.maxDimensions[DimensionWidth].value = maxWidth
		node.style.maxDimensions[DimensionWidth].unit = UnitPercent
		if YGFloatIsUndefined(maxWidth) {
			node.style.maxDimensions[DimensionWidth].unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetMaxWidth gets max width
func YGNodeStyleGetMaxWidth(node *YGNode) YGValue {
	return node.style.maxDimensions[DimensionWidth]
}

// YGNodeStyleSetMaxHeight sets max width
func YGNodeStyleSetMaxHeight(node *YGNode, maxHeight float32) {
	if node.style.maxDimensions[DimensionHeight].value != maxHeight ||
		node.style.maxDimensions[DimensionHeight].unit != UnitPoint {
		node.style.maxDimensions[DimensionHeight].value = maxHeight
		node.style.maxDimensions[DimensionHeight].unit = UnitPoint
		if YGFloatIsUndefined(maxHeight) {
			node.style.maxDimensions[DimensionHeight].unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMaxHeightPercent sets max height percent
func YGNodeStyleSetMaxHeightPercent(node *YGNode, maxHeight float32) {
	if node.style.maxDimensions[DimensionHeight].value != maxHeight ||
		node.style.maxDimensions[DimensionHeight].unit != UnitPercent {
		node.style.maxDimensions[DimensionHeight].value = maxHeight
		node.style.maxDimensions[DimensionHeight].unit = UnitPercent
		if YGFloatIsUndefined(maxHeight) {
			node.style.maxDimensions[DimensionHeight].unit = UnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetMaxHeight gets max height
func YGNodeStyleGetMaxHeight(node *YGNode) YGValue {
	return node.style.maxDimensions[DimensionHeight]
}

// YGNodeStyleSetAspectRatio sets axpect ratio
func YGNodeStyleSetAspectRatio(node *YGNode, aspectRatio float32) {
	if node.style.aspectRatio != aspectRatio {
		node.style.aspectRatio = aspectRatio
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleGetAspectRatio gets aspect ratio
func YGNodeStyleGetAspectRatio(node *YGNode) float32 {
	return node.style.aspectRatio
}

// YGNodeLayoutGetLeft gets left
func YGNodeLayoutGetLeft(node *YGNode) float32 {
	return node.layout.position[EdgeLeft]
}

// YGNodeLayoutGetTop gets top
func YGNodeLayoutGetTop(node *YGNode) float32 {
	return node.layout.position[EdgeTop]
}

// YGNodeLayoutGetRight gets right
func YGNodeLayoutGetRight(node *YGNode) float32 {
	return node.layout.position[EdgeRight]
}

// YGNodeLayoutGetBottom gets bottom
func YGNodeLayoutGetBottom(node *YGNode) float32 {
	return node.layout.position[EdgeBottom]
}

// YGNodeLayoutGetWidth gets width
func YGNodeLayoutGetWidth(node *YGNode) float32 {
	return node.layout.dimensions[DimensionWidth]
}

// YGNodeLayoutGetHeight gets height
func YGNodeLayoutGetHeight(node *YGNode) float32 {
	return node.layout.dimensions[DimensionHeight]
}

// YGNodeLayoutGetDirection gets direction
func YGNodeLayoutGetDirection(node *YGNode) Direction {
	return node.layout.direction
}

// YGNodeLayoutGetHadOverflow gets hadOverflow
func YGNodeLayoutGetHadOverflow(node *YGNode) bool {
	return node.layout.hadOverflow
}

// YGNodeLayoutGetMargin gets margin
func YGNodeLayoutGetMargin(node *YGNode, edge Edge) float32 {
	YGAssertWithNode(node, edge < EdgeEnd, "Cannot get layout properties of multi-edge shorthands")
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
func YGNodeLayoutGetBorder(node *YGNode, edge Edge) float32 {
	YGAssertWithNode(node, edge < EdgeEnd,
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
func YGNodeLayoutGetPadding(node *YGNode, edge Edge) float32 {
	YGAssertWithNode(node, edge < EdgeEnd,
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
