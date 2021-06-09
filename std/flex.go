// build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TFlexPanel interface {
		srx.TComponent
	}

	flex struct {
		*srx.Component
	}
)

var _ TFlexPanel = &flex{}

// -----------------------------------------------------------------------------

func FlexPannelOf(owner srx.TComponent) TFlexPanel {
	ret := &grid{
		Component: srx.NewComponent(owner, js.Create("div").Add("d-flex")),
	}
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
