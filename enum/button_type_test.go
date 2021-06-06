// +build js,wasm

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestButtonType(t *testing.T) {
	assert.Equal(t, "button", Button.String())
	assert.Equal(t, "submit", Submit.String())
	assert.Equal(t, "reset", Reset.String())
	assert.Equal(t, "button", ButtonType(1000).String())

	var typ ButtonType

	typ.SetString("button")
	assert.Equal(t, Button, typ)

	typ.SetString("submit")
	assert.Equal(t, Submit, typ)

	typ.SetString("reset")
	assert.Equal(t, Reset, typ)

	typ.SetString("no button")
	assert.Equal(t, Button, typ)
}
