// +build js,wasm

package bootstrap

import "github.com/dairaga/srx/js"

type (
	BS struct {
		js.Value
	}
)

var bootstrap = js.Global().Get("bootstrap")

func (bs *BS) Dispose() {
	bs.Call("dispose")
}

func getOrCreateInstance(inst string, value js.Wrapper) *BS {
	return &BS{
		Value: bootstrap.Get(inst).Call("getOrCreateInstance", value),
	}
}

func Alert(value js.Wrapper) *BS {
	return getOrCreateInstance("Alert", value)
}

func Collapse(value js.Wrapper) *BS {
	return getOrCreateInstance("Collapse", value)
}
