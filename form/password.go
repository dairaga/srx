// +build js,wasm

package form

import (
	"github.com/dairaga/srx"
)

type (
	TPassword interface {
		TBaseInput
	}
)

// -----------------------------------------------------------------------------

func PasswordOf(owner srx.TComponent) TPassword {
	input := newFormCtrol(owner)
	input.SetType("password")
	return input
}
