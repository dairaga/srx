// +build js,wasm

package srx

import "github.com/dairaga/srx/js"

type (
	TComponent interface {
		TObject

		// Tattoo returns a vdom tag.
		Tattoo() string

		Owner() TComponent
		setOwner(owner TComponent)
		Add(child TComponent)
		Remove(child TComponent)
		Release()
	}

	component struct {
		*object
		owner    TComponent
		children map[string]TComponent
	}
)

var _ TComponent = &component{}

func (com *component) Tattoo() string {
	return com.Attr(srxTattoo)
}

func (com *component) Owner() TComponent {
	return com.owner
}

func (com *component) setOwner(owner TComponent) {
	if com.owner != nil {
		com.owner.Remove(com)
	}
	com.owner = owner
}

func (com *component) Add(child TComponent) {
	if child != nil && child.OK() {
		t := child.Tattoo()
		old, ok := com.children[t]
		if ok && old.OK() {
			old.Release()
		}
		com.children[t] = child
		mem.alloc(child)
	}
}

func (com *component) Remove(child TComponent) {
	delete(com.children, child.Tattoo())
}

func (com *component) Release() {
	for _, v := range com.children {
		v.Release()
	}

	if com.owner != nil {
		com.owner.Remove(com)
	}
	mem.free(com)
	com.object.Release()
}

func newComponent(el *js.Element) *component {
	el.SetAttr(srxTattoo, tattoo())
	return &component{
		object:   newObject(el),
		owner:    nil,
		children: make(map[string]TComponent),
	}
}

func Component(el *js.Element) TComponent {
	return newComponent(el)
}
