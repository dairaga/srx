// +build js,wasm

package std

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
	return i.Attr("value")
}

// -----------------------------------------------------------------------------

func (i *input) SetValue(v string) {
	i.SetAttr("value", v)
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

func InputOf(owner srx.TComponent) TInput {
	ret := &input{
		Component: srx.NewComponent(owner, js.From(`<input>`)),
	}
	if owner != nil {
		owner.Add(ret)
	}

	return ret
}
