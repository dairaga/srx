// +build js,wasm

package form

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TInput interface {
		srx.TComponent
		Type() string
		SetType(t string)

		Name() string
		SetName(n string)

		Value() string
		SetValue(v string)

		ReadOnly() bool
		SetReadOnly(only bool)

		Required() bool
		SetRequired(required bool)

		SetPlaceholder(h string)
	}

	input struct {
		*srx.Component
		input *js.Element
	}
)

var _ TInput = &input{}

// -----------------------------------------------------------------------------

func (i *input) Type() string {
	return i.input.Attr("type")
}

// -----------------------------------------------------------------------------

func (i *input) SetType(t string) {
	i.input.SetAttr("type", t)
}

// -----------------------------------------------------------------------------

func (i *input) Name() string {
	return i.input.Attr("name")
}

// -----------------------------------------------------------------------------

func (i *input) SetName(n string) {
	i.input.SetAttr("name", n)

}

// -----------------------------------------------------------------------------

func (i *input) Value() string {
	return i.input.Attr("value")
}

// -----------------------------------------------------------------------------

func (i *input) SetValue(v string) {
	i.input.SetAttr("value", v)
}

// -----------------------------------------------------------------------------

func (i *input) ReadOnly() bool {
	return i.input.HasAttr("readonly")
}

// -----------------------------------------------------------------------------

func (i *input) SetReadOnly(only bool) {
	if only {
		i.input.SetAttr("readonly", "true")
	} else {
		i.input.RemoveAttr("readonly")
	}
}

// -----------------------------------------------------------------------------

func (i *input) Required() bool {
	return i.input.HasAttr("required")
}

// -----------------------------------------------------------------------------

func (i *input) SetRequired(required bool) {
	if required {
		i.input.SetAttr("required", "true")
	} else {
		i.input.RemoveAttr("required")
	}
}

// -----------------------------------------------------------------------------

func (i *input) SetPlaceholder(h string) {
	i.input.SetAttr("placeholder", h)
}

// -----------------------------------------------------------------------------

func InputOf(owner srx.TComponent) TInput {
	div := js.Create("div")
	in := js.Create("input").Add("form-control")
	div.Append(in)

	ret := &input{
		Component: srx.NewComponent(owner, div),
		input:     in,
	}
	if owner != nil {
		owner.Add(ret)
	}

	return ret
}
