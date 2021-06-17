// +build js,wasm

package srx

import (
	"fmt"

	"github.com/dairaga/srx/js"
)

type (
	TAssist interface {
		TObject
		Description() string
		SetDescription(d string)
	}

	assist struct {
		*object
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
		object: newObject(
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
