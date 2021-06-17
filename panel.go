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

// -----------------------------------------------------------------------------

func PanelOf(owner TComponent) TPanel {
	ret := &panel{
		component: newComponent(owner, js.Create("div")),
	}
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
