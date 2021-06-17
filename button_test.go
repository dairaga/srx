// +build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestButton(t *testing.T) {
	caption := "TestButton"
	btn := Button(Root())
	Root().Append(btn)

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
	color := enum.None
	assert.Equal(t, color, btn.Color())

	color = enum.Primary
	btn.SetColor(color)
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

	/* Size */
	w := btn.(*button).Prop("offsetWidth").Int()
	h := btn.(*button).Prop("offsetHeight").Int()
	assert.NotEqual(t, 0, w)
	assert.NotEqual(t, 0, h)

	btn.SetSize(enum.Large)
	assert.True(t, btn.(*button).Contains(enum.Large.Style("btn")))
	btn.SetSize(enum.N0)
	assert.False(t, btn.(*button).Contains(enum.Large.Style("btn")))
	assert.Equal(t, w, btn.(*button).Prop("offsetWidth").Int())
	assert.Equal(t, h, btn.(*button).Prop("offsetHeight").Int())

	/* Click */
	count := 0
	caption = "Button Clicked"
	btn.OnClick(func(sender TButton, evt js.TEvent) {
		sender.SetCaption(caption)
		count = 1
	})
	btn.Click()
	assert.Equal(t, 1, count)
	assert.Equal(t, caption, btn.Caption())
}
