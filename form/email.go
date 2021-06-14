// +build js,wasm

package form

import (
	"github.com/dairaga/srx"
)

type (
	TEmail interface {
		TBaseInput
	}
)

// -----------------------------------------------------------------------------

func EmailOf(owner srx.TComponent) TEmail {
	input := newInput(owner)
	input.SetType("email")
	return input
}
