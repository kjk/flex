package flex

import "fmt"

func indent(node *Node, n int) {
	for i := 0; i < n; i++ {
		log(node, LogLevelDebug, "  ")
	}
}

func printNumberIfNotUndefinedf(node *Node, str string, number float32) {
	if !FloatIsUndefined(number) {
		log(node, LogLevelDebug, "%s: %g; ", str, number)
	}
}

func printNumberIfNotUndefined(node *Node, str string, number *Value) {
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
func printNumberIfNotAuto(node *Node, str string, number *Value) {
	if number.Unit != UnitAuto {
		printNumberIfNotUndefined(node, str, number)
	}
}

func printEdgeIfNotUndefined(node *Node, str string, edges []Value, edge Edge) {
	printNumberIfNotUndefined(node, str, computedEdgeValue(edges, edge, &ValueUndefined))
}

func printNumberIfNotZero(node *Node, str string, number *Value) {
	if !FloatsEqual(number.Value, 0) {
		printNumberIfNotUndefined(node, str, number)
	}
}

func fourValuesEqual(four []Value) bool {
	return ValueEqual(four[0], four[1]) && ValueEqual(four[0], four[2]) &&
		ValueEqual(four[0], four[3])
}

func printEdges(node *Node, str string, edges []Value) {
	if fourValuesEqual(edges) {
		printNumberIfNotZero(node, str, &edges[EdgeLeft])
	} else {
		for edge := EdgeLeft; edge < EdgeCount; edge++ {
			buf := fmt.Sprintf("%s-%s", str, YGEdgeToString(edge))
			printNumberIfNotZero(node, buf, &edges[edge])
		}
	}
}

func nodePrintInternal(node *Node, options PrintOptions, level int) {
	indent(node, level)
	log(node, LogLevelDebug, "<div ")

	if node.Print != nil {
		node.Print(node)
	}

	if options&PrintOptionsLayout != 0 {
		log(node, LogLevelDebug, "layout=\"")
		log(node, LogLevelDebug, "width: %g; ", node.Layout.dimensions[DimensionWidth])
		log(node, LogLevelDebug, "height: %g; ", node.Layout.dimensions[DimensionHeight])
		log(node, LogLevelDebug, "top: %g; ", node.Layout.Position[EdgeTop])
		log(node, LogLevelDebug, "left: %g;", node.Layout.Position[EdgeLeft])
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

		printNumberIfNotUndefinedf(node, "flex-grow", node.Style.flexGrow)
		printNumberIfNotUndefinedf(node, "flex-shrink", node.Style.flexShrink)
		printNumberIfNotAuto(node, "flex-basis", &node.Style.flexBasis)
		printNumberIfNotUndefinedf(node, "flex", node.Style.flex)

		if node.Style.flexWrap != gYGNodeDefaults.Style.flexWrap {
			log(node, LogLevelDebug, "flexWrap: %s; ", YGWrapToString(node.Style.flexWrap))
		}

		if node.Style.overflow != gYGNodeDefaults.Style.overflow {
			log(node, LogLevelDebug, "overflow: %s; ", YGOverflowToString(node.Style.overflow))
		}

		if node.Style.display != gYGNodeDefaults.Style.display {
			log(node, LogLevelDebug, "display: %s; ", YGDisplayToString(node.Style.display))
		}

		printEdges(node, "margin", node.Style.margin[:])
		printEdges(node, "padding", node.Style.padding[:])
		printEdges(node, "border", node.Style.border[:])

		printNumberIfNotAuto(node, "width", &node.Style.dimensions[DimensionWidth])
		printNumberIfNotAuto(node, "height", &node.Style.dimensions[DimensionHeight])
		printNumberIfNotAuto(node, "max-width", &node.Style.maxDimensions[DimensionWidth])
		printNumberIfNotAuto(node, "max-height", &node.Style.maxDimensions[DimensionHeight])
		printNumberIfNotAuto(node, "min-width", &node.Style.minDimensions[DimensionWidth])
		printNumberIfNotAuto(node, "min-height", &node.Style.minDimensions[DimensionHeight])

		if node.Style.positionType != gYGNodeDefaults.Style.positionType {
			log(node,
				LogLevelDebug,
				"position: %s; ",
				YGPositionTypeToString(node.Style.positionType))
		}

		printEdgeIfNotUndefined(node, "left", node.Style.position[:], EdgeLeft)
		printEdgeIfNotUndefined(node, "right", node.Style.position[:], EdgeRight)
		printEdgeIfNotUndefined(node, "top", node.Style.position[:], EdgeTop)
		printEdgeIfNotUndefined(node, "bottom", node.Style.position[:], EdgeBottom)
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
			nodePrintInternal(YGNodeGetChild(node, i), options, level+1)
		}
		indent(node, level)
		log(node, LogLevelDebug, "\n")
	}
	log(node, LogLevelDebug, "</div>")
}

// NodePrint prints node
func NodePrint(node *Node, options PrintOptions) {
	nodePrintInternal(node, options, 0)
}
