// +build js,wasm

package srx

import "github.com/dairaga/srx/js"

type (
	TCollapse interface {
		TComponent
		Bind(el *js.Element)
	}

	collapse struct {
		*component
	}
)

var _ TCollapse = &collapse{}

func (c *collapse) Bind(el *js.Element) {
	if el.TagName() == "a" {
		el.SetAttr("href", "#"+c.ID())
	} else {
		el.SetAttr("data-bs-target", "#"+c.ID())
	}
	el.SetAttr("data-bs-toggle", "collapse").
		SetAttr("aria-expanded", "false").
		SetAttr("aria-controls", c.ID())
}

func newCollapse(owner TComponent) *collapse {
	el := js.Create("div").Add("collapse")
	ret := &collapse{
		component: newComponent(el),
	}

	bindOwner(owner, ret)
	return ret
}

func Collapse(owner TComponent) TCollapse {
	return newCollapse(owner)
}
