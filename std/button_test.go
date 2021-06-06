// +build js,wasm

package std

import (
	"testing"

	"github.com/dairaga/srx"
	"github.com/stretchr/testify/assert"
)

func TestButtonCaption(t *testing.T) {
	caption := "TestButton"
	btn := ButtonOf(srx.Root())
	srx.Root().Append(btn)

	btn.SetCaption(caption)
	assert.Equal(t, caption, btn.Caption())
	assert.Equal(t, caption, btn.(*button).Text())
}
