package flex

import "fmt"

func YGIndent(node *YGNode, n int) {
	for i := 0; i < n; i++ {
		YGLog(node, YGLogLevelDebug, "  ")
	}
}

func YGPrintNumberIfNotUndefinedf(node *YGNode, str string, number float32) {
	if !YGFloatIsUndefined(number) {
		YGLog(node, YGLogLevelDebug, "%s: %g; ", str, number)
	}
}

func YGPrintNumberIfNotUndefined(node *YGNode, str string, number *YGValue) {
	if number.unit != YGUnitUndefined {
		if number.unit == YGUnitAuto {
			YGLog(node, YGLogLevelDebug, "%s: auto; ", str)
		} else {
			unit := "%"

			if number.unit == YGUnitPoint {
				unit = "px"
			}
			YGLog(node, YGLogLevelDebug, "%s: %g%s; ", str, number.value, unit)
		}
	}
}
func YGPrintNumberIfNotAuto(node *YGNode, str string, number *YGValue) {
	if number.unit != YGUnitAuto {
		YGPrintNumberIfNotUndefined(node, str, number)
	}
}

func YGPrintEdgeIfNotUndefined(node *YGNode, str string, edges []YGValue, edge Edge) {
	YGPrintNumberIfNotUndefined(node, str, YGComputedEdgeValue(edges, edge, &YGValueUndefined))
}

func YGPrintNumberIfNotZero(node *YGNode, str string, number *YGValue) {
	if !YGFloatsEqual(number.value, 0) {
		YGPrintNumberIfNotUndefined(node, str, number)
	}
}

func YGFourValuesEqual(four []YGValue) bool {
	return YGValueEqual(four[0], four[1]) && YGValueEqual(four[0], four[2]) &&
		YGValueEqual(four[0], four[3])
}

func YGPrintEdges(node *YGNode, str string, edges []YGValue) {
	if YGFourValuesEqual(edges) {
		YGPrintNumberIfNotZero(node, str, &edges[EdgeLeft])
	} else {
		for edge := EdgeLeft; edge < EdgeCount; edge++ {
			buf := fmt.Sprintf("%s-%s", str, YGEdgeToString(edge))
			YGPrintNumberIfNotZero(node, buf, &edges[edge])
		}
	}
}

func YGNodePrintInternal(node *YGNode, options YGPrintOptions, level int) {
	YGIndent(node, level)
	YGLog(node, YGLogLevelDebug, "<div ")

	if node.print != nil {
		node.print(node)
	}

	if options&YGPrintOptionsLayout != 0 {
		YGLog(node, YGLogLevelDebug, "layout=\"")
		YGLog(node, YGLogLevelDebug, "width: %g; ", node.layout.dimensions[DimensionWidth])
		YGLog(node, YGLogLevelDebug, "height: %g; ", node.layout.dimensions[DimensionHeight])
		YGLog(node, YGLogLevelDebug, "top: %g; ", node.layout.position[EdgeTop])
		YGLog(node, YGLogLevelDebug, "left: %g;", node.layout.position[EdgeLeft])
		YGLog(node, YGLogLevelDebug, "\" ")
	}

	if options&YGPrintOptionsStyle != 0 {
		YGLog(node, YGLogLevelDebug, "style=\"")
		if node.style.flexDirection != gYGNodeDefaults.style.flexDirection {
			YGLog(node,
				YGLogLevelDebug,
				"flex-direction: %s; ",
				YGFlexDirectionToString(node.style.flexDirection))
		}
		if node.style.justifyContent != gYGNodeDefaults.style.justifyContent {
			YGLog(node,
				YGLogLevelDebug,
				"justify-content: %s; ",
				YGJustifyToString(node.style.justifyContent))
		}
		if node.style.alignItems != gYGNodeDefaults.style.alignItems {
			YGLog(node, YGLogLevelDebug, "align-items: %s; ", YGAlignToString(node.style.alignItems))
		}
		if node.style.alignContent != gYGNodeDefaults.style.alignContent {
			YGLog(node, YGLogLevelDebug, "align-content: %s; ", YGAlignToString(node.style.alignContent))
		}
		if node.style.alignSelf != gYGNodeDefaults.style.alignSelf {
			YGLog(node, YGLogLevelDebug, "align-self: %s; ", YGAlignToString(node.style.alignSelf))
		}

		YGPrintNumberIfNotUndefinedf(node, "flex-grow", node.style.flexGrow)
		YGPrintNumberIfNotUndefinedf(node, "flex-shrink", node.style.flexShrink)
		YGPrintNumberIfNotAuto(node, "flex-basis", &node.style.flexBasis)
		YGPrintNumberIfNotUndefinedf(node, "flex", node.style.flex)

		if node.style.flexWrap != gYGNodeDefaults.style.flexWrap {
			YGLog(node, YGLogLevelDebug, "flexWrap: %s; ", YGWrapToString(node.style.flexWrap))
		}

		if node.style.overflow != gYGNodeDefaults.style.overflow {
			YGLog(node, YGLogLevelDebug, "overflow: %s; ", YGOverflowToString(node.style.overflow))
		}

		if node.style.display != gYGNodeDefaults.style.display {
			YGLog(node, YGLogLevelDebug, "display: %s; ", YGDisplayToString(node.style.display))
		}

		YGPrintEdges(node, "margin", node.style.margin[:])
		YGPrintEdges(node, "padding", node.style.padding[:])
		YGPrintEdges(node, "border", node.style.border[:])

		YGPrintNumberIfNotAuto(node, "width", &node.style.dimensions[DimensionWidth])
		YGPrintNumberIfNotAuto(node, "height", &node.style.dimensions[DimensionHeight])
		YGPrintNumberIfNotAuto(node, "max-width", &node.style.maxDimensions[DimensionWidth])
		YGPrintNumberIfNotAuto(node, "max-height", &node.style.maxDimensions[DimensionHeight])
		YGPrintNumberIfNotAuto(node, "min-width", &node.style.minDimensions[DimensionWidth])
		YGPrintNumberIfNotAuto(node, "min-height", &node.style.minDimensions[DimensionHeight])

		if node.style.positionType != gYGNodeDefaults.style.positionType {
			YGLog(node,
				YGLogLevelDebug,
				"position: %s; ",
				YGPositionTypeToString(node.style.positionType))
		}

		YGPrintEdgeIfNotUndefined(node, "left", node.style.position[:], EdgeLeft)
		YGPrintEdgeIfNotUndefined(node, "right", node.style.position[:], EdgeRight)
		YGPrintEdgeIfNotUndefined(node, "top", node.style.position[:], EdgeTop)
		YGPrintEdgeIfNotUndefined(node, "bottom", node.style.position[:], EdgeBottom)
		YGLog(node, YGLogLevelDebug, "\" ")

		if node.measure != nil {
			YGLog(node, YGLogLevelDebug, "has-custom-measure=\"true\"")
		}
	}
	YGLog(node, YGLogLevelDebug, ">")

	childCount := YGNodeListCount(node.children)
	if options&YGPrintOptionsChildren != 0 && childCount > 0 {
		for i := 0; i < childCount; i++ {
			YGLog(node, YGLogLevelDebug, "\n")
			YGNodePrintInternal(YGNodeGetChild(node, i), options, level+1)
		}
		YGIndent(node, level)
		YGLog(node, YGLogLevelDebug, "\n")
	}
	YGLog(node, YGLogLevelDebug, "</div>")
}

func YGNodePrint(node *YGNode, options YGPrintOptions) {
	YGNodePrintInternal(node, options, 0)
}
