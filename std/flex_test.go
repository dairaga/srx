// +build js,wasm

package std

import (
	"testing"

	"github.com/dairaga/srx"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestFlex(t *testing.T) {

	f := FlexPanelOf(srx.Root())
	f.Ref().SetStyle("height", "100px")
	f1 := PanelOf(f)
	f1.Ref().SetText("item1")
	f.Append(f1)

	f2 := PanelOf(f)
	f2.Ref().SetText("item2")
	f.Append(f2)
	f3 := PanelOf(f)
	f3.Ref().SetText("item3")
	f.Append(f3)

	js.Append(f)

	/* init */
	assert.True(t, f.Ref().Contains("d-flex"))
	assert.True(t, f.Ref().Contains("flex-row"))
	assert.Equal(t, enum.FlexModeRow, f.Mode())
	assert.Equal(t, enum.AlignNone, f.(*flex).h)
	assert.Equal(t, enum.AlignNone, f.(*flex).v)

	/* mode */
	f.SetMode(enum.FlexModeColumn)
	assert.Equal(t, enum.FlexModeColumn, f.Mode())
	assert.False(t, f.Ref().Contains("flex-row"))
	assert.True(t, f.Ref().Contains("flex-column"))

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

	for _, a := range al {
		oldh := f.(*flex).h
		oldv := f.(*flex).v

		f.AlignHorizontal(a)
		f.AlignVertical(a)

		if a.IsHorizontal() {
			assert.True(t, f.Ref().Contains(a.Horizontal()))
			if oldh != enum.AlignNone {
				assert.False(t, f.Ref().Contains(oldh.Horizontal()))
			}
		}

		if a.IsVertical() {
			assert.True(t, f.Ref().Contains(a.Vertical()))
			if oldv != enum.AlignNone {
				assert.False(t, f.Ref().Contains(oldv.Vertical()))
			}
		}

	}
}
