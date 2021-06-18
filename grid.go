// build js,wasm

package srx

import (
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TGridPanel interface {
		TComponent
		//AlignHorizontal(al enum.Align)
		//AlignVertical(al enum.Align)
		AddCell(span enum.Size, children ...TObject) TCell
		TCeller
	}

	grid struct {
		*component
		//h     enum.Align
		//v     enum.Align
		cells []TCell
	}
)

var _ TGridPanel = &grid{}

// -----------------------------------------------------------------------------

//func (g *grid) AlignHorizontal(al enum.Align) {
//	if h := g.h.Horizontal(); h != "" {
//		g.Element.Remove(h)
//	}
//
//	if h := al.Horizontal(); h != "" {
//		g.Element.Add(h)
//	}
//	g.h = al
//}
//
//// -----------------------------------------------------------------------------
//
//func (g *grid) AlignVertical(al enum.Align) {
//	if v := g.v.Vertical(); v != "" {
//		g.Element.Remove(v)
//	}
//	if v := al.Vertical(); v != "" {
//		g.Element.Add(v)
//	}
//	g.v = al
//}

// -----------------------------------------------------------------------------

func (g *grid) AddCell(span enum.Size, children ...TObject) TCell {
	cell := Cell(children...)
	cell.Ref().Add(span.Col())
	g.Element.Append(cell)
	g.cells = append(g.cells, cell)
	return cell
}

// -----------------------------------------------------------------------------

func (g *grid) Append(children ...TObject) {
	g.AddCell(enum.N0, children...)
}

// -----------------------------------------------------------------------------

func (g *grid) Prepend(children ...TObject) {
	g.AddCell(enum.N0, children...)
}

// -----------------------------------------------------------------------------

func (g *grid) Cell(index int) TCell {
	if index >= 0 && index < len(g.cells) {
		return g.cells[index]
	}
	return nil
}

// -----------------------------------------------------------------------------

func (g *grid) Cells() []TCell {
	return g.cells
}

// -----------------------------------------------------------------------------

func (g *grid) CellLen() int {
	return len(g.cells)
}

// -----------------------------------------------------------------------------

func GridPanel(owner TComponent) TGridPanel {
	ret := &grid{
		component: newComponent(owner, js.Create("div").Add("row")),
		//h:         enum.AlignNone,
		//v:         enum.AlignNone,
	}
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
