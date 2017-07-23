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

// ExperimentalFeature defines experimental features
type ExperimentalFeature int

const (
	// ExperimentalFeatureWebFlexBasis is web flex basis
	ExperimentalFeatureWebFlexBasis ExperimentalFeature = iota
)

const (
	experimentalFeatureCount = 1
)

// FlexDirection describes "flex-direction" property
type FlexDirection int

const (
	// FlexDirectionColumn is "column"
	FlexDirectionColumn FlexDirection = iota
	// FlexDirectionColumnReverse is "column-reverse"
	FlexDirectionColumnReverse
	// FlexDirectionRow is "row"
	FlexDirectionRow
	// FlexDirectionRowReverse is "row-reverse"
	FlexDirectionRowReverse
)

// Justify is "justify" property
type Justify int

const (
	// JustifyFlexStart is "flex-start"
	JustifyFlexStart Justify = iota
	// JustifyCenter is "center"
	JustifyCenter
	// JustifyFlexEnd is "flex-end"
	JustifyFlexEnd
	// JustifySpaceBetween is "space-between"
	JustifySpaceBetween
	// JustifySpaceAround is "space-around"
	JustifySpaceAround
)

// LogLevel represents log level
type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
	LogLevelVerbose
	LogLevelFatal
)

// MeasureMode defines measurement mode
type MeasureMode int

const (
	// MeasureModeUndefined is undefined
	MeasureModeUndefined MeasureMode = iota
	// MeasureModeExactly is exactly
	MeasureModeExactly
	// MeasureModeAtMost is at-most
	MeasureModeAtMost
)

const (
	measureModeCount = 3
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
func YGExperimentalFeatureToString(value ExperimentalFeature) string {
	return ""
}

func YGFlexDirectionToString(value FlexDirection) string {
	return ""
}

func YGJustifyToString(value Justify) string {
	return ""
}

func YGLogLevelToString(value LogLevel) string {
	return ""
}

func YGMeasureModeToString(value MeasureMode) string {
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
