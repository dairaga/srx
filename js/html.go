// +build js,wasm

package js

type HTML string

func (h HTML) JSValue() Value {
	tmpl := Create("template")
	tmpl.SetHTML(h)
	return tmpl.Prop("content")
}

func (h HTML) String() string {
	return string(h)
}

func (h HTML) Element() *Element {
	return ElementOf(h.JSValue().Get("children").Index(0))
}
