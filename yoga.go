package flex

import (
	"fmt"
	"os"
)

// Yoga.h

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
	// YGUndefined defines undefined value
	YGUndefined = NAN
	// YGValueUndefined defines undefined YGValue
	YGValueUndefined = YGValue{YGUndefined, YGUnitUndefined}
	// YGValueAuto defines auto YGValue
	YGValueAuto = YGValue{YGUndefined, YGUnitAuto}

	gCurrentGenerationCount = 0

	gPrintTree    = false
	gPrintChanges = false
	gPrintSkips   = false
	gDepth        = 0
)

// YGMeasureFunc describes function for measuring
type YGMeasureFunc func(node YGNodeRef, width float32, widthMode YGMeasureMode, height float32, heightMode YGMeasureMode) YGSize

// YGBaselineFunc describes function for baseline
type YGBaselineFunc func(node YGNodeRef, width float32, height float32) float32

// YGPrintFunc defines function for printing
type YGPrintFunc func(node YGNodeRef)

// YGLogger defines logging function
type YGLogger func(config YGConfigRef,
	node YGNodeRef,
	level YGLogLevel,
	format string, args ...interface{}) int

// Yoga.c

// This value was chosen based on empiracle data. Even the most complicated
// layouts should not require more than 16 entries to fit within the cache.
const YG_MAX_CACHED_RESULT_COUNT = 16

// YGCachedMeasurement describes measurements
type YGCachedMeasurement struct {
	availableWidth    float32
	availableHeight   float32
	widthMeasureMode  YGMeasureMode
	heightMeasureMode YGMeasureMode

	computedWidth  float32
	computedHeight float32
}

// YGLayout describes layout
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
	generationCount     int
	lastParentDirection YGDirection

	nextCachedMeasurementsIndex int
	cachedMeasurements          [YG_MAX_CACHED_RESULT_COUNT]YGCachedMeasurement

	measuredDimensions [2]float32

	cachedLayout YGCachedMeasurement
}

// YGStyle describes a style
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

// YGConfig describes a config
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

var (
	gYGConfigDefaults = YGConfig{
		experimentalFeatures: [YGExperimentalFeatureCount + 1]bool{
			false,
			false,
		},
		useWebDefaults:   false,
		pointScaleFactor: 1.0,
		logger:           YGDefaultLog,
		context:          nil,
	}
)

var (
	YGValueZero = YGValue{value: 0, unit: YGUnitPoint}
)

var (
	leading  = [4]YGEdge{YGEdgeTop, YGEdgeBottom, YGEdgeLeft, YGEdgeRight}
	trailing = [4]YGEdge{YGEdgeBottom, YGEdgeTop, YGEdgeRight, YGEdgeLeft}
	pos      = [4]YGEdge{YGEdgeTop, YGEdgeBottom, YGEdgeLeft, YGEdgeRight}
	dim      = [4]YGDimension{YGDimensionHeight, YGDimensionHeight, YGDimensionWidth, YGDimensionWidth}
)

// YGDefaultLog is default logging function
func YGDefaultLog(config YGConfigRef, node YGNodeRef, level YGLogLevel, format string,
	args ...interface{}) int {
	switch level {
	case YGLogLevelError:
	case YGLogLevelFatal:
		n, _ := fmt.Fprintf(os.Stderr, format, args...)
		return n
	case YGLogLevelWarn:
	case YGLogLevelInfo:
	case YGLogLevelDebug:
	case YGLogLevelVerbose:
	default:
		n, _ := fmt.Printf(format, args...)
		return n
	}
	return 0
}

// YGComputedEdgeValue computes edge value
func YGComputedEdgeValue(edges [YGEdgeCount]YGValue, edge YGEdge, defaultValue *YGValue) *YGValue {
	if edges[edge].unit != YGUnitUndefined {
		return &edges[edge]
	}

	if edge == YGEdgeTop || edge == YGEdgeBottom &&
		edges[YGEdgeVertical].unit != YGUnitUndefined {
		return &edges[YGEdgeVertical]
	}

	if (edge == YGEdgeLeft || edge == YGEdgeRight || edge == YGEdgeStart || edge == YGEdgeEnd) &&
		edges[YGEdgeHorizontal].unit != YGUnitUndefined {
		return &edges[YGEdgeHorizontal]
	}

	if edges[YGEdgeAll].unit != YGUnitUndefined {
		return &edges[YGEdgeAll]
	}

	if edge == YGEdgeStart || edge == YGEdgeEnd {
		return &YGValueUndefined
	}

	return defaultValue
}

// YGResolveValue resolves a value
func YGResolveValue(value *YGValue, parentSize float32) float32 {
	switch value.unit {
	case YGUnitUndefined:
	case YGUnitAuto:
		return YGUndefined
	case YGUnitPoint:
		return value.value
	case YGUnitPercent:
		return value.value * parentSize / 100
	}
	return YGUndefined
}

// YGResolveValueMargin resolves margin value
func YGResolveValueMargin(value *YGValue, parentSize float32) float32 {
	if value.unit == YGUnitAuto {
		return 0
	}
	return YGResolveValue(value, parentSize)
}

var (
	gNodeInstanceCount   int = 0
	gConfigInstanceCount int = 0
)

// YGNodeNewWithConfig creates new node with config
func YGNodeNewWithConfig(config YGConfigRef) YGNodeRef {
	node := YGNode{}
	gNodeInstanceCount++

	node = gYGNodeDefaults
	if config.useWebDefaults {
		node.style.flexDirection = YGFlexDirectionRow
		node.style.alignContent = YGAlignStretch
	}
	node.config = config
	return &node
}

// YGNodeNew creates a new node
func YGNodeNew() YGNodeRef {
	return YGNodeNewWithConfig(&gYGConfigDefaults)
}

// YGNodeFree frees a node
func YGNodeFree(node YGNodeRef) {
	if node.parent != nil {
		YGNodeListDelete(node.parent.children, node)
		node.parent = nil
	}

	childCount := YGNodeGetChildCount(node)
	for i := 0; i < childCount; i++ {
		child := YGNodeGetChild(node, i)
		child.parent = nil
	}

	YGNodeListFree(node.children)
	//gYGFree(node);
	gNodeInstanceCount--
}

// YGNodeGetChildCount returns number of children
func YGNodeGetChildCount(node YGNodeRef) int {
	return YGNodeListCount(node.children)
}

// YGNodeGetChild returns a child
func YGNodeGetChild(node YGNodeRef, index int) YGNodeRef {
	return YGNodeListGet(node.children, index)
}

// YGNodeFreeRecursive frees root recursevily
func YGNodeFreeRecursive(root YGNodeRef) {
	for YGNodeGetChildCount(root) > 0 {
		child := YGNodeGetChild(root, 0)
		YGNodeRemoveChild(root, child)
		YGNodeFreeRecursive(child)
	}
	YGNodeFree(root)
}

// YGNodeRemoveChild removes the child
func YGNodeRemoveChild(node YGNodeRef, child YGNodeRef) {
	if YGNodeListDelete(node.children, child) != nil {
		child.layout = gYGNodeDefaults.layout // layout is no longer valid
		child.parent = nil
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeMarkDirtyInternal marks the node as dirty, internally
func YGNodeMarkDirtyInternal(node YGNodeRef) {
	if !node.isDirty {
		node.isDirty = true
		node.layout.computedFlexBasis = YGUndefined
		if node.parent != nil {
			YGNodeMarkDirtyInternal(node.parent)
		}
	}
}

// YGNodeMarkDirty marks node as dirty. Only valid for nodes with a custom measure function
// set.
// YG knows when to mark all other nodes as dirty but because nodes with
// measure functions
// depends on information not known to YG they must perform this dirty
// marking manually.
func YGNodeMarkDirty(node YGNodeRef) {
	YGAssertWithNode(node, node.measure != nil,
		"Only leaf nodes with custom measure functions should manually mark themselves as dirty")
	YGNodeMarkDirtyInternal(node)
}

// YGNodeIsDirty returns true if node is dirty
func YGNodeIsDirty(node YGNodeRef) bool {
	return node.isDirty
}

// YGAssert asserts that cond is true
func YGAssert(cond bool, format string, args ...interface{}) {
	if !cond {
		panic(format)
	}
}

// YGAssertWithNode assert if cond is not true
func YGAssertWithNode(node YGNodeRef, cond bool, format string, args ...interface{}) {
	YGAssert(cond, format, args...)
}

// YGNodeReset resets a node
func YGNodeReset(node YGNodeRef) {
	YGAssertWithNode(node,
		YGNodeGetChildCount(node) == 0,
		"Cannot reset a node which still has children attached")
	YGAssertWithNode(node, node.parent == nil, "Cannot reset a node still attached to a parent")

	YGNodeListFree(node.children)

	config := node.config
	*node = gYGNodeDefaults
	if config.useWebDefaults {
		node.style.flexDirection = YGFlexDirectionRow
		node.style.alignContent = YGAlignStretch
	}
	node.config = config
}

// YGNodeGetInstanceCount returns node instance count
func YGNodeGetInstanceCount() int {
	return gNodeInstanceCount
}

// YGConfigGetInstanceCount returns config instance count
func YGConfigGetInstanceCount() int {
	return gConfigInstanceCount
}

// YGConfigGetDefault returns default config, only for C#
func YGConfigGetDefault() YGConfigRef {
	return &gYGConfigDefaults
}

// YGConfigNew creates new config
func YGConfigNew() YGConfigRef {
	config := &YGConfig{}
	YGAssert(config != nil, "Could not allocate memory for config")

	gConfigInstanceCount++
	*config = gYGConfigDefaults
	return config
}

// YGConfigFree frees a config
func YGConfigFree(config YGConfigRef) {
	// gYGFree(config)
	gConfigInstanceCount--
}

// YGConfigCopy copies a config
func YGConfigCopy(dest YGConfigRef, src YGConfigRef) {
	*dest = *src
}

// YGFloatIsUndefined returns true if value is undefined
func YGFloatIsUndefined(value float32) bool {
	return IsNaN(value)
}

func YGNodeStyleSetWidth(node YGNodeRef, width float32) {
	if node.style.dimensions[YGDimensionWidth].value != width ||
		node.style.dimensions[YGDimensionWidth].unit != YGUnitPoint {
		node.style.dimensions[YGDimensionWidth].value = width
		unit := YGUnitPoint
		if YGFloatIsUndefined(width) {
			unit = YGUnitAuto
		}
		node.style.dimensions[YGDimensionWidth].unit = unit
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetWidthPercent(node YGNodeRef, width float32) {
	if node.style.dimensions[YGDimensionWidth].value != width ||
		node.style.dimensions[YGDimensionWidth].unit != YGUnitPercent {
		node.style.dimensions[YGDimensionWidth].value = width
		unit := YGUnitPercent
		if YGFloatIsUndefined(width) {
			unit = YGUnitAuto
		}
		node.style.dimensions[YGDimensionWidth].unit = unit
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetWidthAuto(node YGNodeRef) {
	if node.style.dimensions[YGDimensionWidth].unit != YGUnitAuto {
		node.style.dimensions[YGDimensionWidth].value = YGUndefined
		node.style.dimensions[YGDimensionWidth].unit = YGUnitAuto
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetWidth(node YGNodeRef) YGValue {
	return node.style.dimensions[YGDimensionWidth]
}

func YGNodeStyleSetHeight(node YGNodeRef, height float32) {
	if node.style.dimensions[YGDimensionHeight].value != height ||
		node.style.dimensions[YGDimensionHeight].unit != YGUnitPoint {
		node.style.dimensions[YGDimensionHeight].value = height
		unit := YGUnitPoint
		if YGFloatIsUndefined(height) {
			unit = YGUnitAuto
		}
		node.style.dimensions[YGDimensionHeight].unit = unit
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetHeightPercent(node YGNodeRef, height float32) {
	if node.style.dimensions[YGDimensionHeight].value != height ||
		node.style.dimensions[YGDimensionHeight].unit != YGUnitPercent {
		node.style.dimensions[YGDimensionHeight].value = height
		unit := YGUnitPercent
		if YGFloatIsUndefined(height) {
			unit = YGUnitAuto
		}
		node.style.dimensions[YGDimensionHeight].unit = unit
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetHeightAuto(node YGNodeRef) {
	if node.style.dimensions[YGDimensionHeight].unit != YGUnitAuto {
		node.style.dimensions[YGDimensionHeight].value = YGUndefined
		node.style.dimensions[YGDimensionHeight].unit = YGUnitAuto
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetHeight(node YGNodeRef) YGValue {
	return node.style.dimensions[YGDimensionHeight]
}

func YGNodeStyleSetPositionType(node *YGNode, positionType YGPositionType) {
	if node.style.positionType != positionType {
		node.style.positionType = positionType
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetPositionType(node *YGNode) YGPositionType {
	return node.style.positionType
}

func YGNodeStyleSetPosition(node *YGNode, edge YGEdge,
	position float32) {
	if node.style.position[edge].value != position ||
		node.style.position[edge].unit != YGUnitPoint {
		node.style.position[edge].value = position
		unit := YGUnitPoint
		if YGFloatIsUndefined(position) {
			unit = YGUnitUndefined
		}
		node.style.position[edge].unit = unit
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetPositionPercent(node *YGNode, edge YGEdge,
	position float32) {
	if node.style.position[edge].value != position ||
		node.style.position[edge].unit != YGUnitPercent {
		node.style.position[edge].value = position
		unit := YGUnitPercent
		if YGFloatIsUndefined(position) {
			unit = YGUnitUndefined
		}
		node.style.position[edge].unit = unit
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetPosition(node *YGNode, edge YGEdge) YGValue {
	return node.style.position[edge]
}

func YGNodeInsertChild(node *YGNode, child *YGNode, index int) {
	YGAssertWithNode(node, child.parent == nil, "Child already has a parent, it must be removed first.")
	YGAssertWithNode(node, node.measure == nil, "Cannot add child: Nodes with measure functions cannot have children.")

	YGNodeListInsert(&node.children, child, index)
	child.parent = node
	YGNodeMarkDirtyInternal(node)
}

func YGFloatsEqual(a float32, b float32) bool {
	if YGFloatIsUndefined(a) {
		return YGFloatIsUndefined(b)
	}
	return fabs(a-b) < 0.0001
}

func YGValueEqual(a YGValue, b YGValue) bool {
	if a.unit != b.unit {
		return false
	}

	if a.unit == YGUnitUndefined {
		return true
	}

	return fabs(a.value-b.value) < 0.0001
}

func YGResolveDimensions(node *YGNode) {
	for dim := YGDimensionWidth; dim <= YGDimensionHeight; dim++ {
		if node.style.maxDimensions[dim].unit != YGUnitUndefined &&
			YGValueEqual(node.style.maxDimensions[dim], node.style.minDimensions[dim]) {
			node.resolvedDimensions[dim] = &node.style.maxDimensions[dim]
		} else {
			node.resolvedDimensions[dim] = &node.style.dimensions[dim]
		}
	}
}

func YGNodeIsStyleDimDefined(node *YGNode, axis YGFlexDirection, parentSize float32) bool {
	return !(node.resolvedDimensions[dim[axis]].unit == YGUnitAuto ||
		node.resolvedDimensions[dim[axis]].unit == YGUnitUndefined ||
		(node.resolvedDimensions[dim[axis]].unit == YGUnitPoint &&
			node.resolvedDimensions[dim[axis]].value < 0) ||
		(node.resolvedDimensions[dim[axis]].unit == YGUnitPercent &&
			(node.resolvedDimensions[dim[axis]].value < 0 || YGFloatIsUndefined(parentSize))))
}

func YGNodeMarginForAxis(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	return YGNodeLeadingMargin(node, axis, widthSize) + YGNodeTrailingMargin(node, axis, widthSize)
}

func YGFlexDirectionIsRow(flexDirection YGFlexDirection) bool {
	return flexDirection == YGFlexDirectionRow || flexDirection == YGFlexDirectionRowReverse
}

func YGFlexDirectionIsColumn(flexDirection YGFlexDirection) bool {
	return flexDirection == YGFlexDirectionColumn || flexDirection == YGFlexDirectionColumnReverse
}

func YGNodeLeadingMargin(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.margin[YGEdgeStart].unit != YGUnitUndefined {
		return YGResolveValueMargin(&node.style.margin[YGEdgeStart], widthSize)
	}

	return YGResolveValueMargin(YGComputedEdgeValue(node.style.margin, leading[axis], &YGValueZero), widthSize)
}

func YGNodeTrailingMargin(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.margin[YGEdgeEnd].unit != YGUnitUndefined {
		return YGResolveValueMargin(&node.style.margin[YGEdgeEnd], widthSize)
	}

	return YGResolveValueMargin(YGComputedEdgeValue(node.style.margin, trailing[axis], &YGValueZero),
		widthSize)
}

func YGNodeLeadingPadding(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.padding[YGEdgeStart].unit != YGUnitUndefined &&
		YGResolveValue(&node.style.padding[YGEdgeStart], widthSize) >= 0 {
		return YGResolveValue(&node.style.padding[YGEdgeStart], widthSize)
	}

	return fmaxf(YGResolveValue(YGComputedEdgeValue(node.style.padding, leading[axis], &YGValueZero), widthSize), 0)
}

func YGNodeTrailingPadding(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.padding[YGEdgeEnd].unit != YGUnitUndefined &&
		YGResolveValue(&node.style.padding[YGEdgeEnd], widthSize) >= 0 {
		return YGResolveValue(&node.style.padding[YGEdgeEnd], widthSize)
	}

	return fmaxf(YGResolveValue(YGComputedEdgeValue(node.style.padding, trailing[axis], &YGValueZero), widthSize), 0)
}

func YGNodeLeadingBorder(node *YGNode, axis YGFlexDirection) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.border[YGEdgeStart].unit != YGUnitUndefined &&
		node.style.border[YGEdgeStart].value >= 0 {
		return node.style.border[YGEdgeStart].value
	}

	return fmaxf(YGComputedEdgeValue(node.style.border, leading[axis], &YGValueZero).value, 0)
}

func YGNodeTrailingBorder(node *YGNode, axis YGFlexDirection) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.border[YGEdgeEnd].unit != YGUnitUndefined &&
		node.style.border[YGEdgeEnd].value >= 0 {
		return node.style.border[YGEdgeEnd].value
	}

	return fmaxf(YGComputedEdgeValue(node.style.border, trailing[axis], &YGValueZero).value, 0)
}

func YGNodeLeadingPaddingAndBorder(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	return YGNodeLeadingPadding(node, axis, widthSize) + YGNodeLeadingBorder(node, axis)
}

func YGNodeTrailingPaddingAndBorder(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	return YGNodeTrailingPadding(node, axis, widthSize) + YGNodeTrailingBorder(node, axis)
}

func YGNodePaddingAndBorderForAxis(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	return YGNodeLeadingPaddingAndBorder(node, axis, widthSize) +
		YGNodeTrailingPaddingAndBorder(node, axis, widthSize)
}

func YGNodeAlignItem(node *YGNode, child *YGNode) YGAlign {
	align := child.style.alignSelf
	if child.style.alignSelf == YGAlignAuto {
		align = node.style.alignItems
	}
	if align == YGAlignBaseline && YGFlexDirectionIsColumn(node.style.flexDirection) {
		return YGAlignFlexStart
	}
	return align
}

func YGNodeResolveDirection(node *YGNode, parentDirection YGDirection) YGDirection {
	if node.style.direction == YGDirectionInherit {
		if parentDirection > YGDirectionInherit {
			return parentDirection
		}
		return YGDirectionLTR
	}
	return node.style.direction
}

func YGResolveFlexDirection(flexDirection YGFlexDirection, direction YGDirection) YGFlexDirection {
	if direction == YGDirectionRTL {
		if flexDirection == YGFlexDirectionRow {
			return YGFlexDirectionRowReverse
		} else if flexDirection == YGFlexDirectionRowReverse {
			return YGFlexDirectionRow
		}
	}

	return flexDirection
}

func YGFlexDirectionCross(flexDirection YGFlexDirection, direction YGDirection) YGFlexDirection {
	if YGFlexDirectionIsColumn(flexDirection) {
		return YGResolveFlexDirection(YGFlexDirectionRow, direction)
	}
	return YGFlexDirectionColumn
}

func YGResolveFlexGrow(node *YGNode) float32 {
	// Root nodes flexGrow should always be 0
	if node.parent == nil {
		return 0.0
	}
	if !YGFloatIsUndefined(node.style.flexGrow) {
		return node.style.flexGrow
	}
	if !YGFloatIsUndefined(node.style.flex) && node.style.flex > 0 {
		return node.style.flex
	}
	return kDefaultFlexGrow
}

func YGNodeStyleGetFlexGrow(node *YGNode) float32 {
	if YGFloatIsUndefined(node.style.flexGrow) {
		return kDefaultFlexGrow
	}
	return node.style.flexGrow
}

func YGNodeStyleGetFlexShrink(node *YGNode) float32 {
	if YGFloatIsUndefined(node.style.flexShrink) {
		if node.config.useWebDefaults {
			return kWebDefaultFlexShrink
		}
		return kDefaultFlexShrink
	}
	return node.style.flexShrink
}

func YGNodeResolveFlexShrink(node *YGNode) float32 {
	// Root nodes flexShrink should always be 0
	if node.parent == nil {
		return 0.0
	}
	if !YGFloatIsUndefined(node.style.flexShrink) {
		return node.style.flexShrink
	}
	if !node.config.useWebDefaults && !YGFloatIsUndefined(node.style.flex) &&
		node.style.flex < 0 {
		return -node.style.flex
	}
	if node.config.useWebDefaults {
		return kWebDefaultFlexShrink
	}
	return kDefaultFlexShrink
}

func YGNodeResolveFlexBasisPtr(node *YGNode) *YGValue {
	if node.style.flexBasis.unit != YGUnitAuto && node.style.flexBasis.unit != YGUnitUndefined {
		return &node.style.flexBasis
	}
	if !YGFloatIsUndefined(node.style.flex) && node.style.flex > 0 {
		if node.config.useWebDefaults {
			return &YGValueAuto
		}
		return &YGValueZero
	}
	return &YGValueAuto
}

func YGNodeIsFlex(node *YGNode) bool {
	return (node.style.positionType == YGPositionTypeRelative &&
		(YGResolveFlexGrow(node) != 0 || YGNodeResolveFlexShrink(node) != 0))
}

func YGMarginLeadingValue(node *YGNode, axis YGFlexDirection) *YGValue {
	if YGFlexDirectionIsRow(axis) && node.style.margin[YGEdgeStart].unit != YGUnitUndefined {
		return &node.style.margin[YGEdgeStart]
	}
	return &node.style.margin[leading[axis]]
}

func YGMarginTrailingValue(node *YGNode, axis YGFlexDirection) *YGValue {
	if YGFlexDirectionIsRow(axis) && node.style.margin[YGEdgeEnd].unit != YGUnitUndefined {
		return &node.style.margin[YGEdgeEnd]
	}
	return &node.style.margin[trailing[axis]]

}

func YGNodeDimWithMargin(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	return node.layout.measuredDimensions[dim[axis]] + YGNodeLeadingMargin(node, axis, widthSize) +
		YGNodeTrailingMargin(node, axis, widthSize)
}

func YGNodeIsLayoutDimDefined(node *YGNode, axis YGFlexDirection) bool {
	value := node.layout.measuredDimensions[dim[axis]]
	return !YGFloatIsUndefined(value) && value >= 0
}

func YGNodeIsLeadingPosDefined(node *YGNode, axis YGFlexDirection) bool {
	return (YGFlexDirectionIsRow(axis) &&
		YGComputedEdgeValue(node.style.position, YGEdgeStart, &YGValueUndefined).unit !=
			YGUnitUndefined) ||
		YGComputedEdgeValue(node.style.position, leading[axis], &YGValueUndefined).unit !=
			YGUnitUndefined
}

func YGNodeIsTrailingPosDefined(node *YGNode, axis YGFlexDirection) bool {
	return (YGFlexDirectionIsRow(axis) &&
		YGComputedEdgeValue(node.style.position, YGEdgeEnd, &YGValueUndefined).unit !=
			YGUnitUndefined) ||
		YGComputedEdgeValue(node.style.position, trailing[axis], &YGValueUndefined).unit !=
			YGUnitUndefined
}

func YGNodeLeadingPosition(node *YGNode, axis YGFlexDirection, axisSize float32) float32 {
	if YGFlexDirectionIsRow(axis) {
		leadingPosition := YGComputedEdgeValue(node.style.position, YGEdgeStart, &YGValueUndefined)
		if leadingPosition.unit != YGUnitUndefined {
			return YGResolveValue(leadingPosition, axisSize)
		}
	}

	leadingPosition := YGComputedEdgeValue(node.style.position, leading[axis], &YGValueUndefined)

	if leadingPosition.unit == YGUnitUndefined {
		return 0
	}
	return YGResolveValue(leadingPosition, axisSize)
}

func YGNodeTrailingPosition(node *YGNode, axis YGFlexDirection, axisSize float32) float32 {
	if YGFlexDirectionIsRow(axis) {
		trailingPosition := YGComputedEdgeValue(node.style.position, YGEdgeEnd, &YGValueUndefined)
		if trailingPosition.unit != YGUnitUndefined {
			return YGResolveValue(trailingPosition, axisSize)
		}
	}

	trailingPosition := YGComputedEdgeValue(node.style.position, trailing[axis], &YGValueUndefined)

	if trailingPosition.unit == YGUnitUndefined {
		return 0
	}
	return YGResolveValue(trailingPosition, axisSize)
}

// If both left and right are defined, then use left. Otherwise return
// +left or -right depending on which is defined.
func YGNodeRelativePosition(node *YGNode, axis YGFlexDirection, axisSize float32) float32 {
	if YGNodeIsLeadingPosDefined(node, axis) {
		return YGNodeLeadingPosition(node, axis, axisSize)
	}
	return -YGNodeTrailingPosition(node, axis, axisSize)
}

func YGNodeSetPosition(node *YGNode, direction YGDirection, mainSize float32, crossSize float32, parentWidth float32) {

	directionRespectingRoot := YGDirectionLTR
	if node.parent != nil {
		directionRespectingRoot = direction
	}

	mainAxis := YGResolveFlexDirection(node.style.flexDirection, directionRespectingRoot)
	crossAxis := YGFlexDirectionCross(mainAxis, directionRespectingRoot)

	relativePositionMain := YGNodeRelativePosition(node, mainAxis, mainSize)
	relativePositionCross := YGNodeRelativePosition(node, crossAxis, crossSize)

	node.layout.position[leading[mainAxis]] =
		YGNodeLeadingMargin(node, mainAxis, parentWidth) + relativePositionMain
	node.layout.position[trailing[mainAxis]] =
		YGNodeTrailingMargin(node, mainAxis, parentWidth) + relativePositionMain
	node.layout.position[leading[crossAxis]] =
		YGNodeLeadingMargin(node, crossAxis, parentWidth) + relativePositionCross
	node.layout.position[trailing[crossAxis]] =
		YGNodeTrailingMargin(node, crossAxis, parentWidth) +
			relativePositionCross
}

func YGNodePrintInternal(node *YGNode, options YGPrintOptions, level int) {
	// TODO: write me
}

func YGNodePrint(node *YGNode, options YGPrintOptions) {
	YGNodePrintInternal(node, options, 0)
}

const (
	spacer = "                                                            "
)

// YGSpacer returns spacer string
func YGSpacer(level int) string {
	n := len(spacer)
	if level > n {
		level = n
	}
	return spacer[:level]
}

var (
	kMeasureModeNames = [YGMeasureModeCount]string{"UNDEFINED", "EXACTLY", "AT_MOST"}
	kLayoutModeNames  = [YGMeasureModeCount]string{"LAY_UNDEFINED", "LAY_EXACTLY", "LAY_AT_MOST"}
)

func YGMeasureModeName(mode YGMeasureMode, performLayout bool) string {

	if mode >= YGMeasureModeCount {
		return ""
	}

	if performLayout {
		return kLayoutModeNames[mode]
	}
	return kMeasureModeNames[mode]
}

func YGMeasureModeSizeIsExactAndMatchesOldMeasuredSize(sizeMode YGMeasureMode, size float32, lastComputedSize float32) bool {
	return sizeMode == YGMeasureModeExactly && YGFloatsEqual(size, lastComputedSize)
}

func YGMeasureModeOldSizeIsUnspecifiedAndStillFits(sizeMode YGMeasureMode, size float32, lastSizeMode YGMeasureMode, lastComputedSize float32) bool {
	return sizeMode == YGMeasureModeAtMost && lastSizeMode == YGMeasureModeUndefined &&
		(size >= lastComputedSize || YGFloatsEqual(size, lastComputedSize))
}

func YGMeasureModeNewMeasureSizeIsStricterAndStillValid(sizeMode YGMeasureMode, size float32, lastSizeMode YGMeasureMode, lastSize float32, lastComputedSize float32) bool {
	return lastSizeMode == YGMeasureModeAtMost && sizeMode == YGMeasureModeAtMost &&
		lastSize > size && (lastComputedSize <= size || YGFloatsEqual(size, lastComputedSize))
}

func YGConfigSetPointScaleFactor(config YGConfigRef, pixelsInPoint float32) {
	//YGAssertWithConfig(config, pixelsInPoint >= 0, "Scale factor should not be less than zero");

	// We store points for Pixel as we will use it for rounding
	if pixelsInPoint == 0 {
		// Zero is used to skip rounding
		config.pointScaleFactor = 0
	} else {
		config.pointScaleFactor = pixelsInPoint
	}
}

func YGNodeCanUseCachedMeasurement(widthMode YGMeasureMode, width float32, heightMode YGMeasureMode, height float32, lastWidthMode YGMeasureMode, lastWidth float32, lastHeightMode YGMeasureMode, lastHeight float32, lastComputedWidth float32, lastComputedHeight float32, marginRow float32, marginColumn float32, config YGConfigRef) bool {
	if lastComputedHeight < 0 || lastComputedWidth < 0 {
		return false
	}
	useRoundedComparison := config != nil && config.pointScaleFactor != 0
	effectiveWidth := width
	if useRoundedComparison {
		effectiveWidth = YGRoundValueToPixelGrid(width, config.pointScaleFactor, false, false)
	}
	effectiveHeight := height
	if useRoundedComparison {
		effectiveHeight = YGRoundValueToPixelGrid(height, config.pointScaleFactor, false, false)
	}
	effectiveLastWidth := lastWidth
	if useRoundedComparison {
		effectiveLastWidth = YGRoundValueToPixelGrid(lastWidth, config.pointScaleFactor, false, false)
	}
	effectiveLastHeight := lastHeight
	if useRoundedComparison {
		effectiveLastHeight = YGRoundValueToPixelGrid(lastHeight, config.pointScaleFactor, false, false)
	}

	hasSameWidthSpec := lastWidthMode == widthMode && YGFloatsEqual(effectiveLastWidth, effectiveWidth)
	hasSameHeightSpec := lastHeightMode == heightMode && YGFloatsEqual(effectiveLastHeight, effectiveHeight)

	widthIsCompatible :=
		hasSameWidthSpec || YGMeasureModeSizeIsExactAndMatchesOldMeasuredSize(widthMode,
			width-marginRow,
			lastComputedWidth) ||
			YGMeasureModeOldSizeIsUnspecifiedAndStillFits(widthMode,
				width-marginRow,
				lastWidthMode,
				lastComputedWidth) ||
			YGMeasureModeNewMeasureSizeIsStricterAndStillValid(
				widthMode, width-marginRow, lastWidthMode, lastWidth, lastComputedWidth)

	heightIsCompatible :=
		hasSameHeightSpec || YGMeasureModeSizeIsExactAndMatchesOldMeasuredSize(heightMode,
			height-marginColumn,
			lastComputedHeight) ||
			YGMeasureModeOldSizeIsUnspecifiedAndStillFits(heightMode,
				height-marginColumn,
				lastHeightMode,
				lastComputedHeight) ||
			YGMeasureModeNewMeasureSizeIsStricterAndStillValid(
				heightMode, height-marginColumn, lastHeightMode, lastHeight, lastComputedHeight)

	return widthIsCompatible && heightIsCompatible
}

// This is a wrapper around the YGNodelayoutImpl function. It determines
// whether the layout request is redundant and can be skipped.
//
// Parameters:
//  Input parameters are the same as YGNodelayoutImpl (see above)
//  Return parameter is true if layout was performed, false if skipped
func YGLayoutNodeInternal(node *YGNode, availableWidth float32, availableHeight float32,
	parentDirection YGDirection,
	widthMeasureMode YGMeasureMode,
	heightMeasureMode YGMeasureMode,
	parentWidth float32,
	parentHeight float32,
	performLayout bool,
	reason string,
	config YGConfigRef) bool {
	layout := &node.layout

	gDepth++

	needToVisitNode :=
		(node.isDirty && layout.generationCount != gCurrentGenerationCount) ||
			layout.lastParentDirection != parentDirection

	if needToVisitNode {
		// Invalidate the cached results.
		layout.nextCachedMeasurementsIndex = 0
		layout.cachedLayout.widthMeasureMode = YGMeasureMode(-1)
		layout.cachedLayout.heightMeasureMode = YGMeasureMode(-1)
		layout.cachedLayout.computedWidth = -1
		layout.cachedLayout.computedHeight = -1
	}

	var cachedResults *YGCachedMeasurement

	// Determine whether the results are already cached. We maintain a separate
	// cache for layouts and measurements. A layout operation modifies the
	// positions
	// and dimensions for nodes in the subtree. The algorithm assumes that each
	// node
	// gets layed out a maximum of one time per tree layout, but multiple
	// measurements
	// may be required to resolve all of the flex dimensions.
	// We handle nodes with measure functions specially here because they are the
	// most
	// expensive to measure, so it's worth avoiding redundant measurements if at
	// all possible.
	if node.measure != nil {
		marginAxisRow := YGNodeMarginForAxis(node, YGFlexDirectionRow, parentWidth)
		marginAxisColumn := YGNodeMarginForAxis(node, YGFlexDirectionColumn, parentWidth)

		// First, try to use the layout cache.
		if YGNodeCanUseCachedMeasurement(widthMeasureMode,
			availableWidth,
			heightMeasureMode,
			availableHeight,
			layout.cachedLayout.widthMeasureMode,
			layout.cachedLayout.availableWidth,
			layout.cachedLayout.heightMeasureMode,
			layout.cachedLayout.availableHeight,
			layout.cachedLayout.computedWidth,
			layout.cachedLayout.computedHeight,
			marginAxisRow,
			marginAxisColumn,
			config) {
			cachedResults = &layout.cachedLayout
		} else {
			// Try to use the measurement cache.
			for i := 0; i < layout.nextCachedMeasurementsIndex; i++ {
				if YGNodeCanUseCachedMeasurement(widthMeasureMode,
					availableWidth,
					heightMeasureMode,
					availableHeight,
					layout.cachedMeasurements[i].widthMeasureMode,
					layout.cachedMeasurements[i].availableWidth,
					layout.cachedMeasurements[i].heightMeasureMode,
					layout.cachedMeasurements[i].availableHeight,
					layout.cachedMeasurements[i].computedWidth,
					layout.cachedMeasurements[i].computedHeight,
					marginAxisRow,
					marginAxisColumn,
					config) {
					cachedResults = &layout.cachedMeasurements[i]
					break
				}
			}
		}
	} else if performLayout {
		if YGFloatsEqual(layout.cachedLayout.availableWidth, availableWidth) &&
			YGFloatsEqual(layout.cachedLayout.availableHeight, availableHeight) &&
			layout.cachedLayout.widthMeasureMode == widthMeasureMode &&
			layout.cachedLayout.heightMeasureMode == heightMeasureMode {
			cachedResults = &layout.cachedLayout
		}
	} else {
		for i := 0; i < layout.nextCachedMeasurementsIndex; i++ {
			if YGFloatsEqual(layout.cachedMeasurements[i].availableWidth, availableWidth) &&
				YGFloatsEqual(layout.cachedMeasurements[i].availableHeight, availableHeight) &&
				layout.cachedMeasurements[i].widthMeasureMode == widthMeasureMode &&
				layout.cachedMeasurements[i].heightMeasureMode == heightMeasureMode {
				cachedResults = &layout.cachedMeasurements[i]
				break
			}
		}
	}

	if !needToVisitNode && cachedResults != nil {
		layout.measuredDimensions[YGDimensionWidth] = cachedResults.computedWidth
		layout.measuredDimensions[YGDimensionHeight] = cachedResults.computedHeight

		if gPrintChanges && gPrintSkips {
			fmt.Printf("%s%d.{[skipped] ", YGSpacer(gDepth), gDepth)
			if node.print != nil {
				node.print(node)
			}
			fmt.Printf("wm: %s, hm: %s, aw: %f ah: %f => d: (%f, %f) %s\n",
				YGMeasureModeName(widthMeasureMode, performLayout),
				YGMeasureModeName(heightMeasureMode, performLayout),
				availableWidth,
				availableHeight,
				cachedResults.computedWidth,
				cachedResults.computedHeight,
				reason)
		}
	} else {
		if gPrintChanges {
			s := ""
			if needToVisitNode {
				s = "*"
			}
			fmt.Printf("%s%d.{%s", YGSpacer(gDepth), gDepth, s)
			if node.print != nil {
				node.print(node)
			}
			fmt.Printf("wm: %s, hm: %s, aw: %f ah: %f %s\n",
				YGMeasureModeName(widthMeasureMode, performLayout),
				YGMeasureModeName(heightMeasureMode, performLayout),
				availableWidth,
				availableHeight,
				reason)
		}

		YGNodelayoutImpl(node,
			availableWidth,
			availableHeight,
			parentDirection,
			widthMeasureMode,
			heightMeasureMode,
			parentWidth,
			parentHeight,
			performLayout,
			config)

		if gPrintChanges {
			s := ""
			if needToVisitNode {
				s = "*"
			}
			fmt.Printf("%s%d.}%s", YGSpacer(gDepth), gDepth, s)
			if node.print != nil {
				node.print(node)
			}
			fmt.Printf("wm: %s, hm: %s, d: (%f, %f) %s\n",
				YGMeasureModeName(widthMeasureMode, performLayout),
				YGMeasureModeName(heightMeasureMode, performLayout),
				layout.measuredDimensions[YGDimensionWidth],
				layout.measuredDimensions[YGDimensionHeight],
				reason)
		}

		layout.lastParentDirection = parentDirection

		if cachedResults == nil {
			if layout.nextCachedMeasurementsIndex == YG_MAX_CACHED_RESULT_COUNT {
				if gPrintChanges {
					fmt.Printf("Out of cache entries!\n")
				}
				layout.nextCachedMeasurementsIndex = 0
			}

			var newCacheEntry *YGCachedMeasurement
			if performLayout {
				// Use the single layout cache entry.
				newCacheEntry = &layout.cachedLayout
			} else {
				// Allocate a new measurement cache entry.
				newCacheEntry = &layout.cachedMeasurements[layout.nextCachedMeasurementsIndex]
				layout.nextCachedMeasurementsIndex++
			}

			newCacheEntry.availableWidth = availableWidth
			newCacheEntry.availableHeight = availableHeight
			newCacheEntry.widthMeasureMode = widthMeasureMode
			newCacheEntry.heightMeasureMode = heightMeasureMode
			newCacheEntry.computedWidth = layout.measuredDimensions[YGDimensionWidth]
			newCacheEntry.computedHeight = layout.measuredDimensions[YGDimensionHeight]
		}
	}

	if performLayout {
		node.layout.dimensions[YGDimensionWidth] = node.layout.measuredDimensions[YGDimensionWidth]
		node.layout.dimensions[YGDimensionHeight] = node.layout.measuredDimensions[YGDimensionHeight]
		node.hasNewLayout = true
		node.isDirty = false
	}

	gDepth--
	layout.generationCount = gCurrentGenerationCount
	return (needToVisitNode || cachedResults == nil)
}

func YGNodeCalculateLayout(node *YGNode, parentWidth float32, parentHeight float32, parentDirection YGDirection) {
	// Increment the generation count. This will force the recursive routine to
	// visit
	// all dirty nodes at least once. Subsequent visits will be skipped if the
	// input
	// parameters don't change.
	gCurrentGenerationCount++

	YGResolveDimensions(node)

	width := YGUndefined
	widthMeasureMode := YGMeasureModeUndefined
	if YGNodeIsStyleDimDefined(node, YGFlexDirectionRow, parentWidth) {
		width = YGResolveValue(node.resolvedDimensions[dim[YGFlexDirectionRow]], parentWidth) +
			YGNodeMarginForAxis(node, YGFlexDirectionRow, parentWidth)
		widthMeasureMode = YGMeasureModeExactly
	} else if YGResolveValue(&node.style.maxDimensions[YGDimensionWidth], parentWidth) >= 0.0 {
		width = YGResolveValue(&node.style.maxDimensions[YGDimensionWidth], parentWidth)
		widthMeasureMode = YGMeasureModeAtMost
	} else {
		width = parentWidth
		widthMeasureMode = YGMeasureModeExactly
		if YGFloatIsUndefined(width) {
			widthMeasureMode = YGMeasureModeUndefined
		}
	}

	height := YGUndefined
	heightMeasureMode := YGMeasureModeUndefined
	if YGNodeIsStyleDimDefined(node, YGFlexDirectionColumn, parentHeight) {
		height = YGResolveValue(node.resolvedDimensions[dim[YGFlexDirectionColumn]], parentHeight) +
			YGNodeMarginForAxis(node, YGFlexDirectionColumn, parentWidth)
		heightMeasureMode = YGMeasureModeExactly
	} else if YGResolveValue(&node.style.maxDimensions[YGDimensionHeight], parentHeight) >= 0.0 {
		height = YGResolveValue(&node.style.maxDimensions[YGDimensionHeight], parentHeight)
		heightMeasureMode = YGMeasureModeAtMost
	} else {
		height = parentHeight
		heightMeasureMode = YGMeasureModeExactly
		if YGFloatIsUndefined(height) {
			heightMeasureMode = YGMeasureModeUndefined
		}
	}

	if YGLayoutNodeInternal(node,
		width,
		height,
		parentDirection,
		widthMeasureMode,
		heightMeasureMode,
		parentWidth,
		parentHeight,
		true,
		"initial",
		node.config) {
		YGNodeSetPosition(node, node.layout.direction, parentWidth, parentHeight, parentWidth)
		YGRoundToPixelGrid(node, node.config.pointScaleFactor, 0, 0)

		if gPrintTree {
			YGNodePrint(node, YGPrintOptionsLayout|YGPrintOptionsChildren|YGPrintOptionsStyle)
		}
	}
}

func YGRoundValueToPixelGrid(value float32, pointScaleFactor float32, forceCeil bool, forceFloor bool) float32 {
	scaledValue := value * pointScaleFactor
	fractial := fmodf(scaledValue, 1.0)
	if YGFloatsEqual(fractial, 0) {
		// Still remove fractial as fractial could be  extremely small.
		scaledValue = scaledValue - fractial
	} else if forceCeil {
		scaledValue = scaledValue - fractial + 1.0
	} else if forceFloor {
		scaledValue = scaledValue - fractial
	} else {
		var f float32
		if fractial >= 0.5 {
			f = 1.0
		}
		scaledValue = scaledValue - fractial + f
	}
	return scaledValue / pointScaleFactor
}

func YGRoundToPixelGrid(node *YGNode, pointScaleFactor float32, absoluteLeft float32, absoluteTop float32) {
	if pointScaleFactor == 0.0 {
		return
	}

	nodeLeft := node.layout.position[YGEdgeLeft]
	nodeTop := node.layout.position[YGEdgeTop]

	nodeWidth := node.layout.dimensions[YGDimensionWidth]
	nodeHeight := node.layout.dimensions[YGDimensionHeight]

	absoluteNodeLeft := absoluteLeft + nodeLeft
	absoluteNodeTop := absoluteTop + nodeTop

	absoluteNodeRight := absoluteNodeLeft + nodeWidth
	absoluteNodeBottom := absoluteNodeTop + nodeHeight

	// If a node has a custom measure function we never want to round down its size as this could
	// lead to unwanted text truncation.
	textRounding := node.nodeType == YGNodeTypeText

	node.layout.position[YGEdgeLeft] =
		YGRoundValueToPixelGrid(nodeLeft, pointScaleFactor, false, textRounding)
	node.layout.position[YGEdgeTop] =
		YGRoundValueToPixelGrid(nodeTop, pointScaleFactor, false, textRounding)

		// We multiply dimension by scale factor and if the result is close to the whole number, we don't have any fraction
		// To verify if the result is close to whole number we want to check both floor and ceil numbers
	hasFractionalWidth := !YGFloatsEqual(fmodf(nodeWidth*pointScaleFactor, 1.0), 0) &&
		!YGFloatsEqual(fmodf(nodeWidth*pointScaleFactor, 1.0), 1.0)
	hasFractionalHeight := !YGFloatsEqual(fmodf(nodeHeight*pointScaleFactor, 1.0), 0) &&
		!YGFloatsEqual(fmodf(nodeHeight*pointScaleFactor, 1.0), 1.0)

	node.layout.dimensions[YGDimensionWidth] =
		YGRoundValueToPixelGrid(
			absoluteNodeRight,
			pointScaleFactor,
			(textRounding && hasFractionalWidth),
			(textRounding && !hasFractionalWidth)) -
			YGRoundValueToPixelGrid(absoluteNodeLeft, pointScaleFactor, false, textRounding)
	node.layout.dimensions[YGDimensionHeight] =
		YGRoundValueToPixelGrid(
			absoluteNodeBottom,
			pointScaleFactor,
			(textRounding && hasFractionalHeight),
			(textRounding && !hasFractionalHeight)) -
			YGRoundValueToPixelGrid(absoluteNodeTop, pointScaleFactor, false, textRounding)

	childCount := YGNodeListCount(node.children)
	for i := 0; i < childCount; i++ {
		YGRoundToPixelGrid(YGNodeGetChild(node, i), pointScaleFactor, absoluteNodeLeft, absoluteNodeTop)
	}
}
