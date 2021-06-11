// +build js,wasm

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlignment(t *testing.T) {

	data := map[Align][]string{
		AlignStart:    {"start", "justify-content-start", "align-items-start"},
		AlignEnd:      {"end", "justify-content-end", "align-items-end"},
		AlignCenter:   {"center", "justify-content-center", "align-items-center"},
		AlignBetween:  {"between", "justify-content-between", ""},
		AlignAround:   {"around", "justify-content-around", ""},
		AlignEvenly:   {"evenly", "justify-content-evenly", ""},
		AlignBaseline: {"baseline", "", "align-items-baseline"},
		AlignStretch:  {"stretch", "", "align-items-stretch"},
		Align(1000):   {"", "", ""},
	}

	for k, v := range data {
		assert.Equal(t, v[0], k.String())
		assert.Equal(t, v[1], k.Horizontal())
		assert.Equal(t, v[2], k.Vertical())
	}

}
