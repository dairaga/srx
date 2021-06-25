// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/bootstrap"
	"github.com/dairaga/srx/js"
)

type (
	TCollapse interface {
		TComponent
		Bind(el *js.Element)
		Toggle()
		Show()
		Hide()

		OnShowing(func(TCollapse, js.TEvent))
		OnShown(func(TCollapse, js.TEvent))
		OnHiding(func(TCollapse, js.TEvent))
		OnHidden(func(TCollapse, js.TEvent))
	}

	collapse struct {
		*component
		bs *bootstrap.BS
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

func (c *collapse) Toggle() {
	c.bs.Call("toggle")
}

func (c *collapse) Show() {
	c.bs.Call("show")
}

func (c *collapse) Hide() {
	c.bs.Call("hide")
}

func (c *collapse) Release() {
	c.bs.Dispose()
	c.component.Release()
}

func (c *collapse) on(evt string, fn func(TCollapse, js.TEvent)) {
	c.On(evt, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		sender := Lookup(this).(*collapse)
		event := js.EventOf(args[0])
		fn(sender, event)
		return nil
	}))
}

func (c *collapse) OnShowing(fn func(TCollapse, js.TEvent)) {
	// TOFix: bootstrap do not invoke showing callback.
	c.on("show.bs.collapse", fn)
}

func (c *collapse) OnShown(fn func(TCollapse, js.TEvent)) {
	c.on("shown.bs.collapse", fn)
}

func (c *collapse) OnHiding(fn func(TCollapse, js.TEvent)) {
	c.on("hide.bs.collapse", fn)
}

func (c *collapse) OnHidden(fn func(TCollapse, js.TEvent)) {
	c.on("hidden.bs.collapse", fn)
}

func newCollapse(owner TComponent) *collapse {
	el := js.Create("div").Add("collapse")
	ret := &collapse{
		component: newComponent(el),
		bs:        bootstrap.Collapse(el),
	}

	bindOwner(owner, ret)
	return ret
}

func Collapse(owner TComponent) TCollapse {
	return newCollapse(owner)
}
