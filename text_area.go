// +build js,wasm

package srx

import (
	"strconv"

	"github.com/dairaga/srx/js"
)

type (
	TTextarea interface {
		TBaseInput
		SetRows(rows int)
	}

	textarea struct {
		*component
	}
)

var _ TTextarea = &textarea{}

// -----------------------------------------------------------------------------

func (t *textarea) Name() string {
	return t.Attr("name")
}

// -----------------------------------------------------------------------------

func (t *textarea) SetName(n string) {
	t.SetAttr("name", n)

}

// -----------------------------------------------------------------------------

func (t *textarea) Value() string {
	return t.Prop("value").String()
}

// -----------------------------------------------------------------------------

func (t *textarea) SetValue(v string) {
	t.SetProp("value", v)
}

// -----------------------------------------------------------------------------

func (t *textarea) ReadOnly() bool {
	return t.HasAttr("readonly")
}

// -----------------------------------------------------------------------------

func (t *textarea) SetReadOnly(only bool) {
	if only {
		t.SetAttr("readonly", "true")
	} else {
		t.RemoveAttr("readonly")
	}
}

// -----------------------------------------------------------------------------

func (t *textarea) Required() bool {
	return t.HasAttr("required")
}

// -----------------------------------------------------------------------------

func (t *textarea) SetRequired(required bool) {
	if required {
		t.SetAttr("required", "true")
	} else {
		t.RemoveAttr("required")
	}
}

// -----------------------------------------------------------------------------

func (t *textarea) SetPlaceholder(h string) {
	t.SetAttr("placeholder", h)
}

// -----------------------------------------------------------------------------

func (t *textarea) SetRows(rows int) {
	t.SetAttr("rows", strconv.Itoa(rows))
}

// -----------------------------------------------------------------------------

func Textarea(owner TComponent) TTextarea {
	el := js.Create("textarea").Add("form-control").SetAttr("rows", "3")

	ret := &textarea{
		component: newComponent(owner, el),
	}
	if owner != nil {
		owner.Add(ret)
	}

	return ret
}
