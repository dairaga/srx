// build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TButton interface {
		srx.TComponent
		Type() enum.ButtonType
		SetType(typ enum.ButtonType)
		SetCaption(c string)
		Caption() string
	}

	button struct {
		*srx.Component
		caption *srx.Object
	}
)

var _ TButton = &button{}

// -----------------------------------------------------------------------------

func (btn *button) Caption() string {
	return btn.caption.Text()
}

// -----------------------------------------------------------------------------

func (btn *button) SetCaption(c string) {
	btn.caption.SetText(c)
}

// -----------------------------------------------------------------------------

func (btn *button) Type() enum.ButtonType {
	var t enum.ButtonType
	t.SetString(btn.Attr("type"))
	return t
}

// -----------------------------------------------------------------------------

func (btn *button) SetType(typ enum.ButtonType) {
	btn.SetAttr("type", typ.String())
}

// -----------------------------------------------------------------------------

func ButtonOf(owner srx.TComponent) TButton {
	caption := srx.NewObject(js.Create("span"))
	el := js.From(js.HTML(`<button type="button" class="btn"></button>`))
	el.Append(caption)

	btn := &button{
		Component: srx.NewComponent(owner, el),
		caption:   caption,
	}
	if owner != nil {
		owner.Add(btn)
	}
	return btn
}
