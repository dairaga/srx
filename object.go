// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

const (
	srxTattoo = `data-srx`
)

type (
	// TObject represents vdom object.
	// It is a js html element with tattoo for seeking object.
	TObject interface {
		js.Wrapper
		enum.ObjRef

		OK() bool

		Append(children ...TObject)
		Prepend(Children ...TObject)

		// ID returns js attribute id.
		ID() string
		SetID(id string)

		// Disabled returns true if object is disabled, or returns true.
		Disabled() bool
		Disable()
		Enable()

		Hidden() bool
		Hide()
		Show()

		Color() enum.Color
		SetColor(c enum.Color)

		Background() enum.Color
		SetBackground(c enum.Color)

		FlexMode() enum.FlexMode
		SetFlexMode(m enum.FlexMode)
		AlignHorizontal(al enum.Align)
		AlignVertical(al enum.Align)

		FontSize() enum.Size
		SetFontSize(s enum.Size)

		FontWeight() enum.FontWeight
		SetFontWeight(w enum.FontWeight)

		Rounded() enum.RoundedType
		SetRounded(r enum.RoundedType)

		Border() enum.BorderType
		SetBorder(b enum.BorderType)

		BorderSize() enum.Size
		SetBorderSize(s enum.Size)

		BorderColor() enum.Color
		SetBorderColor(c enum.Color)

		Size() enum.Size
		SetSize(s enum.Size)

		SetMargin(pos enum.Pos, size enum.Size)
		RemoveMargin(pos enum.Pos, size enum.Size)

		SetPadding(pos enum.Pos, size enum.Size)
		RemovePadding(pos enum.Pos, size enum.Size)

		Release()
	}

	object struct {
		*js.Element
		size enum.Size

		rounded enum.RoundedType
		color   enum.Color
		bgColor enum.Color

		border      enum.BorderType
		borderSize  enum.Size
		borderColor enum.Color

		flexMode enum.FlexMode
		alH      enum.Align
		alV      enum.Align

		fs enum.Size
		fw enum.FontWeight
	}
)

var _ TObject = &object{}

func (obj *object) Ref() *js.Element {
	return obj.Element
}

func (obj *object) Color() enum.Color {
	return obj.color
}

func (obj *object) SetColor(c enum.Color) {
	if obj.color != c && (c == enum.None || c.ApplyText(obj)) {
		obj.color.UnapplyText(obj)
		obj.color = c
	}
}

func (obj *object) Background() enum.Color {
	return obj.bgColor
}

func (obj *object) SetBackground(c enum.Color) {
	if obj.bgColor != c && (c == enum.None || c.ApplyBackground(obj)) {
		obj.bgColor.UnapplyBackground(obj)
		obj.bgColor = c
	}
}

func (obj *object) FlexMode() enum.FlexMode {
	return obj.flexMode
}

func (obj *object) SetFlexMode(m enum.FlexMode) {
	if obj.flexMode != m && m.Apply(obj) {
		obj.flexMode.Unapply(obj)
		obj.flexMode = m
	}
}

func (obj *object) AlignHorizontal(al enum.Align) {
	if obj.alH != al && (al == enum.AlignNone || al.ApplyHorizontal(obj)) {
		obj.alH.UnapplyHorizontal(obj)
		obj.alH = al
	}
}

func (obj *object) AlignVertical(al enum.Align) {
	if obj.alV != al && (al == enum.AlignNone || al.ApplyVertical(obj)) {
		obj.alV.UnapplyVertical(obj)
		obj.alV = al
	}

}

func (obj *object) FontSize() enum.Size {
	return obj.fs
}

func (obj *object) SetFontSize(s enum.Size) {
	if obj.fs != s && (s == enum.N0 || s.ApplyFont(obj)) {
		obj.fs.UnapplyFont(obj)
		obj.fs = s
	}
}

func (obj *object) FontWeight() enum.FontWeight {
	return obj.fw
}

func (obj *object) SetFontWeight(w enum.FontWeight) {
	if obj.fw != w && (w == enum.FontWeightNone || w.Apply(obj)) {
		obj.fw.Unapply(obj)
		obj.fw = w
	}
}

func (obj *object) Rounded() enum.RoundedType {
	return obj.rounded
}

func (obj *object) SetRounded(r enum.RoundedType) {
	if obj.rounded != r && (r == enum.RoundedNone || r.Apply(obj)) {
		obj.rounded.Unapply(obj)
		obj.rounded = r
	}
}

func (obj *object) Size() enum.Size {
	return obj.size
}

func (obj *object) SetSize(s enum.Size) {
	obj.size = s
}

func (obj *object) Append(children ...TObject) {
	for i := range children {
		obj.Element.Append(children[i])
	}
}

func (obj *object) Prepend(children ...TObject) {
	for i := range children {
		obj.Element.Prepend(children[i])
	}
}

func (obj *object) SetMargin(pos enum.Pos, size enum.Size) {
	obj.Element.Add(pos.Margin(size))
}

func (obj *object) RemoveMargin(pos enum.Pos, size enum.Size) {
	obj.Element.Remove(pos.Margin(size))
}

func (obj *object) SetPadding(pos enum.Pos, size enum.Size) {
	obj.Element.Add(pos.Padding(size))
}

func (obj *object) RemovePadding(pos enum.Pos, size enum.Size) {
	obj.Element.Remove(pos.Padding(size))
}

func (obj *object) Border() enum.BorderType {
	return obj.border
}

func (obj *object) SetBorder(b enum.BorderType) {
	if obj.border != b && (b == enum.BorderNone || b.Apply(obj)) {
		obj.border.Unapply(obj)
		obj.border = b
	}
}

func (obj *object) BorderSize() enum.Size {
	return obj.borderSize
}

func (obj *object) SetBorderSize(s enum.Size) {
	if obj.borderSize != s && (s == enum.N0 || s.ApplyBorder(obj)) {
		obj.borderSize.UnapplyBorder(obj)
		obj.borderSize = s
	}
}

func (obj *object) BorderColor() enum.Color {
	return obj.borderColor
}

func (obj *object) SetBorderColor(c enum.Color) {
	if obj.borderColor != c && (c == enum.None || c.ApplyBorder(obj)) {
		obj.borderColor.UnapplyBorder(obj)
		obj.borderColor = c
	}
}

func newObject(el *js.Element) *object {

	return &object{
		Element: el,
		color:   enum.None,
	}
}

func Object(el *js.Element) TObject {
	return newObject(el)
}
