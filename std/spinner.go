// +build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/el"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TSpinner interface {
		srx.TComponent

		Type() enum.SpinnerType
		SetType(t enum.SpinnerType)

		Color() enum.Color
		SetColor(color enum.Color)

		Description() string
		SetDescription(desc string)
	}

	spinner struct {
		*srx.Component
		assist el.TAssist
		typ    enum.SpinnerType
		color  enum.Color
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

func (s *spinner) Color() enum.Color {
	return s.color
}

// -----------------------------------------------------------------------------

func (s *spinner) SetColor(c enum.Color) {
	s.Element.Replace(s.color.Style("text"), c.Style("text"))
	s.color = c
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

func SpinnerOf(owner srx.TComponent) TSpinner {
	assist := el.AssistOf()
	el := js.From(`<div class="spinner-border text-dark" role="status"></div>`)
	ret := &spinner{
		Component: srx.NewComponent(owner, el),
		assist:    assist,
		color:     enum.Dark,
		typ:       enum.SpinnerBorder,
	}
	ret.Element.Append(assist)
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
