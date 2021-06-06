// +build js,wasm

package std

import (
	"testing"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestButtonCaption(t *testing.T) {
	caption := "TestButton"
	btn := ButtonOf(nil)
	btn.SetCaption(caption)

	js.Append(btn)
	assert.Equal(t, caption, btn.Caption())
	assert.Equal(t, caption, btn.(*button).Text())
}
