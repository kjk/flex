package flex

// Yoga.h

type YGSize struct {
	width  float32
	height float32
}

type YGValue struct {
	value float32
	unit  YGUnit
}

var (
	YGUndefined = NAN
)

var (
	YGValueUndefined = YGValue{YGUndefined, YGUnitUndefined}
	YGValueAuto      = YGValue{YGUndefined, YGUnitAuto}
)

type YGMeasureFunc func(node YGNodeRef, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize

type YGBaselineFunc func(node YGNodeRef, width float32, height float32) float32
type YGPrintFunc func(node YGNodeRef)
type YGLogger func(config YGConfigRef,
	node YGNodeRef,
	level YGLogLevel,
	format string, args ...interface{}) int

// Yoga.c

// This value was chosen based on empiracle data. Even the most complicated
// layouts should not require more than 16 entries to fit within the cache.
const YG_MAX_CACHED_RESULT_COUNT = 16

type YGLayout struct {
	position   [4]float32
	dimensions [2]float32
	margin     [6]float32
	border     [6]float32
	padding    [6]float32
	direction  YGDirection

	computedFlexBasisGeneration uint32
	computedFlexBasis           float32
	hadOverflow                 bool

	// Instead of recomputing the entire layout every single time, we
	// cache some information to break early when nothing changed
	generationCount     uint32
	lastParentDirection YGDirection

	nextCachedMeasurementsIndex uint32
	cachedMeasurements          [YG_MAX_CACHED_RESULT_COUNT]YGCachedMeasurement

	measuredDimensions [2]float32

	cachedLayout YGCachedMeasurement
}

type YGStyle struct {
	direction      YGDirection
	flexDirection  YGFlexDirection
	justifyContent YGJustify
	alignContent   YGAlign
	alignItems     YGAlign
	alignSelf      YGAlign
	positionType   YGPositionType
	flexWrap       YGWrap
	overflow       YGOverflow
	display        YGDisplay
	flex           float32
	flexGrow       float32
	flexShrink     float32
	flexBasis      YGValue
	margin         [YGEdgeCount]YGValue
	position       [YGEdgeCount]YGValue
	padding        [YGEdgeCount]YGValue
	border         [YGEdgeCount]YGValue
	dimensions     [2]YGValue
	minDimensions  [2]YGValue
	maxDimensions  [2]YGValue

	// Yoga specific properties, not compatible with flexbox specification
	aspectRatio float32
}

type YGConfig struct {
	experimentalFeatures      [YGExperimentalFeatureCount + 1]bool
	useWebDefaults            bool
	useLegacyStretchBehaviour bool
	pointScaleFactor          float32
	logger                    YGLogger
	context                   interface{}
}

type YGConfigRef *YGConfig

type YGNode struct {
	style     YGStyle
	layout    YGLayout
	lineIndex uint32

	parent   YGNodeRef
	children YGNodeListRef

	nextChild *YGNode

	measure  YGMeasureFunc
	baseline YGBaselineFunc
	print    YGPrintFunc
	config   YGConfigRef
	context  interface{}

	isDirty      bool
	hasNewLayout bool
	nodeType     YGNodeType

	resolvedDimensions [2]*YGValue
}

type YGNodeRef *YGNode

var (
	YG_UNDEFINED_VALUES = YGValue{
		value: YGUndefined,
		unit:  YGUnitUndefined,
	}

	YG_AUTO_VALUES = YGValue{
		value: YGUndefined,
		unit:  YGUnitAuto,
	}
)

var (
	YG_DEFAULT_EDGE_VALUES_UNIT = [YGEdgeCount]YGValue{
		YG_UNDEFINED_VALUES,
		YG_UNDEFINED_VALUES,
		YG_UNDEFINED_VALUES,
		YG_UNDEFINED_VALUES,
		YG_UNDEFINED_VALUES,
		YG_UNDEFINED_VALUES,
		YG_UNDEFINED_VALUES,
		YG_UNDEFINED_VALUES,
		YG_UNDEFINED_VALUES,
	}
	YG_DEFAULT_DIMENSION_VALUES = [2]float32{
		YGUndefined,
		YGUndefined,
	}

	YG_DEFAULT_DIMENSION_VALUES_UNIT = [2]YGValue{
		YG_UNDEFINED_VALUES,
		YG_UNDEFINED_VALUES,
	}

	YG_DEFAULT_DIMENSION_VALUES_AUTO_UNIT = [2]YGValue{
		YG_AUTO_VALUES,
		YG_AUTO_VALUES,
	}
)

const (
	kDefaultFlexGrow      float32 = 0.0
	kDefaultFlexShrink    float32 = 0.0
	kWebDefaultFlexShrink float32 = 1.0
)

var (
	gYGNodeDefaults = YGNode{
		parent:             nil,
		children:           nil,
		hasNewLayout:       true,
		isDirty:            false,
		nodeType:           YGNodeTypeDefault,
		resolvedDimensions: [2]*YGValue{&YGValueUndefined, &YGValueUndefined},
		style: YGStyle{
			flex:           YGUndefined,
			flexGrow:       YGUndefined,
			flexShrink:     YGUndefined,
			flexBasis:      YG_AUTO_VALUES,
			justifyContent: YGJustifyFlexStart,
			alignItems:     YGAlignStretch,
			alignContent:   YGAlignFlexStart,
			direction:      YGDirectionInherit,
			flexDirection:  YGFlexDirectionColumn,
			overflow:       YGOverflowVisible,
			display:        YGDisplayFlex,
			dimensions:     YG_DEFAULT_DIMENSION_VALUES_AUTO_UNIT,
			minDimensions:  YG_DEFAULT_DIMENSION_VALUES_UNIT,
			maxDimensions:  YG_DEFAULT_DIMENSION_VALUES_UNIT,
			position:       YG_DEFAULT_EDGE_VALUES_UNIT,
			margin:         YG_DEFAULT_EDGE_VALUES_UNIT,
			padding:        YG_DEFAULT_EDGE_VALUES_UNIT,
			border:         YG_DEFAULT_EDGE_VALUES_UNIT,
			aspectRatio:    YGUndefined,
		},
		layout: YGLayout{
			dimensions:                  YG_DEFAULT_DIMENSION_VALUES,
			lastParentDirection:         YGDirection(-1),
			nextCachedMeasurementsIndex: 0,
			computedFlexBasis:           YGUndefined,
			hadOverflow:                 false,
			measuredDimensions:          YG_DEFAULT_DIMENSION_VALUES,
			cachedLayout: YGCachedMeasurement{
				widthMeasureMode:  YGMeasureMode(-1),
				heightMeasureMode: YGMeasureMode(-1),
				computedWidth:     -1,
				computedHeight:    -1,
			},
		},
	}
)
