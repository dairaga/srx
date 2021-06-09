// +build js,wasm

package js

import (
	"strings"
	"syscall/js"
)

type (
	Element struct {
		*EventTarget
	}
)

func (el *Element) OK() bool {
	return el.ref.Truthy() && IsJSObject(el.ref)
}

func (el *Element) Prop(p string) Value {
	if !el.OK() {
		return undefined
	}

	return el.ref.Get(p)
}

func (el *Element) SetProp(p string, x interface{}) *Element {
	if el.OK() {
		el.ref.Set(p, x)
	}
	return el
}

func (el *Element) Attr(a string) string {
	if v := el.Call("getAttribute", a); v.Truthy() {
		return v.String()
	}

	return ""
}

func (el *Element) SetAttr(a, v string) *Element {
	el.Call("setAttribute", a, v)
	return el
}

func (el *Element) RemoveAttr(a string) *Element {
	el.Call("removeAttribute", a)
	return el
}

func (el *Element) HasAttr(a string) bool {
	if v := el.Call("hasAttribute", a); v.Type() == js.TypeBoolean {
		return v.Bool()
	}
	return false
}

func (el *Element) AddAttrs(attrs map[string]string) *Element {
	if el.OK() && IsJSFunc(el.Prop("setAttribute")) {
		for k, v := range attrs {
			el.ref.Call("setAttribute", k, v)
		}
	}
	return el
}

func (el *Element) Call(m string, args ...interface{}) Value {
	if el.OK() {
		return el.ref.Call(m, args...)
	}
	return undefined
}

func (el *Element) TagName() string {
	return strings.ToLower(el.Prop("tagName").String())
}

func (el *Element) ID() string {
	return el.Attr("id")
}

func (el *Element) SetID(id string) {
	el.SetAttr("id", id)
}

func (el *Element) Text() string {
	return el.Prop("innerText").String()
}

func (el *Element) SetText(t string) *Element {
	el.SetProp("innerText", t)
	return el
}

func (el *Element) HTML() HTML {
	return HTML(el.Prop("innerHTML").String())
}

func (el *Element) SetHTML(h HTML) *Element {
	el.SetProp("innerHTML", h.String())
	return el
}

func (el *Element) Append(child interface{}) *Element {
	el.Call("append", child)
	return el
}

func (el *Element) Prepend(child interface{}) *Element {
	el.Call("prepend", child)
	return el
}

func (el *Element) inner(prop, method string, args ...string) Value {
	clz := el.Prop(prop)
	if !(clz.Truthy() && IsJSObject(clz)) {
		return undefined
	}

	size := len(args)
	if size == 0 {
		return clz.Call(method)
	}

	filter := make([]string, 0, size)
	for i := range args {
		if args[i] != "" {
			filter = append(filter, args[i])
		}
	}

	args = filter
	size = len(args)

	switch size {
	case 0:
		return undefined
	case 1:
		return clz.Call(method, args[0])
	case 2:
		return clz.Call(method, args[0], args[1])
	case 3:
		return clz.Call(method, args[0], args[1], args[2])
	case 4:
		return clz.Call(method, args[0], args[1], args[2], args[3])
	case 5:
		return clz.Call(method, args[0], args[1], args[2], args[3], args[4])
	}

	x := make([]interface{}, size)
	for i, str := range args {
		x[i] = str
	}
	return clz.Call(method, x...)
}

func (el *Element) class(method string, args ...string) Value {
	return el.inner("classList", method, args...)
}

func (el *Element) Add(names ...string) *Element {
	el.class("add", names...)
	return el
}

func (el *Element) Remove(names ...string) *Element {
	el.class("remove", names...)
	return el
}

func (el *Element) Toggle(name string) *Element {
	el.class("toggle", name)
	return el
}

func (el *Element) Replace(oldName, newName string) *Element {
	el.class("replace", oldName, newName)
	return el
}

func (el *Element) Contains(name string) bool {
	return el.class("contains", name).Bool()
}

func (el *Element) style(method string, args ...string) Value {
	return el.inner("style", method, args...)
}

func (el *Element) Style(p string) string {
	return el.style("getPropertyValue", p).String()
}

func (el *Element) SetStyle(p, v string) *Element {
	el.style("setProperty", p, v)
	return el
}

func (el *Element) RemoveStyle(p string) *Element {
	el.style("removeProperty", p)
	return el
}

func (el *Element) Query(selectors string) *Element {
	return ElementOf(el.Call("querySelector", selectors))
}

func (el *Element) QueryAll(selectors string) Elements {
	els := el.Call("querySelectorAll", selectors)

	if IsJSObject(els) {
		size := els.Length()

		ret := make([]*Element, size)
		for i := 0; i < size; i++ {
			ret[i] = ElementOf(els.Index(i))
		}
		return ret
	}

	return []*Element{}
}

func (el *Element) Disabled() bool {
	return el.HasAttr("disabled")
}

func (el *Element) Enable() {
	el.RemoveAttr("disabled")
}

func (el *Element) Disable() {
	el.SetAttr("disabled", "true")
}

func (el *Element) Hidden() bool {
	return el.Prop("hidden").Bool()
}

func (el *Element) Hide() {
	el.SetProp("hidden", true)
}

func (el *Element) Show() {
	el.SetProp("hidden", false)
}

func (el *Element) Release() {
	if el.OK() {
		// it maybe removed from DOM by other framework.
		parent := el.Prop("parentNode")
		if IsJSObject(parent) {
			parent.Call("removeChild", el)
		}
		el.EventTarget.Release()
	}
}

func ElementOf(v Value) *Element {
	return &Element{
		EventTarget: EventTargetOf(v),
	}
}

func From(tmpl HTML) *Element {
	return tmpl.Element()
}
