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
		OnClosing(fn func(TAlert))
		OnClosed(fn func(TAlert))
	}

	alert struct {
		*component
		bs        *bootstrap.BS
		onClosing func(TAlert)
		onClosed  func(TAlert)
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

func (a *alert) Release() {
	a.bs.Dispose()
	a.component.Release()
}

func newAlert(owner TComponent) *alert {
	el := js.Create("div").Add("alert").SetAttr("role", "alert")
	ret := &alert{
		component: newComponent(owner, el),
		bs:        bootstrap.Alert(el),
	}

	ret.Ref().On("closed.bs.alert", func(this js.Value, args []js.Value) interface{} {
		return nil
	})

	if owner != nil {
		owner.Add(ret)
	}
	return ret
}

func Alert(owner TComponent) TAlert {
	return newAlert(owner)
}
