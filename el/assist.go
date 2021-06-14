// +build js,wasm

package el

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TAssist interface {
		srx.TObject
		Description() string
		SetDescription(d string)
	}

	assist struct {
		*srx.Object
	}
)

var _ TAssist = &assist{}

// -----------------------------------------------------------------------------

func (a *assist) Description() string {
	return a.Element.Text()
}

// -----------------------------------------------------------------------------

func (a *assist) SetDescription(d string) {
	a.Element.SetText(d)
}

// -----------------------------------------------------------------------------

func Assist() TAssist {
	return &assist{
		Object: srx.NewObject(
			js.From(`<span class="visually-hidden"></span>`),
		),
	}
}
