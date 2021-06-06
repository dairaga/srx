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

	/* Caption */
	btn.SetCaption(caption)
	assert.Equal(t, caption, btn.Caption())
	assert.Equal(t, caption, btn.(*button).Text())

	/* Button Type */
	assert.Equal(t, enum.Button, btn.Type())

	btn.SetType(enum.Submit)
	assert.Equal(t, enum.Submit, btn.Type())

	btn.SetType(enum.Reset)
	assert.Equal(t, enum.Reset, btn.Type())

	btn.SetType(enum.ButtonType(1000))
	assert.Equal(t, enum.Button, btn.Type())
	assert.Equal(t, "button", btn.(*button).Attr("type"))

	/* Value */
	assert.Equal(t, "", btn.Value())

	value := "abc"
	btn.SetValue(value)
	assert.Equal(t, value, btn.Value())

	/* Color */
	color := enum.Primary
	assert.Equal(t, color, btn.Color())
	assert.True(t, btn.(*button).Contains(color.Style("btn")))

	color = enum.Danger
	btn.SetColor(color)
	assert.Equal(t, color, btn.Color())
	assert.True(t, btn.(*button).Contains(color.Style("btn")))

	/* Outline */
	assert.False(t, btn.Outline())
	btn.SetOutline(true)
	assert.True(t, btn.Outline())
	btn.(*button).Contains(color.Style("btn", "outline"))
}
