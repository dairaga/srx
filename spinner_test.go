// +build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestSpinner(t *testing.T) {
	desc := ""
	color := enum.Dark
	typ := enum.SpinnerBorder

	s := Spinner(Root())
	js.Append(s)

	/* init */
	assert.Equal(t, desc, s.Description())
	assert.Equal(t, desc, s.(*spinner).Query(".visually-hidden").Text())

	assert.Equal(t, s.Color(), color)
	assert.True(t, s.(*spinner).Contains(color.Style("text")))

	assert.Equal(t, typ, s.Type())
	assert.True(t, s.(*spinner).Contains(typ.String()))

	/* assist */
	desc = `Loading...`
	s.SetDescription(desc)
	assert.Equal(t, desc, s.Description())
	assert.Equal(t, desc, s.(*spinner).Query(".visually-hidden").Text())

	/* color */
	color = enum.Primary
	s.SetColor(color)
	assert.Equal(t, s.Color(), color)
	assert.True(t, s.(*spinner).Contains(color.Style("text")))

	/* type */
	typ = enum.SpinnerGrow
	s.SetType(typ)
	assert.Equal(t, typ, s.Type())
	assert.True(t, s.(*spinner).Contains(typ.String()))

}
