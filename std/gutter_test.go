// +build js,wasm

package std

import (
	"fmt"
	"testing"

	"github.com/dairaga/srx"
	"github.com/dairaga/srx/el"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestGutter(t *testing.T) {

	g := GutterPanelOf(srx.Root())

	for i := 0; i < 3; i++ {
		caption := el.Caption()
		caption.SetCaption(fmt.Sprintf("item-%d", i+1))
		g.Append(caption)
	}
	js.Append(g)

	/* init */
	assert.True(t, g.Ref().Contains("row"))
	assert.True(t, g.Ref().Contains(enum.N1.Gutter(enum.X)))
	assert.True(t, g.Ref().Contains(enum.N1.Gutter(enum.Y)))
	for _, elm := range g.Ref().QueryAll("div") {
		assert.True(t, elm.Contains(enum.N1PerRow.String()))
	}

	items := []enum.ItemsPerRow{
		enum.N1PerRow,
		enum.N2PerRow,
		enum.N3PerRow,
		enum.N4PerRow,
		enum.N6PerRow,
		enum.N12PerRow,
	}

	size := []enum.Size{
		enum.N0,
		enum.N1,
		enum.N2,
		enum.N3,
		enum.N4,
		enum.N5,
		//enum.N6,
	}

	for i := range items {
		per := items[i]
		count := 12 / int(per)
		g := GutterPanelOf(srx.Root())
		g.SetMargin(enum.Y, enum.N3)
		g.SetPadding(enum.X, enum.N3)

		g.Ref().Add("border")
		g.SetItemsPerRow(per)

		for j := 0; j < count; j++ {
			caption := el.Caption()
			caption.SetCaption(fmt.Sprintf("item-%d", i+1))
			g.Append(caption)
		}
		js.Append(g)

		assert.True(t, g.Ref().Contains("row"))
		assert.True(t, g.Ref().Contains(enum.N1.Gutter(enum.X)))
		assert.True(t, g.Ref().Contains(enum.N1.Gutter(enum.Y)))
		for _, elm := range g.Ref().QueryAll("div") {
			assert.True(t, elm.Contains(per.String()))
		}

		for k := range size {
			s := size[k]
			g.SetGutterSize(s, s)
			assert.True(t, g.Ref().Contains(s.Gutter(enum.X)))
			assert.True(t, g.Ref().Contains(s.Gutter(enum.Y)))
		}
	}
}
