// +build js,wasm

package srx

import "github.com/dairaga/srx/js"

const (
	srxTattoo = `data-srx`
)

type (
	// TObject represents vdom object.
	// It is a js html element with tattoo for seeking object.
	TObject interface {
		js.Wrapper

		OK() bool

		// ID returns js attribute id.
		ID() string

		// Disabled returns true if object is disabled, or returns true.
		Disabled() bool
		Disable()
		Enable()

		Hidden() bool
		Hide()
		Show()

		Append(child TObject)
		Prepend(Child TObject)

		Release()
	}

	Object struct {
		*js.Element
	}
)

var _ TObject = &Object{}

// -----------------------------------------------------------------------------

func (obj *Object) Append(child TObject) {
	obj.Element.Append(child)
}

// -----------------------------------------------------------------------------

func (obj *Object) Prepend(child TObject) {
	obj.Element.Prepend(child)
}

// -----------------------------------------------------------------------------

func NewObject(el *js.Element) *Object {

	return &Object{
		Element: el,
	}
}

// -----------------------------------------------------------------------------

func ObjectOf(el *js.Element) TObject {
	return NewObject(el)
}
