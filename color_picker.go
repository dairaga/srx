// +build js,wasm

package srx

type (
	TColorPicker interface {
		TBaseFormControl
	}
)

func ColorPicker(owner TComponent) TColorPicker {
	input := newInput(owner)
	input.SetType("color")
	input.Ref().Add("form-control-color")
	return input
}
