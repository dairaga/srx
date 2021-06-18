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
		size    enum.Size
		color   enum.Color
		bgColor enum.Color

		flexMode enum.FlexMode
		alH      enum.Align
		alV      enum.Align
	}
)

var _ TObject = &object{}

// -----------------------------------------------------------------------------

func (obj *object) Ref() *js.Element {
	return obj.Element
}

// -----------------------------------------------------------------------------

func (obj *object) Color() enum.Color {
	return obj.color
}

// -----------------------------------------------------------------------------

func (obj *object) SetColor(c enum.Color) {
	if obj.color != c && (c == enum.None || c.ApplyTextColor(obj)) {
		obj.color.UnapplyTextColor(obj)
		obj.color = c
	}
}

// -----------------------------------------------------------------------------

func (obj *object) Background() enum.Color {
	return obj.bgColor
}

// -----------------------------------------------------------------------------

func (obj *object) SetBackground(c enum.Color) {
	if obj.bgColor != c && (c == enum.None || c.ApplyBackground(obj)) {
		obj.bgColor.UnapplyBackground(obj)
		obj.bgColor = c
	}
}

// -----------------------------------------------------------------------------

func (obj *object) FlexMode() enum.FlexMode {
	return obj.flexMode
}

// -----------------------------------------------------------------------------

func (obj *object) SetFlexMode(m enum.FlexMode) {
	if obj.flexMode != m && m.Apply(obj) {
		obj.flexMode.Unapply(obj)
		obj.flexMode = m
	}
}

// -----------------------------------------------------------------------------

func (obj *object) AlignHorizontal(al enum.Align) {
	if obj.alH != al && (al == enum.AlignNone || al.ApplyHorizontal(obj)) {
		obj.alH.UnapplyHorizontal(obj)
		obj.alH = al
	}
}

// -----------------------------------------------------------------------------

func (obj *object) AlignVertical(al enum.Align) {
	if obj.alV != al && (al == enum.AlignNone || al.ApplyVertical(obj)) {
		obj.alV.UnapplyVertical(obj)
		obj.alV = al
	}

}

// -----------------------------------------------------------------------------

func (obj *object) Size() enum.Size {
	return obj.size
}

// -----------------------------------------------------------------------------

func (obj *object) SetSize(s enum.Size) {
	obj.size = s
}

// -----------------------------------------------------------------------------

func (obj *object) Append(children ...TObject) {
	for i := range children {
		obj.Element.Append(children[i])
	}
}

// -----------------------------------------------------------------------------

func (obj *object) Prepend(children ...TObject) {
	for i := range children {
		obj.Element.Prepend(children[i])
	}
}

// -----------------------------------------------------------------------------

func (obj *object) SetMargin(pos enum.Pos, size enum.Size) {
	obj.Element.Add(pos.Margin(size))
}

// -----------------------------------------------------------------------------

func (obj *object) RemoveMargin(pos enum.Pos, size enum.Size) {
	obj.Element.Remove(pos.Margin(size))
}

// -----------------------------------------------------------------------------

func (obj *object) SetPadding(pos enum.Pos, size enum.Size) {
	obj.Element.Add(pos.Padding(size))
}

// -----------------------------------------------------------------------------

func (obj *object) RemovePadding(pos enum.Pos, size enum.Size) {
	obj.Element.Remove(pos.Padding(size))
}

// -----------------------------------------------------------------------------

func newObject(el *js.Element) *object {

	return &object{
		Element: el,
		color:   enum.None,
	}
}

// -----------------------------------------------------------------------------

func Object(el *js.Element) TObject {
	return newObject(el)
}
