// +build js,wasm

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {

	data := map[Color]string{
		Primary:   "primary",
		Secondary: "secondary",
		Success:   "success",
		Info:      "info",
		Warning:   "warning",
		Danger:    "danger",
		Light:     "light",
		Dark:      "dark",
		Link:      "link",
	}

	var typ Color
	for k, v := range data {
		assert.Equal(t, v, k.String())
		typ.SetString(v)
		assert.Equal(t, k, typ)
		assert.Equal(t, "btn-outline-"+k.String(), k.Style("btn", "outline"))
	}

	assert.Equal(t, None.String(), Color(1000).String())
	typ.SetString("other color")
	assert.Equal(t, None, typ)

}
