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

// NodeStyleSetWidth sets width
func NodeStyleSetWidth(node *Node, width float32) {
	dim := &node.Style.Dimensions[DimensionWidth]
	if dim.Value != width || dim.Unit != UnitPoint {
		dim.Value = width
		dim.Unit = UnitPoint
		if FloatIsUndefined(width) {
			dim.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetWidthPercent sets width percent
func NodeStyleSetWidthPercent(node *Node, width float32) {
	dim := &node.Style.Dimensions[DimensionWidth]
	if dim.Value != width || dim.Unit != UnitPercent {
		dim.Value = width
		dim.Unit = UnitPercent
		if FloatIsUndefined(width) {
			dim.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetWidthAuto sets width auto
func NodeStyleSetWidthAuto(node *Node) {
	dim := &node.Style.Dimensions[DimensionWidth]
	if dim.Unit != UnitAuto {
		dim.Value = Undefined
		dim.Unit = UnitAuto
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleGetWidth gets width
func NodeStyleGetWidth(node *Node) Value {
	return node.Style.Dimensions[DimensionWidth]
}

// NodeStyleSetHeight sets height
func NodeStyleSetHeight(node *Node, height float32) {
	dim := &node.Style.Dimensions[DimensionHeight]
	if dim.Value != height || dim.Unit != UnitPoint {
		dim.Value = height
		dim.Unit = UnitPoint
		if FloatIsUndefined(height) {
			dim.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetHeightPercent sets height percent
func NodeStyleSetHeightPercent(node *Node, height float32) {
	dim := &node.Style.Dimensions[DimensionHeight]
	if dim.Value != height || dim.Unit != UnitPercent {
		dim.Value = height
		dim.Unit = UnitPercent
		if FloatIsUndefined(height) {
			dim.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetHeightAuto sets height auto
func NodeStyleSetHeightAuto(node *Node) {
	dim := &node.Style.Dimensions[DimensionHeight]
	if dim.Unit != UnitAuto {
		dim.Value = Undefined
		dim.Unit = UnitAuto
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleGetHeight gets height
func NodeStyleGetHeight(node *Node) Value {
	return node.Style.Dimensions[DimensionHeight]
}

// StyleSetPositionType sets position type
func (node *Node) StyleSetPositionType(positionType PositionType) {
	if node.Style.PositionType != positionType {
		node.Style.PositionType = positionType
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetPosition sets position
func NodeStyleSetPosition(node *Node, edge Edge, position float32) {
	pos := &node.Style.Position[edge]
	if pos.Value != position || pos.Unit != UnitPoint {
		pos.Value = position
		pos.Unit = UnitPoint
		if FloatIsUndefined(position) {
			pos.Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetPositionPercent sets position percent
func NodeStyleSetPositionPercent(node *Node, edge Edge, position float32) {
	pos := &node.Style.Position[edge]
	if pos.Value != position || pos.Unit != UnitPercent {
		pos.Value = position
		pos.Unit = UnitPercent
		if FloatIsUndefined(position) {
			pos.Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// StyleGetPosition gets position
func (node *Node) StyleGetPosition(edge Edge) Value {
	return node.Style.Position[edge]
}

// NodeStyleSetDirection sets direction
func NodeStyleSetDirection(node *Node, direction Direction) {
	if node.Style.Direction != direction {
		node.Style.Direction = direction
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetFlexDirection sets flex directions
func NodeStyleSetFlexDirection(node *Node, flexDirection FlexDirection) {
	if node.Style.FlexDirection != flexDirection {
		node.Style.FlexDirection = flexDirection
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetJustifyContent sets justify content
func NodeStyleSetJustifyContent(node *Node, justifyContent Justify) {
	if node.Style.JustifyContent != justifyContent {
		node.Style.JustifyContent = justifyContent
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetAlignContent sets align content
func NodeStyleSetAlignContent(node *Node, alignContent Align) {
	if node.Style.AlignContent != alignContent {
		node.Style.AlignContent = alignContent
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetAlignItems sets align content
func NodeStyleSetAlignItems(node *Node, alignItems Align) {
	if node.Style.AlignItems != alignItems {
		node.Style.AlignItems = alignItems
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetAlignSelf sets align self
func NodeStyleSetAlignSelf(node *Node, alignSelf Align) {
	if node.Style.AlignSelf != alignSelf {
		node.Style.AlignSelf = alignSelf
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetFlexWrap sets flex wrap
func NodeStyleSetFlexWrap(node *Node, flexWrap Wrap) {
	if node.Style.FlexWrap != flexWrap {
		node.Style.FlexWrap = flexWrap
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetOverflow sets overflow
func NodeStyleSetOverflow(node *Node, overflow Overflow) {
	if node.Style.Overflow != overflow {
		node.Style.Overflow = overflow
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetDisplay sets display
func NodeStyleSetDisplay(node *Node, display Display) {
	if node.Style.Display != display {
		node.Style.Display = display
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetFlex sets flex
func NodeStyleSetFlex(node *Node, flex float32) {
	if node.Style.Flex != flex {
		node.Style.Flex = flex
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleGetFlex gets flex
func NodeStyleGetFlex(node *Node) float32 {
	return node.Style.Flex
}

// NodeStyleSetFlexGrow sets flex grow
func NodeStyleSetFlexGrow(node *Node, flexGrow float32) {
	if node.Style.FlexGrow != flexGrow {
		node.Style.FlexGrow = flexGrow
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetFlexShrink sets flex shrink
func NodeStyleSetFlexShrink(node *Node, flexShrink float32) {
	if node.Style.FlexShrink != flexShrink {
		node.Style.FlexShrink = flexShrink
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetFlexBasis sets flex basis
func NodeStyleSetFlexBasis(node *Node, flexBasis float32) {
	if node.Style.FlexBasis.Value != flexBasis ||
		node.Style.FlexBasis.Unit != UnitPoint {
		node.Style.FlexBasis.Value = flexBasis
		node.Style.FlexBasis.Unit = UnitPoint
		if FloatIsUndefined(flexBasis) {
			node.Style.FlexBasis.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetFlexBasisPercent sets flex basis percent
func NodeStyleSetFlexBasisPercent(node *Node, flexBasis float32) {
	if node.Style.FlexBasis.Value != flexBasis ||
		node.Style.FlexBasis.Unit != UnitPercent {
		node.Style.FlexBasis.Value = flexBasis
		node.Style.FlexBasis.Unit = UnitPercent
		if FloatIsUndefined(flexBasis) {
			node.Style.FlexBasis.Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetFlexBasisAuto sets flex basis auto
func NodeStyleSetFlexBasisAuto(node *Node) {
	if node.Style.FlexBasis.Unit != UnitAuto {
		node.Style.FlexBasis.Value = Undefined
		node.Style.FlexBasis.Unit = UnitAuto
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetMargin sets margin
func NodeStyleSetMargin(node *Node, edge Edge, margin float32) {
	if node.Style.Margin[edge].Value != margin ||
		node.Style.Margin[edge].Unit != UnitPoint {
		node.Style.Margin[edge].Value = margin
		node.Style.Margin[edge].Unit = UnitPoint
		if FloatIsUndefined(margin) {
			node.Style.Margin[edge].Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// YGNodeStyleSetMarginPercent sets margin percent
func YGNodeStyleSetMarginPercent(node *Node, edge Edge, margin float32) {
	if node.Style.Margin[edge].Value != margin ||
		node.Style.Margin[edge].Unit != UnitPercent {
		node.Style.Margin[edge].Value = margin
		node.Style.Margin[edge].Unit = UnitPercent
		if FloatIsUndefined(margin) {
			node.Style.Margin[edge].Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleGetMargin gets margin
func NodeStyleGetMargin(node *Node, edge Edge) Value {
	return node.Style.Margin[edge]
}

// NodeStyleSetMarginAuto sets margin auto
func NodeStyleSetMarginAuto(node *Node, edge Edge) {
	if node.Style.Margin[edge].Unit != UnitAuto {
		node.Style.Margin[edge].Value = Undefined
		node.Style.Margin[edge].Unit = UnitAuto
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetPadding sets padding
func NodeStyleSetPadding(node *Node, edge Edge, padding float32) {
	if node.Style.Padding[edge].Value != padding ||
		node.Style.Padding[edge].Unit != UnitPoint {
		node.Style.Padding[edge].Value = padding
		node.Style.Padding[edge].Unit = UnitPoint
		if FloatIsUndefined(padding) {
			node.Style.Padding[edge].Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetPaddingPercent sets padding percent
func NodeStyleSetPaddingPercent(node *Node, edge Edge, padding float32) {
	if node.Style.Padding[edge].Value != padding ||
		node.Style.Padding[edge].Unit != UnitPercent {
		node.Style.Padding[edge].Value = padding
		node.Style.Padding[edge].Unit = UnitPercent
		if FloatIsUndefined(padding) {
			node.Style.Padding[edge].Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleGetPadding gets padding
func NodeStyleGetPadding(node *Node, edge Edge) Value {
	return node.Style.Padding[edge]
}

// NodeStyleSetBorder sets border
func NodeStyleSetBorder(node *Node, edge Edge, border float32) {
	if node.Style.Border[edge].Value != border ||
		node.Style.Border[edge].Unit != UnitPoint {
		node.Style.Border[edge].Value = border
		node.Style.Border[edge].Unit = UnitPoint
		if FloatIsUndefined(border) {
			node.Style.Border[edge].Unit = UnitUndefined
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleGetBorder gets border
func NodeStyleGetBorder(node *Node, edge Edge) float32 {
	return node.Style.Border[edge].Value
}

// NodeStyleSetMinWidth sets min width
func NodeStyleSetMinWidth(node *Node, minWidth float32) {
	if node.Style.MinDimensions[DimensionWidth].Value != minWidth ||
		node.Style.MinDimensions[DimensionWidth].Unit != UnitPoint {
		node.Style.MinDimensions[DimensionWidth].Value = minWidth
		node.Style.MinDimensions[DimensionWidth].Unit = UnitPoint
		if FloatIsUndefined(minWidth) {
			node.Style.MinDimensions[DimensionWidth].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetMinWidthPercent sets width percent
func NodeStyleSetMinWidthPercent(node *Node, minWidth float32) {
	if node.Style.MinDimensions[DimensionWidth].Value != minWidth ||
		node.Style.MinDimensions[DimensionWidth].Unit != UnitPercent {
		node.Style.MinDimensions[DimensionWidth].Value = minWidth
		node.Style.MinDimensions[DimensionWidth].Unit = UnitPercent
		if FloatIsUndefined(minWidth) {
			node.Style.MinDimensions[DimensionWidth].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleGetMinWidth gets min width
func NodeStyleGetMinWidth(node *Node) Value {
	return node.Style.MinDimensions[DimensionWidth]
}

// NodeStyleSetMinHeight sets min width
func NodeStyleSetMinHeight(node *Node, minHeight float32) {
	if node.Style.MinDimensions[DimensionHeight].Value != minHeight ||
		node.Style.MinDimensions[DimensionHeight].Unit != UnitPoint {
		node.Style.MinDimensions[DimensionHeight].Value = minHeight
		node.Style.MinDimensions[DimensionHeight].Unit = UnitPoint
		if FloatIsUndefined(minHeight) {
			node.Style.MinDimensions[DimensionHeight].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetMinHeightPercent sets min height percent
func NodeStyleSetMinHeightPercent(node *Node, minHeight float32) {
	if node.Style.MinDimensions[DimensionHeight].Value != minHeight ||
		node.Style.MinDimensions[DimensionHeight].Unit != UnitPercent {
		node.Style.MinDimensions[DimensionHeight].Value = minHeight
		node.Style.MinDimensions[DimensionHeight].Unit = UnitPercent
		if FloatIsUndefined(minHeight) {
			node.Style.MinDimensions[DimensionHeight].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleGetMinHeight gets min height
func NodeStyleGetMinHeight(node *Node) Value {
	return node.Style.MinDimensions[DimensionHeight]
}

// NodeStyleSetMaxWidth sets max width
func NodeStyleSetMaxWidth(node *Node, maxWidth float32) {
	if node.Style.MaxDimensions[DimensionWidth].Value != maxWidth ||
		node.Style.MaxDimensions[DimensionWidth].Unit != UnitPoint {
		node.Style.MaxDimensions[DimensionWidth].Value = maxWidth
		node.Style.MaxDimensions[DimensionWidth].Unit = UnitPoint
		if FloatIsUndefined(maxWidth) {
			node.Style.MaxDimensions[DimensionWidth].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetMaxWidthPercent sets max width percent
func NodeStyleSetMaxWidthPercent(node *Node, maxWidth float32) {
	if node.Style.MaxDimensions[DimensionWidth].Value != maxWidth ||
		node.Style.MaxDimensions[DimensionWidth].Unit != UnitPercent {
		node.Style.MaxDimensions[DimensionWidth].Value = maxWidth
		node.Style.MaxDimensions[DimensionWidth].Unit = UnitPercent
		if FloatIsUndefined(maxWidth) {
			node.Style.MaxDimensions[DimensionWidth].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleGetMaxWidth gets max width
func NodeStyleGetMaxWidth(node *Node) Value {
	return node.Style.MaxDimensions[DimensionWidth]
}

// NodeStyleSetMaxHeight sets max width
func NodeStyleSetMaxHeight(node *Node, maxHeight float32) {
	if node.Style.MaxDimensions[DimensionHeight].Value != maxHeight ||
		node.Style.MaxDimensions[DimensionHeight].Unit != UnitPoint {
		node.Style.MaxDimensions[DimensionHeight].Value = maxHeight
		node.Style.MaxDimensions[DimensionHeight].Unit = UnitPoint
		if FloatIsUndefined(maxHeight) {
			node.Style.MaxDimensions[DimensionHeight].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleSetMaxHeightPercent sets max height percent
func NodeStyleSetMaxHeightPercent(node *Node, maxHeight float32) {
	if node.Style.MaxDimensions[DimensionHeight].Value != maxHeight ||
		node.Style.MaxDimensions[DimensionHeight].Unit != UnitPercent {
		node.Style.MaxDimensions[DimensionHeight].Value = maxHeight
		node.Style.MaxDimensions[DimensionHeight].Unit = UnitPercent
		if FloatIsUndefined(maxHeight) {
			node.Style.MaxDimensions[DimensionHeight].Unit = UnitAuto
		}
		nodeMarkDirtyInternal(node)
	}
}

// NodeStyleGetMaxHeight gets max height
func NodeStyleGetMaxHeight(node *Node) Value {
	return node.Style.MaxDimensions[DimensionHeight]
}

// NodeStyleSetAspectRatio sets axpect ratio
func NodeStyleSetAspectRatio(node *Node, aspectRatio float32) {
	if node.Style.AspectRatio != aspectRatio {
		node.Style.AspectRatio = aspectRatio
		nodeMarkDirtyInternal(node)
	}
}

// NodeLayoutGetLeft gets left
func NodeLayoutGetLeft(node *Node) float32 {
	return node.Layout.Position[EdgeLeft]
}

// NodeLayoutGetTop gets top
func NodeLayoutGetTop(node *Node) float32 {
	return node.Layout.Position[EdgeTop]
}

// NodeLayoutGetRight gets right
func NodeLayoutGetRight(node *Node) float32 {
	return node.Layout.Position[EdgeRight]
}

// NodeLayoutGetBottom gets bottom
func NodeLayoutGetBottom(node *Node) float32 {
	return node.Layout.Position[EdgeBottom]
}

// NodeLayoutGetWidth gets width
func NodeLayoutGetWidth(node *Node) float32 {
	return node.Layout.Dimensions[DimensionWidth]
}

// NodeLayoutGetHeight gets height
func NodeLayoutGetHeight(node *Node) float32 {
	return node.Layout.Dimensions[DimensionHeight]
}

// NodeLayoutGetMargin gets margin
func NodeLayoutGetMargin(node *Node, edge Edge) float32 {
	assertWithNode(node, edge < EdgeEnd, "Cannot get layout properties of multi-edge shorthands")
	if edge == EdgeLeft {
		if node.Layout.Direction == DirectionRTL {
			return node.Layout.Margin[EdgeEnd]
		}
		return node.Layout.Margin[EdgeStart]
	}
	if edge == EdgeRight {
		if node.Layout.Direction == DirectionRTL {
			return node.Layout.Margin[EdgeStart]
		}
		return node.Layout.Margin[EdgeEnd]
	}
	return node.Layout.Margin[edge]
}

// NodeLayoutGetBorder gets border
func NodeLayoutGetBorder(node *Node, edge Edge) float32 {
	assertWithNode(node, edge < EdgeEnd,
		"Cannot get layout properties of multi-edge shorthands")
	if edge == EdgeLeft {
		if node.Layout.Direction == DirectionRTL {
			return node.Layout.Border[EdgeEnd]
		}
		return node.Layout.Border[EdgeStart]
	}
	if edge == EdgeRight {
		if node.Layout.Direction == DirectionRTL {
			return node.Layout.Border[EdgeStart]
		}
		return node.Layout.Border[EdgeEnd]
	}
	return node.Layout.Border[edge]
}

// NodeLayoutGetPadding gets padding
func NodeLayoutGetPadding(node *Node, edge Edge) float32 {
	assertWithNode(node, edge < EdgeEnd,
		"Cannot get layout properties of multi-edge shorthands")
	if edge == EdgeLeft {
		if node.Layout.Direction == DirectionRTL {
			return node.Layout.Padding[EdgeEnd]
		}
		return node.Layout.Padding[EdgeStart]
	}
	if edge == EdgeRight {
		if node.Layout.Direction == DirectionRTL {
			return node.Layout.Padding[EdgeStart]
		}
		return node.Layout.Padding[EdgeEnd]
	}
	return node.Layout.Padding[edge]
}
