// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/bootstrap"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TAlert interface {
		TComponent
		Close()
		OnClosing(fn func(TAlert, js.TEvent))
		OnClosed(fn func(TAlert, js.TEvent))
	}

	alert struct {
		*component
		bs        *bootstrap.BS
		onClosing func(TAlert, js.TEvent)
		onClosed  func(TAlert, js.TEvent)
	}
)

func (a *alert) SetColor(c enum.Color) {
	if a.color != c && c.ApplyAlert(a) {
		a.color.UnapplyAlert(a)
		a.color = c
	}
}

func (a *alert) Background() enum.Color {
	return a.Color()
}

func (a *alert) SetBackground(c enum.Color) {
	a.SetColor(c)
}

func (a *alert) Close() {
	a.bs.Call("close")
}

func (a *alert) OnClosing(fn func(TAlert, js.TEvent)) {
	a.onClosing = fn
}

func (a *alert) OnClosed(fn func(TAlert, js.TEvent)) {
	a.onClosed = fn
}

func (a *alert) Release() {
	a.bs.Dispose()
	a.component.Release()
}

func newAlert(owner TComponent) *alert {
	el := js.Create("div").Add("alert", "alert-dismissible", "fade", "show").SetAttr("role", "alert")
	el.Append(js.HTML(`<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`))
	ret := &alert{
		component: newComponent(el),
		bs:        bootstrap.Alert(el),
	}

	ret.Ref().On("closed.bs.alert", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		sender := Lookup(this).(*alert)
		evt := js.EventOf(args[0])

		if sender.onClosed != nil {
			sender.onClosed(sender, evt)
		}
		sender.Release()
		return nil
	}))

	ret.Ref().On("close.bs.alert", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		sender := Lookup(this).(*alert)
		evt := js.EventOf(args[0])
		if sender.onClosing != nil {
			sender.onClosing(sender, evt)
		}
		return nil
	}))

	bindOwner(owner, ret)
	return ret
}

func Alert(owner TComponent) TAlert {
	return newAlert(owner)
}
