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
		AlignHorizontal(al enum.Align)
		AlignVertical(al enum.Align)
		AppendCol(child srx.TComponent, span enum.Size)
	}

	grid struct {
		*srx.Component
		h enum.Align
		v enum.Align
	}
)

var _ TGridPanel = &grid{}

// -----------------------------------------------------------------------------

func (g *grid) AlignHorizontal(al enum.Align) {
	if h := g.h.Horizontal(); h != "" {
		g.Element.Remove(h)
	}

	if h := al.Horizontal(); h != "" {
		g.Element.Add(h)
	}
	g.h = al
}

// -----------------------------------------------------------------------------

func (g *grid) AlignVertical(al enum.Align) {
	if v := g.v.Vertical(); v != "" {
		g.Element.Remove(v)
	}
	if v := al.Vertical(); v != "" {
		g.Element.Add(v)
	}
	g.v = al
}

// -----------------------------------------------------------------------------

func (g *grid) AppendCol(child srx.TComponent, span enum.Size) {
	child.Ref().Add(span.Col())
	g.Element.Append(child)
}

// -----------------------------------------------------------------------------

func GridPannelOf(owner srx.TComponent) TGridPanel {
	ret := &grid{
		Component: srx.NewComponent(owner, js.Create("div").Add("row")),
		h:         enum.AlignNone,
		v:         enum.AlignNone,
	}
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
