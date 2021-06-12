// +build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TGutterPanel interface {
		srx.TComponent
		SetGutterSize(x, y enum.Size)
		SetItemsPerRow(n enum.ItemsPerRow)
	}

	gutter struct {
		*srx.Component
		x     enum.Size
		y     enum.Size
		items enum.ItemsPerRow
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

func (g *gutter) Append(item srx.TObject) {
	item.Ref().Add(g.items.String())
	g.Element.Append(item)
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
