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

	computedFlexBasisGeneration int
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

func YGNodeSetMeasureFunc(node *YGNode, measureFunc YGMeasureFunc) {
	if measureFunc == nil {
		node.measure = nil
		// TODO: t18095186 Move nodeType to opt-in function and mark appropriate places in Litho
		node.nodeType = YGNodeTypeDefault
	} else {
		YGAssertWithNode(
			node,
			YGNodeGetChildCount(node) == 0,
			"Cannot set measure function: Nodes with measure functions cannot have children.")
		node.measure = measureFunc
		// TODO: t18095186 Move nodeType to opt-in function and mark appropriate places in Litho
		node.nodeType = YGNodeTypeText
	}
}

func YGNodeGetMeasureFunc(node *YGNode) YGMeasureFunc {
	return node.measure
}

func YGNodeSetBaselineFunc(node *YGNode, baselineFunc YGBaselineFunc) {
	node.baseline = baselineFunc
}

func YGNodeGetBaselineFunc(node *YGNode) YGBaselineFunc {
	return node.baseline
}

func styleEq(s1, s2 *YGStyle) bool {
	if s1.direction != s2.direction ||
		s1.flexDirection != s2.flexDirection ||
		s1.justifyContent != s2.justifyContent ||
		s1.alignContent != s2.alignContent ||
		s1.alignItems != s2.alignItems ||
		s1.alignSelf != s2.alignSelf ||
		s1.positionType != s2.positionType ||
		s1.flexWrap != s2.flexWrap ||
		s1.overflow != s2.overflow ||
		s1.display != s2.display ||
		s1.flex != s2.flex ||
		s1.flexGrow != s2.flexGrow ||
		s1.flexShrink != s2.flexShrink ||
		s1.flexBasis != s2.flexBasis {
		return false
	}
	for i := 0; i < YGEdgeCount; i++ {
		if s1.margin[i] != s2.margin[i] ||
			s1.position[i] != s2.position[i] ||
			s1.padding[i] != s2.padding[i] ||
			s1.border[i] != s2.border[i] {
			return false
		}
	}
	for i := 0; i < 2; i++ {
		if s1.dimensions[i] != s2.dimensions[i] ||
			s1.minDimensions[i] != s2.minDimensions[i] ||
			s1.maxDimensions[i] != s2.maxDimensions[i] {
			return false
		}
	}
	return true
}

func YGNodeCopyStyle(dstNode *YGNode, srcNode *YGNode) {
	if !styleEq(&dstNode.style, &srcNode.style) {
		dstNode.style = srcNode.style
		YGNodeMarkDirtyInternal(dstNode)
	}
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

func YGNodeGetParent(node *YGNode) *YGNode {
	return node.parent
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

func YGNodeSetContext(node *YGNode, context interface{}) {
	node.context = context
}

func YGNodeGetContext(node *YGNode) interface{} { return node.context }

func YGNodeSetPrintFunc(node *YGNode, printFunc YGPrintFunc) {
	node.print = printFunc
}

func YGNodeGetPrintFunc(node *YGNode) YGPrintFunc { return node.print }

func YGNodeSetHasNewLayout(node *YGNode, hasNewLayout bool) {
	node.hasNewLayout = hasNewLayout
}

func YGNodeGetHasNewLayout(node *YGNode) bool {
	return node.hasNewLayout
}

func YGNodeSetNodeType(node *YGNode, nodeType YGNodeType) {
	node.nodeType = nodeType
}

func YGNodeGetNodeType(node *YGNode) YGNodeType { return node.nodeType }

func YGNodeStyleSetDirection(node *YGNode, direction YGDirection) {
	if node.style.direction != direction {
		node.style.direction = direction
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetDirection(node *YGNode) YGDirection {
	return node.style.direction
}

func YGNodeStyleSetFlexDirection(node *YGNode, flexDirection YGFlexDirection) {
	if node.style.flexDirection != flexDirection {
		node.style.flexDirection = flexDirection
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetFlexDirection(node *YGNode) YGFlexDirection {
	return node.style.flexDirection
}

func YGNodeStyleSetJustifyContent(node *YGNode, justifyContent YGJustify) {
	if node.style.justifyContent != justifyContent {
		node.style.justifyContent = justifyContent
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetJustifyContent(node *YGNode) YGJustify {
	return node.style.justifyContent
}

func YGNodeStyleSetAlignContent(node *YGNode, alignContent YGAlign) {
	if node.style.alignContent != alignContent {
		node.style.alignContent = alignContent
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetAlignContent(node *YGNode) YGAlign {
	return node.style.alignContent
}

func YGNodeStyleSetAlignItems(node *YGNode, alignItems YGAlign) {
	if node.style.alignItems != alignItems {
		node.style.alignItems = alignItems
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetAlignItems(node *YGNode) YGAlign {
	return node.style.alignItems
}

func YGNodeStyleSetAlignSelf(node *YGNode, alignSelf YGAlign) {
	if node.style.alignSelf != alignSelf {
		node.style.alignSelf = alignSelf
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetAlignSelf(node *YGNode) YGAlign {
	return node.style.alignSelf
}

func YGNodeStyleSetFlexWrap(node *YGNode, flexWrap YGWrap) {
	if node.style.flexWrap != flexWrap {
		node.style.flexWrap = flexWrap
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetFlexWrap(node *YGNode) YGWrap {
	return node.style.flexWrap
}

func YGNodeStyleSetOverflow(node *YGNode, overflow YGOverflow) {
	if node.style.overflow != overflow {
		node.style.overflow = overflow
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetOverflow(node *YGNode) YGOverflow {
	return node.style.overflow
}

func YGNodeStyleSetDisplay(node *YGNode, display YGDisplay) {
	if node.style.display != display {
		node.style.display = display
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetDisplay(node *YGNode) YGDisplay {
	return node.style.display
}

func YGNodeStyleSetFlex(node *YGNode, flex float32) {
	if node.style.flex != flex {
		node.style.flex = flex
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetFlex(node *YGNode) float32 { return node.style.flex }

func YGNodeStyleSetFlexGrow(node *YGNode, flexGrow float32) {
	if node.style.flexGrow != flexGrow {
		node.style.flexGrow = flexGrow
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetFlexShrink(node *YGNode, flexShrink float32) {
	if node.style.flexShrink != flexShrink {
		node.style.flexShrink = flexShrink
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetFlexBasis(node *YGNode, flexBasis float32) {
	if node.style.flexBasis.value != flexBasis ||
		node.style.flexBasis.unit != YGUnitPoint {
		node.style.flexBasis.value = flexBasis
		node.style.flexBasis.unit = YGUnitPoint
		if YGFloatIsUndefined(flexBasis) {
			node.style.flexBasis.unit = YGUnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetFlexBasisPercent(node *YGNode, flexBasis float32) {
	if node.style.flexBasis.value != flexBasis ||
		node.style.flexBasis.unit != YGUnitPercent {
		node.style.flexBasis.value = flexBasis
		node.style.flexBasis.unit = YGUnitPercent
		if YGFloatIsUndefined(flexBasis) {
			node.style.flexBasis.unit = YGUnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetFlexBasisAuto(node *YGNode) {
	if node.style.flexBasis.unit != YGUnitAuto {
		node.style.flexBasis.value = YGUndefined
		node.style.flexBasis.unit = YGUnitAuto
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetFlexBasis(node *YGNode) YGValue {
	return node.style.flexBasis
}

func YGNodeStyleSetMargin(node *YGNode, edge YGEdge, margin float32) {
	if node.style.margin[edge].value != margin ||
		node.style.margin[edge].unit != YGUnitPoint {
		node.style.margin[edge].value = margin
		node.style.margin[edge].unit = YGUnitPoint
		if YGFloatIsUndefined(margin) {
			node.style.margin[edge].unit = YGUnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetMarginPercent(node *YGNode, edge YGEdge, margin float32) {
	if node.style.margin[edge].value != margin ||
		node.style.margin[edge].unit != YGUnitPercent {
		node.style.margin[edge].value = margin
		node.style.margin[edge].unit = YGUnitPercent
		if YGFloatIsUndefined(margin) {
			node.style.margin[edge].unit = YGUnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetMargin(node *YGNode, edge YGEdge) YGValue {
	return node.style.margin[edge]
}

func YGNodeStyleSetMarginAuto(node *YGNode, edge YGEdge) {
	if node.style.margin[edge].unit != YGUnitAuto {
		node.style.margin[edge].value = YGUndefined
		node.style.margin[edge].unit = YGUnitAuto
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetPadding(node *YGNode, edge YGEdge, padding float32) {
	if node.style.padding[edge].value != padding ||
		node.style.padding[edge].unit != YGUnitPoint {
		node.style.padding[edge].value = padding
		node.style.padding[edge].unit = YGUnitPoint
		if YGFloatIsUndefined(padding) {
			node.style.padding[edge].unit = YGUnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetPaddingPercent(node *YGNode, edge YGEdge, padding float32) {
	if node.style.padding[edge].value != padding ||
		node.style.padding[edge].unit != YGUnitPercent {
		node.style.padding[edge].value = padding
		node.style.padding[edge].unit = YGUnitPercent
		if YGFloatIsUndefined(padding) {
			node.style.padding[edge].unit = YGUnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetPadding(node *YGNode, edge YGEdge) YGValue {
	return node.style.padding[edge]
}

func YGNodeStyleSetBorder(node *YGNode, edge YGEdge, border float32) {
	if node.style.border[edge].value != border ||
		node.style.border[edge].unit != YGUnitPoint {
		node.style.border[edge].value = border
		node.style.border[edge].unit = YGUnitPoint
		if YGFloatIsUndefined(border) {
			node.style.border[edge].unit = YGUnitUndefined
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetBorder(node *YGNode, edge YGEdge) float32 {
	return node.style.border[edge].value
}

func YGNodeStyleSetMinWidth(node *YGNode, minWidth float32) {
	if node.style.minDimensions[YGDimensionWidth].value != minWidth ||
		node.style.minDimensions[YGDimensionWidth].unit != YGUnitPoint {
		node.style.minDimensions[YGDimensionWidth].value = minWidth
		node.style.minDimensions[YGDimensionWidth].unit = YGUnitPoint
		if YGFloatIsUndefined(minWidth) {
			node.style.minDimensions[YGDimensionWidth].unit = YGUnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetMinWidthPercent(node *YGNode, minWidth float32) {
	if node.style.minDimensions[YGDimensionWidth].value != minWidth ||
		node.style.minDimensions[YGDimensionWidth].unit != YGUnitPercent {
		node.style.minDimensions[YGDimensionWidth].value = minWidth
		node.style.minDimensions[YGDimensionWidth].unit = YGUnitPercent
		if YGFloatIsUndefined(minWidth) {
			node.style.minDimensions[YGDimensionWidth].unit = YGUnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetMinWidth(node *YGNode) YGValue {
	return node.style.minDimensions[YGDimensionWidth]
}

func YGNodeStyleSetMinHeight(node *YGNode, minHeight float32) {
	if node.style.minDimensions[YGDimensionHeight].value != minHeight ||
		node.style.minDimensions[YGDimensionHeight].unit != YGUnitPoint {
		node.style.minDimensions[YGDimensionHeight].value = minHeight
		node.style.minDimensions[YGDimensionHeight].unit = YGUnitPoint
		if YGFloatIsUndefined(minHeight) {
			node.style.minDimensions[YGDimensionHeight].unit = YGUnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetMinHeightPercent(node *YGNode, minHeight float32) {
	if node.style.minDimensions[YGDimensionHeight].value != minHeight ||
		node.style.minDimensions[YGDimensionHeight].unit != YGUnitPercent {
		node.style.minDimensions[YGDimensionHeight].value = minHeight
		node.style.minDimensions[YGDimensionHeight].unit = YGUnitPercent
		if YGFloatIsUndefined(minHeight) {
			node.style.minDimensions[YGDimensionHeight].unit = YGUnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetMinHeight(node *YGNode) YGValue {
	return node.style.minDimensions[YGDimensionHeight]
}

func YGNodeStyleSetMaxWidth(node *YGNode, maxWidth float32) {
	if node.style.maxDimensions[YGDimensionWidth].value != maxWidth ||
		node.style.maxDimensions[YGDimensionWidth].unit != YGUnitPoint {
		node.style.maxDimensions[YGDimensionWidth].value = maxWidth
		node.style.maxDimensions[YGDimensionWidth].unit = YGUnitPoint
		if YGFloatIsUndefined(maxWidth) {
			node.style.maxDimensions[YGDimensionWidth].unit = YGUnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetMaxWidthPercent(node *YGNode, maxWidth float32) {
	if node.style.maxDimensions[YGDimensionWidth].value != maxWidth ||
		node.style.maxDimensions[YGDimensionWidth].unit != YGUnitPercent {
		node.style.maxDimensions[YGDimensionWidth].value = maxWidth
		node.style.maxDimensions[YGDimensionWidth].unit = YGUnitPercent
		if YGFloatIsUndefined(maxWidth) {
			node.style.maxDimensions[YGDimensionWidth].unit = YGUnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetMaxWidth(node *YGNode) YGValue {
	return node.style.maxDimensions[YGDimensionWidth]
}

func YGNodeStyleSetMaxHeight(node *YGNode, maxHeight float32) {
	if node.style.maxDimensions[YGDimensionHeight].value != maxHeight ||
		node.style.maxDimensions[YGDimensionHeight].unit != YGUnitPoint {
		node.style.maxDimensions[YGDimensionHeight].value = maxHeight
		node.style.maxDimensions[YGDimensionHeight].unit = YGUnitPoint
		if YGFloatIsUndefined(maxHeight) {
			node.style.maxDimensions[YGDimensionHeight].unit = YGUnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleSetMaxHeightPercent(node *YGNode, maxHeight float32) {
	if node.style.maxDimensions[YGDimensionHeight].value != maxHeight ||
		node.style.maxDimensions[YGDimensionHeight].unit != YGUnitPercent {
		node.style.maxDimensions[YGDimensionHeight].value = maxHeight
		node.style.maxDimensions[YGDimensionHeight].unit = YGUnitPercent
		if YGFloatIsUndefined(maxHeight) {
			node.style.maxDimensions[YGDimensionHeight].unit = YGUnitAuto
		}
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetMaxHeight(node *YGNode) YGValue {
	return node.style.maxDimensions[YGDimensionHeight]
}

func YGNodeStyleSetAspectRatio(node *YGNode, aspectRatio float32) {
	if node.style.aspectRatio != aspectRatio {
		node.style.aspectRatio = aspectRatio
		YGNodeMarkDirtyInternal(node)
	}
}

func YGNodeStyleGetAspectRatio(node *YGNode) float32 {
	return node.style.aspectRatio
}

func YGNodeLayoutGetLeft(node *YGNode) float32 {
	return node.layout.position[YGEdgeLeft]
}

func YGNodeLayoutGetTop(node *YGNode) float32 {
	return node.layout.position[YGEdgeTop]
}

func YGNodeLayoutGetRight(node *YGNode) float32 {
	return node.layout.position[YGEdgeRight]
}

func YGNodeLayoutGetBottom(node *YGNode) float32 {
	return node.layout.position[YGEdgeBottom]
}

func YGNodeLayoutGetWidth(node *YGNode) float32 {
	return node.layout.dimensions[YGDimensionWidth]
}

func YGNodeLayoutGetHeight(node *YGNode) float32 {
	return node.layout.dimensions[YGDimensionHeight]
}

func YGNodeLayoutGetDirection(node *YGNode) YGDirection {
	return node.layout.direction
}

func YGNodeLayoutGetHadOverflow(node *YGNode) bool {
	return node.layout.hadOverflow
}

func YGNodeLayoutGetMargin(node *YGNode, edge YGEdge) float32 {
	YGAssertWithNode(node, edge < YGEdgeEnd, "Cannot get layout properties of multi-edge shorthands")
	if edge == YGEdgeLeft {
		if node.layout.direction == YGDirectionRTL {
			return node.layout.margin[YGEdgeEnd]
		} else {
			return node.layout.margin[YGEdgeStart]
		}
	}
	if edge == YGEdgeRight {
		if node.layout.direction == YGDirectionRTL {
			return node.layout.margin[YGEdgeStart]
		} else {
			return node.layout.margin[YGEdgeEnd]
		}
	}
	return node.layout.margin[edge]
}

func YGNodeLayoutGetBorder(node *YGNode, edge YGEdge) float32 {
	YGAssertWithNode(node, edge < YGEdgeEnd,
		"Cannot get layout properties of multi-edge shorthands")
	if edge == YGEdgeLeft {
		if node.layout.direction == YGDirectionRTL {
			return node.layout.border[YGEdgeEnd]
		} else {
			return node.layout.border[YGEdgeStart]
		}
	}
	if edge == YGEdgeRight {
		if node.layout.direction == YGDirectionRTL {
			return node.layout.border[YGEdgeStart]
		} else {
			return node.layout.border[YGEdgeEnd]
		}
	}
	return node.layout.border[edge]
}

func YGNodeLayoutGetPadding(node *YGNode, edge YGEdge) float32 {
	YGAssertWithNode(node, edge < YGEdgeEnd,
		"Cannot get layout properties of multi-edge shorthands")
	if edge == YGEdgeLeft {
		if node.layout.direction == YGDirectionRTL {
			return node.layout.padding[YGEdgeEnd]
		} else {
			return node.layout.padding[YGEdgeStart]
		}
	}
	if edge == YGEdgeRight {
		if node.layout.direction == YGDirectionRTL {
			return node.layout.padding[YGEdgeStart]
		} else {
			return node.layout.padding[YGEdgeEnd]
		}
	}
	return node.layout.padding[edge]
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

func YGBaseline(node *YGNode) float32 {
	if node.baseline != nil {
		baseline := node.baseline(node, node.layout.measuredDimensions[YGDimensionWidth], node.layout.measuredDimensions[YGDimensionHeight])
		YGAssertWithNode(node, !YGFloatIsUndefined(baseline), "Expect custom baseline function to not return NaN")
		return baseline
	}

	var baselineChild *YGNode
	childCount := YGNodeGetChildCount(node)
	for i := 0; i < childCount; i++ {
		child := YGNodeGetChild(node, i)
		if child.lineIndex > 0 {
			break
		}
		if child.style.positionType == YGPositionTypeAbsolute {
			continue
		}
		if YGNodeAlignItem(node, child) == YGAlignBaseline {
			baselineChild = child
			break
		}

		if baselineChild == nil {
			baselineChild = child
		}
	}

	if baselineChild == nil {
		return node.layout.measuredDimensions[YGDimensionHeight]
	}

	baseline := YGBaseline(baselineChild)
	return baseline + baselineChild.layout.position[YGEdgeTop]
}

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

func YGNodeComputeFlexBasisForChild(node *YGNode,
	child *YGNode,
	width float32,
	widthMode YGMeasureMode,
	height float32,
	parentWidth float32,
	parentHeight float32,
	heightMode YGMeasureMode,
	direction YGDirection,
	config YGConfigRef) {
	mainAxis := YGResolveFlexDirection(node.style.flexDirection, direction)
	isMainAxisRow := YGFlexDirectionIsRow(mainAxis)
	mainAxisSize := height
	mainAxisParentSize := parentHeight
	if isMainAxisRow {
		mainAxisSize = width
		mainAxisParentSize = parentWidth
	}

	var childWidth float32
	var childHeight float32
	var childWidthMeasureMode YGMeasureMode
	var childHeightMeasureMode YGMeasureMode

	resolvedFlexBasis := YGResolveValue(YGNodeResolveFlexBasisPtr(child), mainAxisParentSize)
	isRowStyleDimDefined := YGNodeIsStyleDimDefined(child, YGFlexDirectionRow, parentWidth)
	isColumnStyleDimDefined := YGNodeIsStyleDimDefined(child, YGFlexDirectionColumn, parentHeight)

	if !YGFloatIsUndefined(resolvedFlexBasis) && !YGFloatIsUndefined(mainAxisSize) {
		if YGFloatIsUndefined(child.layout.computedFlexBasis) ||
			(YGConfigIsExperimentalFeatureEnabled(child.config, YGExperimentalFeatureWebFlexBasis) &&
				child.layout.computedFlexBasisGeneration != gCurrentGenerationCount) {
			child.layout.computedFlexBasis =
				fmaxf(resolvedFlexBasis, YGNodePaddingAndBorderForAxis(child, mainAxis, parentWidth))
		}
	} else if isMainAxisRow && isRowStyleDimDefined {
		// The width is definite, so use that as the flex basis.
		child.layout.computedFlexBasis =
			fmaxf(YGResolveValue(child.resolvedDimensions[YGDimensionWidth], parentWidth),
				YGNodePaddingAndBorderForAxis(child, YGFlexDirectionRow, parentWidth))
	} else if !isMainAxisRow && isColumnStyleDimDefined {
		// The height is definite, so use that as the flex basis.
		child.layout.computedFlexBasis =
			fmaxf(YGResolveValue(child.resolvedDimensions[YGDimensionHeight], parentHeight),
				YGNodePaddingAndBorderForAxis(child, YGFlexDirectionColumn, parentWidth))
	} else {
		// Compute the flex basis and hypothetical main size (i.e. the clamped
		// flex basis).
		childWidth = YGUndefined
		childHeight = YGUndefined
		childWidthMeasureMode = YGMeasureModeUndefined
		childHeightMeasureMode = YGMeasureModeUndefined

		marginRow := YGNodeMarginForAxis(child, YGFlexDirectionRow, parentWidth)
		marginColumn := YGNodeMarginForAxis(child, YGFlexDirectionColumn, parentWidth)

		if isRowStyleDimDefined {
			childWidth =
				YGResolveValue(child.resolvedDimensions[YGDimensionWidth], parentWidth) + marginRow
			childWidthMeasureMode = YGMeasureModeExactly
		}
		if isColumnStyleDimDefined {
			childHeight =
				YGResolveValue(child.resolvedDimensions[YGDimensionHeight], parentHeight) + marginColumn
			childHeightMeasureMode = YGMeasureModeExactly
		}

		// The W3C spec doesn't say anything about the 'overflow' property,
		// but all major browsers appear to implement the following logic.
		if (!isMainAxisRow && node.style.overflow == YGOverflowScroll) ||
			node.style.overflow != YGOverflowScroll {
			if YGFloatIsUndefined(childWidth) && !YGFloatIsUndefined(width) {
				childWidth = width
				childWidthMeasureMode = YGMeasureModeAtMost
			}
		}

		if (isMainAxisRow && node.style.overflow == YGOverflowScroll) ||
			node.style.overflow != YGOverflowScroll {
			if YGFloatIsUndefined(childHeight) && !YGFloatIsUndefined(height) {
				childHeight = height
				childHeightMeasureMode = YGMeasureModeAtMost
			}
		}

		// If child has no defined size in the cross axis and is set to stretch,
		// set the cross
		// axis to be measured exactly with the available inner width
		if !isMainAxisRow && !YGFloatIsUndefined(width) && !isRowStyleDimDefined &&
			widthMode == YGMeasureModeExactly && YGNodeAlignItem(node, child) == YGAlignStretch {
			childWidth = width
			childWidthMeasureMode = YGMeasureModeExactly
		}
		if isMainAxisRow && !YGFloatIsUndefined(height) && !isColumnStyleDimDefined &&
			heightMode == YGMeasureModeExactly && YGNodeAlignItem(node, child) == YGAlignStretch {
			childHeight = height
			childHeightMeasureMode = YGMeasureModeExactly
		}

		if !YGFloatIsUndefined(child.style.aspectRatio) {
			if !isMainAxisRow && childWidthMeasureMode == YGMeasureModeExactly {
				child.layout.computedFlexBasis =
					fmaxf((childWidth-marginRow)/child.style.aspectRatio,
						YGNodePaddingAndBorderForAxis(child, YGFlexDirectionColumn, parentWidth))
				return
			} else if isMainAxisRow && childHeightMeasureMode == YGMeasureModeExactly {
				child.layout.computedFlexBasis =
					fmaxf((childHeight-marginColumn)*child.style.aspectRatio,
						YGNodePaddingAndBorderForAxis(child, YGFlexDirectionRow, parentWidth))
				return
			}
		}

		YGConstrainMaxSizeForMode(
			child, YGFlexDirectionRow, parentWidth, parentWidth, &childWidthMeasureMode, &childWidth)
		YGConstrainMaxSizeForMode(child,
			YGFlexDirectionColumn,
			parentHeight,
			parentWidth,
			&childHeightMeasureMode,
			&childHeight)

		// Measure the child
		YGLayoutNodeInternal(child,
			childWidth,
			childHeight,
			direction,
			childWidthMeasureMode,
			childHeightMeasureMode,
			parentWidth,
			parentHeight,
			false,
			"measure",
			config)

		child.layout.computedFlexBasis =
			fmaxf(child.layout.measuredDimensions[dim[mainAxis]],
				YGNodePaddingAndBorderForAxis(child, mainAxis, parentWidth))
	}

	child.layout.computedFlexBasisGeneration = gCurrentGenerationCount
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

func YGIsBaselineLayout(node *YGNode) bool {
	if YGFlexDirectionIsColumn(node.style.flexDirection) {
		return false
	}
	if node.style.alignItems == YGAlignBaseline {
		return true
	}
	childCount := YGNodeGetChildCount(node)
	for i := 0; i < childCount; i++ {
		child := YGNodeGetChild(node, i)
		if child.style.positionType == YGPositionTypeRelative &&
			child.style.alignSelf == YGAlignBaseline {
			return true
		}
	}

	return false
}

func YGNodeBoundAxisWithinMinAndMax(node *YGNode, axis YGFlexDirection, value float32, axisSize float32) float32 {
	min := YGUndefined
	max := YGUndefined

	if YGFlexDirectionIsColumn(axis) {
		min = YGResolveValue(&node.style.minDimensions[YGDimensionHeight], axisSize)
		max = YGResolveValue(&node.style.maxDimensions[YGDimensionHeight], axisSize)
	} else if YGFlexDirectionIsRow(axis) {
		min = YGResolveValue(&node.style.minDimensions[YGDimensionWidth], axisSize)
		max = YGResolveValue(&node.style.maxDimensions[YGDimensionWidth], axisSize)
	}

	boundValue := value

	if !YGFloatIsUndefined(max) && max >= 0 && boundValue > max {
		boundValue = max
	}

	if !YGFloatIsUndefined(min) && min >= 0 && boundValue < min {
		boundValue = min
	}

	return boundValue
}

// Like YGNodeBoundAxisWithinMinAndMax but also ensures that the value doesn't go
// below the
// padding and border amount.
func YGNodeBoundAxis(node *YGNode, axis YGFlexDirection, value float32, axisSize float32, widthSize float32) float32 {
	return fmaxf(YGNodeBoundAxisWithinMinAndMax(node, axis, value, axisSize),
		YGNodePaddingAndBorderForAxis(node, axis, widthSize))
}

func YGNodeSetChildTrailingPosition(node *YGNode, child *YGNode, axis YGFlexDirection) {
	size := child.layout.measuredDimensions[dim[axis]]
	child.layout.position[trailing[axis]] =
		node.layout.measuredDimensions[dim[axis]] - size - child.layout.position[pos[axis]]
}

func YGConstrainMaxSizeForMode(node *YGNode, axis YGFlexDirection, parentAxisSize float32, parentWidth float32, mode *YGMeasureMode, size *float32) {
	maxSize := YGResolveValue(&node.style.maxDimensions[dim[axis]], parentAxisSize) +
		YGNodeMarginForAxis(node, axis, parentWidth)
	switch *mode {
	case YGMeasureModeExactly:
	case YGMeasureModeAtMost:
		if YGFloatIsUndefined(maxSize) || *size < maxSize {
			// TODO: this is redundant, but what is in original code
			*size = *size
		} else {
			*size = maxSize
		}

		break
	case YGMeasureModeUndefined:
		if !YGFloatIsUndefined(maxSize) {
			*mode = YGMeasureModeAtMost
			*size = maxSize
		}
		break
	}
}

func YGNodeFixedSizeSetMeasuredDimensions(node *YGNode,
	availableWidth float32,
	availableHeight float32,
	widthMeasureMode YGMeasureMode,
	heightMeasureMode YGMeasureMode,
	parentWidth float32,
	parentHeight float32) bool {
	if (widthMeasureMode == YGMeasureModeAtMost && availableWidth <= 0) ||
		(heightMeasureMode == YGMeasureModeAtMost && availableHeight <= 0) ||
		(widthMeasureMode == YGMeasureModeExactly && heightMeasureMode == YGMeasureModeExactly) {
		marginAxisColumn := YGNodeMarginForAxis(node, YGFlexDirectionColumn, parentWidth)
		marginAxisRow := YGNodeMarginForAxis(node, YGFlexDirectionRow, parentWidth)

		width := availableWidth - marginAxisRow
		if YGFloatIsUndefined(availableWidth) || (widthMeasureMode == YGMeasureModeAtMost && availableWidth < 0) {
			width = 0
		}
		node.layout.measuredDimensions[YGDimensionWidth] =
			YGNodeBoundAxis(node, YGFlexDirectionRow, width, parentWidth, parentWidth)

		height := availableHeight - marginAxisColumn
		if YGFloatIsUndefined(availableHeight) || (heightMeasureMode == YGMeasureModeAtMost && availableHeight < 0) {
			height = 0
		}
		node.layout.measuredDimensions[YGDimensionHeight] =
			YGNodeBoundAxis(node, YGFlexDirectionColumn, height, parentHeight, parentWidth)

		return true
	}

	return false
}

// For nodes with no children, use the available values if they were provided,
// or the minimum size as indicated by the padding and border sizes.
func YGNodeEmptyContainerSetMeasuredDimensions(node *YGNode, availableWidth float32, availableHeight float32, widthMeasureMode YGMeasureMode, heightMeasureMode YGMeasureMode, parentWidth float32, parentHeight float32) {
	paddingAndBorderAxisRow := YGNodePaddingAndBorderForAxis(node, YGFlexDirectionRow, parentWidth)
	paddingAndBorderAxisColumn := YGNodePaddingAndBorderForAxis(node, YGFlexDirectionColumn, parentWidth)
	marginAxisRow := YGNodeMarginForAxis(node, YGFlexDirectionRow, parentWidth)
	marginAxisColumn := YGNodeMarginForAxis(node, YGFlexDirectionColumn, parentWidth)

	width := availableWidth - marginAxisRow
	if widthMeasureMode == YGMeasureModeUndefined || widthMeasureMode == YGMeasureModeAtMost {
		width = paddingAndBorderAxisRow
	}
	node.layout.measuredDimensions[YGDimensionWidth] = YGNodeBoundAxis(node, YGFlexDirectionRow, width, parentWidth, parentWidth)

	height := availableHeight - marginAxisColumn
	if heightMeasureMode == YGMeasureModeUndefined || heightMeasureMode == YGMeasureModeAtMost {
		height = paddingAndBorderAxisColumn
	}
	node.layout.measuredDimensions[YGDimensionHeight] = YGNodeBoundAxis(node, YGFlexDirectionColumn, height, parentHeight, parentWidth)
}

func YGNodeWithMeasureFuncSetMeasuredDimensions(node *YGNode, availableWidth float32, availableHeight float32, widthMeasureMode YGMeasureMode, heightMeasureMode YGMeasureMode, parentWidth float32, parentHeight float32) {
	YGAssertWithNode(node, node.measure != nil, "Expected node to have custom measure function")

	paddingAndBorderAxisRow := YGNodePaddingAndBorderForAxis(node, YGFlexDirectionRow, availableWidth)
	paddingAndBorderAxisColumn := YGNodePaddingAndBorderForAxis(node, YGFlexDirectionColumn, availableWidth)
	marginAxisRow := YGNodeMarginForAxis(node, YGFlexDirectionRow, availableWidth)
	marginAxisColumn := YGNodeMarginForAxis(node, YGFlexDirectionColumn, availableWidth)

	// We want to make sure we don't call measure with negative size
	innerWidth := availableWidth
	if !YGFloatIsUndefined(availableWidth) {
		innerWidth = fmaxf(0, availableWidth-marginAxisRow-paddingAndBorderAxisRow)
	}
	innerHeight := availableHeight
	if !YGFloatIsUndefined(availableHeight) {
		innerHeight = fmaxf(0, availableHeight-marginAxisColumn-paddingAndBorderAxisColumn)
	}

	if widthMeasureMode == YGMeasureModeExactly && heightMeasureMode == YGMeasureModeExactly {
		// Don't bother sizing the text if both dimensions are already defined.
		node.layout.measuredDimensions[YGDimensionWidth] = YGNodeBoundAxis(
			node, YGFlexDirectionRow, availableWidth-marginAxisRow, parentWidth, parentWidth)
		node.layout.measuredDimensions[YGDimensionHeight] = YGNodeBoundAxis(
			node, YGFlexDirectionColumn, availableHeight-marginAxisColumn, parentHeight, parentWidth)
	} else {
		// Measure the text under the current raints.
		measuredSize := node.measure(node, innerWidth, widthMeasureMode, innerHeight, heightMeasureMode)

		width := availableWidth - marginAxisRow
		if widthMeasureMode == YGMeasureModeUndefined ||
			widthMeasureMode == YGMeasureModeAtMost {
			width = measuredSize.width + paddingAndBorderAxisRow

		}

		node.layout.measuredDimensions[YGDimensionWidth] = YGNodeBoundAxis(node, YGFlexDirectionRow, width, availableWidth, availableWidth)

		height := availableHeight - marginAxisColumn
		if heightMeasureMode == YGMeasureModeUndefined ||
			heightMeasureMode == YGMeasureModeAtMost {
			height = measuredSize.height + paddingAndBorderAxisColumn

		}

		node.layout.measuredDimensions[YGDimensionHeight] = YGNodeBoundAxis(node, YGFlexDirectionColumn, height, availableHeight, availableWidth)
	}
}

func YGZeroOutLayoutRecursivly(node *YGNode) {
	node.layout.dimensions[YGDimensionHeight] = 0
	node.layout.dimensions[YGDimensionWidth] = 0
	node.layout.position[YGEdgeTop] = 0
	node.layout.position[YGEdgeBottom] = 0
	node.layout.position[YGEdgeLeft] = 0
	node.layout.position[YGEdgeRight] = 0
	node.layout.cachedLayout.availableHeight = 0
	node.layout.cachedLayout.availableWidth = 0
	node.layout.cachedLayout.heightMeasureMode = YGMeasureModeExactly
	node.layout.cachedLayout.widthMeasureMode = YGMeasureModeExactly
	node.layout.cachedLayout.computedWidth = 0
	node.layout.cachedLayout.computedHeight = 0
	node.hasNewLayout = true
	childCount := YGNodeGetChildCount(node)
	for i := 0; i < childCount; i++ {
		child := YGNodeListGet(node.children, i)
		YGZeroOutLayoutRecursivly(child)
	}
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

func YGConfigSetExperimentalFeatureEnabled(config YGConfigRef,
	feature YGExperimentalFeature,
	enabled bool) {
	config.experimentalFeatures[feature] = enabled
}

func YGConfigIsExperimentalFeatureEnabled(config YGConfigRef, feature YGExperimentalFeature) bool {
	return config.experimentalFeatures[feature]
}

func YGConfigSetUseWebDefaults(config YGConfigRef, enabled bool) {
	config.useWebDefaults = enabled
}

func YGConfigSetUseLegacyStretchBehaviour(config YGConfigRef, useLegacyStretchBehaviour bool) {
	config.useLegacyStretchBehaviour = useLegacyStretchBehaviour
}

func YGConfigGetUseWebDefaults(config YGConfigRef) bool {
	return config.useWebDefaults
}

func YGConfigSetContext(config YGConfigRef, context interface{}) {
	config.context = context
}

func YGConfigGetContext(config YGConfigRef) interface{} {
	return config.context
}
