// +build js,wasm

package form

import (
	"github.com/dairaga/srx"
)

type (
	TText interface {
		TBaseInput
	}
)

// -----------------------------------------------------------------------------

func TextOf(owner srx.TComponent) TText {
	input := newInput(owner)
	input.SetType("text")
	return input
}
