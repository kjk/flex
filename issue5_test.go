package flex

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIssue5(t *testing.T) {
	config := NewConfig()
	config.Context = "test"

	// check that "padding" set with EdgeAll is printed out
	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionColumn)
	root.StyleSetHeightPercent(100)

	child := NewNodeWithConfig(config)
	child.StyleSetPadding(EdgeAll, 20)
	root.InsertChild(child, 0)

	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	w := &bytes.Buffer{}
	printer := NewNodePrinter(w, PrintOptionsLayout|PrintOptionsStyle|PrintOptionsChildren)
	printer.Print(root)
	got := w.String()
	exp := `<div layout="width: 40; height: 40; top: 0; left: 0;" style="height: 100%; ">
  <div layout="width: 40; height: 40; top: 0; left: 0;" style="padding: 20px; "></div>
</div>`
	assert.Equal(t, got, exp)
}
