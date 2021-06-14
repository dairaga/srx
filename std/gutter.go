// +build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/el"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TGutterPanel interface {
		srx.TComponent
		SetGutterSize(x, y enum.Size)
		SetItemsPerRow(n enum.ItemsPerRow)
		Insert(items ...srx.TObject) el.TCell
		Cell(index int) el.TCell
	}

	gutter struct {
		*srx.Component
		x     enum.Size
		y     enum.Size
		items enum.ItemsPerRow
		cells []el.TCell
	}
)

var _ TGutterPanel = &gutter{}

// -----------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------

func (g *gutter) SetItemsPerRow(n enum.ItemsPerRow) {
	g.items = n
}

// -----------------------------------------------------------------------------

func (g *gutter) Insert(items ...srx.TObject) el.TCell {
	cell := el.Cell()
	for i := range items {
		cell.Append(items[i])
	}
	cell.Ref().Add(g.items.String())
	g.Element.Append(cell)
	g.cells = append(g.cells, cell)
	return cell
}

// -----------------------------------------------------------------------------

func (g *gutter) Append(children ...srx.TObject) {
	g.Insert(children...)
}

// -----------------------------------------------------------------------------

func (g *gutter) Prepend(children ...srx.TObject) {
	g.Insert(children...)
}

// -----------------------------------------------------------------------------

func (g *gutter) Cell(index int) el.TCell {
	if index >= 0 && index < len(g.cells) {
		return g.cells[index]
	}
	return nil
}

// -----------------------------------------------------------------------------

func GutterPanelOf(owner srx.TComponent) TGutterPanel {
	el := js.Create("div").Add("row", "gx-1", "gy-1")
	ret := &gutter{
		Component: srx.NewComponent(owner, el),
		x:         enum.N1,
		y:         enum.N1,
		items:     enum.N1PerRow,
	}
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
