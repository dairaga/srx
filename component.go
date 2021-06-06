// +build js,wasm

package srx

import "github.com/dairaga/srx/js"

type (
	TComponent interface {
		TObject

		// Tattoo returns a vdom tag.
		Tattoo() string

		Owner() TComponent
		Add(child TComponent)
		Remove(child TComponent)
		Release()
	}

	Component struct {
		*Object
		owner    TComponent
		children map[string]TComponent
	}
)

var _ TComponent = &Component{}

// -----------------------------------------------------------------------------

func (com *Component) Tattoo() string {
	return com.Attr(srxTattoo)
}

// -----------------------------------------------------------------------------

func (com *Component) Owner() TComponent {
	return com.owner
}

// -----------------------------------------------------------------------------

func (com *Component) Add(child TComponent) {
	t := child.Tattoo()
	old, ok := com.children[t]
	if ok && old.OK() {
		old.Release()
	}
	com.children[t] = child
}

// -----------------------------------------------------------------------------

func (com *Component) Remove(child TComponent) {
	delete(com.children, child.Tattoo())
}

// -----------------------------------------------------------------------------

func (com *Component) Release() {
	for _, v := range com.children {
		v.Release()
	}

	if com.owner != nil {
		com.owner.Remove(com)
	}

	com.Object.Release()
}

// -----------------------------------------------------------------------------

func NewComponent(owner TComponent, el *js.Element) *Component {
	el.SetAttr(srxTattoo, tattoo())
	return &Component{
		Object:   NewObject(el),
		owner:    owner,
		children: make(map[string]TComponent),
	}
}

// -----------------------------------------------------------------------------

func ComponentOf(owner TComponent, el *js.Element) TComponent {
	return NewComponent(owner, el)
}
