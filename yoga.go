package flex

import (
	"fmt"
	"os"
)

// YGCachedMeasurement describes measurements
type YGCachedMeasurement struct {
	availableWidth    float32
	availableHeight   float32
	widthMeasureMode  YGMeasureMode
	heightMeasureMode YGMeasureMode

	computedWidth  float32
	computedHeight float32
}

// This value was chosen based on empiracle data. Even the most complicated
// layouts should not require more than 16 entries to fit within the cache.
const YG_MAX_CACHED_RESULT_COUNT = 16

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

// YGNode describes a node
type YGNode struct {
	style     YGStyle
	layout    YGLayout
	lineIndex int

	parent   *YGNode
	children *YGNodeList

	nextChild *YGNode

	measure  YGMeasureFunc
	baseline YGBaselineFunc
	print    YGPrintFunc
	config   *YGConfig
	context  interface{}

	isDirty      bool
	hasNewLayout bool
	nodeType     YGNodeType

	resolvedDimensions [2]*YGValue
}

var (
	YG_UNDEFINED_VALUES = YGValue{
		value: YGUndefined,
		unit:  YGUnitUndefined,
	}

	YG_AUTO_VALUES = YGValue{
		value: YGUndefined,
		unit:  YGUnitAuto,
	}

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
	kDefaultFlexGrow      float32 = 0
	kDefaultFlexShrink    float32 = 0
	kWebDefaultFlexShrink float32 = 1
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

	gYGConfigDefaults = YGConfig{
		experimentalFeatures: [YGExperimentalFeatureCount + 1]bool{
			false,
			false,
		},
		useWebDefaults:   false,
		pointScaleFactor: 1,
		logger:           YGDefaultLog,
		context:          nil,
	}

	YGValueZero = YGValue{value: 0, unit: YGUnitPoint}
)

// YGDefaultLog is default logging function
func YGDefaultLog(config *YGConfig, node *YGNode, level YGLogLevel, format string,
	args ...interface{}) int {
	switch level {
	case YGLogLevelError, YGLogLevelFatal:
		n, _ := fmt.Fprintf(os.Stderr, format, args...)
		return n
	case YGLogLevelWarn, YGLogLevelInfo, YGLogLevelDebug, YGLogLevelVerbose:
		fallthrough
	default:
		n, _ := fmt.Printf(format, args...)
		return n
	}
}

// YGComputedEdgeValue computes edge value
func YGComputedEdgeValue(edges []YGValue, edge YGEdge, defaultValue *YGValue) *YGValue {
	if edges[edge].unit != YGUnitUndefined {
		return &edges[edge]
	}

	isVertEdge := edge == YGEdgeTop || edge == YGEdgeBottom
	if isVertEdge && edges[YGEdgeVertical].unit != YGUnitUndefined {
		return &edges[YGEdgeVertical]
	}

	isHorizEdge := (edge == YGEdgeLeft || edge == YGEdgeRight || edge == YGEdgeStart || edge == YGEdgeEnd)
	if isHorizEdge && edges[YGEdgeHorizontal].unit != YGUnitUndefined {
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
	case YGUnitUndefined, YGUnitAuto:
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
	gNodeInstanceCount   = 0
	gConfigInstanceCount = 0
)

// YGNodeNewWithConfig creates new node with config
func YGNodeNewWithConfig(config *YGConfig) *YGNode {
	node := gYGNodeDefaults
	gNodeInstanceCount++

	if config.useWebDefaults {
		node.style.flexDirection = YGFlexDirectionRow
		node.style.alignContent = YGAlignStretch
	}
	node.config = config
	return &node
}

// YGNodeNew creates a new node
func YGNodeNew() *YGNode {
	return YGNodeNewWithConfig(&gYGConfigDefaults)
}

// YGNodeFree frees a node
func YGNodeFree(node *YGNode) {
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

// YGNodeFreeRecursive frees root recursevily
func YGNodeFreeRecursive(root *YGNode) {
	for YGNodeGetChildCount(root) > 0 {
		child := YGNodeGetChild(root, 0)
		YGNodeRemoveChild(root, child)
		YGNodeFreeRecursive(child)
	}
	YGNodeFree(root)
}

// YGNodeReset resets a node
func YGNodeReset(node *YGNode) {
	YGAssertWithNode(node, YGNodeGetChildCount(node) == 0, "Cannot reset a node which still has children attached")
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
func YGConfigGetDefault() *YGConfig {
	return &gYGConfigDefaults
}

// YGConfigNew creates new config
func YGConfigNew() *YGConfig {
	config := &YGConfig{}
	YGAssert(config != nil, "Could not allocate memory for config")

	gConfigInstanceCount++
	*config = gYGConfigDefaults
	return config
}

// YGConfigFree frees a config
func YGConfigFree(config *YGConfig) {
	gConfigInstanceCount--
}

// YGConfigCopy copies a config
func YGConfigCopy(dest *YGConfig, src *YGConfig) {
	*dest = *src
}

// YGNodeMarkDirtyInternal marks the node as dirty, internally
func YGNodeMarkDirtyInternal(node *YGNode) {
	if !node.isDirty {
		node.isDirty = true
		node.layout.computedFlexBasis = YGUndefined
		if node.parent != nil {
			YGNodeMarkDirtyInternal(node.parent)
		}
	}
}

// YGNodeSetMeasureFunc sets measure function
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

// YGNodeGetMeasureFunc gets measure function
func YGNodeGetMeasureFunc(node *YGNode) YGMeasureFunc {
	return node.measure
}

// YGNodeSetBaselineFunc sets baseline function
func YGNodeSetBaselineFunc(node *YGNode, baselineFunc YGBaselineFunc) {
	node.baseline = baselineFunc
}

// YGNodeGetBaselineFunc gets baseline function
func YGNodeGetBaselineFunc(node *YGNode) YGBaselineFunc {
	return node.baseline
}

// YGNodeInsertChild inserts a child
func YGNodeInsertChild(node *YGNode, child *YGNode, index int) {
	YGAssertWithNode(node, child.parent == nil, "Child already has a parent, it must be removed first.")
	YGAssertWithNode(node, node.measure == nil, "Cannot add child: Nodes with measure functions cannot have children.")

	YGNodeListInsert(&node.children, child, index)
	child.parent = node
	YGNodeMarkDirtyInternal(node)
}

// YGNodeRemoveChild removes the child
func YGNodeRemoveChild(node *YGNode, child *YGNode) {
	if YGNodeListDelete(node.children, child) != nil {
		child.layout = gYGNodeDefaults.layout // layout is no longer valid
		child.parent = nil
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeGetChild returns a child
func YGNodeGetChild(node *YGNode, index int) *YGNode {
	return YGNodeListGet(node.children, index)
}

// YGNodeGetParent gets parent
func YGNodeGetParent(node *YGNode) *YGNode {
	return node.parent
}

// YGNodeGetChildCount returns number of children
func YGNodeGetChildCount(node *YGNode) int {
	return YGNodeListCount(node.children)
}

// YGNodeMarkDirty marks node as dirty
func YGNodeMarkDirty(node *YGNode) {
	YGAssertWithNode(node, node.measure != nil,
		"Only leaf nodes with custom measure functions should manually mark themselves as dirty")
	YGNodeMarkDirtyInternal(node)
}

// YGNodeIsDirty returns true if node is dirty
func YGNodeIsDirty(node *YGNode) bool {
	return node.isDirty
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
	// TODO: for now, always return false
	return false
	//return true
}

// YGNodeCopyStyle copies style
func YGNodeCopyStyle(dstNode *YGNode, srcNode *YGNode) {
	if !styleEq(&dstNode.style, &srcNode.style) {
		dstNode.style = srcNode.style
		YGNodeMarkDirtyInternal(dstNode)
	}
}

// YGResolveFlexGrow resolves flex grow
func YGResolveFlexGrow(node *YGNode) float32 {
	// Root nodes flexGrow should always be 0
	if node.parent == nil {
		return 0
	}
	if !YGFloatIsUndefined(node.style.flexGrow) {
		return node.style.flexGrow
	}
	if !YGFloatIsUndefined(node.style.flex) && node.style.flex > 0 {
		return node.style.flex
	}
	return kDefaultFlexGrow
}

// YGNodeStyleGetFlexGrow gets flex grow
func YGNodeStyleGetFlexGrow(node *YGNode) float32 {
	if YGFloatIsUndefined(node.style.flexGrow) {
		return kDefaultFlexGrow
	}
	return node.style.flexGrow
}

// YGNodeStyleGetFlexShrink gets flex shrink
func YGNodeStyleGetFlexShrink(node *YGNode) float32 {
	if YGFloatIsUndefined(node.style.flexShrink) {
		if node.config.useWebDefaults {
			return kWebDefaultFlexShrink
		}
		return kDefaultFlexShrink
	}
	return node.style.flexShrink
}

// YGNodeResolveFlexShrink resolves flex shrink
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

// YGNodeResolveFlexBasisPtr resolves flex basis ptr
func YGNodeResolveFlexBasisPtr(node *YGNode) *YGValue {
	style := &node.style
	if style.flexBasis.unit != YGUnitAuto && style.flexBasis.unit != YGUnitUndefined {
		return &style.flexBasis
	}
	if !YGFloatIsUndefined(style.flex) && style.flex > 0 {
		if node.config.useWebDefaults {
			return &YGValueAuto
		}
		return &YGValueZero
	}
	return &YGValueAuto
}

// see yoga_props.go

var (
	gCurrentGenerationCount = 0
)

// YGFloatIsUndefined returns true if value is undefined
func YGFloatIsUndefined(value float32) bool {
	return IsNaN(value)
}

// YGValueEqual returns true if values are equal
func YGValueEqual(a YGValue, b YGValue) bool {
	if a.unit != b.unit {
		return false
	}

	if a.unit == YGUnitUndefined {
		return true
	}

	return fabs(a.value-b.value) < 0.0001
}

/// -------------------------- not-yet-arranged

var (
	gPrintTree    = false
	gPrintChanges = false
	gPrintSkips   = false
	gDepth        = 0
)

var (
	leading  = [4]YGEdge{YGEdgeTop, YGEdgeBottom, YGEdgeLeft, YGEdgeRight}
	trailing = [4]YGEdge{YGEdgeBottom, YGEdgeTop, YGEdgeRight, YGEdgeLeft}
	pos      = [4]YGEdge{YGEdgeTop, YGEdgeBottom, YGEdgeLeft, YGEdgeRight}
	dim      = [4]YGDimension{YGDimensionHeight, YGDimensionHeight, YGDimensionWidth, YGDimensionWidth}
)

func YGFloatsEqual(a float32, b float32) bool {
	if YGFloatIsUndefined(a) {
		return YGFloatIsUndefined(b)
	}
	return fabs(a-b) < 0.0001
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
	v := node.resolvedDimensions[dim[axis]]
	isNotDefined := (v.unit == YGUnitAuto ||
		v.unit == YGUnitUndefined ||
		(v.unit == YGUnitPoint && v.value < 0) ||
		(v.unit == YGUnitPercent && (v.value < 0 || YGFloatIsUndefined(parentSize))))
	return !isNotDefined
}

func YGNodeMarginForAxis(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	leading := YGNodeLeadingMargin(node, axis, widthSize)
	trailing := YGNodeTrailingMargin(node, axis, widthSize)
	return leading + trailing
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

	v := YGComputedEdgeValue(node.style.margin[:], leading[axis], &YGValueZero)
	return YGResolveValueMargin(v, widthSize)
}

func YGNodeTrailingMargin(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.margin[YGEdgeEnd].unit != YGUnitUndefined {
		return YGResolveValueMargin(&node.style.margin[YGEdgeEnd], widthSize)
	}

	return YGResolveValueMargin(YGComputedEdgeValue(node.style.margin[:], trailing[axis], &YGValueZero),
		widthSize)
}

func YGNodeLeadingPadding(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.padding[YGEdgeStart].unit != YGUnitUndefined &&
		YGResolveValue(&node.style.padding[YGEdgeStart], widthSize) >= 0 {
		return YGResolveValue(&node.style.padding[YGEdgeStart], widthSize)
	}

	return fmaxf(YGResolveValue(YGComputedEdgeValue(node.style.padding[:], leading[axis], &YGValueZero), widthSize), 0)
}

func YGNodeTrailingPadding(node *YGNode, axis YGFlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.padding[YGEdgeEnd].unit != YGUnitUndefined &&
		YGResolveValue(&node.style.padding[YGEdgeEnd], widthSize) >= 0 {
		return YGResolveValue(&node.style.padding[YGEdgeEnd], widthSize)
	}

	return fmaxf(YGResolveValue(YGComputedEdgeValue(node.style.padding[:], trailing[axis], &YGValueZero), widthSize), 0)
}

func YGNodeLeadingBorder(node *YGNode, axis YGFlexDirection) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.border[YGEdgeStart].unit != YGUnitUndefined &&
		node.style.border[YGEdgeStart].value >= 0 {
		return node.style.border[YGEdgeStart].value
	}

	return fmaxf(YGComputedEdgeValue(node.style.border[:], leading[axis], &YGValueZero).value, 0)
}

func YGNodeTrailingBorder(node *YGNode, axis YGFlexDirection) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.border[YGEdgeEnd].unit != YGUnitUndefined &&
		node.style.border[YGEdgeEnd].value >= 0 {
		return node.style.border[YGEdgeEnd].value
	}

	return fmaxf(YGComputedEdgeValue(node.style.border[:], trailing[axis], &YGValueZero).value, 0)
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
		YGComputedEdgeValue(node.style.position[:], YGEdgeStart, &YGValueUndefined).unit !=
			YGUnitUndefined) ||
		YGComputedEdgeValue(node.style.position[:], leading[axis], &YGValueUndefined).unit !=
			YGUnitUndefined
}

func YGNodeIsTrailingPosDefined(node *YGNode, axis YGFlexDirection) bool {
	return (YGFlexDirectionIsRow(axis) &&
		YGComputedEdgeValue(node.style.position[:], YGEdgeEnd, &YGValueUndefined).unit !=
			YGUnitUndefined) ||
		YGComputedEdgeValue(node.style.position[:], trailing[axis], &YGValueUndefined).unit !=
			YGUnitUndefined
}

func YGNodeLeadingPosition(node *YGNode, axis YGFlexDirection, axisSize float32) float32 {
	if YGFlexDirectionIsRow(axis) {
		leadingPosition := YGComputedEdgeValue(node.style.position[:], YGEdgeStart, &YGValueUndefined)
		if leadingPosition.unit != YGUnitUndefined {
			return YGResolveValue(leadingPosition, axisSize)
		}
	}

	leadingPosition := YGComputedEdgeValue(node.style.position[:], leading[axis], &YGValueUndefined)

	if leadingPosition.unit == YGUnitUndefined {
		return 0
	}
	return YGResolveValue(leadingPosition, axisSize)
}

func YGNodeTrailingPosition(node *YGNode, axis YGFlexDirection, axisSize float32) float32 {
	if YGFlexDirectionIsRow(axis) {
		trailingPosition := YGComputedEdgeValue(node.style.position[:], YGEdgeEnd, &YGValueUndefined)
		if trailingPosition.unit != YGUnitUndefined {
			return YGResolveValue(trailingPosition, axisSize)
		}
	}

	trailingPosition := YGComputedEdgeValue(node.style.position[:], trailing[axis], &YGValueUndefined)

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

func YGIndent(node *YGNode, n int) {
	for i := 0; i < n; i++ {
		YGLog(node, YGLogLevelDebug, "  ")
	}
}

func YGPrintNumberIfNotUndefinedf(node *YGNode, str string, number float32) {
	if !YGFloatIsUndefined(number) {
		YGLog(node, YGLogLevelDebug, "%s: %g; ", str, number)
	}
}

func YGPrintNumberIfNotUndefined(node *YGNode, str string, number *YGValue) {
	if number.unit != YGUnitUndefined {
		if number.unit == YGUnitAuto {
			YGLog(node, YGLogLevelDebug, "%s: auto; ", str)
		} else {
			unit := "%"

			if number.unit == YGUnitPoint {
				unit = "px"
			}
			YGLog(node, YGLogLevelDebug, "%s: %g%s; ", str, number.value, unit)
		}
	}
}

func YGPrintNumberIfNotAuto(node *YGNode, str string, number *YGValue) {
	if number.unit != YGUnitAuto {
		YGPrintNumberIfNotUndefined(node, str, number)
	}
}

func YGPrintEdgeIfNotUndefined(node *YGNode, str string, edges []YGValue, edge YGEdge) {
	YGPrintNumberIfNotUndefined(node, str, YGComputedEdgeValue(edges, edge, &YGValueUndefined))
}

func YGPrintNumberIfNotZero(node *YGNode, str string, number *YGValue) {
	if !YGFloatsEqual(number.value, 0) {
		YGPrintNumberIfNotUndefined(node, str, number)
	}
}

func YGFourValuesEqual(four []YGValue) bool {
	return YGValueEqual(four[0], four[1]) && YGValueEqual(four[0], four[2]) &&
		YGValueEqual(four[0], four[3])
}

func YGPrintEdges(node *YGNode, str string, edges []YGValue) {
	if YGFourValuesEqual(edges) {
		YGPrintNumberIfNotZero(node, str, &edges[YGEdgeLeft])
	} else {
		for edge := YGEdgeLeft; edge < YGEdgeCount; edge++ {
			buf := fmt.Sprintf("%s-%s", str, YGEdgeToString(edge))
			YGPrintNumberIfNotZero(node, buf, &edges[edge])
		}
	}
}

func YGNodePrintInternal(node *YGNode, options YGPrintOptions, level int) {
	YGIndent(node, level)
	YGLog(node, YGLogLevelDebug, "<div ")

	if node.print != nil {
		node.print(node)
	}

	if options&YGPrintOptionsLayout != 0 {
		YGLog(node, YGLogLevelDebug, "layout=\"")
		YGLog(node, YGLogLevelDebug, "width: %g; ", node.layout.dimensions[YGDimensionWidth])
		YGLog(node, YGLogLevelDebug, "height: %g; ", node.layout.dimensions[YGDimensionHeight])
		YGLog(node, YGLogLevelDebug, "top: %g; ", node.layout.position[YGEdgeTop])
		YGLog(node, YGLogLevelDebug, "left: %g;", node.layout.position[YGEdgeLeft])
		YGLog(node, YGLogLevelDebug, "\" ")
	}

	if options&YGPrintOptionsStyle != 0 {
		YGLog(node, YGLogLevelDebug, "style=\"")
		if node.style.flexDirection != gYGNodeDefaults.style.flexDirection {
			YGLog(node,
				YGLogLevelDebug,
				"flex-direction: %s; ",
				YGFlexDirectionToString(node.style.flexDirection))
		}
		if node.style.justifyContent != gYGNodeDefaults.style.justifyContent {
			YGLog(node,
				YGLogLevelDebug,
				"justify-content: %s; ",
				YGJustifyToString(node.style.justifyContent))
		}
		if node.style.alignItems != gYGNodeDefaults.style.alignItems {
			YGLog(node, YGLogLevelDebug, "align-items: %s; ", YGAlignToString(node.style.alignItems))
		}
		if node.style.alignContent != gYGNodeDefaults.style.alignContent {
			YGLog(node, YGLogLevelDebug, "align-content: %s; ", YGAlignToString(node.style.alignContent))
		}
		if node.style.alignSelf != gYGNodeDefaults.style.alignSelf {
			YGLog(node, YGLogLevelDebug, "align-self: %s; ", YGAlignToString(node.style.alignSelf))
		}

		YGPrintNumberIfNotUndefinedf(node, "flex-grow", node.style.flexGrow)
		YGPrintNumberIfNotUndefinedf(node, "flex-shrink", node.style.flexShrink)
		YGPrintNumberIfNotAuto(node, "flex-basis", &node.style.flexBasis)
		YGPrintNumberIfNotUndefinedf(node, "flex", node.style.flex)

		if node.style.flexWrap != gYGNodeDefaults.style.flexWrap {
			YGLog(node, YGLogLevelDebug, "flexWrap: %s; ", YGWrapToString(node.style.flexWrap))
		}

		if node.style.overflow != gYGNodeDefaults.style.overflow {
			YGLog(node, YGLogLevelDebug, "overflow: %s; ", YGOverflowToString(node.style.overflow))
		}

		if node.style.display != gYGNodeDefaults.style.display {
			YGLog(node, YGLogLevelDebug, "display: %s; ", YGDisplayToString(node.style.display))
		}

		YGPrintEdges(node, "margin", node.style.margin[:])
		YGPrintEdges(node, "padding", node.style.padding[:])
		YGPrintEdges(node, "border", node.style.border[:])

		YGPrintNumberIfNotAuto(node, "width", &node.style.dimensions[YGDimensionWidth])
		YGPrintNumberIfNotAuto(node, "height", &node.style.dimensions[YGDimensionHeight])
		YGPrintNumberIfNotAuto(node, "max-width", &node.style.maxDimensions[YGDimensionWidth])
		YGPrintNumberIfNotAuto(node, "max-height", &node.style.maxDimensions[YGDimensionHeight])
		YGPrintNumberIfNotAuto(node, "min-width", &node.style.minDimensions[YGDimensionWidth])
		YGPrintNumberIfNotAuto(node, "min-height", &node.style.minDimensions[YGDimensionHeight])

		if node.style.positionType != gYGNodeDefaults.style.positionType {
			YGLog(node,
				YGLogLevelDebug,
				"position: %s; ",
				YGPositionTypeToString(node.style.positionType))
		}

		YGPrintEdgeIfNotUndefined(node, "left", node.style.position[:], YGEdgeLeft)
		YGPrintEdgeIfNotUndefined(node, "right", node.style.position[:], YGEdgeRight)
		YGPrintEdgeIfNotUndefined(node, "top", node.style.position[:], YGEdgeTop)
		YGPrintEdgeIfNotUndefined(node, "bottom", node.style.position[:], YGEdgeBottom)
		YGLog(node, YGLogLevelDebug, "\" ")

		if node.measure != nil {
			YGLog(node, YGLogLevelDebug, "has-custom-measure=\"true\"")
		}
	}
	YGLog(node, YGLogLevelDebug, ">")

	childCount := YGNodeListCount(node.children)
	if options&YGPrintOptionsChildren != 0 && childCount > 0 {
		for i := 0; i < childCount; i++ {
			YGLog(node, YGLogLevelDebug, "\n")
			YGNodePrintInternal(YGNodeGetChild(node, i), options, level+1)
		}
		YGIndent(node, level)
		YGLog(node, YGLogLevelDebug, "\n")
	}
	YGLog(node, YGLogLevelDebug, "</div>")
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

func YGConfigSetPointScaleFactor(config *YGConfig, pixelsInPoint float32) {
	YGAssertWithConfig(config, pixelsInPoint >= 0, "Scale factor should not be less than zero")

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
	config *YGConfig) {
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

func YGNodeCanUseCachedMeasurement(widthMode YGMeasureMode, width float32, heightMode YGMeasureMode, height float32, lastWidthMode YGMeasureMode, lastWidth float32, lastHeightMode YGMeasureMode, lastHeight float32, lastComputedWidth float32, lastComputedHeight float32, marginRow float32, marginColumn float32, config *YGConfig) bool {
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
	case YGMeasureModeExactly, YGMeasureModeAtMost:
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

func triFloat(useFirst bool, f1, f2 float32) float32 {
	if useFirst {
		return f1
	}
	return f2
}

// This is the main routine that implements a subset of the flexbox layout
// algorithm
// described in the W3C YG documentation: https://www.w3.org/TR/YG3-flexbox/.
//
// Limitations of this algorithm, compared to the full standard:
//  * Display property is always assumed to be 'flex' except for Text nodes,
//  which
//    are assumed to be 'inline-flex'.
//  * The 'zIndex' property (or any form of z ordering) is not supported. Nodes
//  are
//    stacked in document order.
//  * The 'order' property is not supported. The order of flex items is always
//  defined
//    by document order.
//  * The 'visibility' property is always assumed to be 'visible'. Values of
//  'collapse'
//    and 'hidden' are not supported.
//  * There is no support for forced breaks.
//  * It does not support vertical inline directions (top-to-bottom or
//  bottom-to-top text).
//
// Deviations from standard:
//  * Section 4.5 of the spec indicates that all flex items have a default
//  minimum
//    main size. For text blocks, for example, this is the width of the widest
//    word.
//    Calculating the minimum width is expensive, so we forego it and assume a
//    default
//    minimum main size of 0.
//  * Min/Max sizes in the main axis are not honored when resolving flexible
//  lengths.
//  * The spec indicates that the default value for 'flexDirection' is 'row',
//  but
//    the algorithm below assumes a default of 'column'.
//
// Input parameters:
//    - node: current node to be sized and layed out
//    - availableWidth & availableHeight: available size to be used for sizing
//    the node
//      or YGUndefined if the size is not available; interpretation depends on
//      layout
//      flags
//    - parentDirection: the inline (text) direction within the parent
//    (left-to-right or
//      right-to-left)
//    - widthMeasureMode: indicates the sizing rules for the width (see below
//    for explanation)
//    - heightMeasureMode: indicates the sizing rules for the height (see below
//    for explanation)
//    - performLayout: specifies whether the caller is interested in just the
//    dimensions
//      of the node or it requires the entire node and its subtree to be layed
//      out
//      (with final positions)
//
// Details:
//    This routine is called recursively to lay out subtrees of flexbox
//    elements. It uses the
//    information in node.style, which is treated as a read-only input. It is
//    responsible for
//    setting the layout.direction and layout.measuredDimensions fields for the
//    input node as well
//    as the layout.position and layout.lineIndex fields for its child nodes.
//    The
//    layout.measuredDimensions field includes any border or padding for the
//    node but does
//    not include margins.
//
//    The spec describes four different layout modes: "fill available", "max
//    content", "min
//    content",
//    and "fit content". Of these, we don't use "min content" because we don't
//    support default
//    minimum main sizes (see above for details). Each of our measure modes maps
//    to a layout mode
//    from the spec (https://www.w3.org/TR/YG3-sizing/#terms):
//      - YGMeasureModeUndefined: max content
//      - YGMeasureModeExactly: fill available
//      - YGMeasureModeAtMost: fit content
//
//    When calling YGNodelayoutImpl and YGLayoutNodeInternal, if the caller passes
//    an available size of
//    undefined then it must also pass a measure mode of YGMeasureModeUndefined
//    in that dimension.
func YGNodelayoutImpl(node *YGNode, availableWidth float32, availableHeight float32,
	parentDirection YGDirection, widthMeasureMode YGMeasureMode,
	heightMeasureMode YGMeasureMode, parentWidth float32, parentHeight float32,
	performLayout bool, config *YGConfig) {
	// YGAssertWithNode(node, YGFloatIsUndefined(availableWidth) ? widthMeasureMode == YGMeasureModeUndefined : true, "availableWidth is indefinite so widthMeasureMode must be YGMeasureModeUndefined");
	//YGAssertWithNode(node, YGFloatIsUndefined(availableHeight) ? heightMeasureMode == YGMeasureModeUndefined : true, "availableHeight is indefinite so heightMeasureMode must be YGMeasureModeUndefined");

	// Set the resolved resolution in the node's layout.
	direction := YGNodeResolveDirection(node, parentDirection)
	node.layout.direction = direction

	flexRowDirection := YGResolveFlexDirection(YGFlexDirectionRow, direction)
	flexColumnDirection := YGResolveFlexDirection(YGFlexDirectionColumn, direction)

	node.layout.margin[YGEdgeStart] = YGNodeLeadingMargin(node, flexRowDirection, parentWidth)
	node.layout.margin[YGEdgeEnd] = YGNodeTrailingMargin(node, flexRowDirection, parentWidth)
	node.layout.margin[YGEdgeTop] = YGNodeLeadingMargin(node, flexColumnDirection, parentWidth)
	node.layout.margin[YGEdgeBottom] = YGNodeTrailingMargin(node, flexColumnDirection, parentWidth)

	node.layout.border[YGEdgeStart] = YGNodeLeadingBorder(node, flexRowDirection)
	node.layout.border[YGEdgeEnd] = YGNodeTrailingBorder(node, flexRowDirection)
	node.layout.border[YGEdgeTop] = YGNodeLeadingBorder(node, flexColumnDirection)
	node.layout.border[YGEdgeBottom] = YGNodeTrailingBorder(node, flexColumnDirection)

	node.layout.padding[YGEdgeStart] = YGNodeLeadingPadding(node, flexRowDirection, parentWidth)
	node.layout.padding[YGEdgeEnd] = YGNodeTrailingPadding(node, flexRowDirection, parentWidth)
	node.layout.padding[YGEdgeTop] = YGNodeLeadingPadding(node, flexColumnDirection, parentWidth)
	node.layout.padding[YGEdgeBottom] = YGNodeTrailingPadding(node, flexColumnDirection, parentWidth)

	if node.measure != nil {
		YGNodeWithMeasureFuncSetMeasuredDimensions(node, availableWidth, availableHeight, widthMeasureMode, heightMeasureMode, parentWidth, parentHeight)
		return
	}

	childCount := YGNodeListCount(node.children)
	if childCount == 0 {
		YGNodeEmptyContainerSetMeasuredDimensions(node, availableWidth, availableHeight, widthMeasureMode, heightMeasureMode, parentWidth, parentHeight)
		return
	}

	// If we're not being asked to perform a full layout we can skip the algorithm if we already know
	// the size
	if !performLayout && YGNodeFixedSizeSetMeasuredDimensions(node, availableWidth, availableHeight, widthMeasureMode, heightMeasureMode, parentWidth, parentHeight) {
		return
	}

	// Reset layout flags, as they could have changed.
	node.layout.hadOverflow = false

	// STEP 1: CALCULATE VALUES FOR REMAINDER OF ALGORITHM
	mainAxis := YGResolveFlexDirection(node.style.flexDirection, direction)
	crossAxis := YGFlexDirectionCross(mainAxis, direction)
	isMainAxisRow := YGFlexDirectionIsRow(mainAxis)
	justifyContent := node.style.justifyContent
	isNodeFlexWrap := node.style.flexWrap != YGWrapNoWrap

	mainAxisParentSize := parentHeight
	crossAxisParentSize := parentWidth
	if isMainAxisRow {
		mainAxisParentSize = parentWidth
		crossAxisParentSize = parentHeight
	}

	var firstAbsoluteChild *YGNode
	var currentAbsoluteChild *YGNode

	leadingPaddingAndBorderMain := YGNodeLeadingPaddingAndBorder(node, mainAxis, parentWidth)
	trailingPaddingAndBorderMain := YGNodeTrailingPaddingAndBorder(node, mainAxis, parentWidth)
	leadingPaddingAndBorderCross := YGNodeLeadingPaddingAndBorder(node, crossAxis, parentWidth)
	paddingAndBorderAxisMain := YGNodePaddingAndBorderForAxis(node, mainAxis, parentWidth)
	paddingAndBorderAxisCross := YGNodePaddingAndBorderForAxis(node, crossAxis, parentWidth)

	measureModeMainDim := heightMeasureMode
	measureModeCrossDim := widthMeasureMode

	if isMainAxisRow {
		measureModeMainDim = widthMeasureMode
		measureModeCrossDim = heightMeasureMode
	}

	paddingAndBorderAxisRow := paddingAndBorderAxisCross
	if isMainAxisRow {
		paddingAndBorderAxisRow = paddingAndBorderAxisMain
	}

	paddingAndBorderAxisColumn := paddingAndBorderAxisMain
	if isMainAxisRow {
		paddingAndBorderAxisColumn = paddingAndBorderAxisCross
	}

	marginAxisRow := YGNodeMarginForAxis(node, YGFlexDirectionRow, parentWidth)
	marginAxisColumn := YGNodeMarginForAxis(node, YGFlexDirectionColumn, parentWidth)

	// STEP 2: DETERMINE AVAILABLE SIZE IN MAIN AND CROSS DIRECTIONS
	minInnerWidth := YGResolveValue(&node.style.minDimensions[YGDimensionWidth], parentWidth) - marginAxisRow -
		paddingAndBorderAxisRow
	maxInnerWidth := YGResolveValue(&node.style.maxDimensions[YGDimensionWidth], parentWidth) - marginAxisRow -
		paddingAndBorderAxisRow
	minInnerHeight := YGResolveValue(&node.style.minDimensions[YGDimensionHeight], parentHeight) -
		marginAxisColumn - paddingAndBorderAxisColumn
	maxInnerHeight := YGResolveValue(&node.style.maxDimensions[YGDimensionHeight], parentHeight) -
		marginAxisColumn - paddingAndBorderAxisColumn

	minInnerMainDim := minInnerHeight
	if isMainAxisRow {
		minInnerMainDim = minInnerWidth
	}
	maxInnerMainDim := maxInnerHeight
	if isMainAxisRow {
		maxInnerMainDim = maxInnerWidth
	}

	// Max dimension overrides predefined dimension value; Min dimension in turn overrides both of the
	// above
	availableInnerWidth := availableWidth - marginAxisRow - paddingAndBorderAxisRow
	if !YGFloatIsUndefined(availableInnerWidth) {
		// We want to make sure our available width does not violate min and max raints
		availableInnerWidth = fmaxf(fminf(availableInnerWidth, maxInnerWidth), minInnerWidth)
	}

	availableInnerHeight := availableHeight - marginAxisColumn - paddingAndBorderAxisColumn
	if !YGFloatIsUndefined(availableInnerHeight) {
		// We want to make sure our available height does not violate min and max raints
		availableInnerHeight = fmaxf(fminf(availableInnerHeight, maxInnerHeight), minInnerHeight)
	}

	availableInnerMainDim := availableInnerHeight
	if isMainAxisRow {
		availableInnerMainDim = availableInnerWidth
	}
	availableInnerCrossDim := availableInnerWidth
	if isMainAxisRow {
		availableInnerCrossDim = availableInnerHeight
	}

	// If there is only one child with flexGrow + flexShrink it means we can set the
	// computedFlexBasis to 0 instead of measuring and shrinking / flexing the child to exactly
	// match the remaining space
	var singleFlexChild *YGNode
	if measureModeMainDim == YGMeasureModeExactly {
		for i := 0; i < childCount; i++ {
			child := YGNodeGetChild(node, i)
			if singleFlexChild != nil {
				if YGNodeIsFlex(child) {
					// There is already a flexible child, abort.
					singleFlexChild = nil
					break
				}
			} else if YGResolveFlexGrow(child) > 0 && YGNodeResolveFlexShrink(child) > 0 {
				singleFlexChild = child
			}
		}
	}

	var totalOuterFlexBasis float32

	// STEP 3: DETERMINE FLEX BASIS FOR EACH ITEM
	for i := 0; i < childCount; i++ {
		child := YGNodeListGet(node.children, i)
		if child.style.display == YGDisplayNone {
			YGZeroOutLayoutRecursivly(child)
			child.hasNewLayout = true
			child.isDirty = false
			continue
		}
		YGResolveDimensions(child)
		if performLayout {
			// Set the initial position (relative to the parent).
			childDirection := YGNodeResolveDirection(child, direction)
			YGNodeSetPosition(child,
				childDirection,
				availableInnerMainDim,
				availableInnerCrossDim,
				availableInnerWidth)
		}

		// Absolute-positioned children don't participate in flex layout. Add them
		// to a list that we can process later.
		if child.style.positionType == YGPositionTypeAbsolute {
			// Store a private linked list of absolutely positioned children
			// so that we can efficiently traverse them later.
			if firstAbsoluteChild == nil {
				firstAbsoluteChild = child
			}
			if currentAbsoluteChild != nil {
				currentAbsoluteChild.nextChild = child
			}
			currentAbsoluteChild = child
			child.nextChild = nil
		} else {
			if child == singleFlexChild {
				child.layout.computedFlexBasisGeneration = gCurrentGenerationCount
				child.layout.computedFlexBasis = 0
			} else {
				YGNodeComputeFlexBasisForChild(node,
					child,
					availableInnerWidth,
					widthMeasureMode,
					availableInnerHeight,
					availableInnerWidth,
					availableInnerHeight,
					heightMeasureMode,
					direction,
					config)
			}
		}

		totalOuterFlexBasis +=
			child.layout.computedFlexBasis + YGNodeMarginForAxis(child, mainAxis, availableInnerWidth)

	}

	flexBasisOverflows := totalOuterFlexBasis > availableInnerMainDim
	if measureModeMainDim == YGMeasureModeUndefined {
		flexBasisOverflows = false
	}
	if isNodeFlexWrap && flexBasisOverflows && measureModeMainDim == YGMeasureModeAtMost {
		measureModeMainDim = YGMeasureModeExactly
	}

	// STEP 4: COLLECT FLEX ITEMS INTO FLEX LINES

	// Indexes of children that represent the first and last items in the line.
	startOfLineIndex := 0
	endOfLineIndex := 0

	// Number of lines.
	lineCount := 0

	// Accumulated cross dimensions of all lines so far.
	var totalLineCrossDim float32

	// Max main dimension of all the lines.
	var maxLineMainDim float32

	for endOfLineIndex < childCount {
		// Number of items on the currently line. May be different than the
		// difference
		// between start and end indicates because we skip over absolute-positioned
		// items.
		itemsOnLine := 0

		// sizeConsumedOnCurrentLine is accumulation of the dimensions and margin
		// of all the children on the current line. This will be used in order to
		// either set the dimensions of the node if none already exist or to compute
		// the remaining space left for the flexible children.
		var sizeConsumedOnCurrentLine float32
		var sizeConsumedOnCurrentLineIncludingMinConstraint float32

		var totalFlexGrowFactors float32
		var totalFlexShrinkScaledFactors float32

		// Maintain a linked list of the child nodes that can shrink and/or grow.
		var firstRelativeChild *YGNode
		var currentRelativeChild *YGNode

		// Add items to the current line until it's full or we run out of items.
		for i := startOfLineIndex; i < childCount; i++ {
			child := YGNodeListGet(node.children, i)
			if child.style.display == YGDisplayNone {
				continue
			}
			child.lineIndex = lineCount

			if child.style.positionType != YGPositionTypeAbsolute {
				childMarginMainAxis := YGNodeMarginForAxis(child, mainAxis, availableInnerWidth)
				flexBasisWithMaxConstraints := fminf(YGResolveValue(&child.style.maxDimensions[dim[mainAxis]], mainAxisParentSize), child.layout.computedFlexBasis)
				flexBasisWithMinAndMaxConstraints := fmaxf(YGResolveValue(&child.style.minDimensions[dim[mainAxis]], mainAxisParentSize), flexBasisWithMaxConstraints)

				// If this is a multi-line flow and this item pushes us over the
				// available size, we've
				// hit the end of the current line. Break out of the loop and lay out
				// the current line.
				if sizeConsumedOnCurrentLineIncludingMinConstraint+flexBasisWithMinAndMaxConstraints+
					childMarginMainAxis >
					availableInnerMainDim &&
					isNodeFlexWrap && itemsOnLine > 0 {
					break
				}

				sizeConsumedOnCurrentLineIncludingMinConstraint +=
					flexBasisWithMinAndMaxConstraints + childMarginMainAxis
				sizeConsumedOnCurrentLine += flexBasisWithMinAndMaxConstraints + childMarginMainAxis
				itemsOnLine++

				if YGNodeIsFlex(child) {
					totalFlexGrowFactors += YGResolveFlexGrow(child)

					// Unlike the grow factor, the shrink factor is scaled relative to the child dimension.
					totalFlexShrinkScaledFactors +=
						-YGNodeResolveFlexShrink(child) * child.layout.computedFlexBasis
				}

				// Store a private linked list of children that need to be layed out.
				if firstRelativeChild == nil {
					firstRelativeChild = child
				}
				if currentRelativeChild != nil {
					currentRelativeChild.nextChild = child
				}
				currentRelativeChild = child
				child.nextChild = nil
			}
			endOfLineIndex++
		}

		// The total flex factor needs to be floored to 1.
		if totalFlexGrowFactors > 0 && totalFlexGrowFactors < 1 {
			totalFlexGrowFactors = 1
		}

		// The total flex shrink factor needs to be floored to 1.
		if totalFlexShrinkScaledFactors > 0 && totalFlexShrinkScaledFactors < 1 {
			totalFlexShrinkScaledFactors = 1
		}

		// If we don't need to measure the cross axis, we can skip the entire flex
		// step.
		canSkipFlex := !performLayout && measureModeCrossDim == YGMeasureModeExactly

		// In order to position the elements in the main axis, we have two
		// controls. The space between the beginning and the first element
		// and the space between each two elements.
		var leadingMainDim float32
		var betweenMainDim float32

		// STEP 5: RESOLVING FLEXIBLE LENGTHS ON MAIN AXIS
		// Calculate the remaining available space that needs to be allocated.
		// If the main dimension size isn't known, it is computed based on
		// the line length, so there's no more space left to distribute.

		// If we don't measure with exact main dimension we want to ensure we don't violate min and max
		if measureModeMainDim != YGMeasureModeExactly {
			if !YGFloatIsUndefined(minInnerMainDim) && sizeConsumedOnCurrentLine < minInnerMainDim {
				availableInnerMainDim = minInnerMainDim
			} else if !YGFloatIsUndefined(maxInnerMainDim) &&
				sizeConsumedOnCurrentLine > maxInnerMainDim {
				availableInnerMainDim = maxInnerMainDim
			} else {
				if !node.config.useLegacyStretchBehaviour &&
					(totalFlexGrowFactors == 0 || YGResolveFlexGrow(node) == 0) {
					// If we don't have any children to flex or we can't flex the node itself,
					// space we've used is all space we need. Root node also should be shrunk to minimum
					availableInnerMainDim = sizeConsumedOnCurrentLine
				}
			}
		}

		var remainingFreeSpace float32
		if !YGFloatIsUndefined(availableInnerMainDim) {
			remainingFreeSpace = availableInnerMainDim - sizeConsumedOnCurrentLine
		} else if sizeConsumedOnCurrentLine < 0 {
			// availableInnerMainDim is indefinite which means the node is being sized based on its
			// content.
			// sizeConsumedOnCurrentLine is negative which means the node will allocate 0 points for
			// its content. Consequently, remainingFreeSpace is 0 - sizeConsumedOnCurrentLine.
			remainingFreeSpace = -sizeConsumedOnCurrentLine
		}

		originalRemainingFreeSpace := remainingFreeSpace
		var deltaFreeSpace float32

		if !canSkipFlex {
			var childFlexBasis float32
			var flexShrinkScaledFactor float32
			var flexGrowFactor float32
			var baseMainSize float32
			var boundMainSize float32

			// Do two passes over the flex items to figure out how to distribute the
			// remaining space.
			// The first pass finds the items whose min/max raints trigger,
			// freezes them at those
			// sizes, and excludes those sizes from the remaining space. The second
			// pass sets the size
			// of each flexible item. It distributes the remaining space amongst the
			// items whose min/max
			// raints didn't trigger in pass 1. For the other items, it sets
			// their sizes by forcing
			// their min/max raints to trigger again.
			//
			// This two pass approach for resolving min/max raints deviates from
			// the spec. The
			// spec (https://www.w3.org/TR/YG-flexbox-1/#resolve-flexible-lengths)
			// describes a process
			// that needs to be repeated a variable number of times. The algorithm
			// implemented here
			// won't handle all cases but it was simpler to implement and it mitigates
			// performance
			// concerns because we know exactly how many passes it'll do.

			// First pass: detect the flex items whose min/max raints trigger
			var deltaFlexShrinkScaledFactors float32
			var deltaFlexGrowFactors float32
			currentRelativeChild = firstRelativeChild
			for currentRelativeChild != nil {
				childFlexBasis =
					fminf(YGResolveValue(&currentRelativeChild.style.maxDimensions[dim[mainAxis]],
						mainAxisParentSize),
						fmaxf(YGResolveValue(&currentRelativeChild.style.minDimensions[dim[mainAxis]],
							mainAxisParentSize),
							currentRelativeChild.layout.computedFlexBasis))

				if remainingFreeSpace < 0 {
					flexShrinkScaledFactor = -YGNodeResolveFlexShrink(currentRelativeChild) * childFlexBasis

					// Is this child able to shrink?
					if flexShrinkScaledFactor != 0 {
						baseMainSize =
							childFlexBasis +
								remainingFreeSpace/totalFlexShrinkScaledFactors*flexShrinkScaledFactor
						boundMainSize = YGNodeBoundAxis(currentRelativeChild,
							mainAxis,
							baseMainSize,
							availableInnerMainDim,
							availableInnerWidth)
						if baseMainSize != boundMainSize {
							// By excluding this item's size and flex factor from remaining,
							// this item's
							// min/max raints should also trigger in the second pass
							// resulting in the
							// item's size calculation being identical in the first and second
							// passes.
							deltaFreeSpace -= boundMainSize - childFlexBasis
							deltaFlexShrinkScaledFactors -= flexShrinkScaledFactor
						}
					}
				} else if remainingFreeSpace > 0 {
					flexGrowFactor = YGResolveFlexGrow(currentRelativeChild)

					// Is this child able to grow?
					if flexGrowFactor != 0 {
						baseMainSize =
							childFlexBasis + remainingFreeSpace/totalFlexGrowFactors*flexGrowFactor
						boundMainSize = YGNodeBoundAxis(currentRelativeChild,
							mainAxis,
							baseMainSize,
							availableInnerMainDim,
							availableInnerWidth)

						if baseMainSize != boundMainSize {
							// By excluding this item's size and flex factor from remaining,
							// this item's
							// min/max raints should also trigger in the second pass
							// resulting in the
							// item's size calculation being identical in the first and second
							// passes.
							deltaFreeSpace -= boundMainSize - childFlexBasis
							deltaFlexGrowFactors -= flexGrowFactor
						}
					}
				}

				currentRelativeChild = currentRelativeChild.nextChild
			}

			totalFlexShrinkScaledFactors += deltaFlexShrinkScaledFactors
			totalFlexGrowFactors += deltaFlexGrowFactors
			remainingFreeSpace += deltaFreeSpace

			// Second pass: resolve the sizes of the flexible items
			deltaFreeSpace = 0
			currentRelativeChild = firstRelativeChild
			for currentRelativeChild != nil {
				childFlexBasis =
					fminf(YGResolveValue(&currentRelativeChild.style.maxDimensions[dim[mainAxis]],
						mainAxisParentSize),
						fmaxf(YGResolveValue(&currentRelativeChild.style.minDimensions[dim[mainAxis]],
							mainAxisParentSize),
							currentRelativeChild.layout.computedFlexBasis))
				updatedMainSize := childFlexBasis

				if remainingFreeSpace < 0 {
					flexShrinkScaledFactor = -YGNodeResolveFlexShrink(currentRelativeChild) * childFlexBasis
					// Is this child able to shrink?
					if flexShrinkScaledFactor != 0 {
						var childSize float32

						if totalFlexShrinkScaledFactors == 0 {
							childSize = childFlexBasis + flexShrinkScaledFactor
						} else {
							childSize =
								childFlexBasis +
									(remainingFreeSpace/totalFlexShrinkScaledFactors)*flexShrinkScaledFactor
						}

						updatedMainSize = YGNodeBoundAxis(currentRelativeChild,
							mainAxis,
							childSize,
							availableInnerMainDim,
							availableInnerWidth)
					}
				} else if remainingFreeSpace > 0 {
					flexGrowFactor = YGResolveFlexGrow(currentRelativeChild)

					// Is this child able to grow?
					if flexGrowFactor != 0 {
						updatedMainSize =
							YGNodeBoundAxis(currentRelativeChild,
								mainAxis,
								childFlexBasis+
									remainingFreeSpace/totalFlexGrowFactors*flexGrowFactor,
								availableInnerMainDim,
								availableInnerWidth)
					}
				}

				deltaFreeSpace -= updatedMainSize - childFlexBasis

				marginMain := YGNodeMarginForAxis(currentRelativeChild, mainAxis, availableInnerWidth)
				marginCross := YGNodeMarginForAxis(currentRelativeChild, crossAxis, availableInnerWidth)

				var childCrossSize float32
				childMainSize := updatedMainSize + marginMain
				var childCrossMeasureMode YGMeasureMode
				childMainMeasureMode := YGMeasureModeExactly

				if !YGFloatIsUndefined(availableInnerCrossDim) &&
					!YGNodeIsStyleDimDefined(currentRelativeChild, crossAxis, availableInnerCrossDim) &&
					measureModeCrossDim == YGMeasureModeExactly &&
					!(isNodeFlexWrap && flexBasisOverflows) &&
					YGNodeAlignItem(node, currentRelativeChild) == YGAlignStretch {
					childCrossSize = availableInnerCrossDim
					childCrossMeasureMode = YGMeasureModeExactly
				} else if !YGNodeIsStyleDimDefined(currentRelativeChild,
					crossAxis,
					availableInnerCrossDim) {
					childCrossSize = availableInnerCrossDim
					childCrossMeasureMode = YGMeasureModeAtMost
					if YGFloatIsUndefined(childCrossSize) {
						childCrossMeasureMode = YGMeasureModeUndefined
					}
				} else {
					childCrossSize = YGResolveValue(currentRelativeChild.resolvedDimensions[dim[crossAxis]],
						availableInnerCrossDim) +
						marginCross
					isLoosePercentageMeasurement := currentRelativeChild.resolvedDimensions[dim[crossAxis]].unit == YGUnitPercent &&
						measureModeCrossDim != YGMeasureModeExactly
					childCrossMeasureMode = YGMeasureModeExactly
					if YGFloatIsUndefined(childCrossSize) || isLoosePercentageMeasurement {
						childCrossMeasureMode = YGMeasureModeUndefined
					}
				}

				if !YGFloatIsUndefined(currentRelativeChild.style.aspectRatio) {
					v := (childMainSize - marginMain) * currentRelativeChild.style.aspectRatio
					if isMainAxisRow {
						v = (childMainSize - marginMain) / currentRelativeChild.style.aspectRatio
					}
					childCrossSize = fmaxf(v, YGNodePaddingAndBorderForAxis(currentRelativeChild, crossAxis, availableInnerWidth))
					childCrossMeasureMode = YGMeasureModeExactly

					// Parent size raint should have higher priority than flex
					if YGNodeIsFlex(currentRelativeChild) {
						childCrossSize = fminf(childCrossSize-marginCross, availableInnerCrossDim)
						childMainSize = marginMain
						if isMainAxisRow {
							childMainSize += childCrossSize * currentRelativeChild.style.aspectRatio
						} else {
							childMainSize += childCrossSize / currentRelativeChild.style.aspectRatio
						}
					}

					childCrossSize += marginCross
				}

				YGConstrainMaxSizeForMode(currentRelativeChild,
					mainAxis,
					availableInnerMainDim,
					availableInnerWidth,
					&childMainMeasureMode,
					&childMainSize)
				YGConstrainMaxSizeForMode(currentRelativeChild,
					crossAxis,
					availableInnerCrossDim,
					availableInnerWidth,
					&childCrossMeasureMode,
					&childCrossSize)

				requiresStretchLayout := !YGNodeIsStyleDimDefined(currentRelativeChild, crossAxis, availableInnerCrossDim) &&
					YGNodeAlignItem(node, currentRelativeChild) == YGAlignStretch

				childWidth := childCrossSize
				if isMainAxisRow {
					childWidth = childMainSize
				}
				childHeight := childCrossSize
				if !isMainAxisRow {
					childHeight = childMainSize
				}

				childWidthMeasureMode := childCrossMeasureMode
				if isMainAxisRow {
					childWidthMeasureMode = childMainMeasureMode
				}
				childHeightMeasureMode := childCrossMeasureMode
				if !isMainAxisRow {
					childHeightMeasureMode = childMainMeasureMode
				}

				// Recursively call the layout algorithm for this child with the updated
				// main size.
				YGLayoutNodeInternal(currentRelativeChild,
					childWidth,
					childHeight,
					direction,
					childWidthMeasureMode,
					childHeightMeasureMode,
					availableInnerWidth,
					availableInnerHeight,
					performLayout && !requiresStretchLayout,
					"flex",
					config)
				if currentRelativeChild.layout.hadOverflow {
					node.layout.hadOverflow = true
				}

				currentRelativeChild = currentRelativeChild.nextChild
			}
		}

		remainingFreeSpace = originalRemainingFreeSpace + deltaFreeSpace
		if remainingFreeSpace < 0 {
			node.layout.hadOverflow = true
		}

		// STEP 6: MAIN-AXIS JUSTIFICATION & CROSS-AXIS SIZE DETERMINATION

		// At this point, all the children have their dimensions set in the main
		// axis.
		// Their dimensions are also set in the cross axis with the exception of
		// items
		// that are aligned "stretch". We need to compute these stretch values and
		// set the final positions.

		// If we are using "at most" rules in the main axis. Calculate the remaining space when
		// raint by the min size defined for the main axis.

		if measureModeMainDim == YGMeasureModeAtMost && remainingFreeSpace > 0 {
			if node.style.minDimensions[dim[mainAxis]].unit != YGUnitUndefined &&
				YGResolveValue(&node.style.minDimensions[dim[mainAxis]], mainAxisParentSize) >= 0 {
				remainingFreeSpace =
					fmaxf(0,
						YGResolveValue(&node.style.minDimensions[dim[mainAxis]], mainAxisParentSize)-
							(availableInnerMainDim-remainingFreeSpace))
			} else {
				remainingFreeSpace = 0
			}
		}

		numberOfAutoMarginsOnCurrentLine := 0
		for i := startOfLineIndex; i < endOfLineIndex; i++ {
			child := YGNodeListGet(node.children, i)
			if child.style.positionType == YGPositionTypeRelative {
				if YGMarginLeadingValue(child, mainAxis).unit == YGUnitAuto {
					numberOfAutoMarginsOnCurrentLine++
				}
				if YGMarginTrailingValue(child, mainAxis).unit == YGUnitAuto {
					numberOfAutoMarginsOnCurrentLine++
				}
			}
		}

		if numberOfAutoMarginsOnCurrentLine == 0 {
			switch justifyContent {
			case YGJustifyCenter:
				leadingMainDim = remainingFreeSpace / 2
			case YGJustifyFlexEnd:
				leadingMainDim = remainingFreeSpace
			case YGJustifySpaceBetween:
				if itemsOnLine > 1 {
					betweenMainDim = fmaxf(remainingFreeSpace, 0) / float32(itemsOnLine-1)
				} else {
					betweenMainDim = 0
				}
			case YGJustifySpaceAround:
				// Space on the edges is half of the space between elements
				betweenMainDim = remainingFreeSpace / float32(itemsOnLine)
				leadingMainDim = betweenMainDim / 2
			case YGJustifyFlexStart:
			}
		}

		mainDim := leadingPaddingAndBorderMain + leadingMainDim
		var crossDim float32

		for i := startOfLineIndex; i < endOfLineIndex; i++ {
			child := YGNodeListGet(node.children, i)
			if child.style.display == YGDisplayNone {
				continue
			}
			if child.style.positionType == YGPositionTypeAbsolute &&
				YGNodeIsLeadingPosDefined(child, mainAxis) {
				if performLayout {
					// In case the child is position absolute and has left/top being
					// defined, we override the position to whatever the user said
					// (and margin/border).
					child.layout.position[pos[mainAxis]] =
						YGNodeLeadingPosition(child, mainAxis, availableInnerMainDim) +
							YGNodeLeadingBorder(node, mainAxis) +
							YGNodeLeadingMargin(child, mainAxis, availableInnerWidth)
				}
			} else {
				// Now that we placed the element, we need to update the variables.
				// We need to do that only for relative elements. Absolute elements
				// do not take part in that phase.
				if child.style.positionType == YGPositionTypeRelative {
					if YGMarginLeadingValue(child, mainAxis).unit == YGUnitAuto {
						mainDim += remainingFreeSpace / float32(numberOfAutoMarginsOnCurrentLine)
					}

					if performLayout {
						child.layout.position[pos[mainAxis]] += mainDim
					}

					if YGMarginTrailingValue(child, mainAxis).unit == YGUnitAuto {
						mainDim += remainingFreeSpace / float32(numberOfAutoMarginsOnCurrentLine)
					}

					if canSkipFlex {
						// If we skipped the flex step, then we can't rely on the
						// measuredDims because
						// they weren't computed. This means we can't call YGNodeDimWithMargin.
						mainDim += betweenMainDim + YGNodeMarginForAxis(child, mainAxis, availableInnerWidth) +
							child.layout.computedFlexBasis
						crossDim = availableInnerCrossDim
					} else {
						// The main dimension is the sum of all the elements dimension plus the spacing.
						mainDim += betweenMainDim + YGNodeDimWithMargin(child, mainAxis, availableInnerWidth)

						// The cross dimension is the max of the elements dimension since
						// there can only be one element in that cross dimension.
						crossDim = fmaxf(crossDim, YGNodeDimWithMargin(child, crossAxis, availableInnerWidth))
					}
				} else if performLayout {
					child.layout.position[pos[mainAxis]] +=
						YGNodeLeadingBorder(node, mainAxis) + leadingMainDim
				}
			}
		}

		mainDim += trailingPaddingAndBorderMain

		containerCrossAxis := availableInnerCrossDim
		if measureModeCrossDim == YGMeasureModeUndefined ||
			measureModeCrossDim == YGMeasureModeAtMost {
			// Compute the cross axis from the max cross dimension of the children.
			containerCrossAxis = YGNodeBoundAxis(node,
				crossAxis,
				crossDim+paddingAndBorderAxisCross,
				crossAxisParentSize,
				parentWidth) -
				paddingAndBorderAxisCross
		}

		// If there's no flex wrap, the cross dimension is defined by the container.
		if !isNodeFlexWrap && measureModeCrossDim == YGMeasureModeExactly {
			crossDim = availableInnerCrossDim
		}

		// Clamp to the min/max size specified on the container.
		crossDim = YGNodeBoundAxis(node,
			crossAxis,
			crossDim+paddingAndBorderAxisCross,
			crossAxisParentSize,
			parentWidth) -
			paddingAndBorderAxisCross

		// STEP 7: CROSS-AXIS ALIGNMENT
		// We can skip child alignment if we're just measuring the container.
		if performLayout {
			for i := startOfLineIndex; i < endOfLineIndex; i++ {
				child := YGNodeListGet(node.children, i)
				if child.style.display == YGDisplayNone {
					continue
				}
				if child.style.positionType == YGPositionTypeAbsolute {
					// If the child is absolutely positioned and has a
					// top/left/bottom/right
					// set, override all the previously computed positions to set it
					// correctly.
					if YGNodeIsLeadingPosDefined(child, crossAxis) {
						child.layout.position[pos[crossAxis]] =
							YGNodeLeadingPosition(child, crossAxis, availableInnerCrossDim) +
								YGNodeLeadingBorder(node, crossAxis) +
								YGNodeLeadingMargin(child, crossAxis, availableInnerWidth)
					} else {
						child.layout.position[pos[crossAxis]] =
							YGNodeLeadingBorder(node, crossAxis) +
								YGNodeLeadingMargin(child, crossAxis, availableInnerWidth)
					}
				} else {
					leadingCrossDim := leadingPaddingAndBorderCross

					// For a relative children, we're either using alignItems (parent) or
					// alignSelf (child) in order to determine the position in the cross
					// axis
					alignItem := YGNodeAlignItem(node, child)

					// If the child uses align stretch, we need to lay it out one more
					// time, this time
					// forcing the cross-axis size to be the computed cross size for the
					// current line.
					if alignItem == YGAlignStretch &&
						YGMarginLeadingValue(child, crossAxis).unit != YGUnitAuto &&
						YGMarginTrailingValue(child, crossAxis).unit != YGUnitAuto {
						// If the child defines a definite size for its cross axis, there's
						// no need to stretch.
						if !YGNodeIsStyleDimDefined(child, crossAxis, availableInnerCrossDim) {
							childMainSize := child.layout.measuredDimensions[dim[mainAxis]]
							childCrossSize := crossDim
							if !YGFloatIsUndefined(child.style.aspectRatio) {
								childCrossSize = YGNodeMarginForAxis(child, crossAxis, availableInnerWidth)
								if isMainAxisRow {
									childCrossSize += childMainSize / child.style.aspectRatio
								} else {
									childCrossSize += childMainSize * child.style.aspectRatio
								}
							}

							childMainSize += YGNodeMarginForAxis(child, mainAxis, availableInnerWidth)

							childMainMeasureMode := YGMeasureModeExactly
							childCrossMeasureMode := YGMeasureModeExactly
							YGConstrainMaxSizeForMode(child,
								mainAxis,
								availableInnerMainDim,
								availableInnerWidth,
								&childMainMeasureMode,
								&childMainSize)
							YGConstrainMaxSizeForMode(child,
								crossAxis,
								availableInnerCrossDim,
								availableInnerWidth,
								&childCrossMeasureMode,
								&childCrossSize)

							childWidth := childCrossSize
							if isMainAxisRow {
								childWidth = childMainSize
							}
							childHeight := childCrossSize
							if !isMainAxisRow {
								childHeight = childMainSize
							}

							childWidthMeasureMode := YGMeasureModeExactly
							if YGFloatIsUndefined(childWidth) {
								childWidthMeasureMode = YGMeasureModeUndefined
							}

							childHeightMeasureMode := YGMeasureModeExactly
							if YGFloatIsUndefined(childHeight) {
								childHeightMeasureMode = YGMeasureModeUndefined
							}

							YGLayoutNodeInternal(child,
								childWidth,
								childHeight,
								direction,
								childWidthMeasureMode,
								childHeightMeasureMode,
								availableInnerWidth,
								availableInnerHeight,
								true,
								"stretch",
								config)
						}
					} else {
						remainingCrossDim := containerCrossAxis - YGNodeDimWithMargin(child, crossAxis, availableInnerWidth)

						if YGMarginLeadingValue(child, crossAxis).unit == YGUnitAuto &&
							YGMarginTrailingValue(child, crossAxis).unit == YGUnitAuto {
							leadingCrossDim += fmaxf(0, remainingCrossDim/2)
						} else if YGMarginTrailingValue(child, crossAxis).unit == YGUnitAuto {
							// No-Op
						} else if YGMarginLeadingValue(child, crossAxis).unit == YGUnitAuto {
							leadingCrossDim += fmaxf(0, remainingCrossDim)
						} else if alignItem == YGAlignFlexStart {
							// No-Op
						} else if alignItem == YGAlignCenter {
							leadingCrossDim += remainingCrossDim / 2
						} else {
							leadingCrossDim += remainingCrossDim
						}
					}
					// And we apply the position
					child.layout.position[pos[crossAxis]] += totalLineCrossDim + leadingCrossDim
				}
			}
		}

		totalLineCrossDim += crossDim
		maxLineMainDim = fmaxf(maxLineMainDim, mainDim)

		lineCount++
		startOfLineIndex = endOfLineIndex

	}

	// STEP 8: MULTI-LINE CONTENT ALIGNMENT
	if performLayout && (lineCount > 1 || YGIsBaselineLayout(node)) &&
		!YGFloatIsUndefined(availableInnerCrossDim) {
		remainingAlignContentDim := availableInnerCrossDim - totalLineCrossDim

		var crossDimLead float32
		currentLead := leadingPaddingAndBorderCross

		switch node.style.alignContent {
		case YGAlignFlexEnd:
			currentLead += remainingAlignContentDim
		case YGAlignCenter:
			currentLead += remainingAlignContentDim / 2
		case YGAlignStretch:
			if availableInnerCrossDim > totalLineCrossDim {
				crossDimLead = remainingAlignContentDim / float32(lineCount)
			}
		case YGAlignSpaceAround:
			if availableInnerCrossDim > totalLineCrossDim {
				currentLead += remainingAlignContentDim / float32(2*lineCount)
				if lineCount > 1 {
					crossDimLead = remainingAlignContentDim / float32(lineCount)
				}
			} else {
				currentLead += remainingAlignContentDim / 2
			}
		case YGAlignSpaceBetween:
			if availableInnerCrossDim > totalLineCrossDim && lineCount > 1 {
				crossDimLead = remainingAlignContentDim / float32(lineCount-1)
			}
		case YGAlignAuto:
		case YGAlignFlexStart:
		case YGAlignBaseline:
		}

		endIndex := 0
		for i := 0; i < lineCount; i++ {
			startIndex := endIndex
			var ii int

			// compute the line's height and find the endIndex
			var lineHeight float32
			var maxAscentForCurrentLine float32
			var maxDescentForCurrentLine float32
			for ii = startIndex; ii < childCount; ii++ {
				child := YGNodeListGet(node.children, ii)
				if child.style.display == YGDisplayNone {
					continue
				}
				if child.style.positionType == YGPositionTypeRelative {
					if child.lineIndex != i {
						break
					}
					if YGNodeIsLayoutDimDefined(child, crossAxis) {
						lineHeight = fmaxf(lineHeight,
							child.layout.measuredDimensions[dim[crossAxis]]+
								YGNodeMarginForAxis(child, crossAxis, availableInnerWidth))
					}
					if YGNodeAlignItem(node, child) == YGAlignBaseline {
						ascent := YGBaseline(child) + YGNodeLeadingMargin(child, YGFlexDirectionColumn, availableInnerWidth)
						descent := child.layout.measuredDimensions[YGDimensionHeight] + YGNodeMarginForAxis(child, YGFlexDirectionColumn, availableInnerWidth) - ascent
						maxAscentForCurrentLine = fmaxf(maxAscentForCurrentLine, ascent)
						maxDescentForCurrentLine = fmaxf(maxDescentForCurrentLine, descent)
						lineHeight = fmaxf(lineHeight, maxAscentForCurrentLine+maxDescentForCurrentLine)
					}
				}
			}
			endIndex = ii
			lineHeight += crossDimLead

			if performLayout {
				for ii = startIndex; ii < endIndex; ii++ {
					child := YGNodeListGet(node.children, ii)
					if child.style.display == YGDisplayNone {
						continue
					}
					if child.style.positionType == YGPositionTypeRelative {
						switch YGNodeAlignItem(node, child) {
						case YGAlignFlexStart:
							{
								child.layout.position[pos[crossAxis]] =
									currentLead + YGNodeLeadingMargin(child, crossAxis, availableInnerWidth)
							}
						case YGAlignFlexEnd:
							{
								child.layout.position[pos[crossAxis]] =
									currentLead + lineHeight -
										YGNodeTrailingMargin(child, crossAxis, availableInnerWidth) -
										child.layout.measuredDimensions[dim[crossAxis]]
							}
						case YGAlignCenter:
							{
								childHeight := child.layout.measuredDimensions[dim[crossAxis]]
								child.layout.position[pos[crossAxis]] = currentLead + (lineHeight-childHeight)/2
							}
						case YGAlignStretch:
							{
								child.layout.position[pos[crossAxis]] =
									currentLead + YGNodeLeadingMargin(child, crossAxis, availableInnerWidth)

								// Remeasure child with the line height as it as been only measured with the
								// parents height yet.
								if !YGNodeIsStyleDimDefined(child, crossAxis, availableInnerCrossDim) {
									childWidth := lineHeight
									if isMainAxisRow {
										childWidth = child.layout.measuredDimensions[YGDimensionWidth] +
											YGNodeMarginForAxis(child, mainAxis, availableInnerWidth)
									}

									childHeight := lineHeight
									if !isMainAxisRow {
										childHeight = child.layout.measuredDimensions[YGDimensionHeight] +
											YGNodeMarginForAxis(child, crossAxis, availableInnerWidth)
									}

									if !(YGFloatsEqual(childWidth,
										child.layout.measuredDimensions[YGDimensionWidth]) &&
										YGFloatsEqual(childHeight,
											child.layout.measuredDimensions[YGDimensionHeight])) {
										YGLayoutNodeInternal(child,
											childWidth,
											childHeight,
											direction,
											YGMeasureModeExactly,
											YGMeasureModeExactly,
											availableInnerWidth,
											availableInnerHeight,
											true,
											"multiline-stretch",
											config)
									}
								}
							}
						case YGAlignBaseline:
							{
								child.layout.position[YGEdgeTop] =
									currentLead + maxAscentForCurrentLine - YGBaseline(child) +
										YGNodeLeadingPosition(child, YGFlexDirectionColumn, availableInnerCrossDim)
							}
						case YGAlignAuto:
						case YGAlignSpaceBetween:
						case YGAlignSpaceAround:
						}
					}
				}
			}

			currentLead += lineHeight
		}
	}

	// STEP 9: COMPUTING FINAL DIMENSIONS
	node.layout.measuredDimensions[YGDimensionWidth] = YGNodeBoundAxis(
		node, YGFlexDirectionRow, availableWidth-marginAxisRow, parentWidth, parentWidth)
	node.layout.measuredDimensions[YGDimensionHeight] = YGNodeBoundAxis(
		node, YGFlexDirectionColumn, availableHeight-marginAxisColumn, parentHeight, parentWidth)

	// If the user didn't specify a width or height for the node, set the
	// dimensions based on the children.
	if measureModeMainDim == YGMeasureModeUndefined ||
		(node.style.overflow != YGOverflowScroll && measureModeMainDim == YGMeasureModeAtMost) {
		// Clamp the size to the min/max size, if specified, and make sure it
		// doesn't go below the padding and border amount.
		node.layout.measuredDimensions[dim[mainAxis]] =
			YGNodeBoundAxis(node, mainAxis, maxLineMainDim, mainAxisParentSize, parentWidth)
	} else if measureModeMainDim == YGMeasureModeAtMost &&
		node.style.overflow == YGOverflowScroll {
		node.layout.measuredDimensions[dim[mainAxis]] = fmaxf(
			fminf(availableInnerMainDim+paddingAndBorderAxisMain,
				YGNodeBoundAxisWithinMinAndMax(node, mainAxis, maxLineMainDim, mainAxisParentSize)),
			paddingAndBorderAxisMain)
	}

	if measureModeCrossDim == YGMeasureModeUndefined ||
		(node.style.overflow != YGOverflowScroll && measureModeCrossDim == YGMeasureModeAtMost) {
		// Clamp the size to the min/max size, if specified, and make sure it
		// doesn't go below the padding and border amount.
		node.layout.measuredDimensions[dim[crossAxis]] =
			YGNodeBoundAxis(node,
				crossAxis,
				totalLineCrossDim+paddingAndBorderAxisCross,
				crossAxisParentSize,
				parentWidth)
	} else if measureModeCrossDim == YGMeasureModeAtMost &&
		node.style.overflow == YGOverflowScroll {
		node.layout.measuredDimensions[dim[crossAxis]] =
			fmaxf(fminf(availableInnerCrossDim+paddingAndBorderAxisCross,
				YGNodeBoundAxisWithinMinAndMax(node,
					crossAxis,
					totalLineCrossDim+paddingAndBorderAxisCross,
					crossAxisParentSize)),
				paddingAndBorderAxisCross)
	}

	// As we only wrapped in normal direction yet, we need to reverse the positions on wrap-reverse.
	if performLayout && node.style.flexWrap == YGWrapWrapReverse {
		for i := 0; i < childCount; i++ {
			child := YGNodeGetChild(node, i)
			if child.style.positionType == YGPositionTypeRelative {
				child.layout.position[pos[crossAxis]] = node.layout.measuredDimensions[dim[crossAxis]] -
					child.layout.position[pos[crossAxis]] -
					child.layout.measuredDimensions[dim[crossAxis]]
			}
		}
	}

	if performLayout {
		// STEP 10: SIZING AND POSITIONING ABSOLUTE CHILDREN
		for currentAbsoluteChild = firstAbsoluteChild; currentAbsoluteChild != nil; currentAbsoluteChild = currentAbsoluteChild.nextChild {
			mode := measureModeCrossDim
			if isMainAxisRow {
				mode = measureModeMainDim
			}

			YGNodeAbsoluteLayoutChild(node,
				currentAbsoluteChild,
				availableInnerWidth,
				mode,
				availableInnerHeight,
				direction,
				config)
		}

		// STEP 11: SETTING TRAILING POSITIONS FOR CHILDREN
		needsMainTrailingPos := mainAxis == YGFlexDirectionRowReverse || mainAxis == YGFlexDirectionColumnReverse
		needsCrossTrailingPos := crossAxis == YGFlexDirectionRowReverse || crossAxis == YGFlexDirectionColumnReverse

		// Set trailing position if necessary.
		if needsMainTrailingPos || needsCrossTrailingPos {
			for i := 0; i < childCount; i++ {
				child := YGNodeListGet(node.children, i)
				if child.style.display == YGDisplayNone {
					continue
				}
				if needsMainTrailingPos {
					YGNodeSetChildTrailingPosition(node, child, mainAxis)
				}

				if needsCrossTrailingPos {
					YGNodeSetChildTrailingPosition(node, child, crossAxis)
				}
			}
		}
	}
}

func YGNodeAbsoluteLayoutChild(node *YGNode, child *YGNode, width float32, widthMode YGMeasureMode, height float32, direction YGDirection, config *YGConfig) {
	mainAxis := YGResolveFlexDirection(node.style.flexDirection, direction)
	crossAxis := YGFlexDirectionCross(mainAxis, direction)
	isMainAxisRow := YGFlexDirectionIsRow(mainAxis)

	childWidth := YGUndefined
	childHeight := YGUndefined
	childWidthMeasureMode := YGMeasureModeUndefined
	childHeightMeasureMode := YGMeasureModeUndefined

	marginRow := YGNodeMarginForAxis(child, YGFlexDirectionRow, width)
	marginColumn := YGNodeMarginForAxis(child, YGFlexDirectionColumn, width)

	if YGNodeIsStyleDimDefined(child, YGFlexDirectionRow, width) {
		childWidth = YGResolveValue(child.resolvedDimensions[YGDimensionWidth], width) + marginRow
	} else {
		// If the child doesn't have a specified width, compute the width based
		// on the left/right
		// offsets if they're defined.
		if YGNodeIsLeadingPosDefined(child, YGFlexDirectionRow) &&
			YGNodeIsTrailingPosDefined(child, YGFlexDirectionRow) {
			childWidth = node.layout.measuredDimensions[YGDimensionWidth] -
				(YGNodeLeadingBorder(node, YGFlexDirectionRow) +
					YGNodeTrailingBorder(node, YGFlexDirectionRow)) -
				(YGNodeLeadingPosition(child, YGFlexDirectionRow, width) +
					YGNodeTrailingPosition(child, YGFlexDirectionRow, width))
			childWidth = YGNodeBoundAxis(child, YGFlexDirectionRow, childWidth, width, width)
		}
	}

	if YGNodeIsStyleDimDefined(child, YGFlexDirectionColumn, height) {
		childHeight =
			YGResolveValue(child.resolvedDimensions[YGDimensionHeight], height) + marginColumn
	} else {
		// If the child doesn't have a specified height, compute the height
		// based on the top/bottom
		// offsets if they're defined.
		if YGNodeIsLeadingPosDefined(child, YGFlexDirectionColumn) &&
			YGNodeIsTrailingPosDefined(child, YGFlexDirectionColumn) {
			childHeight = node.layout.measuredDimensions[YGDimensionHeight] -
				(YGNodeLeadingBorder(node, YGFlexDirectionColumn) +
					YGNodeTrailingBorder(node, YGFlexDirectionColumn)) -
				(YGNodeLeadingPosition(child, YGFlexDirectionColumn, height) +
					YGNodeTrailingPosition(child, YGFlexDirectionColumn, height))
			childHeight = YGNodeBoundAxis(child, YGFlexDirectionColumn, childHeight, height, width)
		}
	}

	// Exactly one dimension needs to be defined for us to be able to do aspect ratio
	// calculation. One dimension being the anchor and the other being flexible.
	if YGFloatIsUndefined(childWidth) != YGFloatIsUndefined(childHeight) {
		if !YGFloatIsUndefined(child.style.aspectRatio) {
			if YGFloatIsUndefined(childWidth) {
				childWidth =
					marginRow + fmaxf((childHeight-marginColumn)*child.style.aspectRatio,
						YGNodePaddingAndBorderForAxis(child, YGFlexDirectionColumn, width))
			} else if YGFloatIsUndefined(childHeight) {
				childHeight =
					marginColumn + fmaxf((childWidth-marginRow)/child.style.aspectRatio,
						YGNodePaddingAndBorderForAxis(child, YGFlexDirectionRow, width))
			}
		}
	}

	// If we're still missing one or the other dimension, measure the content.
	if YGFloatIsUndefined(childWidth) || YGFloatIsUndefined(childHeight) {
		childWidthMeasureMode = YGMeasureModeExactly
		if YGFloatIsUndefined(childWidth) {
			childWidthMeasureMode = YGMeasureModeUndefined
		}
		childHeightMeasureMode = YGMeasureModeExactly
		if YGFloatIsUndefined(childHeight) {
			childHeightMeasureMode = YGMeasureModeUndefined
		}

		// If the size of the parent is defined then try to rain the absolute child to that size
		// as well. This allows text within the absolute child to wrap to the size of its parent.
		// This is the same behavior as many browsers implement.
		if !isMainAxisRow && YGFloatIsUndefined(childWidth) && widthMode != YGMeasureModeUndefined &&
			width > 0 {
			childWidth = width
			childWidthMeasureMode = YGMeasureModeAtMost
		}

		YGLayoutNodeInternal(child,
			childWidth,
			childHeight,
			direction,
			childWidthMeasureMode,
			childHeightMeasureMode,
			childWidth,
			childHeight,
			false,
			"abs-measure",
			config)
		childWidth = child.layout.measuredDimensions[YGDimensionWidth] +
			YGNodeMarginForAxis(child, YGFlexDirectionRow, width)
		childHeight = child.layout.measuredDimensions[YGDimensionHeight] +
			YGNodeMarginForAxis(child, YGFlexDirectionColumn, width)
	}

	YGLayoutNodeInternal(child,
		childWidth,
		childHeight,
		direction,
		YGMeasureModeExactly,
		YGMeasureModeExactly,
		childWidth,
		childHeight,
		true,
		"abs-layout",
		config)

	if YGNodeIsTrailingPosDefined(child, mainAxis) && !YGNodeIsLeadingPosDefined(child, mainAxis) {
		axisSize := height
		if isMainAxisRow {
			axisSize = width
		}
		child.layout.position[leading[mainAxis]] = node.layout.measuredDimensions[dim[mainAxis]] -
			child.layout.measuredDimensions[dim[mainAxis]] -
			YGNodeTrailingBorder(node, mainAxis) -
			YGNodeTrailingMargin(child, mainAxis, width) -
			YGNodeTrailingPosition(child, mainAxis, axisSize)
	} else if !YGNodeIsLeadingPosDefined(child, mainAxis) &&
		node.style.justifyContent == YGJustifyCenter {
		child.layout.position[leading[mainAxis]] = (node.layout.measuredDimensions[dim[mainAxis]] -
			child.layout.measuredDimensions[dim[mainAxis]]) /
			2.0
	} else if !YGNodeIsLeadingPosDefined(child, mainAxis) &&
		node.style.justifyContent == YGJustifyFlexEnd {
		child.layout.position[leading[mainAxis]] = (node.layout.measuredDimensions[dim[mainAxis]] -
			child.layout.measuredDimensions[dim[mainAxis]])
	}

	if YGNodeIsTrailingPosDefined(child, crossAxis) &&
		!YGNodeIsLeadingPosDefined(child, crossAxis) {
		axisSize := height
		if isMainAxisRow {
			axisSize = width
		}

		child.layout.position[leading[crossAxis]] = node.layout.measuredDimensions[dim[crossAxis]] -
			child.layout.measuredDimensions[dim[crossAxis]] -
			YGNodeTrailingBorder(node, crossAxis) -
			YGNodeTrailingMargin(child, crossAxis, width) -
			YGNodeTrailingPosition(child, crossAxis, axisSize)
	} else if !YGNodeIsLeadingPosDefined(child, crossAxis) &&
		YGNodeAlignItem(node, child) == YGAlignCenter {
		child.layout.position[leading[crossAxis]] =
			(node.layout.measuredDimensions[dim[crossAxis]] -
				child.layout.measuredDimensions[dim[crossAxis]]) /
				2.0
	} else if !YGNodeIsLeadingPosDefined(child, crossAxis) &&
		((YGNodeAlignItem(node, child) == YGAlignFlexEnd) != (node.style.flexWrap == YGWrapWrapReverse)) {
		child.layout.position[leading[crossAxis]] = (node.layout.measuredDimensions[dim[crossAxis]] -
			child.layout.measuredDimensions[dim[crossAxis]])
	}
}

// This is a wrapper around the YGNodelayoutImpl function. It determines
// whether the layout request is redundant and can be skipped.
//
// Parameters:
//  Input parameters are the same as YGNodelayoutImpl (see above)
//  Return parameter is true if layout was performed, false if skipped
func YGLayoutNodeInternal(node *YGNode, availableWidth float32, availableHeight float32,
	parentDirection YGDirection, widthMeasureMode YGMeasureMode,
	heightMeasureMode YGMeasureMode, parentWidth float32, parentHeight float32,
	performLayout bool, reason string, config *YGConfig) bool {
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

func calcStartWidth(node *YGNode, parentWidth float32) (float32, YGMeasureMode) {
	if YGNodeIsStyleDimDefined(node, YGFlexDirectionRow, parentWidth) {
		width := YGResolveValue(node.resolvedDimensions[dim[YGFlexDirectionRow]], parentWidth)
		margin := YGNodeMarginForAxis(node, YGFlexDirectionRow, parentWidth)
		return width + margin, YGMeasureModeExactly
	}
	if YGResolveValue(&node.style.maxDimensions[YGDimensionWidth], parentWidth) >= 0.0 {
		width := YGResolveValue(&node.style.maxDimensions[YGDimensionWidth], parentWidth)
		return width, YGMeasureModeAtMost
	}

	width := parentWidth
	widthMeasureMode := YGMeasureModeExactly
	if YGFloatIsUndefined(width) {
		widthMeasureMode = YGMeasureModeUndefined
	}
	return width, widthMeasureMode
}
func calcStartHeight(node *YGNode, parentWidth, parentHeight float32) (float32, YGMeasureMode) {
	if YGNodeIsStyleDimDefined(node, YGFlexDirectionColumn, parentHeight) {
		height := YGResolveValue(node.resolvedDimensions[dim[YGFlexDirectionColumn]], parentHeight)
		margin := YGNodeMarginForAxis(node, YGFlexDirectionColumn, parentWidth)
		return height + margin, YGMeasureModeExactly
	}
	if YGResolveValue(&node.style.maxDimensions[YGDimensionHeight], parentHeight) >= 0 {
		height := YGResolveValue(&node.style.maxDimensions[YGDimensionHeight], parentHeight)
		return height, YGMeasureModeAtMost
	}
	height := parentHeight
	heightMeasureMode := YGMeasureModeExactly
	if YGFloatIsUndefined(height) {
		heightMeasureMode = YGMeasureModeUndefined
	}
	return height, heightMeasureMode
}

func YGNodeCalculateLayout(node *YGNode, parentWidth float32, parentHeight float32, parentDirection YGDirection) {
	// Increment the generation count. This will force the recursive routine to
	// visit
	// all dirty nodes at least once. Subsequent visits will be skipped if the
	// input
	// parameters don't change.
	gCurrentGenerationCount++

	YGResolveDimensions(node)

	width, widthMeasureMode := calcStartWidth(node, parentWidth)
	height, heightMeasureMode := calcStartHeight(node, parentWidth, parentHeight)

	if YGLayoutNodeInternal(node, width, height, parentDirection,
		widthMeasureMode, heightMeasureMode, parentWidth, parentHeight,
		true, "initial", node.config) {
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

	node.layout.position[YGEdgeLeft] = YGRoundValueToPixelGrid(nodeLeft, pointScaleFactor, false, textRounding)
	node.layout.position[YGEdgeTop] = YGRoundValueToPixelGrid(nodeTop, pointScaleFactor, false, textRounding)

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

func YGConfigSetExperimentalFeatureEnabled(config *YGConfig, feature YGExperimentalFeature, enabled bool) {
	config.experimentalFeatures[feature] = enabled
}

func YGConfigIsExperimentalFeatureEnabled(config *YGConfig, feature YGExperimentalFeature) bool {
	return config.experimentalFeatures[feature]
}

func YGConfigSetUseWebDefaults(config *YGConfig, enabled bool) {
	config.useWebDefaults = enabled
}

func YGConfigSetUseLegacyStretchBehaviour(config *YGConfig, useLegacyStretchBehaviour bool) {
	config.useLegacyStretchBehaviour = useLegacyStretchBehaviour
}

func YGConfigGetUseWebDefaults(config *YGConfig) bool {
	return config.useWebDefaults
}

func YGConfigSetContext(config *YGConfig, context interface{}) {
	config.context = context
}

func YGConfigGetContext(config *YGConfig) interface{} {
	return config.context
}

func YGLog(node *YGNode, level YGLogLevel, format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// YGAssert asserts that cond is true
func YGAssert(cond bool, format string, args ...interface{}) {
	if !cond {
		panic(format)
	}
}

// YGAssertWithNode assert if cond is not true
func YGAssertWithNode(node *YGNode, cond bool, format string, args ...interface{}) {
	YGAssert(cond, format, args...)
}

func YGAssertWithConfig(config *YGConfig, condition bool, message string) {
	if !condition {
		panic(message)
	}
}
