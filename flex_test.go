// +build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestFlex(t *testing.T) {

	f := Panel(Root())
	f.Ref().SetStyle("height", "100px")
	f.SetFlexMode(enum.FlexModeRow)

	f1 := Panel(f)
	f1.Ref().SetText("item1")
	f.Append(f1)

	f2 := Panel(f)
	f2.Ref().SetText("item2")
	f.Append(f2)
	f3 := Panel(f)
	f3.Ref().SetText("item3")
	f.Append(f3)

	js.Append(f)

	/* init */
	assert.True(t, f.Ref().Contains("d-flex"))
	assert.True(t, f.Ref().Contains("flex-row"))
	assert.Equal(t, enum.FlexModeRow, f.FlexMode())
	assert.Equal(t, enum.AlignNone, f.(*panel).alH)
	assert.Equal(t, enum.AlignNone, f.(*panel).alH)

	/* mode */
	f.SetFlexMode(enum.FlexModeColumn)
	assert.Equal(t, enum.FlexModeColumn, f.FlexMode())
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
		oldh := f.(*panel).alH
		oldv := f.(*panel).alV

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
