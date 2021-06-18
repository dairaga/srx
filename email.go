// +build js,wasm

package srx

type (
	TEmail interface {
		TBaseInput
	}
)

func Email(owner TComponent) TEmail {
	input := newInput(owner)
	input.SetType("email")
	return input
}
