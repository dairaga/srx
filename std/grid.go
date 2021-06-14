// build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/el"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TGridPanel interface {
		srx.TComponent
		AlignHorizontal(al enum.Align)
		AlignVertical(al enum.Align)
		AddCell(span enum.Size, children ...srx.TObject) el.TCell
		el.TCeller
	}

	grid struct {
		*srx.Component
		h     enum.Align
		v     enum.Align
		cells []el.TCell
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

func (g *grid) AddCell(span enum.Size, children ...srx.TObject) el.TCell {
	cell := el.Cell(children...)
	cell.Ref().Add(span.Col())
	g.Element.Append(cell)
	g.cells = append(g.cells, cell)
	return cell
}

// -----------------------------------------------------------------------------

func (g *grid) Append(children ...srx.TObject) {
	g.AddCell(enum.N0, children...)
}

// -----------------------------------------------------------------------------

func (g *grid) Prepend(children ...srx.TObject) {
	g.AddCell(enum.N0, children...)
}

// -----------------------------------------------------------------------------

func (g *grid) Cell(index int) el.TCell {
	if index >= 0 && index < len(g.cells) {
		return g.cells[index]
	}
	return nil
}

// -----------------------------------------------------------------------------

func (g *grid) Cells() []el.TCell {
	return g.cells
}

// -----------------------------------------------------------------------------

func (g *grid) CellLen() int {
	return len(g.cells)
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
