// build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestObjectFontSize(t *testing.T) {
	panel := Panel(Root())
	js.Append(panel)

	panel.Append(Caption("test font size"))

	for i := enum.N1; i <= enum.N6; i++ {
		old := panel.FontSize()
		panel.SetFontSize(i)
		assert.True(t, panel.Ref().Contains("fs-"+i.String()))
		if old != enum.N0 {
			assert.False(t, panel.Ref().Contains("fs-"+old.String()))
		}
	}
}

func TestObjectBorder(t *testing.T) {
	panel := Panel(Root())
	js.Append(panel)

	data := []enum.BorderType{
		enum.Border,
		enum.BorderTop,
		enum.BorderEnd,
		enum.BorderBottom,
		enum.BorderStart,
		enum.Border0,
		enum.BorderTop0,
		enum.BorderEnd0,
		enum.BorderBottom0,
		enum.BorderStart0,
		enum.BorderNone,
	}

	for _, v := range data {
		old := panel.Border()
		panel.SetBorder(v)
		if v != enum.BorderNone {
			assert.True(t, panel.Ref().Contains(v.String()))
		}
		if old != enum.BorderNone {
			assert.False(t, panel.Ref().Contains(old.String()))
		}
	}
}

func TestFontWeight(t *testing.T) {
	caption := Caption("test font weight")

	js.Append(caption)

	data := []enum.FontWeight{
		enum.FontWeightBold,
		enum.FontWeightBolder,
		enum.FontWeightNormal,
		enum.FontWeightLight,
		enum.FontWeightLighter,
		enum.FontWeightNone,
	}

	for _, fw := range data {
		old := caption.FontWeight()
		caption.SetFontWeight(fw)
		if fw != enum.FontWeightNone {
			assert.True(t, caption.Ref().Contains(fw.String()))
		}
		if old != enum.FontWeightNone {
			assert.False(t, caption.Ref().Contains(old.String()))
		}
	}

}

func TestObjectItalic(t *testing.T) {
	caption := Caption("Test Italic")
	js.Append(caption)

	assert.False(t, caption.Italic())

	caption.SetItalic(true)
	assert.True(t, caption.Italic())
	assert.True(t, caption.Ref().Contains("fst-italic"))

	caption.SetItalic(false)
	assert.False(t, caption.Italic())
	assert.False(t, caption.Ref().Contains("fst-italic"))
}

func TestObjectDecoration(t *testing.T) {
	caption := Caption("Test Decoration")
	js.Append(caption)
	assert.Equal(t, enum.DecorationReset, caption.Decoration())

	data := []enum.Decoration{
		enum.DecorationNone,
		enum.DecorationUnderline,
		enum.DecorationLineThrough,
		enum.DecorationReset,
	}

	for _, d := range data {
		old := caption.Decoration()
		caption.SetDecoration(d)
		if d != enum.DecorationReset {
			assert.True(t, caption.Ref().Contains(d.String()))
		}
		assert.Equal(t, d, caption.Decoration())
		if old != enum.DecorationReset {
			assert.False(t, caption.Ref().Contains(old.String()))
		}
	}
}
