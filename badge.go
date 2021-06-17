// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TBadge interface {
		TComponent

		Type() enum.RoundType
		SetType(t enum.RoundType)

		Value() string
		SetValue(v string)

		Description() string
		SetDescription(desc string)
	}

	badge struct {
		*component
		caption TCaption
		assist  TAssist
		//color   enum.Color
		typ enum.RoundType
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
	return b.Background()
}

// -----------------------------------------------------------------------------

func (b *badge) SetColor(c enum.Color) {
	b.SetBackground(c)
}

// -----------------------------------------------------------------------------

func (b *badge) Background() enum.Color {
	return b.bgColor
}

// -----------------------------------------------------------------------------

func (b *badge) SetBackground(c enum.Color) {
	if !c.IsTheme() {
		return
	}
	switch b.bgColor {
	case enum.Warning, enum.Info, enum.Light:
		enum.Dark.UnapplyTextColor(b)
		fallthrough
	default:
		b.bgColor.UnapplyBackground(b)
	}
	switch c {
	case enum.Warning, enum.Info, enum.Light:
		b.component.SetColor(enum.Dark)
		fallthrough
	default:
		b.component.SetBackground(c)
	}
}

// -----------------------------------------------------------------------------

func Badge(owner TComponent) TBadge {
	caption := Caption()
	assist := Assist()
	el := js.From(`<span class="badge"></span>`)
	b := &badge{
		component: newComponent(owner, el),
		caption:   caption,
		assist:    assist,
		typ:       enum.RoundNone,
	}
	b.Element.Append(caption)
	b.Element.Append(assist)
	b.SetColor(enum.Secondary)

	if owner != nil {
		owner.Add(b)
	}
	return b
}