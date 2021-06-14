// +build js,wasm

package el

import (
	"fmt"

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

func newAssist(content string) *assist {
	ret := &assist{
		Object: srx.NewObject(
			js.From(`<span class="visually-hidden"></span>`),
		),
	}
	ret.SetDescription(content)
	return ret
}

// -----------------------------------------------------------------------------

func Assist(a ...interface{}) TAssist {
	return newAssist(fmt.Sprint(a...))
}

// -----------------------------------------------------------------------------

func Assistf(format string, a ...interface{}) TAssist {
	return newAssist(fmt.Sprintf(format, a...))
}
