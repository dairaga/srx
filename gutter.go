// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TGutterPanel interface {
		TComponent
		SetGutterSize(x, y enum.Size)
		SetItemsPerRow(n enum.ItemsPerRow)
		AddCell(items ...TObject) TCell

		TCeller
	}

	gutter struct {
		*component
		x     enum.Size
		y     enum.Size
		items enum.ItemsPerRow
		cells []TCell
	}
)

var _ TGutterPanel = &gutter{}

func (g *gutter) SetGutterSize(x, y enum.Size) {
	if x >= enum.N0 && x <= enum.N5 {
		g.Element.Replace(g.x.Gutter(enum.X), x.Gutter(enum.X))
		g.x = x
	}

	if y >= enum.N0 && y <= enum.N5 {
		g.Element.Replace(g.y.Gutter(enum.Y), y.Gutter(enum.Y))
		g.y = y
	}
}

func (g *gutter) SetItemsPerRow(n enum.ItemsPerRow) {
	g.items = n
}

func (g *gutter) AddCell(items ...TObject) TCell {
	cell := Cell(items...)
	cell.Ref().Add(g.items.String())
	g.Element.Append(cell)
	g.cells = append(g.cells, cell)
	return cell
}

func (g *gutter) Append(children ...TObject) {
	g.AddCell(children...)
}

func (g *gutter) Prepend(children ...TObject) {
	g.AddCell(children...)
}

func (g *gutter) Cell(index int) TCell {
	if index >= 0 && index < len(g.cells) {
		return g.cells[index]
	}
	return nil
}

func (g *gutter) Cells() []TCell {
	return g.cells
}

func (g *gutter) CellLen() int {
	return len(g.cells)
}

func GutterPanel(owner TComponent) TGutterPanel {
	el := js.Create("div").Add("row", "gx-1", "gy-1")
	ret := &gutter{
		component: newComponent(owner, el),
		x:         enum.N1,
		y:         enum.N1,
		items:     enum.N1PerRow,
	}
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
