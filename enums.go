package flex

// port of YGEnums.h

type YGAlign int

const (
	YGAlignAuto YGAlign = iota
	YGAlignFlexStart
	YGAlignCenter
	YGAlignFlexEnd
	YGAlignStretch
	YGAlignBaseline
	YGAlignSpaceBetween
	YGAlignSpaceAround
)

type YGDimension int

const (
	YGDimensionWidth YGDimension = iota
	YGDimensionHeight
)

type YGDirection int

const (
	YGDirectionInherit YGDirection = iota
	YGDirectionLTR
	YGDirectionRTL
)

type YGDisplay int

const (
	YGDisplayFlex YGDisplay = iota
	YGDisplayNone
)

type YGEdge int

const (
	YGEdgeLeft YGEdge = iota
	YGEdgeTop
	YGEdgeRight
	YGEdgeBottom
	YGEdgeStart
	YGEdgeEnd
	YGEdgeHorizontal
	YGEdgeVertical
	YGEdgeAll
)

const (
	YGEdgeCount = 9
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

func YGAlignToString(value YGAlign) string {
	return ""
}

func YGDimensionToString(value YGDimension) string {
	return ""
}

func YGDirectionToString(value YGDirection) string {
	return ""
}

func YGDisplayToString(value YGDisplay) string {
	return ""
}

func YGEdgeToString(value YGEdge) string {
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
