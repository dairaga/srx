// +build js,wasm

package srx

type (
	TPassword interface {
		TBaseInput
	}
)

func Password(owner TComponent) TPassword {
	input := newInput(owner)
	input.SetType("password")
	return input
}
