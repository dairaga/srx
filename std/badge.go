// +build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/el"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TBadge interface {
		srx.TComponent

		Type() enum.RoundType
		SetType(t enum.RoundType)

		Color() enum.Color
		SetColor(color enum.Color)

		Value() string
		SetValue(v string)

		Description() string
		SetDescription(desc string)
	}

	badge struct {
		*srx.Component
		caption el.TCaption
		assist  el.TAssist
		color   enum.Color
		typ     enum.RoundType
	}
)

var _ TBadge = &badge{}

// -----------------------------------------------------------------------------

func (b *badge) Type() enum.RoundType {
	return b.typ
}

// -----------------------------------------------------------------------------

func (b *badge) SetType(t enum.RoundType) {
	t.Replace(b.Element, b.typ)
	b.typ = t
}

// -----------------------------------------------------------------------------

func (b *badge) Value() string {
	return b.caption.Caption()
}

// -----------------------------------------------------------------------------

func (b *badge) SetValue(v string) {
	b.caption.SetCaption(v)
}

// -----------------------------------------------------------------------------

func (b *badge) Description() string {
	return b.assist.Description()
}

// -----------------------------------------------------------------------------

func (b *badge) SetDescription(d string) {
	b.assist.SetDescription(d)
}

// -----------------------------------------------------------------------------

func (b *badge) Color() enum.Color {
	return b.color
}

// -----------------------------------------------------------------------------

func (b *badge) SetColor(c enum.Color) {
	switch b.color {
	case enum.Warning, enum.Info, enum.Light:
		b.Element.Remove("text-dark", b.color.Style("bg"))
	default:
		b.Element.Remove(b.color.Style("bg"))
	}
	switch c {
	case enum.Warning, enum.Info, enum.Light:
		b.Element.Add("text-dark", c.Style("bg"))
	default:
		b.Element.Add(c.Style("bg"))
	}
	b.color = c
}

// -----------------------------------------------------------------------------

func BadgeOf(owner srx.TComponent) TBadge {
	caption := el.Caption()
	assist := el.Assist()
	el := js.From(`<span class="badge bg-secondary"></span>`)
	b := &badge{
		Component: srx.NewComponent(owner, el),
		caption:   caption,
		assist:    assist,
		color:     enum.Secondary,
		typ:       enum.RoundNone,
	}
	b.Element.Append(caption)
	b.Element.Append(assist)

	if owner != nil {
		owner.Add(b)
	}
	return b
}
