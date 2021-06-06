// +build js,wasm

package std

import (
	"testing"

	"github.com/dairaga/srx"
	"github.com/dairaga/srx/enum"
	"github.com/stretchr/testify/assert"
)

func TestButtonCaption(t *testing.T) {
	caption := "TestButton"
	btn := ButtonOf(srx.Root())
	srx.Root().Append(btn)

	btn.SetCaption(caption)
	assert.Equal(t, caption, btn.Caption())
	assert.Equal(t, caption, btn.(*button).Text())

	assert.Equal(t, enum.Button, btn.Type())

	btn.SetType(enum.Submit)
	assert.Equal(t, enum.Submit, btn.Type())

	btn.SetType(enum.ButtonType(1000))
	assert.Equal(t, enum.Button, btn.Type())
}
