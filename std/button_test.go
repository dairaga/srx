// +build js,wasm

package std

import (
	"testing"

	"github.com/dairaga/srx"
	"github.com/dairaga/srx/enum"
	"github.com/stretchr/testify/assert"
)

func TestButton(t *testing.T) {
	caption := "TestButton"
	btn := ButtonOf(srx.Root())
	srx.Root().Append(btn)

	btn.SetCaption(caption)
	assert.Equal(t, caption, btn.Caption())
	assert.Equal(t, caption, btn.(*button).Text())

	assert.Equal(t, enum.Button, btn.Type())

	btn.SetType(enum.Submit)
	assert.Equal(t, enum.Submit, btn.Type())

	btn.SetType(enum.Reset)
	assert.Equal(t, enum.Reset, btn.Type())

	btn.SetType(enum.ButtonType(1000))
	assert.Equal(t, enum.Button, btn.Type())
	assert.Equal(t, "button", btn.(*button).Attr("type"))

	assert.Equal(t, "", btn.Value())

	value := "abc"
	btn.SetValue(value)
	assert.Equal(t, value, btn.Value())

	color := enum.Primary
	assert.Equal(t, color, btn.Color())
	btn.(*button).Contains(color.Style("btn"))

	color = enum.Danger
	btn.SetColor(color)
	assert.Equal(t, color, btn.Color())
	btn.(*button).Contains(color.Style("btn"))
}
