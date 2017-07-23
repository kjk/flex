package flex

// port of YGEnums.h

// Align describes align flex attribute
type Align int

const (
	// AlignAuto is "auto"
	AlignAuto Align = iota
	// AlignFlexStart is "flex-start"
	AlignFlexStart
	// AlignCenter if "center"
	AlignCenter
	// AlignFlexEnd is "flex-end"
	AlignFlexEnd
	// AlignStretch is "strech"
	AlignStretch
	// AlignBaseline is "baseline"
	AlignBaseline
	// AlignSpaceBetween is "space-between"
	AlignSpaceBetween
	// AlignSpaceAround is "space-around"
	AlignSpaceAround
)

// Dimension represents dimention
type Dimension int

const (
	// DimensionWidth is width
	DimensionWidth Dimension = iota
	// DimensionHeight is height
	DimensionHeight
)

// Direction represents right-to-left or left-to-right direction
type Direction int

const (
	// DirectionInherit is "inherit"
	DirectionInherit Direction = iota
	// DirectionLTR is "ltr"
	DirectionLTR
	// DirectionRTL is "rtl"
	DirectionRTL
)

// Display is "display" property
type Display int

const (
	// DisplayFlex is "flex"
	DisplayFlex Display = iota
	// DisplayNone is "none"
	DisplayNone
)

// Edge represents an edge
type Edge int

const (
	// EdgeLeft is left edge
	EdgeLeft Edge = iota
	// EdgeTop is top edge
	EdgeTop
	// EdgeRight is right edge
	EdgeRight
	// EdgeBottom is bottom edge
	EdgeBottom
	// EdgeStart is start edge
	EdgeStart
	// EdgeEnd is end edge
	EdgeEnd
	// EdgeHorizontal is horizontal edge
	EdgeHorizontal
	// EdgeVertical is vertical edge
	EdgeVertical
	// EdgeAll is all edge
	EdgeAll
)

const (
	// EdgeCount is count of edges
	EdgeCount = 9
)

type YGExperimentalFeature int

const (
	YGExperimentalFeatureWebFlexBasis YGExperimentalFeature = iota
)

const (
	YGExperimentalFeatureCount = 1
)

type YGFlexDirection int

const (
	YGFlexDirectionColumn YGFlexDirection = iota
	YGFlexDirectionColumnReverse
	YGFlexDirectionRow
	YGFlexDirectionRowReverse
)

type YGJustify int

const (
	YGJustifyFlexStart YGJustify = iota
	YGJustifyCenter
	YGJustifyFlexEnd
	YGJustifySpaceBetween
	YGJustifySpaceAround
)

type YGLogLevel int

const (
	YGLogLevelError YGLogLevel = iota
	YGLogLevelWarn
	YGLogLevelInfo
	YGLogLevelDebug
	YGLogLevelVerbose
	YGLogLevelFatal
)

type YGMeasureMode int

const (
	YGMeasureModeUndefined YGMeasureMode = iota
	YGMeasureModeExactly
	YGMeasureModeAtMost
)

const (
	YGMeasureModeCount = 3
)

type YGNodeType int

const (
	YGNodeTypeDefault YGNodeType = iota
	YGNodeTypeText
)

type YGOverflow int

const (
	YGOverflowVisible YGOverflow = iota
	YGOverflowHidden
	YGOverflowScroll
)

type YGPositionType int

const (
	YGPositionTypeRelative YGPositionType = iota
	YGPositionTypeAbsolute
)

type YGPrintOptions int

const (
	YGPrintOptionsLayout YGPrintOptions = 1 << iota
	YGPrintOptionsStyle
	YGPrintOptionsChildren
)

type YGUnit int

const (
	YGUnitUndefined YGUnit = iota
	YGUnitPoint
	YGUnitPercent
	YGUnitAuto
)

type YGWrap int

const (
	YGWrapNoWrap YGWrap = iota
	YGWrapWrap
	YGWrapWrapReverse
)

func YGWrapToString(value YGWrap) string {
	return ""
}

func YGAlignToString(value Align) string {
	return ""
}

func YGDimensionToString(value Dimension) string {
	return ""
}

func YGDirectionToString(value Direction) string {
	return ""
}

func YGDisplayToString(value Display) string {
	return ""
}

func YGEdgeToString(value Edge) string {
	return ""
}
func YGExperimentalFeatureToString(value YGExperimentalFeature) string {
	return ""
}

func YGFlexDirectionToString(value YGFlexDirection) string {
	return ""
}

func YGJustifyToString(value YGJustify) string {
	return ""
}

func YGLogLevelToString(value YGLogLevel) string {
	return ""
}

func YGMeasureModeToString(value YGMeasureMode) string {
	return ""
}

func YGNodeTypeToString(value YGNodeType) string {
	return ""
}

func YGOverflowToString(value YGOverflow) string {
	return ""
}

func YGPositionTypeToString(value YGPositionType) string {
	return ""
}

func YGPrintOptionsToString(value YGPrintOptions) string {
	return ""
}

func YGUnitToString(value YGUnit) string {
	return ""
}
