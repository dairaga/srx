// +build js,wasm

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {

	data := map[Size]string{
		Small:  "sm",
		Medium: "md",
		Large:  "lg",
		Extra:  "xl",
	}

	var typ Size
	for k, v := range data {
		assert.Equal(t, v, k.String())
		assert.Equal(t, "btn-"+k.String(), k.Style("btn"))
		typ.SetString(v)
		assert.Equal(t, k, typ)

	}

	assert.Equal(t, Zero.String(), Size(1000).String())
	typ.SetString("other color")
	assert.Equal(t, Zero, typ)

}
