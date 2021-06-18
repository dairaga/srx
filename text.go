// +build js,wasm

package srx

type (
	TText interface {
		TBaseInput
	}
)

// -----------------------------------------------------------------------------

func Text(owner TComponent) TText {
	input := newInput(owner)
	input.SetType("text")
	return input
}
