// +build js,wasm

package form

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TBaseFormControl interface {
		srx.TComponent
		Name() string
		SetName(n string)

		Value() string
		SetValue(v string)

		ReadOnly() bool
		SetReadOnly(only bool)

		Required() bool
		SetRequired(required bool)
	}

	TBaseInput interface {
		TBaseFormControl
		SetPlaceholder(h string)
	}

	TInput interface {
		TBaseInput
		Type() string
		SetType(t string)
	}

	input struct {
		*srx.Component
	}
)

var _ TInput = &input{}

// -----------------------------------------------------------------------------

func (i *input) Type() string {
	return i.Attr("type")
}

// -----------------------------------------------------------------------------

func (i *input) SetType(t string) {
	i.SetAttr("type", t)
}

// -----------------------------------------------------------------------------

func (i *input) Name() string {
	return i.Attr("name")
}

// -----------------------------------------------------------------------------

func (i *input) SetName(n string) {
	i.SetAttr("name", n)

}

// -----------------------------------------------------------------------------

func (i *input) Value() string {
	return i.Element.Prop("value").String()
}

// -----------------------------------------------------------------------------

func (i *input) SetValue(v string) {
	i.Element.SetProp("value", v)
}

// -----------------------------------------------------------------------------

func (i *input) ReadOnly() bool {
	return i.HasAttr("readonly")
}

// -----------------------------------------------------------------------------

func (i *input) SetReadOnly(only bool) {
	if only {
		i.SetAttr("readonly", "true")
	} else {
		i.RemoveAttr("readonly")
	}
}

// -----------------------------------------------------------------------------

func (i *input) Required() bool {
	return i.HasAttr("required")
}

// -----------------------------------------------------------------------------

func (i *input) SetRequired(required bool) {
	if required {
		i.SetAttr("required", "true")
	} else {
		i.RemoveAttr("required")
	}
}

// -----------------------------------------------------------------------------

func (i *input) SetPlaceholder(h string) {
	i.SetAttr("placeholder", h)
}

// -----------------------------------------------------------------------------

// TODO: refactor function name: newFormControl, newInput
func newInput(owner srx.TComponent) *input {
	return newFormControl(owner, "input", "form-control")
}

// -----------------------------------------------------------------------------

func newFormControl(owner srx.TComponent, tagName, class string) *input {
	el := js.Create(tagName).Add(class)

	ret := &input{
		Component: srx.NewComponent(owner, el),
	}
	if owner != nil {
		owner.Add(ret)
	}

	return ret
}

// -----------------------------------------------------------------------------

func InputOf(owner srx.TComponent) TInput {
	return newInput(owner)
}
