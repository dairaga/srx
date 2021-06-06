// build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TButton interface {
		srx.TComponent
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

func ButtonOf(owner srx.TComponent) TButton {
	caption := srx.NewObject(js.Create("span"))
	el := js.From(js.HTML(`<button type="button" class="btn"></button>`))
	el.Append(caption)

	btn := &button{
		Component: srx.NewComponent(owner, el),
		caption:   caption,
	}

	return btn
}
