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
		//enum.BorderNone,
	}

	for _, v := range data {
		old := panel.Border()
		panel.SetBorder(v)
		assert.True(t, panel.Ref().Contains(v.String()))
		if old != enum.BorderNone {
			assert.False(t, panel.Ref().Contains(old.String()))
		}
	}
}
