// +build js,wasm

package std

import (
	"testing"

	"github.com/dairaga/srx"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {

	g := GridPannelOf(srx.Root())
	g.Ref().SetStyle("height", "100px")
	c1 := PanelOf(g)
	c1.Ref().SetText("c1")
	c2 := PanelOf(g)
	c2.Ref().SetText("c2")
	c3 := PanelOf(g)
	c3.Ref().SetText("c3")

	g.AppendCol(c1, enum.N2)
	g.AppendCol(c2, enum.N3)
	g.AppendCol(c3, enum.N4)
	js.Append(g)

	/* init */
	assert.True(t, g.Ref().Contains("row"))
	assert.True(t, c1.Ref().Contains("col-2"))
	assert.True(t, c2.Ref().Contains("col-3"))
	assert.True(t, c3.Ref().Contains("col-4"))

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
