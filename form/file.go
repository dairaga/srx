// +build js,wasm

package form

import (
	"github.com/dairaga/srx"
)

type (
	TFile interface {
		TBaseInput
	}
)

// -----------------------------------------------------------------------------

func FileOf(owner srx.TComponent) TFile {
	input := newInput(owner)
	input.SetType("file")
	return input
}
