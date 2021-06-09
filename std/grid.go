// build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TGridPanel interface {
		srx.TComponent
		AppendCol(child srx.TComponent, span enum.Size)
	}

	grid struct {
		*srx.Component
	}
)

var _ TGridPanel = &grid{}

// -----------------------------------------------------------------------------

func (g *grid) AppendCol(child srx.TComponent, span enum.Size) {
	child.Ref().Add(span.Col())
	g.Element.Append(child)
}

// -----------------------------------------------------------------------------

func GridPannelOf(owner srx.TComponent) TGridPanel {
	ret := &grid{
		Component: srx.NewComponent(owner, js.Create("div").Add("row")),
	}
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
