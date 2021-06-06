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
		Value() string
		SetValue(v string)
		Color() enum.Color
		SetColor(c enum.Color)
		Outline() bool
		SetOutline(outline bool)
	}

	button struct {
		*srx.Component
		caption *srx.Object
		color   enum.Color
		outline bool
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
	old := btn.color.Style("btn")
	if btn.outline {
		old = btn.color.Style("btn", "outline")
	}

	new := color.Style("btn")
	if outline {
		new = color.Style("btn", "outline")
	}
	btn.Replace(old, new)
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

func ButtonOf(owner srx.TComponent) TButton {
	caption := srx.NewObject(js.Create("span"))
	el := js.From(js.HTML(`<button type="button" class="btn btn-primary"></button>`))
	el.Append(caption)

	btn := &button{
		Component: srx.NewComponent(owner, el),
		caption:   caption,
		color:     enum.Primary,
	}

	if owner != nil {
		owner.Add(btn)
	}
	return btn
}
