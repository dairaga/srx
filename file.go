// +build js,wasm

package srx

type (
	TFile interface {
		TBaseInput
	}
)

// -----------------------------------------------------------------------------

func File(owner TComponent) TFile {
	input := newInput(owner)
	input.SetType("file")
	return input
}
