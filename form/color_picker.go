// +build js,wasm

package form

import (
	"github.com/dairaga/srx"
)

type (
	TColorPicker interface {
		TBaseFormControl
	}
)

// -----------------------------------------------------------------------------

func ColorPickerOf(owner srx.TComponent) TColorPicker {
	input := newInput(owner)
	input.SetType("color")
	input.Ref().Add("form-control-color")
	return input
}
