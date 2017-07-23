package flex

import "fmt"

func YGIndent(node *YGNode, n int) {
	for i := 0; i < n; i++ {
		YGLog(node, LogLevelDebug, "  ")
	}
}

func YGPrintNumberIfNotUndefinedf(node *YGNode, str string, number float32) {
	if !YGFloatIsUndefined(number) {
		YGLog(node, LogLevelDebug, "%s: %g; ", str, number)
	}
}

func YGPrintNumberIfNotUndefined(node *YGNode, str string, number *YGValue) {
	if number.unit != UnitUndefined {
		if number.unit == UnitAuto {
			YGLog(node, LogLevelDebug, "%s: auto; ", str)
		} else {
			unit := "%"

			if number.unit == UnitPoint {
				unit = "px"
			}
			YGLog(node, LogLevelDebug, "%s: %g%s; ", str, number.value, unit)
		}
	}
}
func YGPrintNumberIfNotAuto(node *YGNode, str string, number *YGValue) {
	if number.unit != UnitAuto {
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

func YGNodePrintInternal(node *YGNode, options PrintOptions, level int) {
	YGIndent(node, level)
	YGLog(node, LogLevelDebug, "<div ")

	if node.print != nil {
		node.print(node)
	}

	if options&PrintOptionsLayout != 0 {
		YGLog(node, LogLevelDebug, "layout=\"")
		YGLog(node, LogLevelDebug, "width: %g; ", node.layout.dimensions[DimensionWidth])
		YGLog(node, LogLevelDebug, "height: %g; ", node.layout.dimensions[DimensionHeight])
		YGLog(node, LogLevelDebug, "top: %g; ", node.layout.position[EdgeTop])
		YGLog(node, LogLevelDebug, "left: %g;", node.layout.position[EdgeLeft])
		YGLog(node, LogLevelDebug, "\" ")
	}

	if options&PrintOptionsStyle != 0 {
		YGLog(node, LogLevelDebug, "style=\"")
		if node.style.flexDirection != gYGNodeDefaults.style.flexDirection {
			YGLog(node,
				LogLevelDebug,
				"flex-direction: %s; ",
				YGFlexDirectionToString(node.style.flexDirection))
		}
		if node.style.justifyContent != gYGNodeDefaults.style.justifyContent {
			YGLog(node,
				LogLevelDebug,
				"justify-content: %s; ",
				YGJustifyToString(node.style.justifyContent))
		}
		if node.style.alignItems != gYGNodeDefaults.style.alignItems {
			YGLog(node, LogLevelDebug, "align-items: %s; ", YGAlignToString(node.style.alignItems))
		}
		if node.style.alignContent != gYGNodeDefaults.style.alignContent {
			YGLog(node, LogLevelDebug, "align-content: %s; ", YGAlignToString(node.style.alignContent))
		}
		if node.style.alignSelf != gYGNodeDefaults.style.alignSelf {
			YGLog(node, LogLevelDebug, "align-self: %s; ", YGAlignToString(node.style.alignSelf))
		}

		YGPrintNumberIfNotUndefinedf(node, "flex-grow", node.style.flexGrow)
		YGPrintNumberIfNotUndefinedf(node, "flex-shrink", node.style.flexShrink)
		YGPrintNumberIfNotAuto(node, "flex-basis", &node.style.flexBasis)
		YGPrintNumberIfNotUndefinedf(node, "flex", node.style.flex)

		if node.style.flexWrap != gYGNodeDefaults.style.flexWrap {
			YGLog(node, LogLevelDebug, "flexWrap: %s; ", YGWrapToString(node.style.flexWrap))
		}

		if node.style.overflow != gYGNodeDefaults.style.overflow {
			YGLog(node, LogLevelDebug, "overflow: %s; ", YGOverflowToString(node.style.overflow))
		}

		if node.style.display != gYGNodeDefaults.style.display {
			YGLog(node, LogLevelDebug, "display: %s; ", YGDisplayToString(node.style.display))
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
				LogLevelDebug,
				"position: %s; ",
				YGPositionTypeToString(node.style.positionType))
		}

		YGPrintEdgeIfNotUndefined(node, "left", node.style.position[:], EdgeLeft)
		YGPrintEdgeIfNotUndefined(node, "right", node.style.position[:], EdgeRight)
		YGPrintEdgeIfNotUndefined(node, "top", node.style.position[:], EdgeTop)
		YGPrintEdgeIfNotUndefined(node, "bottom", node.style.position[:], EdgeBottom)
		YGLog(node, LogLevelDebug, "\" ")

		if node.measure != nil {
			YGLog(node, LogLevelDebug, "has-custom-measure=\"true\"")
		}
	}
	YGLog(node, LogLevelDebug, ">")

	childCount := YGNodeListCount(node.children)
	if options&PrintOptionsChildren != 0 && childCount > 0 {
		for i := 0; i < childCount; i++ {
			YGLog(node, LogLevelDebug, "\n")
			YGNodePrintInternal(YGNodeGetChild(node, i), options, level+1)
		}
		YGIndent(node, level)
		YGLog(node, LogLevelDebug, "\n")
	}
	YGLog(node, LogLevelDebug, "</div>")
}

func YGNodePrint(node *YGNode, options PrintOptions) {
	YGNodePrintInternal(node, options, 0)
}
