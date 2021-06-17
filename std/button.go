// build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/el"
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
		Value() string
		SetValue(v string)
		Color() enum.Color
		SetColor(c enum.Color)
		Outline() bool
		SetOutline(outline bool)

		Click()
		OnClick(fn func(TButton, js.TEvent))
	}

	button struct {
		*srx.Component
		caption el.TCaption
		//color   enum.Color
		outline bool
	}
)

var _ TButton = &button{}

// -----------------------------------------------------------------------------

func (btn *button) Caption() string {
	return btn.caption.Caption()
}

// -----------------------------------------------------------------------------

func (btn *button) SetCaption(c string) {
	btn.caption.SetCaption(c)
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

func (btn *button) Value() string {
	return btn.Attr("value")
}

// -----------------------------------------------------------------------------

func (btn *button) SetValue(v string) {
	btn.SetAttr("value", v)
}

// -----------------------------------------------------------------------------

func (btn *button) Color() enum.Color {
	return btn.color
}

// -----------------------------------------------------------------------------

func (btn *button) setOutlineColor(outline bool, color enum.Color) {
	if btn.outline {
		btn.color.UnapplyOutlineButton(btn)
	} else {
		btn.color.UnapplyButton(btn)
	}

	if outline {
		color.ApplyOutlineButton(btn)
	} else {
		color.ApplyButton(btn)
	}

	btn.outline = outline
	btn.color = color
}

// -----------------------------------------------------------------------------

func (btn *button) SetColor(c enum.Color) {
	btn.setOutlineColor(btn.outline, c)
}

// -----------------------------------------------------------------------------

func (btn *button) Outline() bool {
	return btn.outline
}

// -----------------------------------------------------------------------------

func (btn *button) SetOutline(outline bool) {
	btn.setOutlineColor(outline, btn.color)
}

// -----------------------------------------------------------------------------

func (btn *button) SetSize(s enum.Size) {
	btn.Element.Remove(btn.Size().Style("btn"))

	if s >= enum.Small && s <= enum.Large {
		new := s.Style("btn")
		if new != "" {
			btn.Element.Add(new)
		}

	}
	btn.Object.SetSize(s)
}

// -----------------------------------------------------------------------------

func (btn *button) Click() {
	btn.Element.Call("click")
}

// -----------------------------------------------------------------------------

func (btn *button) OnClick(fn func(TButton, js.TEvent)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		sender := srx.Lookup(this).(TButton)
		evt := js.EventOf(args[0])
		fn(sender, evt)
		return nil
	})
	btn.EventTarget.On("click", cb)
}

// -----------------------------------------------------------------------------

func ButtonOf(owner srx.TComponent) TButton {
	caption := el.Caption()
	el := js.From(js.HTML(`<button type="button" class="btn"></button>`))
	el.Append(caption)

	btn := &button{
		Component: srx.NewComponent(owner, el),
		caption:   caption,
		color:     enum.None,
	}

	if owner != nil {
		owner.Add(btn)
	}
	return btn
}
