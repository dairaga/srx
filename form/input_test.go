// +build js,wasm

package form

import (
	"testing"

	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	typ := ""
	name := ""
	value := ""
	readonly := false
	required := false

	input := InputOf(srx.Root())
	js.Append(input)

	/* init */
	assert.Equal(t, typ, input.Type())
	assert.Equal(t, name, input.Name())
	assert.Equal(t, value, input.Value())
	assert.Equal(t, readonly, input.ReadOnly())
	assert.Equal(t, required, input.Required())

	typ = "text"
	name = "test_input"
	value = "test_input_value"
	readonly = true
	required = true

	input.SetType(typ)
	input.SetName(name)
	input.SetValue(value)
	input.SetReadOnly(readonly)
	input.SetRequired(required)

	assert.Equal(t, typ, input.Type())
	assert.Equal(t, name, input.Name())
	assert.Equal(t, value, input.Value())
	assert.Equal(t, readonly, input.ReadOnly())
	assert.Equal(t, required, input.Required())
}
