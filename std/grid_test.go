// +build js,wasm

package std

import (
	"testing"

	"github.com/dairaga/srx"
	"github.com/dairaga/srx/el"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {

	g := GridPannelOf(srx.Root())
	g.Ref().SetStyle("height", "100px")
	g.AddCell(enum.N2, el.Caption("grid-span-2"))
	g.AddCell(enum.N3, el.Caption("grid-span-3"))
	g.AddCell(enum.N4, el.Caption("grid-span-4"))
	js.Append(g)

	/* init */
	assert.True(t, g.Ref().Contains("row"))
	assert.True(t, g.Cell(0).Ref().Contains("col-2"))
	assert.True(t, g.Cell(1).Ref().Contains("col-3"))
	assert.True(t, g.Cell(2).Ref().Contains("col-4"))

	al := []enum.Align{
		enum.AlignNone,
		enum.AlignStart,
		enum.AlignEnd,
		enum.AlignCenter,
		enum.AlignBetween,
		enum.AlignAround,
		enum.AlignEvenly,
		enum.AlignBaseline,
		enum.AlignStretch,
	}

	old := enum.AlignNone
	for _, v := range al {
		if v.IsHorizontal() {
			old = g.(*grid).h
			g.AlignHorizontal(v)
			assert.True(t, g.Ref().Contains(v.Horizontal()))
			if old != enum.AlignNone {
				assert.False(t, g.Ref().Contains(old.Horizontal()))
			}
		} else if v.IsVertical() {
			old = g.(*grid).v
			g.AlignVertical(v)
			assert.True(t, g.Ref().Contains(v.Vertical()))
			if old != enum.AlignNone {
				assert.False(t, g.Ref().Contains(old.Vertical()))
			}
		}

	}
}
