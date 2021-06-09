// +build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TPanel interface {
		srx.TComponent
	}

	panel struct {
		*srx.Component
	}
)

var _ TPanel = &panel{}

// -----------------------------------------------------------------------------

func PanelOf(owner srx.TComponent) TPanel {
	ret := &panel{
		Component: srx.NewComponent(owner, js.Create("div")),
	}
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
