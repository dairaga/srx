// build js,wasm

package srx

import (
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TButton interface {
		TComponent
		Type() enum.ButtonType
		SetType(typ enum.ButtonType)
		SetCaption(c string)
		Caption() string
		Value() string
		SetValue(v string)

		Outline() bool
		SetOutline(outline bool)

		Click()
		OnClick(fn func(TButton, js.TEvent))
	}

	button struct {
		*component
		caption TCaption
		color   enum.Color
		outline bool
	}
)

var _ TButton = &button{}

func (btn *button) Caption() string {
	return btn.caption.Caption()
}

func (btn *button) SetCaption(c string) {
	btn.caption.SetCaption(c)
}

func (btn *button) Type() enum.ButtonType {
	var t enum.ButtonType
	t.SetString(btn.Attr("type"))
	return t
}

func (btn *button) SetType(typ enum.ButtonType) {
	btn.SetAttr("type", typ.String())
}

func (btn *button) Value() string {
	return btn.Attr("value")
}

func (btn *button) SetValue(v string) {
	btn.SetAttr("value", v)
}

func (btn *button) setOutlineColor(outline bool, color enum.Color) {
	handled := false
	if outline {
		handled = color.ApplyOutlineButton(btn)
	} else {
		handled = color.ApplyButton(btn)
	}

	if !handled {
		return
	}

	if btn.outline {
		btn.color.UnapplyOutlineButton(btn)
	} else {
		btn.color.UnapplyButton(btn)
	}

	btn.outline = outline
	btn.color = color
	btn.bgColor = color
}

func (btn *button) Color() enum.Color {
	return btn.color
}

func (btn *button) SetColor(c enum.Color) {
	btn.setOutlineColor(btn.outline, c)
}

func (btn *button) Background() enum.Color {
	return btn.Color()
}

func (btn *button) SetBackground(c enum.Color) {
	btn.SetColor(c)
}

func (btn *button) Outline() bool {
	return btn.outline
}

func (btn *button) SetOutline(outline bool) {
	btn.setOutlineColor(outline, btn.color)
}

func (btn *button) SetSize(s enum.Size) {
	btn.Element.Remove(btn.Size().Style("btn"))

	if s >= enum.Small && s <= enum.Large {
		new := s.Style("btn")
		if new != "" {
			btn.Element.Add(new)
		}

	}
	btn.object.SetSize(s)
}

func (btn *button) Click() {
	btn.Element.Call("click")
}

func (btn *button) OnClick(fn func(TButton, js.TEvent)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		sender := Lookup(this).(TButton)
		evt := js.EventOf(args[0])
		fn(sender, evt)
		return nil
	})
	btn.EventTarget.On("click", cb)
}

func newButton(owner TComponent) *button {
	caption := Caption()
	el := js.From(js.HTML(`<button type="button" class="btn"></button>`))
	el.Append(caption)

	btn := &button{
		component: newComponent(el),
		caption:   caption,
		color:     enum.None,
	}

	bindOwner(owner, btn)
	return btn
}

func Button(owner TComponent) TButton {
	return newButton(owner)
}
