package flex

var (
	// YGUndefined defines undefined value
	YGUndefined = NAN
)

// YGSize describes size
type YGSize struct {
	width  float32
	height float32
}

// YGValue describes value
type YGValue struct {
	value float32
	unit  YGUnit
}

var (
	// YGValueUndefined defines undefined YGValue
	YGValueUndefined = YGValue{YGUndefined, YGUnitUndefined}
	// YGValueAuto defines auto YGValue
	YGValueAuto = YGValue{YGUndefined, YGUnitAuto}
)

// YGMeasureFunc describes function for measuring
type YGMeasureFunc func(node *YGNode, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) YGSize

// YGBaselineFunc describes function for baseline
type YGBaselineFunc func(node *YGNode, width float32, height float32) float32

// YGPrintFunc defines function for printing
type YGPrintFunc func(node *YGNode)

// YGLogger defines logging function
type YGLogger func(config *YGConfig, node *YGNode, level LogLevel, format string, args ...interface{}) int
