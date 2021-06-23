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

func Alert(value js.Wrapper) *BS {
	val := bootstrap.Get("Alert").Call("getOrCreateInstance", value)
	return &BS{Value: val}
}
