package flex

import (
	"fmt"
	"os"
)

// CachedMeasurement describes measurements
type CachedMeasurement struct {
	availableWidth    float32
	availableHeight   float32
	widthMeasureMode  MeasureMode
	heightMeasureMode MeasureMode

	computedWidth  float32
	computedHeight float32
}

// This value was chosen based on empiracle data. Even the most complicated
// layouts should not require more than 16 entries to fit within the cache.
const YG_MAX_CACHED_RESULT_COUNT = 16

// Layout describes layout
type Layout struct {
	position   [4]float32
	dimensions [2]float32
	margin     [6]float32
	border     [6]float32
	padding    [6]float32
	direction  Direction

	computedFlexBasisGeneration int
	computedFlexBasis           float32
	hadOverflow                 bool

	// Instead of recomputing the entire layout every single time, we
	// cache some information to break early when nothing changed
	generationCount     int
	lastParentDirection Direction

	nextCachedMeasurementsIndex int
	cachedMeasurements          [YG_MAX_CACHED_RESULT_COUNT]CachedMeasurement

	measuredDimensions [2]float32

	cachedLayout CachedMeasurement
}

// Style describes a style
type Style struct {
	direction      Direction
	flexDirection  FlexDirection
	justifyContent Justify
	alignContent   Align
	alignItems     Align
	alignSelf      Align
	positionType   PositionType
	flexWrap       Wrap
	overflow       Overflow
	display        Display
	flex           float32
	flexGrow       float32
	flexShrink     float32
	flexBasis      Value
	margin         [EdgeCount]Value
	position       [EdgeCount]Value
	padding        [EdgeCount]Value
	border         [EdgeCount]Value
	dimensions     [2]Value
	minDimensions  [2]Value
	maxDimensions  [2]Value

	// Yoga specific properties, not compatible with flexbox specification
	aspectRatio float32
}

// Config describes a config
type Config struct {
	experimentalFeatures      [experimentalFeatureCount + 1]bool
	useWebDefaults            bool
	useLegacyStretchBehaviour bool
	pointScaleFactor          float32
	logger                    Logger
	context                   interface{}
}

// Node describes a node
type Node struct {
	style     Style
	layout    Layout
	lineIndex int

	parent   *Node
	children *YGNodeList

	nextChild *Node

	measure  MeasureFunc
	baseline BaselineFunc
	print    PrintFunc
	config   *Config
	context  interface{}

	isDirty      bool
	hasNewLayout bool
	nodeType     NodeType

	resolvedDimensions [2]*Value
}

var (
	YG_UNDEFINED_VALUES = Value{
		Value: Undefined,
		Unit:  UnitUndefined,
	}

	YG_AUTO_VALUES = Value{
		Value: Undefined,
		Unit:  UnitAuto,
	}

	YG_DEFAULT_EDGE_VALUES_UNIT = [EdgeCount]Value{
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
		Undefined,
		Undefined,
	}

	YG_DEFAULT_DIMENSION_VALUES_UNIT = [2]Value{
		YG_UNDEFINED_VALUES,
		YG_UNDEFINED_VALUES,
	}

	YG_DEFAULT_DIMENSION_VALUES_AUTO_UNIT = [2]Value{
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
	gYGNodeDefaults = Node{
		parent:             nil,
		children:           nil,
		hasNewLayout:       true,
		isDirty:            false,
		nodeType:           NodeTypeDefault,
		resolvedDimensions: [2]*Value{&ValueUndefined, &ValueUndefined},
		style: Style{
			flex:           Undefined,
			flexGrow:       Undefined,
			flexShrink:     Undefined,
			flexBasis:      YG_AUTO_VALUES,
			justifyContent: JustifyFlexStart,
			alignItems:     AlignStretch,
			alignContent:   AlignFlexStart,
			direction:      DirectionInherit,
			flexDirection:  FlexDirectionColumn,
			overflow:       OverflowVisible,
			display:        DisplayFlex,
			dimensions:     YG_DEFAULT_DIMENSION_VALUES_AUTO_UNIT,
			minDimensions:  YG_DEFAULT_DIMENSION_VALUES_UNIT,
			maxDimensions:  YG_DEFAULT_DIMENSION_VALUES_UNIT,
			position:       YG_DEFAULT_EDGE_VALUES_UNIT,
			margin:         YG_DEFAULT_EDGE_VALUES_UNIT,
			padding:        YG_DEFAULT_EDGE_VALUES_UNIT,
			border:         YG_DEFAULT_EDGE_VALUES_UNIT,
			aspectRatio:    Undefined,
		},
		layout: Layout{
			dimensions:                  YG_DEFAULT_DIMENSION_VALUES,
			lastParentDirection:         Direction(-1),
			nextCachedMeasurementsIndex: 0,
			computedFlexBasis:           Undefined,
			hadOverflow:                 false,
			measuredDimensions:          YG_DEFAULT_DIMENSION_VALUES,
			cachedLayout: CachedMeasurement{
				widthMeasureMode:  MeasureMode(-1),
				heightMeasureMode: MeasureMode(-1),
				computedWidth:     -1,
				computedHeight:    -1,
			},
		},
	}

	gYGConfigDefaults = Config{
		experimentalFeatures: [experimentalFeatureCount + 1]bool{
			false,
			false,
		},
		useWebDefaults:   false,
		pointScaleFactor: 1,
		logger:           YGDefaultLog,
		context:          nil,
	}

	// ValueZero defines a zero value
	ValueZero = Value{Value: 0, Unit: UnitPoint}
)

func valueEq(v1, v2 Value) bool {
	if v1.Unit != v2.Unit {
		return false
	}
	return feq(v1.Value, v2.Value)
}

// YGDefaultLog is default logging function
func YGDefaultLog(config *Config, node *Node, level LogLevel, format string,
	args ...interface{}) int {
	switch level {
	case LogLevelError, LogLevelFatal:
		n, _ := fmt.Fprintf(os.Stderr, format, args...)
		return n
	case LogLevelWarn, LogLevelInfo, LogLevelDebug, LogLevelVerbose:
		fallthrough
	default:
		n, _ := fmt.Printf(format, args...)
		return n
	}
}

// YGComputedEdgeValue computes edge value
func YGComputedEdgeValue(edges []Value, edge Edge, defaultValue *Value) *Value {
	if edges[edge].Unit != UnitUndefined {
		return &edges[edge]
	}

	isVertEdge := edge == EdgeTop || edge == EdgeBottom
	if isVertEdge && edges[EdgeVertical].Unit != UnitUndefined {
		return &edges[EdgeVertical]
	}

	isHorizEdge := (edge == EdgeLeft || edge == EdgeRight || edge == EdgeStart || edge == EdgeEnd)
	if isHorizEdge && edges[EdgeHorizontal].Unit != UnitUndefined {
		return &edges[EdgeHorizontal]
	}

	if edges[EdgeAll].Unit != UnitUndefined {
		return &edges[EdgeAll]
	}

	if edge == EdgeStart || edge == EdgeEnd {
		return &ValueUndefined
	}

	return defaultValue
}

// YGResolveValue resolves a value
func YGResolveValue(value *Value, parentSize float32) float32 {
	switch value.Unit {
	case UnitUndefined, UnitAuto:
		return Undefined
	case UnitPoint:
		return value.Value
	case UnitPercent:
		return value.Value * parentSize / 100
	}
	return Undefined
}

// YGResolveValueMargin resolves margin value
func YGResolveValueMargin(value *Value, parentSize float32) float32 {
	if value.Unit == UnitAuto {
		return 0
	}
	return YGResolveValue(value, parentSize)
}

// YGNodeNewWithConfig creates new node with config
func YGNodeNewWithConfig(config *Config) *Node {
	node := gYGNodeDefaults

	if config.useWebDefaults {
		node.style.flexDirection = FlexDirectionRow
		node.style.alignContent = AlignStretch
	}
	node.config = config
	return &node
}

// YGNodeNew creates a new node
func YGNodeNew() *Node {
	return YGNodeNewWithConfig(&gYGConfigDefaults)
}

// YGNodeReset resets a node
func YGNodeReset(node *Node) {
	YGAssertWithNode(node, YGNodeGetChildCount(node) == 0, "Cannot reset a node which still has children attached")
	YGAssertWithNode(node, node.parent == nil, "Cannot reset a node still attached to a parent")

	YGNodeListFree(node.children)

	config := node.config
	*node = gYGNodeDefaults
	if config.useWebDefaults {
		node.style.flexDirection = FlexDirectionRow
		node.style.alignContent = AlignStretch
	}
	node.config = config
}

// YGConfigGetDefault returns default config, only for C#
func YGConfigGetDefault() *Config {
	return &gYGConfigDefaults
}

// YGConfigNew creates new config
func YGConfigNew() *Config {
	config := &Config{}
	YGAssert(config != nil, "Could not allocate memory for config")

	*config = gYGConfigDefaults
	return config
}

// YGConfigCopy copies a config
func YGConfigCopy(dest *Config, src *Config) {
	*dest = *src
}

// YGNodeMarkDirtyInternal marks the node as dirty, internally
func YGNodeMarkDirtyInternal(node *Node) {
	if !node.isDirty {
		node.isDirty = true
		node.layout.computedFlexBasis = Undefined
		if node.parent != nil {
			YGNodeMarkDirtyInternal(node.parent)
		}
	}
}

// YGNodeSetMeasureFunc sets measure function
func YGNodeSetMeasureFunc(node *Node, measureFunc MeasureFunc) {
	if measureFunc == nil {
		node.measure = nil
		// TODO: t18095186 Move nodeType to opt-in function and mark appropriate places in Litho
		node.nodeType = NodeTypeDefault
	} else {
		YGAssertWithNode(
			node,
			YGNodeGetChildCount(node) == 0,
			"Cannot set measure function: Nodes with measure functions cannot have children.")
		node.measure = measureFunc
		// TODO: t18095186 Move nodeType to opt-in function and mark appropriate places in Litho
		node.nodeType = NodeTypeText
	}
}

// YGNodeGetMeasureFunc gets measure function
func YGNodeGetMeasureFunc(node *Node) MeasureFunc {
	return node.measure
}

// YGNodeSetBaselineFunc sets baseline function
func YGNodeSetBaselineFunc(node *Node, baselineFunc BaselineFunc) {
	node.baseline = baselineFunc
}

// YGNodeGetBaselineFunc gets baseline function
func YGNodeGetBaselineFunc(node *Node) BaselineFunc {
	return node.baseline
}

// YGNodeInsertChild inserts a child
func YGNodeInsertChild(node *Node, child *Node, index int) {
	YGAssertWithNode(node, child.parent == nil, "Child already has a parent, it must be removed first.")
	YGAssertWithNode(node, node.measure == nil, "Cannot add child: Nodes with measure functions cannot have children.")

	YGNodeListInsert(&node.children, child, index)
	child.parent = node
	YGNodeMarkDirtyInternal(node)
}

// YGNodeRemoveChild removes the child
func YGNodeRemoveChild(node *Node, child *Node) {
	if YGNodeListDelete(node.children, child) != nil {
		child.layout = gYGNodeDefaults.layout // layout is no longer valid
		child.parent = nil
		YGNodeMarkDirtyInternal(node)
	}
}

// YGNodeGetChild returns a child
func YGNodeGetChild(node *Node, index int) *Node {
	return YGNodeListGet(node.children, index)
}

// YGNodeGetParent gets parent
func YGNodeGetParent(node *Node) *Node {
	return node.parent
}

// YGNodeGetChildCount returns number of children
func YGNodeGetChildCount(node *Node) int {
	return YGNodeListCount(node.children)
}

// YGNodeMarkDirty marks node as dirty
func YGNodeMarkDirty(node *Node) {
	YGAssertWithNode(node, node.measure != nil,
		"Only leaf nodes with custom measure functions should manually mark themselves as dirty")
	YGNodeMarkDirtyInternal(node)
}

// YGNodeIsDirty returns true if node is dirty
func YGNodeIsDirty(node *Node) bool {
	return node.isDirty
}

func styleEq(s1, s2 *Style) bool {
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
		!feq(s1.flex, s2.flex) ||
		!feq(s1.flexGrow, s2.flexGrow) ||
		!feq(s1.flexShrink, s2.flexShrink) ||
		!valueEq(s1.flexBasis, s2.flexBasis) {
		return false
	}
	for i := 0; i < EdgeCount; i++ {
		if !valueEq(s1.margin[i], s2.margin[i]) ||
			!valueEq(s1.position[i], s2.position[i]) ||
			!valueEq(s1.padding[i], s2.padding[i]) ||
			!valueEq(s1.border[i], s2.border[i]) {
			return false
		}
	}
	for i := 0; i < 2; i++ {
		if !valueEq(s1.dimensions[i], s2.dimensions[i]) ||
			!valueEq(s1.minDimensions[i], s2.minDimensions[i]) ||
			!valueEq(s1.maxDimensions[i], s2.maxDimensions[i]) {
			return false
		}
	}
	return true
}

// YGNodeCopyStyle copies style
func YGNodeCopyStyle(dstNode *Node, srcNode *Node) {
	if !styleEq(&dstNode.style, &srcNode.style) {
		dstNode.style = srcNode.style
		YGNodeMarkDirtyInternal(dstNode)
	}
}

// YGResolveFlexGrow resolves flex grow
func YGResolveFlexGrow(node *Node) float32 {
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
func YGNodeStyleGetFlexGrow(node *Node) float32 {
	if YGFloatIsUndefined(node.style.flexGrow) {
		return kDefaultFlexGrow
	}
	return node.style.flexGrow
}

// YGNodeStyleGetFlexShrink gets flex shrink
func YGNodeStyleGetFlexShrink(node *Node) float32 {
	if YGFloatIsUndefined(node.style.flexShrink) {
		if node.config.useWebDefaults {
			return kWebDefaultFlexShrink
		}
		return kDefaultFlexShrink
	}
	return node.style.flexShrink
}

// YGNodeResolveFlexShrink resolves flex shrink
func YGNodeResolveFlexShrink(node *Node) float32 {
	// Root nodes flexShrink should always be 0
	if node.parent == nil {
		return 0
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
func YGNodeResolveFlexBasisPtr(node *Node) *Value {
	style := &node.style
	if style.flexBasis.Unit != UnitAuto && style.flexBasis.Unit != UnitUndefined {
		return &style.flexBasis
	}
	if !YGFloatIsUndefined(style.flex) && style.flex > 0 {
		if node.config.useWebDefaults {
			return &ValueAuto
		}
		return &ValueZero
	}
	return &ValueAuto
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
func YGValueEqual(a Value, b Value) bool {
	if a.Unit != b.Unit {
		return false
	}

	if a.Unit == UnitUndefined {
		return true
	}

	return fabs(a.Value-b.Value) < 0.0001
}

// YGResolveDimensions resolves dimensions
func YGResolveDimensions(node *Node) {
	for dim := DimensionWidth; dim <= DimensionHeight; dim++ {
		if node.style.maxDimensions[dim].Unit != UnitUndefined &&
			YGValueEqual(node.style.maxDimensions[dim], node.style.minDimensions[dim]) {
			node.resolvedDimensions[dim] = &node.style.maxDimensions[dim]
		} else {
			node.resolvedDimensions[dim] = &node.style.dimensions[dim]
		}
	}
}

// YGFloatsEqual returns true if floats are approx. equal
func YGFloatsEqual(a float32, b float32) bool {
	if YGFloatIsUndefined(a) {
		return YGFloatIsUndefined(b)
	}
	return fabs(a-b) < 0.0001
}

// see print.go

var (
	leading  = [4]Edge{EdgeTop, EdgeBottom, EdgeLeft, EdgeRight}
	trailing = [4]Edge{EdgeBottom, EdgeTop, EdgeRight, EdgeLeft}
	pos      = [4]Edge{EdgeTop, EdgeBottom, EdgeLeft, EdgeRight}
	dim      = [4]Dimension{DimensionHeight, DimensionHeight, DimensionWidth, DimensionWidth}
)

func init() {
	leading[FlexDirectionColumn] = EdgeTop
	leading[FlexDirectionColumnReverse] = EdgeBottom
	leading[FlexDirectionRow] = EdgeLeft
	leading[FlexDirectionRowReverse] = EdgeRight

	trailing[FlexDirectionColumn] = EdgeBottom
	trailing[FlexDirectionColumnReverse] = EdgeTop
	trailing[FlexDirectionRow] = EdgeRight
	trailing[FlexDirectionRowReverse] = EdgeLeft

	pos[FlexDirectionColumn] = EdgeTop
	pos[FlexDirectionColumnReverse] = EdgeBottom
	pos[FlexDirectionRow] = EdgeLeft
	pos[FlexDirectionRowReverse] = EdgeRight

	dim[FlexDirectionColumn] = DimensionHeight
	dim[FlexDirectionColumnReverse] = DimensionHeight
	dim[FlexDirectionRow] = DimensionWidth
	dim[FlexDirectionRowReverse] = DimensionWidth
}

// YGFlexDirectionIsRow return true if direction is row
func YGFlexDirectionIsRow(flexDirection FlexDirection) bool {
	return flexDirection == FlexDirectionRow || flexDirection == FlexDirectionRowReverse
}

// YGFlexDirectionIsColumn returns true if direction is column
func YGFlexDirectionIsColumn(flexDirection FlexDirection) bool {
	return flexDirection == FlexDirectionColumn || flexDirection == FlexDirectionColumnReverse
}

// YGNodeLeadingMargin returns leading margin
func YGNodeLeadingMargin(node *Node, axis FlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.margin[EdgeStart].Unit != UnitUndefined {
		return YGResolveValueMargin(&node.style.margin[EdgeStart], widthSize)
	}

	v := YGComputedEdgeValue(node.style.margin[:], leading[axis], &ValueZero)
	return YGResolveValueMargin(v, widthSize)
}

// YGNodeTrailingMargin returns trailing margin
func YGNodeTrailingMargin(node *Node, axis FlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.margin[EdgeEnd].Unit != UnitUndefined {
		return YGResolveValueMargin(&node.style.margin[EdgeEnd], widthSize)
	}

	return YGResolveValueMargin(YGComputedEdgeValue(node.style.margin[:], trailing[axis], &ValueZero),
		widthSize)
}

// YGNodeLeadingPadding returns leading padding
func YGNodeLeadingPadding(node *Node, axis FlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.padding[EdgeStart].Unit != UnitUndefined &&
		YGResolveValue(&node.style.padding[EdgeStart], widthSize) >= 0 {
		return YGResolveValue(&node.style.padding[EdgeStart], widthSize)
	}

	return fmaxf(YGResolveValue(YGComputedEdgeValue(node.style.padding[:], leading[axis], &ValueZero), widthSize), 0)
}

// YGNodeTrailingPadding returns trailing padding
func YGNodeTrailingPadding(node *Node, axis FlexDirection, widthSize float32) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.padding[EdgeEnd].Unit != UnitUndefined &&
		YGResolveValue(&node.style.padding[EdgeEnd], widthSize) >= 0 {
		return YGResolveValue(&node.style.padding[EdgeEnd], widthSize)
	}

	return fmaxf(YGResolveValue(YGComputedEdgeValue(node.style.padding[:], trailing[axis], &ValueZero), widthSize), 0)
}

// YGNodeLeadingBorder returns trailing border
func YGNodeLeadingBorder(node *Node, axis FlexDirection) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.border[EdgeStart].Unit != UnitUndefined &&
		node.style.border[EdgeStart].Value >= 0 {
		return node.style.border[EdgeStart].Value
	}

	return fmaxf(YGComputedEdgeValue(node.style.border[:], leading[axis], &ValueZero).Value, 0)
}

// YGNodeTrailingBorder returns trailing border
func YGNodeTrailingBorder(node *Node, axis FlexDirection) float32 {
	if YGFlexDirectionIsRow(axis) && node.style.border[EdgeEnd].Unit != UnitUndefined &&
		node.style.border[EdgeEnd].Value >= 0 {
		return node.style.border[EdgeEnd].Value
	}

	return fmaxf(YGComputedEdgeValue(node.style.border[:], trailing[axis], &ValueZero).Value, 0)
}

// YGNodeLeadingPaddingAndBorder returns leading padding and border
func YGNodeLeadingPaddingAndBorder(node *Node, axis FlexDirection, widthSize float32) float32 {
	return YGNodeLeadingPadding(node, axis, widthSize) + YGNodeLeadingBorder(node, axis)
}

// YGNodeTrailingPaddingAndBorder returns trailing padding and border
func YGNodeTrailingPaddingAndBorder(node *Node, axis FlexDirection, widthSize float32) float32 {
	return YGNodeTrailingPadding(node, axis, widthSize) + YGNodeTrailingBorder(node, axis)
}

// YGNodeMarginForAxis returns margin for axis
func YGNodeMarginForAxis(node *Node, axis FlexDirection, widthSize float32) float32 {
	leading := YGNodeLeadingMargin(node, axis, widthSize)
	trailing := YGNodeTrailingMargin(node, axis, widthSize)
	return leading + trailing
}

// YGNodePaddingAndBorderForAxis returns padding and border for axis
func YGNodePaddingAndBorderForAxis(node *Node, axis FlexDirection, widthSize float32) float32 {
	return YGNodeLeadingPaddingAndBorder(node, axis, widthSize) +
		YGNodeTrailingPaddingAndBorder(node, axis, widthSize)
}

// YGNodeAlignItem returns align item
func YGNodeAlignItem(node *Node, child *Node) Align {
	align := child.style.alignSelf
	if child.style.alignSelf == AlignAuto {
		align = node.style.alignItems
	}
	if align == AlignBaseline && YGFlexDirectionIsColumn(node.style.flexDirection) {
		return AlignFlexStart
	}
	return align
}

// YGNodeResolveDirection resolves direction
func YGNodeResolveDirection(node *Node, parentDirection Direction) Direction {
	if node.style.direction == DirectionInherit {
		if parentDirection > DirectionInherit {
			return parentDirection
		}
		return DirectionLTR
	}
	return node.style.direction
}

// YGBaseline retuns baseline
func YGBaseline(node *Node) float32 {
	if node.baseline != nil {
		baseline := node.baseline(node, node.layout.measuredDimensions[DimensionWidth], node.layout.measuredDimensions[DimensionHeight])
		YGAssertWithNode(node, !YGFloatIsUndefined(baseline), "Expect custom baseline function to not return NaN")
		return baseline
	}

	var baselineChild *Node
	childCount := YGNodeGetChildCount(node)
	for i := 0; i < childCount; i++ {
		child := YGNodeGetChild(node, i)
		if child.lineIndex > 0 {
			break
		}
		if child.style.positionType == PositionTypeAbsolute {
			continue
		}
		if YGNodeAlignItem(node, child) == AlignBaseline {
			baselineChild = child
			break
		}

		if baselineChild == nil {
			baselineChild = child
		}
	}

	if baselineChild == nil {
		return node.layout.measuredDimensions[DimensionHeight]
	}

	baseline := YGBaseline(baselineChild)
	return baseline + baselineChild.layout.position[EdgeTop]
}

// YGResolveFlexDirection resolves flex direction
func YGResolveFlexDirection(flexDirection FlexDirection, direction Direction) FlexDirection {
	if direction == DirectionRTL {
		if flexDirection == FlexDirectionRow {
			return FlexDirectionRowReverse
		} else if flexDirection == FlexDirectionRowReverse {
			return FlexDirectionRow
		}
	}
	return flexDirection
}

// YGFlexDirectionCross gets flex direction cross
func YGFlexDirectionCross(flexDirection FlexDirection, direction Direction) FlexDirection {
	if YGFlexDirectionIsColumn(flexDirection) {
		return YGResolveFlexDirection(FlexDirectionRow, direction)
	}
	return FlexDirectionColumn
}

// YGNodeIsFlex returns true if node is flex
func YGNodeIsFlex(node *Node) bool {
	return (node.style.positionType == PositionTypeRelative &&
		(YGResolveFlexGrow(node) != 0 || YGNodeResolveFlexShrink(node) != 0))
}

// YGIsBaselineLayout returns true if it's baseline layout
func YGIsBaselineLayout(node *Node) bool {
	if YGFlexDirectionIsColumn(node.style.flexDirection) {
		return false
	}
	if node.style.alignItems == AlignBaseline {
		return true
	}
	childCount := YGNodeGetChildCount(node)
	for i := 0; i < childCount; i++ {
		child := YGNodeGetChild(node, i)
		if child.style.positionType == PositionTypeRelative &&
			child.style.alignSelf == AlignBaseline {
			return true
		}
	}

	return false
}

// YGNodeDimWithMargin returns dimension with margin
func YGNodeDimWithMargin(node *Node, axis FlexDirection, widthSize float32) float32 {
	return node.layout.measuredDimensions[dim[axis]] + YGNodeLeadingMargin(node, axis, widthSize) +
		YGNodeTrailingMargin(node, axis, widthSize)
}

// YGNodeIsStyleDimDefined returns true if dimension is defined
func YGNodeIsStyleDimDefined(node *Node, axis FlexDirection, parentSize float32) bool {
	v := node.resolvedDimensions[dim[axis]]
	isNotDefined := (v.Unit == UnitAuto ||
		v.Unit == UnitUndefined ||
		(v.Unit == UnitPoint && v.Value < 0) ||
		(v.Unit == UnitPercent && (v.Value < 0 || YGFloatIsUndefined(parentSize))))
	return !isNotDefined
}

// YGNodeIsLayoutDimDefined returns true if layout dimension is defined
func YGNodeIsLayoutDimDefined(node *Node, axis FlexDirection) bool {
	value := node.layout.measuredDimensions[dim[axis]]
	return !YGFloatIsUndefined(value) && value >= 0
}

// YGNodeIsLeadingPosDefined returns true if leading position is defined
func YGNodeIsLeadingPosDefined(node *Node, axis FlexDirection) bool {
	return (YGFlexDirectionIsRow(axis) &&
		YGComputedEdgeValue(node.style.position[:], EdgeStart, &ValueUndefined).Unit !=
			UnitUndefined) ||
		YGComputedEdgeValue(node.style.position[:], leading[axis], &ValueUndefined).Unit !=
			UnitUndefined
}

// YGNodeIsTrailingPosDefined returns true if trailing pos is defined
func YGNodeIsTrailingPosDefined(node *Node, axis FlexDirection) bool {
	return (YGFlexDirectionIsRow(axis) &&
		YGComputedEdgeValue(node.style.position[:], EdgeEnd, &ValueUndefined).Unit !=
			UnitUndefined) ||
		YGComputedEdgeValue(node.style.position[:], trailing[axis], &ValueUndefined).Unit !=
			UnitUndefined
}

// YGNodeLeadingPosition returns leading position
func YGNodeLeadingPosition(node *Node, axis FlexDirection, axisSize float32) float32 {
	if YGFlexDirectionIsRow(axis) {
		leadingPosition := YGComputedEdgeValue(node.style.position[:], EdgeStart, &ValueUndefined)
		if leadingPosition.Unit != UnitUndefined {
			return YGResolveValue(leadingPosition, axisSize)
		}
	}

	leadingPosition := YGComputedEdgeValue(node.style.position[:], leading[axis], &ValueUndefined)

	if leadingPosition.Unit == UnitUndefined {
		return 0
	}
	return YGResolveValue(leadingPosition, axisSize)
}

// YGNodeTrailingPosition returns trailing position
func YGNodeTrailingPosition(node *Node, axis FlexDirection, axisSize float32) float32 {
	if YGFlexDirectionIsRow(axis) {
		trailingPosition := YGComputedEdgeValue(node.style.position[:], EdgeEnd, &ValueUndefined)
		if trailingPosition.Unit != UnitUndefined {
			return YGResolveValue(trailingPosition, axisSize)
		}
	}

	trailingPosition := YGComputedEdgeValue(node.style.position[:], trailing[axis], &ValueUndefined)

	if trailingPosition.Unit == UnitUndefined {
		return 0
	}
	return YGResolveValue(trailingPosition, axisSize)
}

// YGNodeBoundAxisWithinMinAndMax returns bounds axis
func YGNodeBoundAxisWithinMinAndMax(node *Node, axis FlexDirection, value float32, axisSize float32) float32 {
	min := Undefined
	max := Undefined

	if YGFlexDirectionIsColumn(axis) {
		min = YGResolveValue(&node.style.minDimensions[DimensionHeight], axisSize)
		max = YGResolveValue(&node.style.maxDimensions[DimensionHeight], axisSize)
	} else if YGFlexDirectionIsRow(axis) {
		min = YGResolveValue(&node.style.minDimensions[DimensionWidth], axisSize)
		max = YGResolveValue(&node.style.maxDimensions[DimensionWidth], axisSize)
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

// YGMarginLeadingValue returns margin leading value
func YGMarginLeadingValue(node *Node, axis FlexDirection) *Value {
	if YGFlexDirectionIsRow(axis) && node.style.margin[EdgeStart].Unit != UnitUndefined {
		return &node.style.margin[EdgeStart]
	}
	return &node.style.margin[leading[axis]]
}

// YGMarginTrailingValue returns trailing value
func YGMarginTrailingValue(node *Node, axis FlexDirection) *Value {
	if YGFlexDirectionIsRow(axis) && node.style.margin[EdgeEnd].Unit != UnitUndefined {
		return &node.style.margin[EdgeEnd]
	}
	return &node.style.margin[trailing[axis]]

}

// YGNodeBoundAxis is like YGNodeBoundAxisWithinMinAndMax but also ensures that
// the value doesn't go below the padding and border amount.
func YGNodeBoundAxis(node *Node, axis FlexDirection, value float32, axisSize float32, widthSize float32) float32 {
	return fmaxf(YGNodeBoundAxisWithinMinAndMax(node, axis, value, axisSize),
		YGNodePaddingAndBorderForAxis(node, axis, widthSize))
}

// YGNodeSetChildTrailingPosition sets child's trailing position
func YGNodeSetChildTrailingPosition(node *Node, child *Node, axis FlexDirection) {
	size := child.layout.measuredDimensions[dim[axis]]
	child.layout.position[trailing[axis]] =
		node.layout.measuredDimensions[dim[axis]] - size - child.layout.position[pos[axis]]
}

// YGNodeRelativePosition gets relative position.
// If both left and right are defined, then use left. Otherwise return
// +left or -right depending on which is defined.
func YGNodeRelativePosition(node *Node, axis FlexDirection, axisSize float32) float32 {
	if YGNodeIsLeadingPosDefined(node, axis) {
		return YGNodeLeadingPosition(node, axis, axisSize)
	}
	return -YGNodeTrailingPosition(node, axis, axisSize)
}

// YGConstrainMaxSizeForMode contrains max size for node
func YGConstrainMaxSizeForMode(node *Node, axis FlexDirection, parentAxisSize float32, parentWidth float32, mode *MeasureMode, size *float32) {
	maxSize := YGResolveValue(&node.style.maxDimensions[dim[axis]], parentAxisSize) +
		YGNodeMarginForAxis(node, axis, parentWidth)
	switch *mode {
	case MeasureModeExactly, MeasureModeAtMost:
		if YGFloatIsUndefined(maxSize) || *size < maxSize {
			// TODO: this is redundant, but what is in original code
			//*size = *size
		} else {
			*size = maxSize
		}

		break
	case MeasureModeUndefined:
		if !YGFloatIsUndefined(maxSize) {
			*mode = MeasureModeAtMost
			*size = maxSize
		}
		break
	}
}

// YGNodeSetPosition sets position
func YGNodeSetPosition(node *Node, direction Direction, mainSize float32, crossSize float32, parentWidth float32) {
	/* Root nodes should be always layouted as LTR, so we don't return negative values. */
	directionRespectingRoot := DirectionLTR
	if node.parent != nil {
		directionRespectingRoot = direction
	}

	mainAxis := YGResolveFlexDirection(node.style.flexDirection, directionRespectingRoot)
	crossAxis := YGFlexDirectionCross(mainAxis, directionRespectingRoot)

	relativePositionMain := YGNodeRelativePosition(node, mainAxis, mainSize)
	relativePositionCross := YGNodeRelativePosition(node, crossAxis, crossSize)

	pos := &node.layout.position
	pos[leading[mainAxis]] = YGNodeLeadingMargin(node, mainAxis, parentWidth) + relativePositionMain
	pos[trailing[mainAxis]] = YGNodeTrailingMargin(node, mainAxis, parentWidth) + relativePositionMain
	pos[leading[crossAxis]] = YGNodeLeadingMargin(node, crossAxis, parentWidth) + relativePositionCross
	pos[trailing[crossAxis]] = YGNodeTrailingMargin(node, crossAxis, parentWidth) + relativePositionCross
}

// YGNodeComputeFlexBasisForChild computes flex basis for child
func YGNodeComputeFlexBasisForChild(node *Node,
	child *Node,
	width float32,
	widthMode MeasureMode,
	height float32,
	parentWidth float32,
	parentHeight float32,
	heightMode MeasureMode,
	direction Direction,
	config *Config) {
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
	var childWidthMeasureMode MeasureMode
	var childHeightMeasureMode MeasureMode

	resolvedFlexBasis := YGResolveValue(YGNodeResolveFlexBasisPtr(child), mainAxisParentSize)
	isRowStyleDimDefined := YGNodeIsStyleDimDefined(child, FlexDirectionRow, parentWidth)
	isColumnStyleDimDefined := YGNodeIsStyleDimDefined(child, FlexDirectionColumn, parentHeight)

	if !YGFloatIsUndefined(resolvedFlexBasis) && !YGFloatIsUndefined(mainAxisSize) {
		if YGFloatIsUndefined(child.layout.computedFlexBasis) ||
			(YGConfigIsExperimentalFeatureEnabled(child.config, ExperimentalFeatureWebFlexBasis) &&
				child.layout.computedFlexBasisGeneration != gCurrentGenerationCount) {
			child.layout.computedFlexBasis =
				fmaxf(resolvedFlexBasis, YGNodePaddingAndBorderForAxis(child, mainAxis, parentWidth))
		}
	} else if isMainAxisRow && isRowStyleDimDefined {
		// The width is definite, so use that as the flex basis.
		child.layout.computedFlexBasis =
			fmaxf(YGResolveValue(child.resolvedDimensions[DimensionWidth], parentWidth),
				YGNodePaddingAndBorderForAxis(child, FlexDirectionRow, parentWidth))
	} else if !isMainAxisRow && isColumnStyleDimDefined {
		// The height is definite, so use that as the flex basis.
		child.layout.computedFlexBasis =
			fmaxf(YGResolveValue(child.resolvedDimensions[DimensionHeight], parentHeight),
				YGNodePaddingAndBorderForAxis(child, FlexDirectionColumn, parentWidth))
	} else {
		// Compute the flex basis and hypothetical main size (i.e. the clamped
		// flex basis).
		childWidth = Undefined
		childHeight = Undefined
		childWidthMeasureMode = MeasureModeUndefined
		childHeightMeasureMode = MeasureModeUndefined

		marginRow := YGNodeMarginForAxis(child, FlexDirectionRow, parentWidth)
		marginColumn := YGNodeMarginForAxis(child, FlexDirectionColumn, parentWidth)

		if isRowStyleDimDefined {
			childWidth =
				YGResolveValue(child.resolvedDimensions[DimensionWidth], parentWidth) + marginRow
			childWidthMeasureMode = MeasureModeExactly
		}
		if isColumnStyleDimDefined {
			childHeight =
				YGResolveValue(child.resolvedDimensions[DimensionHeight], parentHeight) + marginColumn
			childHeightMeasureMode = MeasureModeExactly
		}

		// The W3C spec doesn't say anything about the 'overflow' property,
		// but all major browsers appear to implement the following logic.
		if (!isMainAxisRow && node.style.overflow == OverflowScroll) ||
			node.style.overflow != OverflowScroll {
			if YGFloatIsUndefined(childWidth) && !YGFloatIsUndefined(width) {
				childWidth = width
				childWidthMeasureMode = MeasureModeAtMost
			}
		}

		if (isMainAxisRow && node.style.overflow == OverflowScroll) ||
			node.style.overflow != OverflowScroll {
			if YGFloatIsUndefined(childHeight) && !YGFloatIsUndefined(height) {
				childHeight = height
				childHeightMeasureMode = MeasureModeAtMost
			}
		}

		// If child has no defined size in the cross axis and is set to stretch,
		// set the cross
		// axis to be measured exactly with the available inner width
		if !isMainAxisRow && !YGFloatIsUndefined(width) && !isRowStyleDimDefined &&
			widthMode == MeasureModeExactly && YGNodeAlignItem(node, child) == AlignStretch {
			childWidth = width
			childWidthMeasureMode = MeasureModeExactly
		}
		if isMainAxisRow && !YGFloatIsUndefined(height) && !isColumnStyleDimDefined &&
			heightMode == MeasureModeExactly && YGNodeAlignItem(node, child) == AlignStretch {
			childHeight = height
			childHeightMeasureMode = MeasureModeExactly
		}

		if !YGFloatIsUndefined(child.style.aspectRatio) {
			if !isMainAxisRow && childWidthMeasureMode == MeasureModeExactly {
				child.layout.computedFlexBasis =
					fmaxf((childWidth-marginRow)/child.style.aspectRatio,
						YGNodePaddingAndBorderForAxis(child, FlexDirectionColumn, parentWidth))
				return
			} else if isMainAxisRow && childHeightMeasureMode == MeasureModeExactly {
				child.layout.computedFlexBasis =
					fmaxf((childHeight-marginColumn)*child.style.aspectRatio,
						YGNodePaddingAndBorderForAxis(child, FlexDirectionRow, parentWidth))
				return
			}
		}

		YGConstrainMaxSizeForMode(
			child, FlexDirectionRow, parentWidth, parentWidth, &childWidthMeasureMode, &childWidth)
		YGConstrainMaxSizeForMode(child,
			FlexDirectionColumn,
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

// YGNodeAbsoluteLayoutChild calculates absolute child layout
func YGNodeAbsoluteLayoutChild(node *Node, child *Node, width float32, widthMode MeasureMode, height float32, direction Direction, config *Config) {
	mainAxis := YGResolveFlexDirection(node.style.flexDirection, direction)
	crossAxis := YGFlexDirectionCross(mainAxis, direction)
	isMainAxisRow := YGFlexDirectionIsRow(mainAxis)

	childWidth := Undefined
	childHeight := Undefined
	childWidthMeasureMode := MeasureModeUndefined
	childHeightMeasureMode := MeasureModeUndefined

	marginRow := YGNodeMarginForAxis(child, FlexDirectionRow, width)
	marginColumn := YGNodeMarginForAxis(child, FlexDirectionColumn, width)

	if YGNodeIsStyleDimDefined(child, FlexDirectionRow, width) {
		childWidth = YGResolveValue(child.resolvedDimensions[DimensionWidth], width) + marginRow
	} else {
		// If the child doesn't have a specified width, compute the width based
		// on the left/right
		// offsets if they're defined.
		if YGNodeIsLeadingPosDefined(child, FlexDirectionRow) &&
			YGNodeIsTrailingPosDefined(child, FlexDirectionRow) {
			childWidth = node.layout.measuredDimensions[DimensionWidth] -
				(YGNodeLeadingBorder(node, FlexDirectionRow) +
					YGNodeTrailingBorder(node, FlexDirectionRow)) -
				(YGNodeLeadingPosition(child, FlexDirectionRow, width) +
					YGNodeTrailingPosition(child, FlexDirectionRow, width))
			childWidth = YGNodeBoundAxis(child, FlexDirectionRow, childWidth, width, width)
		}
	}

	if YGNodeIsStyleDimDefined(child, FlexDirectionColumn, height) {
		childHeight =
			YGResolveValue(child.resolvedDimensions[DimensionHeight], height) + marginColumn
	} else {
		// If the child doesn't have a specified height, compute the height
		// based on the top/bottom
		// offsets if they're defined.
		if YGNodeIsLeadingPosDefined(child, FlexDirectionColumn) &&
			YGNodeIsTrailingPosDefined(child, FlexDirectionColumn) {
			childHeight = node.layout.measuredDimensions[DimensionHeight] -
				(YGNodeLeadingBorder(node, FlexDirectionColumn) +
					YGNodeTrailingBorder(node, FlexDirectionColumn)) -
				(YGNodeLeadingPosition(child, FlexDirectionColumn, height) +
					YGNodeTrailingPosition(child, FlexDirectionColumn, height))
			childHeight = YGNodeBoundAxis(child, FlexDirectionColumn, childHeight, height, width)
		}
	}

	// Exactly one dimension needs to be defined for us to be able to do aspect ratio
	// calculation. One dimension being the anchor and the other being flexible.
	if YGFloatIsUndefined(childWidth) != YGFloatIsUndefined(childHeight) {
		if !YGFloatIsUndefined(child.style.aspectRatio) {
			if YGFloatIsUndefined(childWidth) {
				childWidth =
					marginRow + fmaxf((childHeight-marginColumn)*child.style.aspectRatio,
						YGNodePaddingAndBorderForAxis(child, FlexDirectionColumn, width))
			} else if YGFloatIsUndefined(childHeight) {
				childHeight =
					marginColumn + fmaxf((childWidth-marginRow)/child.style.aspectRatio,
						YGNodePaddingAndBorderForAxis(child, FlexDirectionRow, width))
			}
		}
	}

	// If we're still missing one or the other dimension, measure the content.
	if YGFloatIsUndefined(childWidth) || YGFloatIsUndefined(childHeight) {
		childWidthMeasureMode = MeasureModeExactly
		if YGFloatIsUndefined(childWidth) {
			childWidthMeasureMode = MeasureModeUndefined
		}
		childHeightMeasureMode = MeasureModeExactly
		if YGFloatIsUndefined(childHeight) {
			childHeightMeasureMode = MeasureModeUndefined
		}

		// If the size of the parent is defined then try to rain the absolute child to that size
		// as well. This allows text within the absolute child to wrap to the size of its parent.
		// This is the same behavior as many browsers implement.
		if !isMainAxisRow && YGFloatIsUndefined(childWidth) && widthMode != MeasureModeUndefined &&
			width > 0 {
			childWidth = width
			childWidthMeasureMode = MeasureModeAtMost
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
		childWidth = child.layout.measuredDimensions[DimensionWidth] +
			YGNodeMarginForAxis(child, FlexDirectionRow, width)
		childHeight = child.layout.measuredDimensions[DimensionHeight] +
			YGNodeMarginForAxis(child, FlexDirectionColumn, width)
	}

	YGLayoutNodeInternal(child,
		childWidth,
		childHeight,
		direction,
		MeasureModeExactly,
		MeasureModeExactly,
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
		node.style.justifyContent == JustifyCenter {
		child.layout.position[leading[mainAxis]] = (node.layout.measuredDimensions[dim[mainAxis]] -
			child.layout.measuredDimensions[dim[mainAxis]]) /
			2.0
	} else if !YGNodeIsLeadingPosDefined(child, mainAxis) &&
		node.style.justifyContent == JustifyFlexEnd {
		child.layout.position[leading[mainAxis]] = (node.layout.measuredDimensions[dim[mainAxis]] -
			child.layout.measuredDimensions[dim[mainAxis]])
	}

	if YGNodeIsTrailingPosDefined(child, crossAxis) &&
		!YGNodeIsLeadingPosDefined(child, crossAxis) {
		axisSize := width
		if isMainAxisRow {
			axisSize = height
		}

		child.layout.position[leading[crossAxis]] = node.layout.measuredDimensions[dim[crossAxis]] -
			child.layout.measuredDimensions[dim[crossAxis]] -
			YGNodeTrailingBorder(node, crossAxis) -
			YGNodeTrailingMargin(child, crossAxis, width) -
			YGNodeTrailingPosition(child, crossAxis, axisSize)
	} else if !YGNodeIsLeadingPosDefined(child, crossAxis) &&
		YGNodeAlignItem(node, child) == AlignCenter {
		child.layout.position[leading[crossAxis]] =
			(node.layout.measuredDimensions[dim[crossAxis]] -
				child.layout.measuredDimensions[dim[crossAxis]]) /
				2.0
	} else if !YGNodeIsLeadingPosDefined(child, crossAxis) &&
		((YGNodeAlignItem(node, child) == AlignFlexEnd) != (node.style.flexWrap == WrapWrapReverse)) {
		child.layout.position[leading[crossAxis]] = (node.layout.measuredDimensions[dim[crossAxis]] -
			child.layout.measuredDimensions[dim[crossAxis]])
	}
}

// YGNodeWithMeasureFuncSetMeasuredDimensions sets measure dimensions for node with measure func
func YGNodeWithMeasureFuncSetMeasuredDimensions(node *Node, availableWidth float32, availableHeight float32, widthMeasureMode MeasureMode, heightMeasureMode MeasureMode, parentWidth float32, parentHeight float32) {
	YGAssertWithNode(node, node.measure != nil, "Expected node to have custom measure function")

	paddingAndBorderAxisRow := YGNodePaddingAndBorderForAxis(node, FlexDirectionRow, availableWidth)
	paddingAndBorderAxisColumn := YGNodePaddingAndBorderForAxis(node, FlexDirectionColumn, availableWidth)
	marginAxisRow := YGNodeMarginForAxis(node, FlexDirectionRow, availableWidth)
	marginAxisColumn := YGNodeMarginForAxis(node, FlexDirectionColumn, availableWidth)

	// We want to make sure we don't call measure with negative size
	innerWidth := fmaxf(0, availableWidth-marginAxisRow-paddingAndBorderAxisRow)
	if YGFloatIsUndefined(availableWidth) {
		innerWidth = availableWidth
	}
	innerHeight := fmaxf(0, availableHeight-marginAxisColumn-paddingAndBorderAxisColumn)
	if YGFloatIsUndefined(availableHeight) {
		innerHeight = availableHeight
	}

	if widthMeasureMode == MeasureModeExactly && heightMeasureMode == MeasureModeExactly {
		// Don't bother sizing the text if both dimensions are already defined.
		node.layout.measuredDimensions[DimensionWidth] = YGNodeBoundAxis(
			node, FlexDirectionRow, availableWidth-marginAxisRow, parentWidth, parentWidth)
		node.layout.measuredDimensions[DimensionHeight] = YGNodeBoundAxis(
			node, FlexDirectionColumn, availableHeight-marginAxisColumn, parentHeight, parentWidth)
	} else {
		// Measure the text under the current raints.
		measuredSize := node.measure(node, innerWidth, widthMeasureMode, innerHeight, heightMeasureMode)

		width := availableWidth - marginAxisRow
		if widthMeasureMode == MeasureModeUndefined ||
			widthMeasureMode == MeasureModeAtMost {
			width = measuredSize.Width + paddingAndBorderAxisRow

		}

		node.layout.measuredDimensions[DimensionWidth] = YGNodeBoundAxis(node, FlexDirectionRow, width, availableWidth, availableWidth)

		height := availableHeight - marginAxisColumn
		if heightMeasureMode == MeasureModeUndefined ||
			heightMeasureMode == MeasureModeAtMost {
			height = measuredSize.Height + paddingAndBorderAxisColumn
		}

		node.layout.measuredDimensions[DimensionHeight] = YGNodeBoundAxis(node, FlexDirectionColumn, height, availableHeight, availableWidth)
	}
}

// YGNodeEmptyContainerSetMeasuredDimensions sets measure dimensions for empty container
// For nodes with no children, use the available values if they were provided,
// or the minimum size as indicated by the padding and border sizes.
func YGNodeEmptyContainerSetMeasuredDimensions(node *Node, availableWidth float32, availableHeight float32, widthMeasureMode MeasureMode, heightMeasureMode MeasureMode, parentWidth float32, parentHeight float32) {
	paddingAndBorderAxisRow := YGNodePaddingAndBorderForAxis(node, FlexDirectionRow, parentWidth)
	paddingAndBorderAxisColumn := YGNodePaddingAndBorderForAxis(node, FlexDirectionColumn, parentWidth)
	marginAxisRow := YGNodeMarginForAxis(node, FlexDirectionRow, parentWidth)
	marginAxisColumn := YGNodeMarginForAxis(node, FlexDirectionColumn, parentWidth)

	width := availableWidth - marginAxisRow
	if widthMeasureMode == MeasureModeUndefined || widthMeasureMode == MeasureModeAtMost {
		width = paddingAndBorderAxisRow
	}
	node.layout.measuredDimensions[DimensionWidth] = YGNodeBoundAxis(node, FlexDirectionRow, width, parentWidth, parentWidth)

	height := availableHeight - marginAxisColumn
	if heightMeasureMode == MeasureModeUndefined || heightMeasureMode == MeasureModeAtMost {
		height = paddingAndBorderAxisColumn
	}
	node.layout.measuredDimensions[DimensionHeight] = YGNodeBoundAxis(node, FlexDirectionColumn, height, parentHeight, parentWidth)
}

// YGNodeFixedSizeSetMeasuredDimensions sets fixed size measure dimensions
func YGNodeFixedSizeSetMeasuredDimensions(node *Node,
	availableWidth float32,
	availableHeight float32,
	widthMeasureMode MeasureMode,
	heightMeasureMode MeasureMode,
	parentWidth float32,
	parentHeight float32) bool {
	if (widthMeasureMode == MeasureModeAtMost && availableWidth <= 0) ||
		(heightMeasureMode == MeasureModeAtMost && availableHeight <= 0) ||
		(widthMeasureMode == MeasureModeExactly && heightMeasureMode == MeasureModeExactly) {
		marginAxisColumn := YGNodeMarginForAxis(node, FlexDirectionColumn, parentWidth)
		marginAxisRow := YGNodeMarginForAxis(node, FlexDirectionRow, parentWidth)

		width := availableWidth - marginAxisRow
		if YGFloatIsUndefined(availableWidth) || (widthMeasureMode == MeasureModeAtMost && availableWidth < 0) {
			width = 0
		}
		node.layout.measuredDimensions[DimensionWidth] =
			YGNodeBoundAxis(node, FlexDirectionRow, width, parentWidth, parentWidth)

		height := availableHeight - marginAxisColumn
		if YGFloatIsUndefined(availableHeight) || (heightMeasureMode == MeasureModeAtMost && availableHeight < 0) {
			height = 0
		}
		node.layout.measuredDimensions[DimensionHeight] =
			YGNodeBoundAxis(node, FlexDirectionColumn, height, parentHeight, parentWidth)

		return true
	}

	return false
}

// YGZeroOutLayoutRecursivly zeros out layout recursively
func YGZeroOutLayoutRecursivly(node *Node) {
	node.layout.dimensions[DimensionHeight] = 0
	node.layout.dimensions[DimensionWidth] = 0
	node.layout.position[EdgeTop] = 0
	node.layout.position[EdgeBottom] = 0
	node.layout.position[EdgeLeft] = 0
	node.layout.position[EdgeRight] = 0
	node.layout.cachedLayout.availableHeight = 0
	node.layout.cachedLayout.availableWidth = 0
	node.layout.cachedLayout.heightMeasureMode = MeasureModeExactly
	node.layout.cachedLayout.widthMeasureMode = MeasureModeExactly
	node.layout.cachedLayout.computedWidth = 0
	node.layout.cachedLayout.computedHeight = 0
	node.hasNewLayout = true
	childCount := YGNodeGetChildCount(node)
	for i := 0; i < childCount; i++ {
		child := YGNodeListGet(node.children, i)
		YGZeroOutLayoutRecursivly(child)
	}
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
//      or Undefined if the size is not available; interpretation depends on
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
func YGNodelayoutImpl(node *Node, availableWidth float32, availableHeight float32,
	parentDirection Direction, widthMeasureMode MeasureMode,
	heightMeasureMode MeasureMode, parentWidth float32, parentHeight float32,
	performLayout bool, config *Config) {
	// YGAssertWithNode(node, YGFloatIsUndefined(availableWidth) ? widthMeasureMode == YGMeasureModeUndefined : true, "availableWidth is indefinite so widthMeasureMode must be YGMeasureModeUndefined");
	//YGAssertWithNode(node, YGFloatIsUndefined(availableHeight) ? heightMeasureMode == YGMeasureModeUndefined : true, "availableHeight is indefinite so heightMeasureMode must be YGMeasureModeUndefined");

	// Set the resolved resolution in the node's layout.
	direction := YGNodeResolveDirection(node, parentDirection)
	node.layout.direction = direction

	flexRowDirection := YGResolveFlexDirection(FlexDirectionRow, direction)
	flexColumnDirection := YGResolveFlexDirection(FlexDirectionColumn, direction)

	node.layout.margin[EdgeStart] = YGNodeLeadingMargin(node, flexRowDirection, parentWidth)
	node.layout.margin[EdgeEnd] = YGNodeTrailingMargin(node, flexRowDirection, parentWidth)
	node.layout.margin[EdgeTop] = YGNodeLeadingMargin(node, flexColumnDirection, parentWidth)
	node.layout.margin[EdgeBottom] = YGNodeTrailingMargin(node, flexColumnDirection, parentWidth)

	node.layout.border[EdgeStart] = YGNodeLeadingBorder(node, flexRowDirection)
	node.layout.border[EdgeEnd] = YGNodeTrailingBorder(node, flexRowDirection)
	node.layout.border[EdgeTop] = YGNodeLeadingBorder(node, flexColumnDirection)
	node.layout.border[EdgeBottom] = YGNodeTrailingBorder(node, flexColumnDirection)

	node.layout.padding[EdgeStart] = YGNodeLeadingPadding(node, flexRowDirection, parentWidth)
	node.layout.padding[EdgeEnd] = YGNodeTrailingPadding(node, flexRowDirection, parentWidth)
	node.layout.padding[EdgeTop] = YGNodeLeadingPadding(node, flexColumnDirection, parentWidth)
	node.layout.padding[EdgeBottom] = YGNodeTrailingPadding(node, flexColumnDirection, parentWidth)

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
	isNodeFlexWrap := node.style.flexWrap != WrapNoWrap

	mainAxisParentSize := parentHeight
	crossAxisParentSize := parentWidth
	if isMainAxisRow {
		mainAxisParentSize = parentWidth
		crossAxisParentSize = parentHeight
	}

	var firstAbsoluteChild *Node
	var currentAbsoluteChild *Node

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
	paddingAndBorderAxisColumn := paddingAndBorderAxisMain
	if isMainAxisRow {
		paddingAndBorderAxisRow = paddingAndBorderAxisMain
		paddingAndBorderAxisColumn = paddingAndBorderAxisCross
	}

	marginAxisRow := YGNodeMarginForAxis(node, FlexDirectionRow, parentWidth)
	marginAxisColumn := YGNodeMarginForAxis(node, FlexDirectionColumn, parentWidth)

	// STEP 2: DETERMINE AVAILABLE SIZE IN MAIN AND CROSS DIRECTIONS
	minInnerWidth := YGResolveValue(&node.style.minDimensions[DimensionWidth], parentWidth) - marginAxisRow -
		paddingAndBorderAxisRow
	maxInnerWidth := YGResolveValue(&node.style.maxDimensions[DimensionWidth], parentWidth) - marginAxisRow -
		paddingAndBorderAxisRow
	minInnerHeight := YGResolveValue(&node.style.minDimensions[DimensionHeight], parentHeight) -
		marginAxisColumn - paddingAndBorderAxisColumn
	maxInnerHeight := YGResolveValue(&node.style.maxDimensions[DimensionHeight], parentHeight) -
		marginAxisColumn - paddingAndBorderAxisColumn

	minInnerMainDim := minInnerHeight
	maxInnerMainDim := maxInnerHeight
	if isMainAxisRow {
		minInnerMainDim = minInnerWidth
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
	availableInnerCrossDim := availableInnerWidth
	if isMainAxisRow {
		availableInnerMainDim = availableInnerWidth
		availableInnerCrossDim = availableInnerHeight
	}

	// If there is only one child with flexGrow + flexShrink it means we can set the
	// computedFlexBasis to 0 instead of measuring and shrinking / flexing the child to exactly
	// match the remaining space
	var singleFlexChild *Node
	if measureModeMainDim == MeasureModeExactly {
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
		if child.style.display == DisplayNone {
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
		if child.style.positionType == PositionTypeAbsolute {
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
	if measureModeMainDim == MeasureModeUndefined {
		flexBasisOverflows = false
	}
	if isNodeFlexWrap && flexBasisOverflows && measureModeMainDim == MeasureModeAtMost {
		measureModeMainDim = MeasureModeExactly
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
		var firstRelativeChild *Node
		var currentRelativeChild *Node

		// Add items to the current line until it's full or we run out of items.
		for i := startOfLineIndex; i < childCount; i++ {
			child := YGNodeListGet(node.children, i)
			if child.style.display == DisplayNone {
				endOfLineIndex++
				continue
			}
			child.lineIndex = lineCount

			if child.style.positionType != PositionTypeAbsolute {
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
		canSkipFlex := !performLayout && measureModeCrossDim == MeasureModeExactly

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
		if measureModeMainDim != MeasureModeExactly {
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
				var childCrossMeasureMode MeasureMode
				childMainMeasureMode := MeasureModeExactly

				if !YGFloatIsUndefined(availableInnerCrossDim) &&
					!YGNodeIsStyleDimDefined(currentRelativeChild, crossAxis, availableInnerCrossDim) &&
					measureModeCrossDim == MeasureModeExactly &&
					!(isNodeFlexWrap && flexBasisOverflows) &&
					YGNodeAlignItem(node, currentRelativeChild) == AlignStretch {
					childCrossSize = availableInnerCrossDim
					childCrossMeasureMode = MeasureModeExactly
				} else if !YGNodeIsStyleDimDefined(currentRelativeChild,
					crossAxis,
					availableInnerCrossDim) {
					childCrossSize = availableInnerCrossDim
					childCrossMeasureMode = MeasureModeAtMost
					if YGFloatIsUndefined(childCrossSize) {
						childCrossMeasureMode = MeasureModeUndefined
					}
				} else {
					childCrossSize = YGResolveValue(currentRelativeChild.resolvedDimensions[dim[crossAxis]],
						availableInnerCrossDim) +
						marginCross
					isLoosePercentageMeasurement := currentRelativeChild.resolvedDimensions[dim[crossAxis]].Unit == UnitPercent &&
						measureModeCrossDim != MeasureModeExactly
					childCrossMeasureMode = MeasureModeExactly
					if YGFloatIsUndefined(childCrossSize) || isLoosePercentageMeasurement {
						childCrossMeasureMode = MeasureModeUndefined
					}
				}

				if !YGFloatIsUndefined(currentRelativeChild.style.aspectRatio) {
					v := (childMainSize - marginMain) * currentRelativeChild.style.aspectRatio
					if isMainAxisRow {
						v = (childMainSize - marginMain) / currentRelativeChild.style.aspectRatio
					}
					childCrossSize = fmaxf(v, YGNodePaddingAndBorderForAxis(currentRelativeChild, crossAxis, availableInnerWidth))
					childCrossMeasureMode = MeasureModeExactly

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
					YGNodeAlignItem(node, currentRelativeChild) == AlignStretch

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

		if measureModeMainDim == MeasureModeAtMost && remainingFreeSpace > 0 {
			if node.style.minDimensions[dim[mainAxis]].Unit != UnitUndefined &&
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
			if child.style.positionType == PositionTypeRelative {
				if YGMarginLeadingValue(child, mainAxis).Unit == UnitAuto {
					numberOfAutoMarginsOnCurrentLine++
				}
				if YGMarginTrailingValue(child, mainAxis).Unit == UnitAuto {
					numberOfAutoMarginsOnCurrentLine++
				}
			}
		}

		if numberOfAutoMarginsOnCurrentLine == 0 {
			switch justifyContent {
			case JustifyCenter:
				leadingMainDim = remainingFreeSpace / 2
			case JustifyFlexEnd:
				leadingMainDim = remainingFreeSpace
			case JustifySpaceBetween:
				if itemsOnLine > 1 {
					betweenMainDim = fmaxf(remainingFreeSpace, 0) / float32(itemsOnLine-1)
				} else {
					betweenMainDim = 0
				}
			case JustifySpaceAround:
				// Space on the edges is half of the space between elements
				betweenMainDim = remainingFreeSpace / float32(itemsOnLine)
				leadingMainDim = betweenMainDim / 2
			case JustifyFlexStart:
			}
		}

		mainDim := leadingPaddingAndBorderMain + leadingMainDim
		var crossDim float32

		for i := startOfLineIndex; i < endOfLineIndex; i++ {
			child := YGNodeListGet(node.children, i)
			if child.style.display == DisplayNone {
				continue
			}
			if child.style.positionType == PositionTypeAbsolute &&
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
				if child.style.positionType == PositionTypeRelative {
					if YGMarginLeadingValue(child, mainAxis).Unit == UnitAuto {
						mainDim += remainingFreeSpace / float32(numberOfAutoMarginsOnCurrentLine)
					}

					if performLayout {
						child.layout.position[pos[mainAxis]] += mainDim
					}

					if YGMarginTrailingValue(child, mainAxis).Unit == UnitAuto {
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
		if measureModeCrossDim == MeasureModeUndefined ||
			measureModeCrossDim == MeasureModeAtMost {
			// Compute the cross axis from the max cross dimension of the children.
			containerCrossAxis = YGNodeBoundAxis(node,
				crossAxis,
				crossDim+paddingAndBorderAxisCross,
				crossAxisParentSize,
				parentWidth) -
				paddingAndBorderAxisCross
		}

		// If there's no flex wrap, the cross dimension is defined by the container.
		if !isNodeFlexWrap && measureModeCrossDim == MeasureModeExactly {
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
				if child.style.display == DisplayNone {
					continue
				}
				if child.style.positionType == PositionTypeAbsolute {
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
					if alignItem == AlignStretch &&
						YGMarginLeadingValue(child, crossAxis).Unit != UnitAuto &&
						YGMarginTrailingValue(child, crossAxis).Unit != UnitAuto {
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

							childMainMeasureMode := MeasureModeExactly
							childCrossMeasureMode := MeasureModeExactly
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

							childWidthMeasureMode := MeasureModeExactly
							if YGFloatIsUndefined(childWidth) {
								childWidthMeasureMode = MeasureModeUndefined
							}

							childHeightMeasureMode := MeasureModeExactly
							if YGFloatIsUndefined(childHeight) {
								childHeightMeasureMode = MeasureModeUndefined
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

						if YGMarginLeadingValue(child, crossAxis).Unit == UnitAuto &&
							YGMarginTrailingValue(child, crossAxis).Unit == UnitAuto {
							leadingCrossDim += fmaxf(0, remainingCrossDim/2)
						} else if YGMarginTrailingValue(child, crossAxis).Unit == UnitAuto {
							// No-Op
						} else if YGMarginLeadingValue(child, crossAxis).Unit == UnitAuto {
							leadingCrossDim += fmaxf(0, remainingCrossDim)
						} else if alignItem == AlignFlexStart {
							// No-Op
						} else if alignItem == AlignCenter {
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
		case AlignFlexEnd:
			currentLead += remainingAlignContentDim
		case AlignCenter:
			currentLead += remainingAlignContentDim / 2
		case AlignStretch:
			if availableInnerCrossDim > totalLineCrossDim {
				crossDimLead = remainingAlignContentDim / float32(lineCount)
			}
		case AlignSpaceAround:
			if availableInnerCrossDim > totalLineCrossDim {
				currentLead += remainingAlignContentDim / float32(2*lineCount)
				if lineCount > 1 {
					crossDimLead = remainingAlignContentDim / float32(lineCount)
				}
			} else {
				currentLead += remainingAlignContentDim / 2
			}
		case AlignSpaceBetween:
			if availableInnerCrossDim > totalLineCrossDim && lineCount > 1 {
				crossDimLead = remainingAlignContentDim / float32(lineCount-1)
			}
		case AlignAuto:
		case AlignFlexStart:
		case AlignBaseline:
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
				if child.style.display == DisplayNone {
					continue
				}
				if child.style.positionType == PositionTypeRelative {
					if child.lineIndex != i {
						break
					}
					if YGNodeIsLayoutDimDefined(child, crossAxis) {
						lineHeight = fmaxf(lineHeight,
							child.layout.measuredDimensions[dim[crossAxis]]+
								YGNodeMarginForAxis(child, crossAxis, availableInnerWidth))
					}
					if YGNodeAlignItem(node, child) == AlignBaseline {
						ascent := YGBaseline(child) + YGNodeLeadingMargin(child, FlexDirectionColumn, availableInnerWidth)
						descent := child.layout.measuredDimensions[DimensionHeight] + YGNodeMarginForAxis(child, FlexDirectionColumn, availableInnerWidth) - ascent
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
					if child.style.display == DisplayNone {
						continue
					}
					if child.style.positionType == PositionTypeRelative {
						switch YGNodeAlignItem(node, child) {
						case AlignFlexStart:
							{
								child.layout.position[pos[crossAxis]] =
									currentLead + YGNodeLeadingMargin(child, crossAxis, availableInnerWidth)
							}
						case AlignFlexEnd:
							{
								child.layout.position[pos[crossAxis]] =
									currentLead + lineHeight -
										YGNodeTrailingMargin(child, crossAxis, availableInnerWidth) -
										child.layout.measuredDimensions[dim[crossAxis]]
							}
						case AlignCenter:
							{
								childHeight := child.layout.measuredDimensions[dim[crossAxis]]
								child.layout.position[pos[crossAxis]] = currentLead + (lineHeight-childHeight)/2
							}
						case AlignStretch:
							{
								child.layout.position[pos[crossAxis]] =
									currentLead + YGNodeLeadingMargin(child, crossAxis, availableInnerWidth)

								// Remeasure child with the line height as it as been only measured with the
								// parents height yet.
								if !YGNodeIsStyleDimDefined(child, crossAxis, availableInnerCrossDim) {
									childWidth := lineHeight
									if isMainAxisRow {
										childWidth = child.layout.measuredDimensions[DimensionWidth] +
											YGNodeMarginForAxis(child, mainAxis, availableInnerWidth)
									}

									childHeight := lineHeight
									if !isMainAxisRow {
										childHeight = child.layout.measuredDimensions[DimensionHeight] +
											YGNodeMarginForAxis(child, crossAxis, availableInnerWidth)
									}

									if !(YGFloatsEqual(childWidth,
										child.layout.measuredDimensions[DimensionWidth]) &&
										YGFloatsEqual(childHeight,
											child.layout.measuredDimensions[DimensionHeight])) {
										YGLayoutNodeInternal(child,
											childWidth,
											childHeight,
											direction,
											MeasureModeExactly,
											MeasureModeExactly,
											availableInnerWidth,
											availableInnerHeight,
											true,
											"multiline-stretch",
											config)
									}
								}
							}
						case AlignBaseline:
							{
								child.layout.position[EdgeTop] =
									currentLead + maxAscentForCurrentLine - YGBaseline(child) +
										YGNodeLeadingPosition(child, FlexDirectionColumn, availableInnerCrossDim)
							}
						case AlignAuto:
						case AlignSpaceBetween:
						case AlignSpaceAround:
						}
					}
				}
			}

			currentLead += lineHeight
		}
	}

	// STEP 9: COMPUTING FINAL DIMENSIONS
	node.layout.measuredDimensions[DimensionWidth] = YGNodeBoundAxis(
		node, FlexDirectionRow, availableWidth-marginAxisRow, parentWidth, parentWidth)
	node.layout.measuredDimensions[DimensionHeight] = YGNodeBoundAxis(
		node, FlexDirectionColumn, availableHeight-marginAxisColumn, parentHeight, parentWidth)

	// If the user didn't specify a width or height for the node, set the
	// dimensions based on the children.
	if measureModeMainDim == MeasureModeUndefined ||
		(node.style.overflow != OverflowScroll && measureModeMainDim == MeasureModeAtMost) {
		// Clamp the size to the min/max size, if specified, and make sure it
		// doesn't go below the padding and border amount.
		node.layout.measuredDimensions[dim[mainAxis]] =
			YGNodeBoundAxis(node, mainAxis, maxLineMainDim, mainAxisParentSize, parentWidth)
	} else if measureModeMainDim == MeasureModeAtMost &&
		node.style.overflow == OverflowScroll {
		node.layout.measuredDimensions[dim[mainAxis]] = fmaxf(
			fminf(availableInnerMainDim+paddingAndBorderAxisMain,
				YGNodeBoundAxisWithinMinAndMax(node, mainAxis, maxLineMainDim, mainAxisParentSize)),
			paddingAndBorderAxisMain)
	}

	if measureModeCrossDim == MeasureModeUndefined ||
		(node.style.overflow != OverflowScroll && measureModeCrossDim == MeasureModeAtMost) {
		// Clamp the size to the min/max size, if specified, and make sure it
		// doesn't go below the padding and border amount.
		node.layout.measuredDimensions[dim[crossAxis]] =
			YGNodeBoundAxis(node,
				crossAxis,
				totalLineCrossDim+paddingAndBorderAxisCross,
				crossAxisParentSize,
				parentWidth)
	} else if measureModeCrossDim == MeasureModeAtMost &&
		node.style.overflow == OverflowScroll {
		node.layout.measuredDimensions[dim[crossAxis]] =
			fmaxf(fminf(availableInnerCrossDim+paddingAndBorderAxisCross,
				YGNodeBoundAxisWithinMinAndMax(node,
					crossAxis,
					totalLineCrossDim+paddingAndBorderAxisCross,
					crossAxisParentSize)),
				paddingAndBorderAxisCross)
	}

	// As we only wrapped in normal direction yet, we need to reverse the positions on wrap-reverse.
	if performLayout && node.style.flexWrap == WrapWrapReverse {
		for i := 0; i < childCount; i++ {
			child := YGNodeGetChild(node, i)
			if child.style.positionType == PositionTypeRelative {
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
		needsMainTrailingPos := mainAxis == FlexDirectionRowReverse || mainAxis == FlexDirectionColumnReverse
		needsCrossTrailingPos := crossAxis == FlexDirectionRowReverse || crossAxis == FlexDirectionColumnReverse

		// Set trailing position if necessary.
		if needsMainTrailingPos || needsCrossTrailingPos {
			for i := 0; i < childCount; i++ {
				child := YGNodeListGet(node.children, i)
				if child.style.display == DisplayNone {
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

var (
	gDepth        = 0
	gPrintTree    = false
	gPrintChanges = false
	gPrintSkips   = false
)

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
	kMeasureModeNames = [measureModeCount]string{"UNDEFINED", "EXACTLY", "AT_MOST"}
	kLayoutModeNames  = [measureModeCount]string{"LAY_UNDEFINED", "LAY_EXACTLY", "LAY_AT_MOST"}
)

// YGMeasureModeName returns name of measure mode
func YGMeasureModeName(mode MeasureMode, performLayout bool) string {

	if mode >= measureModeCount {
		return ""
	}

	if performLayout {
		return kLayoutModeNames[mode]
	}
	return kMeasureModeNames[mode]
}

// YGMeasureModeSizeIsExactAndMatchesOldMeasuredSize retruns true if is exact
func YGMeasureModeSizeIsExactAndMatchesOldMeasuredSize(sizeMode MeasureMode, size float32, lastComputedSize float32) bool {
	return sizeMode == MeasureModeExactly && YGFloatsEqual(size, lastComputedSize)
}

// YGMeasureModeOldSizeIsUnspecifiedAndStillFits returns true if fits
func YGMeasureModeOldSizeIsUnspecifiedAndStillFits(sizeMode MeasureMode, size float32, lastSizeMode MeasureMode, lastComputedSize float32) bool {
	return sizeMode == MeasureModeAtMost && lastSizeMode == MeasureModeUndefined &&
		(size >= lastComputedSize || YGFloatsEqual(size, lastComputedSize))
}

// YGMeasureModeNewMeasureSizeIsStricterAndStillValid returns true if is strict and valid
func YGMeasureModeNewMeasureSizeIsStricterAndStillValid(sizeMode MeasureMode, size float32, lastSizeMode MeasureMode, lastSize float32, lastComputedSize float32) bool {
	return lastSizeMode == MeasureModeAtMost && sizeMode == MeasureModeAtMost &&
		lastSize > size && (lastComputedSize <= size || YGFloatsEqual(size, lastComputedSize))
}

// YGRoundValueToPixelGrid rounds value to pixel grid
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

// YGNodeCanUseCachedMeasurement returns true if can use cached measurement
func YGNodeCanUseCachedMeasurement(widthMode MeasureMode, width float32, heightMode MeasureMode, height float32, lastWidthMode MeasureMode, lastWidth float32, lastHeightMode MeasureMode, lastHeight float32, lastComputedWidth float32, lastComputedHeight float32, marginRow float32, marginColumn float32, config *Config) bool {
	if lastComputedHeight < 0 || lastComputedWidth < 0 {
		return false
	}
	useRoundedComparison := config != nil && config.pointScaleFactor != 0
	effectiveWidth := width
	effectiveHeight := height
	effectiveLastWidth := lastWidth
	effectiveLastHeight := lastHeight

	if useRoundedComparison {
		effectiveWidth = YGRoundValueToPixelGrid(width, config.pointScaleFactor, false, false)
		effectiveHeight = YGRoundValueToPixelGrid(height, config.pointScaleFactor, false, false)
		effectiveLastWidth = YGRoundValueToPixelGrid(lastWidth, config.pointScaleFactor, false, false)
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

// YGLayoutNodeInternal is a wrapper around the YGNodelayoutImpl function. It determines
// whether the layout request is redundant and can be skipped.
//
// Parameters:
//  Input parameters are the same as YGNodelayoutImpl (see above)
//  Return parameter is true if layout was performed, false if skipped
func YGLayoutNodeInternal(node *Node, availableWidth float32, availableHeight float32,
	parentDirection Direction, widthMeasureMode MeasureMode,
	heightMeasureMode MeasureMode, parentWidth float32, parentHeight float32,
	performLayout bool, reason string, config *Config) bool {
	layout := &node.layout

	gDepth++

	needToVisitNode :=
		(node.isDirty && layout.generationCount != gCurrentGenerationCount) ||
			layout.lastParentDirection != parentDirection

	if needToVisitNode {
		// Invalidate the cached results.
		layout.nextCachedMeasurementsIndex = 0
		layout.cachedLayout.widthMeasureMode = MeasureMode(-1)
		layout.cachedLayout.heightMeasureMode = MeasureMode(-1)
		layout.cachedLayout.computedWidth = -1
		layout.cachedLayout.computedHeight = -1
	}

	var cachedResults *CachedMeasurement

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
		marginAxisRow := YGNodeMarginForAxis(node, FlexDirectionRow, parentWidth)
		marginAxisColumn := YGNodeMarginForAxis(node, FlexDirectionColumn, parentWidth)

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
		layout.measuredDimensions[DimensionWidth] = cachedResults.computedWidth
		layout.measuredDimensions[DimensionHeight] = cachedResults.computedHeight

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
				layout.measuredDimensions[DimensionWidth],
				layout.measuredDimensions[DimensionHeight],
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

			var newCacheEntry *CachedMeasurement
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
			newCacheEntry.computedWidth = layout.measuredDimensions[DimensionWidth]
			newCacheEntry.computedHeight = layout.measuredDimensions[DimensionHeight]
		}
	}

	if performLayout {
		node.layout.dimensions[DimensionWidth] = node.layout.measuredDimensions[DimensionWidth]
		node.layout.dimensions[DimensionHeight] = node.layout.measuredDimensions[DimensionHeight]
		node.hasNewLayout = true
		node.isDirty = false
	}

	gDepth--
	layout.generationCount = gCurrentGenerationCount
	return needToVisitNode || cachedResults == nil
}

// YGConfigSetPointScaleFactor sets scale factor
func YGConfigSetPointScaleFactor(config *Config, pixelsInPoint float32) {
	YGAssertWithConfig(config, pixelsInPoint >= 0, "Scale factor should not be less than zero")

	// We store points for Pixel as we will use it for rounding
	if pixelsInPoint == 0 {
		// Zero is used to skip rounding
		config.pointScaleFactor = 0
	} else {
		config.pointScaleFactor = pixelsInPoint
	}
}

// YGRoundToPixelGrid rounds to pixel grid
func YGRoundToPixelGrid(node *Node, pointScaleFactor float32, absoluteLeft float32, absoluteTop float32) {
	if pointScaleFactor == 0.0 {
		return
	}

	nodeLeft := node.layout.position[EdgeLeft]
	nodeTop := node.layout.position[EdgeTop]

	nodeWidth := node.layout.dimensions[DimensionWidth]
	nodeHeight := node.layout.dimensions[DimensionHeight]

	absoluteNodeLeft := absoluteLeft + nodeLeft
	absoluteNodeTop := absoluteTop + nodeTop

	absoluteNodeRight := absoluteNodeLeft + nodeWidth
	absoluteNodeBottom := absoluteNodeTop + nodeHeight

	// If a node has a custom measure function we never want to round down its size as this could
	// lead to unwanted text truncation.
	textRounding := node.nodeType == NodeTypeText

	node.layout.position[EdgeLeft] = YGRoundValueToPixelGrid(nodeLeft, pointScaleFactor, false, textRounding)
	node.layout.position[EdgeTop] = YGRoundValueToPixelGrid(nodeTop, pointScaleFactor, false, textRounding)

	// We multiply dimension by scale factor and if the result is close to the whole number, we don't have any fraction
	// To verify if the result is close to whole number we want to check both floor and ceil numbers
	hasFractionalWidth := !YGFloatsEqual(fmodf(nodeWidth*pointScaleFactor, 1), 0) &&
		!YGFloatsEqual(fmodf(nodeWidth*pointScaleFactor, 1), 1)
	hasFractionalHeight := !YGFloatsEqual(fmodf(nodeHeight*pointScaleFactor, 1), 0) &&
		!YGFloatsEqual(fmodf(nodeHeight*pointScaleFactor, 1), 1)

	node.layout.dimensions[DimensionWidth] =
		YGRoundValueToPixelGrid(
			absoluteNodeRight,
			pointScaleFactor,
			(textRounding && hasFractionalWidth),
			(textRounding && !hasFractionalWidth)) -
			YGRoundValueToPixelGrid(absoluteNodeLeft, pointScaleFactor, false, textRounding)
	node.layout.dimensions[DimensionHeight] =
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

func calcStartWidth(node *Node, parentWidth float32) (float32, MeasureMode) {
	if YGNodeIsStyleDimDefined(node, FlexDirectionRow, parentWidth) {
		width := YGResolveValue(node.resolvedDimensions[dim[FlexDirectionRow]], parentWidth)
		margin := YGNodeMarginForAxis(node, FlexDirectionRow, parentWidth)
		return width + margin, MeasureModeExactly
	}
	if YGResolveValue(&node.style.maxDimensions[DimensionWidth], parentWidth) >= 0.0 {
		width := YGResolveValue(&node.style.maxDimensions[DimensionWidth], parentWidth)
		return width, MeasureModeAtMost
	}

	width := parentWidth
	widthMeasureMode := MeasureModeExactly
	if YGFloatIsUndefined(width) {
		widthMeasureMode = MeasureModeUndefined
	}
	return width, widthMeasureMode
}
func calcStartHeight(node *Node, parentWidth, parentHeight float32) (float32, MeasureMode) {
	if YGNodeIsStyleDimDefined(node, FlexDirectionColumn, parentHeight) {
		height := YGResolveValue(node.resolvedDimensions[dim[FlexDirectionColumn]], parentHeight)
		margin := YGNodeMarginForAxis(node, FlexDirectionColumn, parentWidth)
		return height + margin, MeasureModeExactly
	}
	if YGResolveValue(&node.style.maxDimensions[DimensionHeight], parentHeight) >= 0 {
		height := YGResolveValue(&node.style.maxDimensions[DimensionHeight], parentHeight)
		return height, MeasureModeAtMost
	}
	height := parentHeight
	heightMeasureMode := MeasureModeExactly
	if YGFloatIsUndefined(height) {
		heightMeasureMode = MeasureModeUndefined
	}
	return height, heightMeasureMode
}

// YGNodeCalculateLayout sets
func YGNodeCalculateLayout(node *Node, parentWidth float32, parentHeight float32, parentDirection Direction) {
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
			YGNodePrint(node, PrintOptionsLayout|PrintOptionsChildren|PrintOptionsStyle)
		}
	}
}

// YGConfigSetExperimentalFeatureEnabled enables experimental feature
func YGConfigSetExperimentalFeatureEnabled(config *Config, feature ExperimentalFeature, enabled bool) {
	config.experimentalFeatures[feature] = enabled
}

// YGConfigIsExperimentalFeatureEnabled returns if experimental feature is enabled
func YGConfigIsExperimentalFeatureEnabled(config *Config, feature ExperimentalFeature) bool {
	return config.experimentalFeatures[feature]
}

// YGConfigSetUseWebDefaults sets useWebDefaults
func YGConfigSetUseWebDefaults(config *Config, enabled bool) {
	config.useWebDefaults = enabled
}

// YGConfigSetUseLegacyStretchBehaviour sets legacy stretch behavior
func YGConfigSetUseLegacyStretchBehaviour(config *Config, useLegacyStretchBehaviour bool) {
	config.useLegacyStretchBehaviour = useLegacyStretchBehaviour
}

// YGConfigGetUseWebDefaults gets useWebDefaults
func YGConfigGetUseWebDefaults(config *Config) bool {
	return config.useWebDefaults
}

// YGConfigSetContext sets context
func YGConfigSetContext(config *Config, context interface{}) {
	config.context = context
}

// YGConfigGetContext gets context
func YGConfigGetContext(config *Config) interface{} {
	return config.context
}

// YGLog logs
func YGLog(node *Node, level LogLevel, format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// YGAssert asserts that cond is true
func YGAssert(cond bool, format string, args ...interface{}) {
	if !cond {
		panic(format)
	}
}

// YGAssertWithNode assert if cond is not true
func YGAssertWithNode(node *Node, cond bool, format string, args ...interface{}) {
	YGAssert(cond, format, args...)
}

// YGAssertWithConfig asserts with config
func YGAssertWithConfig(config *Config, condition bool, message string) {
	if !condition {
		panic(message)
	}
}
