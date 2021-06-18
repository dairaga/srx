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
