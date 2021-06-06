// +build js,wasm

package srx

import "github.com/dairaga/srx/js"

var root = NewComponent(nil, js.ElementOf(js.Global().Get("document").Get("body")))

func Root() TComponent {
	return root
}
