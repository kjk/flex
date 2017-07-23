package flex

import "fmt"

func YGIndent(node *Node, n int) {
	for i := 0; i < n; i++ {
		log(node, LogLevelDebug, "  ")
	}
}

func YGPrintNumberIfNotUndefinedf(node *Node, str string, number float32) {
	if !YGFloatIsUndefined(number) {
		log(node, LogLevelDebug, "%s: %g; ", str, number)
	}
}

func YGPrintNumberIfNotUndefined(node *Node, str string, number *Value) {
	if number.Unit != UnitUndefined {
		if number.Unit == UnitAuto {
			log(node, LogLevelDebug, "%s: auto; ", str)
		} else {
			unit := "%"

			if number.Unit == UnitPoint {
				unit = "px"
			}
			log(node, LogLevelDebug, "%s: %g%s; ", str, number.Value, unit)
		}
	}
}
func YGPrintNumberIfNotAuto(node *Node, str string, number *Value) {
	if number.Unit != UnitAuto {
		YGPrintNumberIfNotUndefined(node, str, number)
	}
}

func YGPrintEdgeIfNotUndefined(node *Node, str string, edges []Value, edge Edge) {
	YGPrintNumberIfNotUndefined(node, str, computedEdgeValue(edges, edge, &ValueUndefined))
}

func YGPrintNumberIfNotZero(node *Node, str string, number *Value) {
	if !YGFloatsEqual(number.Value, 0) {
		YGPrintNumberIfNotUndefined(node, str, number)
	}
}

func YGFourValuesEqual(four []Value) bool {
	return YGValueEqual(four[0], four[1]) && YGValueEqual(four[0], four[2]) &&
		YGValueEqual(four[0], four[3])
}

func YGPrintEdges(node *Node, str string, edges []Value) {
	if YGFourValuesEqual(edges) {
		YGPrintNumberIfNotZero(node, str, &edges[EdgeLeft])
	} else {
		for edge := EdgeLeft; edge < EdgeCount; edge++ {
			buf := fmt.Sprintf("%s-%s", str, YGEdgeToString(edge))
			YGPrintNumberIfNotZero(node, buf, &edges[edge])
		}
	}
}

func YGNodePrintInternal(node *Node, options PrintOptions, level int) {
	YGIndent(node, level)
	log(node, LogLevelDebug, "<div ")

	if node.Print != nil {
		node.Print(node)
	}

	if options&PrintOptionsLayout != 0 {
		log(node, LogLevelDebug, "layout=\"")
		log(node, LogLevelDebug, "width: %g; ", node.Layout.dimensions[DimensionWidth])
		log(node, LogLevelDebug, "height: %g; ", node.Layout.dimensions[DimensionHeight])
		log(node, LogLevelDebug, "top: %g; ", node.Layout.position[EdgeTop])
		log(node, LogLevelDebug, "left: %g;", node.Layout.position[EdgeLeft])
		log(node, LogLevelDebug, "\" ")
	}

	if options&PrintOptionsStyle != 0 {
		log(node, LogLevelDebug, "style=\"")
		if node.Style.flexDirection != gYGNodeDefaults.Style.flexDirection {
			log(node,
				LogLevelDebug,
				"flex-direction: %s; ",
				YGFlexDirectionToString(node.Style.flexDirection))
		}
		if node.Style.justifyContent != gYGNodeDefaults.Style.justifyContent {
			log(node,
				LogLevelDebug,
				"justify-content: %s; ",
				YGJustifyToString(node.Style.justifyContent))
		}
		if node.Style.alignItems != gYGNodeDefaults.Style.alignItems {
			log(node, LogLevelDebug, "align-items: %s; ", YGAlignToString(node.Style.alignItems))
		}
		if node.Style.alignContent != gYGNodeDefaults.Style.alignContent {
			log(node, LogLevelDebug, "align-content: %s; ", YGAlignToString(node.Style.alignContent))
		}
		if node.Style.alignSelf != gYGNodeDefaults.Style.alignSelf {
			log(node, LogLevelDebug, "align-self: %s; ", YGAlignToString(node.Style.alignSelf))
		}

		YGPrintNumberIfNotUndefinedf(node, "flex-grow", node.Style.flexGrow)
		YGPrintNumberIfNotUndefinedf(node, "flex-shrink", node.Style.flexShrink)
		YGPrintNumberIfNotAuto(node, "flex-basis", &node.Style.flexBasis)
		YGPrintNumberIfNotUndefinedf(node, "flex", node.Style.flex)

		if node.Style.flexWrap != gYGNodeDefaults.Style.flexWrap {
			log(node, LogLevelDebug, "flexWrap: %s; ", YGWrapToString(node.Style.flexWrap))
		}

		if node.Style.overflow != gYGNodeDefaults.Style.overflow {
			log(node, LogLevelDebug, "overflow: %s; ", YGOverflowToString(node.Style.overflow))
		}

		if node.Style.display != gYGNodeDefaults.Style.display {
			log(node, LogLevelDebug, "display: %s; ", YGDisplayToString(node.Style.display))
		}

		YGPrintEdges(node, "margin", node.Style.margin[:])
		YGPrintEdges(node, "padding", node.Style.padding[:])
		YGPrintEdges(node, "border", node.Style.border[:])

		YGPrintNumberIfNotAuto(node, "width", &node.Style.dimensions[DimensionWidth])
		YGPrintNumberIfNotAuto(node, "height", &node.Style.dimensions[DimensionHeight])
		YGPrintNumberIfNotAuto(node, "max-width", &node.Style.maxDimensions[DimensionWidth])
		YGPrintNumberIfNotAuto(node, "max-height", &node.Style.maxDimensions[DimensionHeight])
		YGPrintNumberIfNotAuto(node, "min-width", &node.Style.minDimensions[DimensionWidth])
		YGPrintNumberIfNotAuto(node, "min-height", &node.Style.minDimensions[DimensionHeight])

		if node.Style.positionType != gYGNodeDefaults.Style.positionType {
			log(node,
				LogLevelDebug,
				"position: %s; ",
				YGPositionTypeToString(node.Style.positionType))
		}

		YGPrintEdgeIfNotUndefined(node, "left", node.Style.position[:], EdgeLeft)
		YGPrintEdgeIfNotUndefined(node, "right", node.Style.position[:], EdgeRight)
		YGPrintEdgeIfNotUndefined(node, "top", node.Style.position[:], EdgeTop)
		YGPrintEdgeIfNotUndefined(node, "bottom", node.Style.position[:], EdgeBottom)
		log(node, LogLevelDebug, "\" ")

		if node.Measure != nil {
			log(node, LogLevelDebug, "has-custom-measure=\"true\"")
		}
	}
	log(node, LogLevelDebug, ">")

	childCount := YGNodeListCount(node.Children)
	if options&PrintOptionsChildren != 0 && childCount > 0 {
		for i := 0; i < childCount; i++ {
			log(node, LogLevelDebug, "\n")
			YGNodePrintInternal(YGNodeGetChild(node, i), options, level+1)
		}
		YGIndent(node, level)
		log(node, LogLevelDebug, "\n")
	}
	log(node, LogLevelDebug, "</div>")
}

func YGNodePrint(node *Node, options PrintOptions) {
	YGNodePrintInternal(node, options, 0)
}
