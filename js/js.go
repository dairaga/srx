// +build js,wasm

package js

import "syscall/js"

type (
	Value   = js.Value
	Wrapper = js.Wrapper
	Type    = js.Type
	Func    = js.Func
)

var (
	ValueOf = js.ValueOf
	FuncOf  = js.FuncOf

	undefined = js.Undefined()

	global   = js.Global()
	document = global.Get("document")
	body     = document.Get("body")
)

func Global() Value {
	return global
}

func IsJSObject(v Wrapper) bool {
	typ := v.JSValue().Type()
	return typ == js.TypeFunction || typ == js.TypeObject
}

func IsJSFunc(v Wrapper) bool {
	return v.JSValue().Type() == js.TypeFunction
}

func Append(x interface{}) {
	body.Call("append", x)
}

func Create(tag string) *Element {
	return ElementOf(document.Call("createElement", tag))
}

func Query(selector string) *Element {
	return ElementOf(document.Call("querySelector", selector))
}
