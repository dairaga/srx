// +build js,wasm

package srx

import "github.com/dairaga/srx/js"

var root = newComponent(js.ElementOf(js.Global().Get("document").Get("body")))

func Root() TComponent {
	return root
}

func bindOwner(owner, child TComponent) {
	if owner == nil {
		owner = root
	}
	child.setOwner(owner)
	owner.Add(child)
}

func Append(children ...TObject) {
	root.Append(children...)
}
