// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/js"
)

type (
	TBaseFormControl interface {
		TComponent
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
		*component
	}
)

var _ TInput = &input{}

// -----------------------------------------------------------------------------

func (i *input) Type() string {
	return i.Prop("type").String()
}

// -----------------------------------------------------------------------------

func (i *input) SetType(t string) {
	i.SetProp("type", t)
}

// -----------------------------------------------------------------------------

func (i *input) Name() string {
	return i.Prop("name").String()
}

// -----------------------------------------------------------------------------

func (i *input) SetName(n string) {
	//i.SetAttr("name", n)
	i.SetProp("name", n)
}

// -----------------------------------------------------------------------------

func (i *input) Value() string {
	return i.Prop("value").String()
}

// -----------------------------------------------------------------------------

func (i *input) SetValue(v string) {
	i.SetProp("value", v)
}

// -----------------------------------------------------------------------------

func (i *input) ReadOnly() bool {
	return i.Prop("readOnly").Bool()
}

// -----------------------------------------------------------------------------

func (i *input) SetReadOnly(only bool) {
	i.SetProp("readOnly", only)
}

// -----------------------------------------------------------------------------

func (i *input) Required() bool {
	return i.Prop("required").Bool()
}

// -----------------------------------------------------------------------------

func (i *input) SetRequired(required bool) {
	i.SetProp("required", required)
}

// -----------------------------------------------------------------------------

func (i *input) SetPlaceholder(h string) {
	i.SetProp("placeholder", h)
}

// -----------------------------------------------------------------------------

func newInput(owner TComponent) *input {
	return newFormControl(owner, "input", "form-control")
}

// -----------------------------------------------------------------------------

func newFormControl(owner TComponent, tagName, class string) *input {
	el := js.Create(tagName).Add(class)

	ret := &input{
		component: newComponent(owner, el),
	}
	if owner != nil {
		owner.Add(ret)
	}

	return ret
}

// -----------------------------------------------------------------------------

func Input(owner TComponent) TInput {
	return newInput(owner)
}
