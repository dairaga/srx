// +build js,wasm

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {

	data := map[Size]string{
		None:   "",
		Small:  "sm",
		Medium: "md",
		Large:  "lg",
		Extra:  "xl",
	}

	var typ Size
	for k, v := range data {
		assert.Equal(t, v, k.String())
		typ.SetString(v)
		assert.Equal(t, k, typ)
		if k != None {
			assert.Equal(t, "btn-"+k.String(), k.Style("btn"))
		} else {
			assert.Equal(t, "", k.Style("btn"))
		}

	}

	assert.Equal(t, None.String(), Size(1000).String())
	typ.SetString("other color")
	assert.Equal(t, None, typ)

}
