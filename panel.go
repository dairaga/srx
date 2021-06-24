// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/js"
)

type (
	TPanel interface {
		TComponent
	}

	panel struct {
		*component
	}
)

var _ TPanel = &panel{}

func newPanel(owner TComponent) *panel {
	ret := &panel{
		component: newComponent(js.Create("div")),
	}
	bindOwner(owner, ret)
	return ret
}

func Panel(owner TComponent) TPanel {
	return newPanel(owner)
}
