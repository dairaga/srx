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
		Ref() *js.Element

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

		// TODO: apply Bootstrap utility on TObject.
		TextColor() enum.Color
		SetTextColor(c enum.Color)

		Size() enum.Size
		SetSize(s enum.Size)

		SetMargin(pos enum.Pos, size enum.Size)
		RemoveMargin(pos enum.Pos, size enum.Size)

		SetPadding(pos enum.Pos, size enum.Size)
		RemovePadding(pos enum.Pos, size enum.Size)

		Release()
	}

	Object struct {
		*js.Element
		size enum.Size
		text enum.Color
	}
)

var _ TObject = &Object{}

// -----------------------------------------------------------------------------

func (obj *Object) Ref() *js.Element {
	return obj.Element
}

// -----------------------------------------------------------------------------

func (obj *Object) Size() enum.Size {
	return obj.size
}

// -----------------------------------------------------------------------------

func (obj *Object) SetSize(s enum.Size) {
	obj.size = s
}

// -----------------------------------------------------------------------------

func (obj *Object) TextColor() enum.Color {
	return obj.text
}

// -----------------------------------------------------------------------------

func (obj *Object) SetTextColor(c enum.Color) {
	obj.Element.Remove(obj.text.TextColor())
	obj.Element.Add(c.TextColor())
	obj.text = c
}

// -----------------------------------------------------------------------------

func (obj *Object) Append(children ...TObject) {
	for i := range children {
		obj.Element.Append(children[i])
	}
}

// -----------------------------------------------------------------------------

func (obj *Object) Prepend(children ...TObject) {
	for i := range children {
		obj.Element.Prepend(children[i])
	}
}

// -----------------------------------------------------------------------------

func (obj *Object) SetMargin(pos enum.Pos, size enum.Size) {
	obj.Element.Add(pos.Margin(size))
}

// -----------------------------------------------------------------------------

func (obj *Object) RemoveMargin(pos enum.Pos, size enum.Size) {
	obj.Element.Remove(pos.Margin(size))
}

// -----------------------------------------------------------------------------

func (obj *Object) SetPadding(pos enum.Pos, size enum.Size) {
	obj.Element.Add(pos.Padding(size))
}

// -----------------------------------------------------------------------------

func (obj *Object) RemovePadding(pos enum.Pos, size enum.Size) {
	obj.Element.Remove(pos.Padding(size))
}

// -----------------------------------------------------------------------------

func NewObject(el *js.Element) *Object {

	return &Object{
		Element: el,
		text:    enum.None,
	}
}

// -----------------------------------------------------------------------------

func ObjectOf(el *js.Element) TObject {
	return NewObject(el)
}
