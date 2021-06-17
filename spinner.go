// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TSpinner interface {
		TComponent

		Type() enum.SpinnerType
		SetType(t enum.SpinnerType)

		Description() string
		SetDescription(desc string)
	}

	spinner struct {
		*component
		assist TAssist
		typ    enum.SpinnerType
	}
)

var _ TSpinner = &spinner{}

// -----------------------------------------------------------------------------

func (s *spinner) Type() enum.SpinnerType {
	return s.typ
}

// -----------------------------------------------------------------------------

func (s *spinner) SetType(t enum.SpinnerType) {
	s.Element.Replace(s.typ.String(), t.String())
	s.typ = t
}

// -----------------------------------------------------------------------------

func (s *spinner) Description() string {
	return s.assist.Description()
}

// -----------------------------------------------------------------------------

func (s *spinner) SetDescription(d string) {
	s.assist.SetDescription(d)
}

// -----------------------------------------------------------------------------

func Spinner(owner TComponent) TSpinner {
	assist := Assist()
	el := js.From(`<div class="text-dark spinner-border" role="status"></div>`)
	ret := &spinner{
		component: newComponent(owner, el),
		assist:    assist,
	}
	ret.Element.Append(assist)
	ret.color = enum.Dark
	ret.typ = enum.SpinnerBorder
	//ret.SetColor(enum.Dark)
	//ret.SetType(enum.SpinnerBorder)

	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
