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
		log(node, LogLevelDebug, "width: %g; ", node.Layout.Dimensions[DimensionWidth])
		log(node, LogLevelDebug, "height: %g; ", node.Layout.Dimensions[DimensionHeight])
		log(node, LogLevelDebug, "top: %g; ", node.Layout.Position[EdgeTop])
		log(node, LogLevelDebug, "left: %g;", node.Layout.Position[EdgeLeft])
		log(node, LogLevelDebug, "\" ")
	}

	if options&PrintOptionsStyle != 0 {
		log(node, LogLevelDebug, "style=\"")
		if node.Style.FlexDirection != nodeDefaults.Style.FlexDirection {
			log(node,
				LogLevelDebug,
				"flex-direction: %s; ",
				YGFlexDirectionToString(node.Style.FlexDirection))
		}
		if node.Style.JustifyContent != nodeDefaults.Style.JustifyContent {
			log(node,
				LogLevelDebug,
				"justify-content: %s; ",
				YGJustifyToString(node.Style.JustifyContent))
		}
		if node.Style.AlignItems != nodeDefaults.Style.AlignItems {
			log(node, LogLevelDebug, "align-items: %s; ", YGAlignToString(node.Style.AlignItems))
		}
		if node.Style.AlignContent != nodeDefaults.Style.AlignContent {
			log(node, LogLevelDebug, "align-content: %s; ", YGAlignToString(node.Style.AlignContent))
		}
		if node.Style.AlignSelf != nodeDefaults.Style.AlignSelf {
			log(node, LogLevelDebug, "align-self: %s; ", YGAlignToString(node.Style.AlignSelf))
		}

		printNumberIfNotUndefinedf(node, "flex-grow", node.Style.FlexGrow)
		printNumberIfNotUndefinedf(node, "flex-shrink", node.Style.FlexShrink)
		printNumberIfNotAuto(node, "flex-basis", &node.Style.FlexBasis)
		printNumberIfNotUndefinedf(node, "flex", node.Style.Flex)

		if node.Style.FlexWrap != nodeDefaults.Style.FlexWrap {
			log(node, LogLevelDebug, "flexWrap: %s; ", YGWrapToString(node.Style.FlexWrap))
		}

		if node.Style.Overflow != nodeDefaults.Style.Overflow {
			log(node, LogLevelDebug, "overflow: %s; ", YGOverflowToString(node.Style.Overflow))
		}

		if node.Style.Display != nodeDefaults.Style.Display {
			log(node, LogLevelDebug, "display: %s; ", YGDisplayToString(node.Style.Display))
		}

		printEdges(node, "margin", node.Style.Margin[:])
		printEdges(node, "padding", node.Style.Padding[:])
		printEdges(node, "border", node.Style.Border[:])

		printNumberIfNotAuto(node, "width", &node.Style.Dimensions[DimensionWidth])
		printNumberIfNotAuto(node, "height", &node.Style.Dimensions[DimensionHeight])
		printNumberIfNotAuto(node, "max-width", &node.Style.MaxDimensions[DimensionWidth])
		printNumberIfNotAuto(node, "max-height", &node.Style.MaxDimensions[DimensionHeight])
		printNumberIfNotAuto(node, "min-width", &node.Style.MinDimensions[DimensionWidth])
		printNumberIfNotAuto(node, "min-height", &node.Style.MinDimensions[DimensionHeight])

		if node.Style.PositionType != nodeDefaults.Style.PositionType {
			log(node,
				LogLevelDebug,
				"position: %s; ",
				YGPositionTypeToString(node.Style.PositionType))
		}

		printEdgeIfNotUndefined(node, "left", node.Style.Position[:], EdgeLeft)
		printEdgeIfNotUndefined(node, "right", node.Style.Position[:], EdgeRight)
		printEdgeIfNotUndefined(node, "top", node.Style.Position[:], EdgeTop)
		printEdgeIfNotUndefined(node, "bottom", node.Style.Position[:], EdgeBottom)
		log(node, LogLevelDebug, "\" ")

		if node.Measure != nil {
			log(node, LogLevelDebug, "has-custom-measure=\"true\"")
		}
	}
	log(node, LogLevelDebug, ">")

	childCount := len(node.Children)
	if options&PrintOptionsChildren != 0 && childCount > 0 {
		for i := 0; i < childCount; i++ {
			log(node, LogLevelDebug, "\n")
			nodePrintInternal(node.Children[i], options, level+1)
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
